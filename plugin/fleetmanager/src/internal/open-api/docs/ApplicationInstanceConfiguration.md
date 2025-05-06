# ApplicationInstanceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the application instance configuration | [readonly] 
**ApplicationInstanceId** | **string** | The ID of the application instance | [readonly] 
**ApplicationBuildConfigurationId** | **string** | The ID of the application build configuration | [readonly] 
**ConfigPath** | **string** | Path of the configuration | [readonly] 
**ConfigName** | **string** | Name of the configuration | [readonly] 
**ConfigContent** | **string** | Content of the configuration | [readonly] 
**UpdatedAt** | **int32** | UnixTimestamp | [readonly] 
**Active** | **int32** | Status of the application instance configuration | [readonly] 

## Methods

### NewApplicationInstanceConfiguration

`func NewApplicationInstanceConfiguration(id string, applicationInstanceId string, applicationBuildConfigurationId string, configPath string, configName string, configContent string, updatedAt int32, active int32, ) *ApplicationInstanceConfiguration`

NewApplicationInstanceConfiguration instantiates a new ApplicationInstanceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceConfigurationWithDefaults

`func NewApplicationInstanceConfigurationWithDefaults() *ApplicationInstanceConfiguration`

NewApplicationInstanceConfigurationWithDefaults instantiates a new ApplicationInstanceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationInstanceConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationInstanceConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationInstanceConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetApplicationInstanceId

`func (o *ApplicationInstanceConfiguration) GetApplicationInstanceId() string`

GetApplicationInstanceId returns the ApplicationInstanceId field if non-nil, zero value otherwise.

### GetApplicationInstanceIdOk

`func (o *ApplicationInstanceConfiguration) GetApplicationInstanceIdOk() (*string, bool)`

GetApplicationInstanceIdOk returns a tuple with the ApplicationInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInstanceId

`func (o *ApplicationInstanceConfiguration) SetApplicationInstanceId(v string)`

SetApplicationInstanceId sets ApplicationInstanceId field to given value.


### GetApplicationBuildConfigurationId

`func (o *ApplicationInstanceConfiguration) GetApplicationBuildConfigurationId() string`

GetApplicationBuildConfigurationId returns the ApplicationBuildConfigurationId field if non-nil, zero value otherwise.

### GetApplicationBuildConfigurationIdOk

`func (o *ApplicationInstanceConfiguration) GetApplicationBuildConfigurationIdOk() (*string, bool)`

GetApplicationBuildConfigurationIdOk returns a tuple with the ApplicationBuildConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBuildConfigurationId

`func (o *ApplicationInstanceConfiguration) SetApplicationBuildConfigurationId(v string)`

SetApplicationBuildConfigurationId sets ApplicationBuildConfigurationId field to given value.


### GetConfigPath

`func (o *ApplicationInstanceConfiguration) GetConfigPath() string`

GetConfigPath returns the ConfigPath field if non-nil, zero value otherwise.

### GetConfigPathOk

`func (o *ApplicationInstanceConfiguration) GetConfigPathOk() (*string, bool)`

GetConfigPathOk returns a tuple with the ConfigPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigPath

`func (o *ApplicationInstanceConfiguration) SetConfigPath(v string)`

SetConfigPath sets ConfigPath field to given value.


### GetConfigName

`func (o *ApplicationInstanceConfiguration) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *ApplicationInstanceConfiguration) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *ApplicationInstanceConfiguration) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.


### GetConfigContent

`func (o *ApplicationInstanceConfiguration) GetConfigContent() string`

GetConfigContent returns the ConfigContent field if non-nil, zero value otherwise.

### GetConfigContentOk

`func (o *ApplicationInstanceConfiguration) GetConfigContentOk() (*string, bool)`

GetConfigContentOk returns a tuple with the ConfigContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContent

`func (o *ApplicationInstanceConfiguration) SetConfigContent(v string)`

SetConfigContent sets ConfigContent field to given value.


### GetUpdatedAt

`func (o *ApplicationInstanceConfiguration) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApplicationInstanceConfiguration) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApplicationInstanceConfiguration) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetActive

`func (o *ApplicationInstanceConfiguration) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationInstanceConfiguration) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationInstanceConfiguration) SetActive(v int32)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


