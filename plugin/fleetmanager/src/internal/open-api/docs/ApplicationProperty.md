# ApplicationProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The application property ID | [readonly] 
**PropertyType** | **int32** | The application property type. For a list of application property types see [&#x60;GET /v3/application/property/type&#x60;](#/Application/get_v3_application_property_type) | 
**PropertyKey** | **string** | The application property key. Only hyphens (-), underscores (_), lowercase characters and numbers are allowed. Keys must start with a lowercase character | 
**PropertyValue** | Pointer to **string** | The application property value | [optional] 

## Methods

### NewApplicationProperty

`func NewApplicationProperty(id string, propertyType int32, propertyKey string, ) *ApplicationProperty`

NewApplicationProperty instantiates a new ApplicationProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPropertyWithDefaults

`func NewApplicationPropertyWithDefaults() *ApplicationProperty`

NewApplicationPropertyWithDefaults instantiates a new ApplicationProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationProperty) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationProperty) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationProperty) SetId(v string)`

SetId sets Id field to given value.


### GetPropertyType

`func (o *ApplicationProperty) GetPropertyType() int32`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *ApplicationProperty) GetPropertyTypeOk() (*int32, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *ApplicationProperty) SetPropertyType(v int32)`

SetPropertyType sets PropertyType field to given value.


### GetPropertyKey

`func (o *ApplicationProperty) GetPropertyKey() string`

GetPropertyKey returns the PropertyKey field if non-nil, zero value otherwise.

### GetPropertyKeyOk

`func (o *ApplicationProperty) GetPropertyKeyOk() (*string, bool)`

GetPropertyKeyOk returns a tuple with the PropertyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyKey

`func (o *ApplicationProperty) SetPropertyKey(v string)`

SetPropertyKey sets PropertyKey field to given value.


### GetPropertyValue

`func (o *ApplicationProperty) GetPropertyValue() string`

GetPropertyValue returns the PropertyValue field if non-nil, zero value otherwise.

### GetPropertyValueOk

`func (o *ApplicationProperty) GetPropertyValueOk() (*string, bool)`

GetPropertyValueOk returns a tuple with the PropertyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValue

`func (o *ApplicationProperty) SetPropertyValue(v string)`

SetPropertyValue sets PropertyValue field to given value.

### HasPropertyValue

`func (o *ApplicationProperty) HasPropertyValue() bool`

HasPropertyValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


