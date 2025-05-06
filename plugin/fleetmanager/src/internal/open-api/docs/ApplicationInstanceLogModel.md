# ApplicationInstanceLogModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationInstanceId** | **string** | ID of Application instance | [readonly] 
**FleetId** | **string** | The application instance&#39;s fleet ID | [readonly] 
**ExitCode** | **int32** | The code given back by the server when the application instance was killed | [readonly] 
**StdOutLog** | **string** | Log output from the application instances from the server | [readonly] 
**CreatedAt** | **int32** | Timestamp when the log was created | [readonly] 

## Methods

### NewApplicationInstanceLogModel

`func NewApplicationInstanceLogModel(applicationInstanceId string, fleetId string, exitCode int32, stdOutLog string, createdAt int32, ) *ApplicationInstanceLogModel`

NewApplicationInstanceLogModel instantiates a new ApplicationInstanceLogModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceLogModelWithDefaults

`func NewApplicationInstanceLogModelWithDefaults() *ApplicationInstanceLogModel`

NewApplicationInstanceLogModelWithDefaults instantiates a new ApplicationInstanceLogModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationInstanceId

`func (o *ApplicationInstanceLogModel) GetApplicationInstanceId() string`

GetApplicationInstanceId returns the ApplicationInstanceId field if non-nil, zero value otherwise.

### GetApplicationInstanceIdOk

`func (o *ApplicationInstanceLogModel) GetApplicationInstanceIdOk() (*string, bool)`

GetApplicationInstanceIdOk returns a tuple with the ApplicationInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInstanceId

`func (o *ApplicationInstanceLogModel) SetApplicationInstanceId(v string)`

SetApplicationInstanceId sets ApplicationInstanceId field to given value.


### GetFleetId

`func (o *ApplicationInstanceLogModel) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *ApplicationInstanceLogModel) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *ApplicationInstanceLogModel) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetExitCode

`func (o *ApplicationInstanceLogModel) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ApplicationInstanceLogModel) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ApplicationInstanceLogModel) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.


### GetStdOutLog

`func (o *ApplicationInstanceLogModel) GetStdOutLog() string`

GetStdOutLog returns the StdOutLog field if non-nil, zero value otherwise.

### GetStdOutLogOk

`func (o *ApplicationInstanceLogModel) GetStdOutLogOk() (*string, bool)`

GetStdOutLogOk returns a tuple with the StdOutLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdOutLog

`func (o *ApplicationInstanceLogModel) SetStdOutLog(v string)`

SetStdOutLog sets StdOutLog field to given value.


### GetCreatedAt

`func (o *ApplicationInstanceLogModel) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationInstanceLogModel) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationInstanceLogModel) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


