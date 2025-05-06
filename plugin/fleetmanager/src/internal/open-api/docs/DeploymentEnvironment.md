# DeploymentEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Deployment environment ID | [readonly] 
**Name** | **string** | The public name of your deployment environment | 
**CreatedAt** | **int32** | UnixTimestamp | [readonly] 

## Methods

### NewDeploymentEnvironment

`func NewDeploymentEnvironment(id string, name string, createdAt int32, ) *DeploymentEnvironment`

NewDeploymentEnvironment instantiates a new DeploymentEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentEnvironmentWithDefaults

`func NewDeploymentEnvironmentWithDefaults() *DeploymentEnvironment`

NewDeploymentEnvironmentWithDefaults instantiates a new DeploymentEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentEnvironment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentEnvironment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentEnvironment) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DeploymentEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentEnvironment) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *DeploymentEnvironment) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentEnvironment) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentEnvironment) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


