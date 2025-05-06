# ApplicationBuildProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The application build property ID | [readonly] 
**PropertyType** | **int32** | The application build property type. For a list of application property types see [&#x60;GET /v3/application/property/type&#x60;](#/Application/get_v3_application_property_type) | [readonly] 
**PropertyKey** | **string** | The application build property key. Only hyphens (-), underscores (_), lowercase characters and numbers are allowed. Keys must start with a lowercase character | [readonly] 
**PropertyValue** | **string** | The application build property value | 

## Methods

### NewApplicationBuildProperty

`func NewApplicationBuildProperty(id string, propertyType int32, propertyKey string, propertyValue string, ) *ApplicationBuildProperty`

NewApplicationBuildProperty instantiates a new ApplicationBuildProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationBuildPropertyWithDefaults

`func NewApplicationBuildPropertyWithDefaults() *ApplicationBuildProperty`

NewApplicationBuildPropertyWithDefaults instantiates a new ApplicationBuildProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationBuildProperty) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationBuildProperty) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationBuildProperty) SetId(v string)`

SetId sets Id field to given value.


### GetPropertyType

`func (o *ApplicationBuildProperty) GetPropertyType() int32`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *ApplicationBuildProperty) GetPropertyTypeOk() (*int32, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *ApplicationBuildProperty) SetPropertyType(v int32)`

SetPropertyType sets PropertyType field to given value.


### GetPropertyKey

`func (o *ApplicationBuildProperty) GetPropertyKey() string`

GetPropertyKey returns the PropertyKey field if non-nil, zero value otherwise.

### GetPropertyKeyOk

`func (o *ApplicationBuildProperty) GetPropertyKeyOk() (*string, bool)`

GetPropertyKeyOk returns a tuple with the PropertyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyKey

`func (o *ApplicationBuildProperty) SetPropertyKey(v string)`

SetPropertyKey sets PropertyKey field to given value.


### GetPropertyValue

`func (o *ApplicationBuildProperty) GetPropertyValue() string`

GetPropertyValue returns the PropertyValue field if non-nil, zero value otherwise.

### GetPropertyValueOk

`func (o *ApplicationBuildProperty) GetPropertyValueOk() (*string, bool)`

GetPropertyValueOk returns a tuple with the PropertyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValue

`func (o *ApplicationBuildProperty) SetPropertyValue(v string)`

SetPropertyValue sets PropertyValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


