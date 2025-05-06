# DeploymentRegionProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Deployment region property ID | [readonly] 
**ApplicationPropertyId** | **string** | The application property ID that this regional property overrides | 
**PropertyValue** | **string** | The overridden value of the application property | 

## Methods

### NewDeploymentRegionProperty

`func NewDeploymentRegionProperty(id string, applicationPropertyId string, propertyValue string, ) *DeploymentRegionProperty`

NewDeploymentRegionProperty instantiates a new DeploymentRegionProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentRegionPropertyWithDefaults

`func NewDeploymentRegionPropertyWithDefaults() *DeploymentRegionProperty`

NewDeploymentRegionPropertyWithDefaults instantiates a new DeploymentRegionProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentRegionProperty) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentRegionProperty) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentRegionProperty) SetId(v string)`

SetId sets Id field to given value.


### GetApplicationPropertyId

`func (o *DeploymentRegionProperty) GetApplicationPropertyId() string`

GetApplicationPropertyId returns the ApplicationPropertyId field if non-nil, zero value otherwise.

### GetApplicationPropertyIdOk

`func (o *DeploymentRegionProperty) GetApplicationPropertyIdOk() (*string, bool)`

GetApplicationPropertyIdOk returns a tuple with the ApplicationPropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPropertyId

`func (o *DeploymentRegionProperty) SetApplicationPropertyId(v string)`

SetApplicationPropertyId sets ApplicationPropertyId field to given value.


### GetPropertyValue

`func (o *DeploymentRegionProperty) GetPropertyValue() string`

GetPropertyValue returns the PropertyValue field if non-nil, zero value otherwise.

### GetPropertyValueOk

`func (o *DeploymentRegionProperty) GetPropertyValueOk() (*string, bool)`

GetPropertyValueOk returns a tuple with the PropertyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValue

`func (o *DeploymentRegionProperty) SetPropertyValue(v string)`

SetPropertyValue sets PropertyValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


