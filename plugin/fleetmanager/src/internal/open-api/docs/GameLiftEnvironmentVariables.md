# GameLiftEnvironmentVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | of game lift environment variable | 
**Type** | **string** | of game lift environment variable: - awsgl_endpoint - awsgl_region - awsgl_token - awsgl_hostname - awsgl_fleetid - raw_value | 
**Name** | **NullableString** | of game lift environment variable, e.g. \&quot;AWS_TOKEN_ID\&quot; | 

## Methods

### NewGameLiftEnvironmentVariables

`func NewGameLiftEnvironmentVariables(id NullableString, type_ string, name NullableString, ) *GameLiftEnvironmentVariables`

NewGameLiftEnvironmentVariables instantiates a new GameLiftEnvironmentVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameLiftEnvironmentVariablesWithDefaults

`func NewGameLiftEnvironmentVariablesWithDefaults() *GameLiftEnvironmentVariables`

NewGameLiftEnvironmentVariablesWithDefaults instantiates a new GameLiftEnvironmentVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GameLiftEnvironmentVariables) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameLiftEnvironmentVariables) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameLiftEnvironmentVariables) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *GameLiftEnvironmentVariables) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GameLiftEnvironmentVariables) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *GameLiftEnvironmentVariables) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameLiftEnvironmentVariables) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameLiftEnvironmentVariables) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *GameLiftEnvironmentVariables) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GameLiftEnvironmentVariables) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GameLiftEnvironmentVariables) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *GameLiftEnvironmentVariables) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GameLiftEnvironmentVariables) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


