# ApplicationInstanceStopMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodId** | Pointer to **int32** | Stop method type, list of types can be found in: [GET /v3/application/stopMethod](game-publisher#/Application/getApplicationStopMethods) | [optional] 
**Timeout** | Pointer to **int32** | Stop method timeout in seconds, default value is 1800 seconds. | [optional] 

## Methods

### NewApplicationInstanceStopMethod

`func NewApplicationInstanceStopMethod() *ApplicationInstanceStopMethod`

NewApplicationInstanceStopMethod instantiates a new ApplicationInstanceStopMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceStopMethodWithDefaults

`func NewApplicationInstanceStopMethodWithDefaults() *ApplicationInstanceStopMethod`

NewApplicationInstanceStopMethodWithDefaults instantiates a new ApplicationInstanceStopMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodId

`func (o *ApplicationInstanceStopMethod) GetMethodId() int32`

GetMethodId returns the MethodId field if non-nil, zero value otherwise.

### GetMethodIdOk

`func (o *ApplicationInstanceStopMethod) GetMethodIdOk() (*int32, bool)`

GetMethodIdOk returns a tuple with the MethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodId

`func (o *ApplicationInstanceStopMethod) SetMethodId(v int32)`

SetMethodId sets MethodId field to given value.

### HasMethodId

`func (o *ApplicationInstanceStopMethod) HasMethodId() bool`

HasMethodId returns a boolean if a field has been set.

### GetTimeout

`func (o *ApplicationInstanceStopMethod) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ApplicationInstanceStopMethod) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ApplicationInstanceStopMethod) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ApplicationInstanceStopMethod) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


