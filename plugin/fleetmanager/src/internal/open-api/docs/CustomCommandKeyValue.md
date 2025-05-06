# CustomCommandKeyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Custom command key that will be send. Max length 50 characters | 
**Value** | **string** | Custom Command value that will be send. Max length 512 characters | 

## Methods

### NewCustomCommandKeyValue

`func NewCustomCommandKeyValue(key string, value string, ) *CustomCommandKeyValue`

NewCustomCommandKeyValue instantiates a new CustomCommandKeyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomCommandKeyValueWithDefaults

`func NewCustomCommandKeyValueWithDefaults() *CustomCommandKeyValue`

NewCustomCommandKeyValueWithDefaults instantiates a new CustomCommandKeyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CustomCommandKeyValue) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomCommandKeyValue) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomCommandKeyValue) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *CustomCommandKeyValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomCommandKeyValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomCommandKeyValue) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


