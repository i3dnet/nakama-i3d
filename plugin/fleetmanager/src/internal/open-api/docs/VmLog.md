# VmLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudInstanceId** | **string** |  | [readonly] 
**Title** | **string** |  | [readonly] 
**Log** | **string** |  | [readonly] 
**Timestamp** | **int32** | Unix timestamp | [readonly] 

## Methods

### NewVmLog

`func NewVmLog(cloudInstanceId string, title string, log string, timestamp int32, ) *VmLog`

NewVmLog instantiates a new VmLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmLogWithDefaults

`func NewVmLogWithDefaults() *VmLog`

NewVmLogWithDefaults instantiates a new VmLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudInstanceId

`func (o *VmLog) GetCloudInstanceId() string`

GetCloudInstanceId returns the CloudInstanceId field if non-nil, zero value otherwise.

### GetCloudInstanceIdOk

`func (o *VmLog) GetCloudInstanceIdOk() (*string, bool)`

GetCloudInstanceIdOk returns a tuple with the CloudInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInstanceId

`func (o *VmLog) SetCloudInstanceId(v string)`

SetCloudInstanceId sets CloudInstanceId field to given value.


### GetTitle

`func (o *VmLog) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VmLog) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VmLog) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetLog

`func (o *VmLog) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *VmLog) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *VmLog) SetLog(v string)`

SetLog sets Log field to given value.


### GetTimestamp

`func (o *VmLog) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VmLog) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VmLog) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


