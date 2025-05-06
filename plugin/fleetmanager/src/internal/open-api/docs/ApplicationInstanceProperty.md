# ApplicationInstanceProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The application instance property ID | [readonly] 
**PropertyType** | **int32** | The application instance property type. For a list of application property types see [&#x60;GET /v3/application/property/type&#x60;](#/Application/get_v3_application_property_type) | [readonly] 
**PropertyKey** | **string** | The application instance property key. Only hyphens (-), underscores (_), lowercase characters and numbers are allowed. Keys must start with a lowercase character | 
**PropertyValue** | **string** | The application instance property value | 

## Methods

### NewApplicationInstanceProperty

`func NewApplicationInstanceProperty(id string, propertyType int32, propertyKey string, propertyValue string, ) *ApplicationInstanceProperty`

NewApplicationInstanceProperty instantiates a new ApplicationInstanceProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstancePropertyWithDefaults

`func NewApplicationInstancePropertyWithDefaults() *ApplicationInstanceProperty`

NewApplicationInstancePropertyWithDefaults instantiates a new ApplicationInstanceProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationInstanceProperty) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationInstanceProperty) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationInstanceProperty) SetId(v string)`

SetId sets Id field to given value.


### GetPropertyType

`func (o *ApplicationInstanceProperty) GetPropertyType() int32`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *ApplicationInstanceProperty) GetPropertyTypeOk() (*int32, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *ApplicationInstanceProperty) SetPropertyType(v int32)`

SetPropertyType sets PropertyType field to given value.


### GetPropertyKey

`func (o *ApplicationInstanceProperty) GetPropertyKey() string`

GetPropertyKey returns the PropertyKey field if non-nil, zero value otherwise.

### GetPropertyKeyOk

`func (o *ApplicationInstanceProperty) GetPropertyKeyOk() (*string, bool)`

GetPropertyKeyOk returns a tuple with the PropertyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyKey

`func (o *ApplicationInstanceProperty) SetPropertyKey(v string)`

SetPropertyKey sets PropertyKey field to given value.


### GetPropertyValue

`func (o *ApplicationInstanceProperty) GetPropertyValue() string`

GetPropertyValue returns the PropertyValue field if non-nil, zero value otherwise.

### GetPropertyValueOk

`func (o *ApplicationInstanceProperty) GetPropertyValueOk() (*string, bool)`

GetPropertyValueOk returns a tuple with the PropertyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValue

`func (o *ApplicationInstanceProperty) SetPropertyValue(v string)`

SetPropertyValue sets PropertyValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


