# ApplicationInstanceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusId** | **int32** | Status code of an application instance | [readonly] 
**StatusDescription** | **string** | Application instance status name | [readonly] 
**Description** | **string** | Application instance status description | [readonly] 
**IsDependency** | **int32** | If isDependency&#x3D;1, status code represent for application instance dependency installer and uninstaller, in case of isDependency&#x3D;0 status code represents application instance game and utility type | [readonly] 

## Methods

### NewApplicationInstanceStatus

`func NewApplicationInstanceStatus(statusId int32, statusDescription string, description string, isDependency int32, ) *ApplicationInstanceStatus`

NewApplicationInstanceStatus instantiates a new ApplicationInstanceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceStatusWithDefaults

`func NewApplicationInstanceStatusWithDefaults() *ApplicationInstanceStatus`

NewApplicationInstanceStatusWithDefaults instantiates a new ApplicationInstanceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusId

`func (o *ApplicationInstanceStatus) GetStatusId() int32`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *ApplicationInstanceStatus) GetStatusIdOk() (*int32, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *ApplicationInstanceStatus) SetStatusId(v int32)`

SetStatusId sets StatusId field to given value.


### GetStatusDescription

`func (o *ApplicationInstanceStatus) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *ApplicationInstanceStatus) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *ApplicationInstanceStatus) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.


### GetDescription

`func (o *ApplicationInstanceStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationInstanceStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationInstanceStatus) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsDependency

`func (o *ApplicationInstanceStatus) GetIsDependency() int32`

GetIsDependency returns the IsDependency field if non-nil, zero value otherwise.

### GetIsDependencyOk

`func (o *ApplicationInstanceStatus) GetIsDependencyOk() (*int32, bool)`

GetIsDependencyOk returns a tuple with the IsDependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDependency

`func (o *ApplicationInstanceStatus) SetIsDependency(v int32)`

SetIsDependency sets IsDependency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


