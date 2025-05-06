# ApplicationBuildConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Application build configuration ID | [readonly] 
**ConfigPath** | **string** | Configuration file path | 
**ConfigName** | **string** | Configuration file name | 
**ConfigContent** | **string** | Configuration file contents | 
**CreatedAt** | **int32** | UnixTimestamp | [readonly] 

## Methods

### NewApplicationBuildConfiguration

`func NewApplicationBuildConfiguration(id string, configPath string, configName string, configContent string, createdAt int32, ) *ApplicationBuildConfiguration`

NewApplicationBuildConfiguration instantiates a new ApplicationBuildConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationBuildConfigurationWithDefaults

`func NewApplicationBuildConfigurationWithDefaults() *ApplicationBuildConfiguration`

NewApplicationBuildConfigurationWithDefaults instantiates a new ApplicationBuildConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationBuildConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationBuildConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationBuildConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetConfigPath

`func (o *ApplicationBuildConfiguration) GetConfigPath() string`

GetConfigPath returns the ConfigPath field if non-nil, zero value otherwise.

### GetConfigPathOk

`func (o *ApplicationBuildConfiguration) GetConfigPathOk() (*string, bool)`

GetConfigPathOk returns a tuple with the ConfigPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigPath

`func (o *ApplicationBuildConfiguration) SetConfigPath(v string)`

SetConfigPath sets ConfigPath field to given value.


### GetConfigName

`func (o *ApplicationBuildConfiguration) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *ApplicationBuildConfiguration) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *ApplicationBuildConfiguration) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.


### GetConfigContent

`func (o *ApplicationBuildConfiguration) GetConfigContent() string`

GetConfigContent returns the ConfigContent field if non-nil, zero value otherwise.

### GetConfigContentOk

`func (o *ApplicationBuildConfiguration) GetConfigContentOk() (*string, bool)`

GetConfigContentOk returns a tuple with the ConfigContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContent

`func (o *ApplicationBuildConfiguration) SetConfigContent(v string)`

SetConfigContent sets ConfigContent field to given value.


### GetCreatedAt

`func (o *ApplicationBuildConfiguration) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationBuildConfiguration) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationBuildConfiguration) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


