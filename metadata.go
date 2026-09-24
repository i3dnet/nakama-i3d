package fleetmanager

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// cloneMetadata retains the concrete Go values used for provider routing and
// callback delivery while detaching mutable JSON-visible data from the caller.
func cloneMetadata(metadata map[string]any) (map[string]any, error) {
	if _, err := json.Marshal(metadata); err != nil {
		return nil, err
	}
	if err := validateMetadataCopy(reflect.ValueOf(metadata), make(map[metadataReference]bool)); err != nil {
		return nil, err
	}
	return copyMetadataValue(reflect.ValueOf(metadata), make(map[metadataReference]reflect.Value)).Interface().(map[string]any), nil
}

type metadataReference struct {
	typ     reflect.Type
	pointer uintptr
	length  int
}

func copyMetadataValue(value reflect.Value, seen map[metadataReference]reflect.Value) reflect.Value {
	switch value.Kind() {
	case reflect.Interface:
		if value.IsNil() {
			return reflect.Zero(value.Type())
		}
		copied := reflect.New(value.Type()).Elem()
		copied.Set(copyMetadataValue(value.Elem(), seen))
		return copied
	case reflect.Map, reflect.Slice, reflect.Pointer:
		if value.IsNil() {
			return reflect.Zero(value.Type())
		}
		key := metadataReference{typ: value.Type(), pointer: uintptr(value.UnsafePointer())}
		if value.Kind() == reflect.Slice {
			key.length = value.Len()
		}
		if copied, ok := seen[key]; ok {
			return copied
		}
		var copied reflect.Value
		switch value.Kind() {
		case reflect.Map:
			copied = reflect.MakeMapWithSize(value.Type(), value.Len())
		case reflect.Slice:
			copied = reflect.MakeSlice(value.Type(), value.Len(), value.Len())
		case reflect.Pointer:
			copied = reflect.New(value.Type().Elem()).Convert(value.Type())
		}
		seen[key] = copied
		switch value.Kind() {
		case reflect.Map:
			entries := value.MapRange()
			for entries.Next() {
				copied.SetMapIndex(copyMetadataValue(entries.Key(), seen), copyMetadataValue(entries.Value(), seen))
			}
		case reflect.Slice:
			for i := 0; i < value.Len(); i++ {
				copied.Index(i).Set(copyMetadataValue(value.Index(i), seen))
			}
		case reflect.Pointer:
			copied.Elem().Set(copyMetadataValue(value.Elem(), seen))
		}
		return copied
	case reflect.Array:
		copied := reflect.New(value.Type()).Elem()
		for i := 0; i < value.Len(); i++ {
			copied.Index(i).Set(copyMetadataValue(value.Index(i), seen))
		}
		return copied
	case reflect.Struct:
		copied := reflect.New(value.Type()).Elem()
		copied.Set(value)
		for i := 0; i < value.NumField(); i++ {
			if copied.Field(i).CanSet() {
				copied.Field(i).Set(copyMetadataValue(value.Field(i), seen))
			}
		}
		return copied
	default:
		return value
	}
}

func validateMetadataCopy(value reflect.Value, seen map[metadataReference]bool) error {
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		return fmt.Errorf("metadata type %s cannot be copied safely", value.Type())
	case reflect.Map, reflect.Slice, reflect.Pointer:
		if value.IsNil() {
			return nil
		}
		key := metadataReference{typ: value.Type(), pointer: uintptr(value.UnsafePointer())}
		if value.Kind() == reflect.Slice {
			key.length = value.Len()
		}
		if seen[key] {
			return nil
		}
		seen[key] = true
	}
	switch value.Kind() {
	case reflect.Interface, reflect.Pointer:
		if !value.IsNil() {
			return validateMetadataCopy(value.Elem(), seen)
		}
	case reflect.Map:
		entries := value.MapRange()
		for entries.Next() {
			if err := validateMetadataCopy(entries.Key(), seen); err != nil {
				return err
			}
			if err := validateMetadataCopy(entries.Value(), seen); err != nil {
				return err
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			if err := validateMetadataCopy(value.Index(i), seen); err != nil {
				return err
			}
		}
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			if field.PkgPath != "" {
				if metadataTypeHasReferences(field.Type) {
					return fmt.Errorf("metadata type %s has mutable unexported state", value.Type())
				}
				continue
			}
			if err := validateMetadataCopy(value.Field(i), seen); err != nil {
				return err
			}
		}
	}
	return nil
}
func metadataTypeHasReferences(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Map, reflect.Slice, reflect.Pointer, reflect.Interface, reflect.Chan, reflect.Func, reflect.UnsafePointer:
		return true
	case reflect.Array:
		return metadataTypeHasReferences(typ.Elem())
	case reflect.Struct:
		for i := 0; i < typ.NumField(); i++ {
			if metadataTypeHasReferences(typ.Field(i).Type) {
				return true
			}
		}
	}
	return false
}
