# DeploymentContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This deployment container&#39;s ID | [readonly] 
**MarkedForDeletion** | **int32** | If set to 1, all application instances will be gracefully removed inside the container. Afterwards this container will be set to inactive | [readonly] 
**ContainerLocations** | [**[]DeploymentContainerLocation**](DeploymentContainerLocation.md) |  | 

## Methods

### NewDeploymentContainer

`func NewDeploymentContainer(id string, markedForDeletion int32, containerLocations []DeploymentContainerLocation, ) *DeploymentContainer`

NewDeploymentContainer instantiates a new DeploymentContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentContainerWithDefaults

`func NewDeploymentContainerWithDefaults() *DeploymentContainer`

NewDeploymentContainerWithDefaults instantiates a new DeploymentContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentContainer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentContainer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentContainer) SetId(v string)`

SetId sets Id field to given value.


### GetMarkedForDeletion

`func (o *DeploymentContainer) GetMarkedForDeletion() int32`

GetMarkedForDeletion returns the MarkedForDeletion field if non-nil, zero value otherwise.

### GetMarkedForDeletionOk

`func (o *DeploymentContainer) GetMarkedForDeletionOk() (*int32, bool)`

GetMarkedForDeletionOk returns a tuple with the MarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForDeletion

`func (o *DeploymentContainer) SetMarkedForDeletion(v int32)`

SetMarkedForDeletion sets MarkedForDeletion field to given value.


### GetContainerLocations

`func (o *DeploymentContainer) GetContainerLocations() []DeploymentContainerLocation`

GetContainerLocations returns the ContainerLocations field if non-nil, zero value otherwise.

### GetContainerLocationsOk

`func (o *DeploymentContainer) GetContainerLocationsOk() (*[]DeploymentContainerLocation, bool)`

GetContainerLocationsOk returns a tuple with the ContainerLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerLocations

`func (o *DeploymentContainer) SetContainerLocations(v []DeploymentContainerLocation)`

SetContainerLocations sets ContainerLocations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


