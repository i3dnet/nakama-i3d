# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Application ID | [readonly] 
**Type** | **int32** | The list of application type can be found in: GET /application/type | 
**ManagementProtocol** | **int32** | The protocol used to manage the application (defaults to None when not set): * 0: None * 1: A2S * 2: ARCUS | 
**Name** | **string** | The public name of your application | 
**WebsiteUrl** | Pointer to **string** | The url to your application&#39;s website | [optional] 
**Description** | Pointer to **string** | A short description of your application | [optional] 
**DeveloperName** | Pointer to **string** | The name of the application&#39;s developer | [optional] 
**PublisherName** | Pointer to **string** | The name of the application&#39;s publisher | [optional] 
**StopDefaultOsGroup** | Pointer to **int32** | The default operating system group being used for the stop method. If you have a Windows and Linux ApplicationBuild you can overwrite this on the ApplicationBuild level. Must be one of: * 1: Windows * 2: Linux (default) | [optional] 
**StopMethod** | Pointer to **int32** | The stop method that will be used on restart or destroy of an application instance. The default is hard kill, list of types can be found in: [GET /v3/application/stopMethod](game-publisher#/Application/getApplicationStopMethods) | [optional] 
**StopTimeout** | Pointer to **int32** | The timeout that will be used on restart or destroy of an application instance, in seconds. The default is 0 (no timeout) | [optional] 
**CreatedAt** | **int32** | UnixTimestamp | [readonly] 
**InUse** | **int32** | If the application currently in use or not | [readonly] 

## Methods

### NewApplication

`func NewApplication(id string, type_ int32, managementProtocol int32, name string, createdAt int32, inUse int32, ) *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Application) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Application) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Application) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Application) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Application) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Application) SetType(v int32)`

SetType sets Type field to given value.


### GetManagementProtocol

`func (o *Application) GetManagementProtocol() int32`

GetManagementProtocol returns the ManagementProtocol field if non-nil, zero value otherwise.

### GetManagementProtocolOk

`func (o *Application) GetManagementProtocolOk() (*int32, bool)`

GetManagementProtocolOk returns a tuple with the ManagementProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementProtocol

`func (o *Application) SetManagementProtocol(v int32)`

SetManagementProtocol sets ManagementProtocol field to given value.


### GetName

`func (o *Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Application) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Application) SetName(v string)`

SetName sets Name field to given value.


### GetWebsiteUrl

`func (o *Application) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *Application) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *Application) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *Application) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetDescription

`func (o *Application) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Application) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Application) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Application) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeveloperName

`func (o *Application) GetDeveloperName() string`

GetDeveloperName returns the DeveloperName field if non-nil, zero value otherwise.

### GetDeveloperNameOk

`func (o *Application) GetDeveloperNameOk() (*string, bool)`

GetDeveloperNameOk returns a tuple with the DeveloperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperName

`func (o *Application) SetDeveloperName(v string)`

SetDeveloperName sets DeveloperName field to given value.

### HasDeveloperName

`func (o *Application) HasDeveloperName() bool`

HasDeveloperName returns a boolean if a field has been set.

### GetPublisherName

`func (o *Application) GetPublisherName() string`

GetPublisherName returns the PublisherName field if non-nil, zero value otherwise.

### GetPublisherNameOk

`func (o *Application) GetPublisherNameOk() (*string, bool)`

GetPublisherNameOk returns a tuple with the PublisherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherName

`func (o *Application) SetPublisherName(v string)`

SetPublisherName sets PublisherName field to given value.

### HasPublisherName

`func (o *Application) HasPublisherName() bool`

HasPublisherName returns a boolean if a field has been set.

### GetStopDefaultOsGroup

`func (o *Application) GetStopDefaultOsGroup() int32`

GetStopDefaultOsGroup returns the StopDefaultOsGroup field if non-nil, zero value otherwise.

### GetStopDefaultOsGroupOk

`func (o *Application) GetStopDefaultOsGroupOk() (*int32, bool)`

GetStopDefaultOsGroupOk returns a tuple with the StopDefaultOsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopDefaultOsGroup

`func (o *Application) SetStopDefaultOsGroup(v int32)`

SetStopDefaultOsGroup sets StopDefaultOsGroup field to given value.

### HasStopDefaultOsGroup

`func (o *Application) HasStopDefaultOsGroup() bool`

HasStopDefaultOsGroup returns a boolean if a field has been set.

### GetStopMethod

`func (o *Application) GetStopMethod() int32`

GetStopMethod returns the StopMethod field if non-nil, zero value otherwise.

### GetStopMethodOk

`func (o *Application) GetStopMethodOk() (*int32, bool)`

GetStopMethodOk returns a tuple with the StopMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopMethod

`func (o *Application) SetStopMethod(v int32)`

SetStopMethod sets StopMethod field to given value.

### HasStopMethod

`func (o *Application) HasStopMethod() bool`

HasStopMethod returns a boolean if a field has been set.

### GetStopTimeout

`func (o *Application) GetStopTimeout() int32`

GetStopTimeout returns the StopTimeout field if non-nil, zero value otherwise.

### GetStopTimeoutOk

`func (o *Application) GetStopTimeoutOk() (*int32, bool)`

GetStopTimeoutOk returns a tuple with the StopTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTimeout

`func (o *Application) SetStopTimeout(v int32)`

SetStopTimeout sets StopTimeout field to given value.

### HasStopTimeout

`func (o *Application) HasStopTimeout() bool`

HasStopTimeout returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Application) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Application) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Application) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetInUse

`func (o *Application) GetInUse() int32`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *Application) GetInUseOk() (*int32, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *Application) SetInUse(v int32)`

SetInUse sets InUse field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


