# HostActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | The incoming data usage | 
**Timestamp** | **int32** | The corresponding date | 

## Methods

### NewHostActionResponse

`func NewHostActionResponse(result string, timestamp int32, ) *HostActionResponse`

NewHostActionResponse instantiates a new HostActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostActionResponseWithDefaults

`func NewHostActionResponseWithDefaults() *HostActionResponse`

NewHostActionResponseWithDefaults instantiates a new HostActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *HostActionResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *HostActionResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *HostActionResponse) SetResult(v string)`

SetResult sets Result field to given value.


### GetTimestamp

`func (o *HostActionResponse) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HostActionResponse) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HostActionResponse) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


