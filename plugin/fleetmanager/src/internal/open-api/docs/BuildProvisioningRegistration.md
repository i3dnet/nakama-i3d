# BuildProvisioningRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the build provisioning registration | [readonly] 
**Username** | **string** | Username the login name for the build provisioning registration (can be i3D.net docker or i3D.net storage) | [readonly] 
**Port** | **int32** | Port for the login to the build provisioning registration | [readonly] 
**Name** | **string** | Unique name of the build provisioning registration | [readonly] 
**OriginDomainName** | **string** | Origin domain name of the build provisioning registration | [readonly] 
**CdnDomainName** | **string** | CDN domain name of the build provisioning registration, where the application instance will be downloaded by the hostagent | [readonly] 
**Headers** | [**[]BuildProvisioningHeaders**](BuildProvisioningHeaders.md) | Array of headers used to download the file from the storage or get access to the docker repository | [readonly] 
**ContractStart** | **int32** | Start UNIX timestamp when the registration starts for the customer | [readonly] 
**ContractEnd** | **int32** | End UNIX timestamp when the registration ends for the customer | [readonly] 
**Type** | **int32** | The type of build provisioning registration storage that is being used [GET /v3/buildProvisioning/storage/registration/types](game-publisher#/BuildProvisioningRegistration/get_v3_buildProvisioning_storage_registration_types) | [readonly] 
**CreatedAt** | **int32** | UNIX timestamp when the build provisioning registration is created | [readonly] 
**ChangedAt** | **int32** | UNIX timestamp when the build provisioning registration is changed | [readonly] 

## Methods

### NewBuildProvisioningRegistration

`func NewBuildProvisioningRegistration(id string, username string, port int32, name string, originDomainName string, cdnDomainName string, headers []BuildProvisioningHeaders, contractStart int32, contractEnd int32, type_ int32, createdAt int32, changedAt int32, ) *BuildProvisioningRegistration`

NewBuildProvisioningRegistration instantiates a new BuildProvisioningRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildProvisioningRegistrationWithDefaults

`func NewBuildProvisioningRegistrationWithDefaults() *BuildProvisioningRegistration`

NewBuildProvisioningRegistrationWithDefaults instantiates a new BuildProvisioningRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BuildProvisioningRegistration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BuildProvisioningRegistration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BuildProvisioningRegistration) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *BuildProvisioningRegistration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BuildProvisioningRegistration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BuildProvisioningRegistration) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPort

`func (o *BuildProvisioningRegistration) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BuildProvisioningRegistration) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BuildProvisioningRegistration) SetPort(v int32)`

SetPort sets Port field to given value.


### GetName

`func (o *BuildProvisioningRegistration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuildProvisioningRegistration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuildProvisioningRegistration) SetName(v string)`

SetName sets Name field to given value.


### GetOriginDomainName

`func (o *BuildProvisioningRegistration) GetOriginDomainName() string`

GetOriginDomainName returns the OriginDomainName field if non-nil, zero value otherwise.

### GetOriginDomainNameOk

`func (o *BuildProvisioningRegistration) GetOriginDomainNameOk() (*string, bool)`

GetOriginDomainNameOk returns a tuple with the OriginDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDomainName

`func (o *BuildProvisioningRegistration) SetOriginDomainName(v string)`

SetOriginDomainName sets OriginDomainName field to given value.


### GetCdnDomainName

`func (o *BuildProvisioningRegistration) GetCdnDomainName() string`

GetCdnDomainName returns the CdnDomainName field if non-nil, zero value otherwise.

### GetCdnDomainNameOk

`func (o *BuildProvisioningRegistration) GetCdnDomainNameOk() (*string, bool)`

GetCdnDomainNameOk returns a tuple with the CdnDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnDomainName

`func (o *BuildProvisioningRegistration) SetCdnDomainName(v string)`

SetCdnDomainName sets CdnDomainName field to given value.


### GetHeaders

`func (o *BuildProvisioningRegistration) GetHeaders() []BuildProvisioningHeaders`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BuildProvisioningRegistration) GetHeadersOk() (*[]BuildProvisioningHeaders, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BuildProvisioningRegistration) SetHeaders(v []BuildProvisioningHeaders)`

SetHeaders sets Headers field to given value.


### GetContractStart

`func (o *BuildProvisioningRegistration) GetContractStart() int32`

GetContractStart returns the ContractStart field if non-nil, zero value otherwise.

### GetContractStartOk

`func (o *BuildProvisioningRegistration) GetContractStartOk() (*int32, bool)`

GetContractStartOk returns a tuple with the ContractStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStart

`func (o *BuildProvisioningRegistration) SetContractStart(v int32)`

SetContractStart sets ContractStart field to given value.


### GetContractEnd

`func (o *BuildProvisioningRegistration) GetContractEnd() int32`

GetContractEnd returns the ContractEnd field if non-nil, zero value otherwise.

### GetContractEndOk

`func (o *BuildProvisioningRegistration) GetContractEndOk() (*int32, bool)`

GetContractEndOk returns a tuple with the ContractEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractEnd

`func (o *BuildProvisioningRegistration) SetContractEnd(v int32)`

SetContractEnd sets ContractEnd field to given value.


### GetType

`func (o *BuildProvisioningRegistration) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BuildProvisioningRegistration) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BuildProvisioningRegistration) SetType(v int32)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *BuildProvisioningRegistration) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BuildProvisioningRegistration) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BuildProvisioningRegistration) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetChangedAt

`func (o *BuildProvisioningRegistration) GetChangedAt() int32`

GetChangedAt returns the ChangedAt field if non-nil, zero value otherwise.

### GetChangedAtOk

`func (o *BuildProvisioningRegistration) GetChangedAtOk() (*int32, bool)`

GetChangedAtOk returns a tuple with the ChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAt

`func (o *BuildProvisioningRegistration) SetChangedAt(v int32)`

SetChangedAt sets ChangedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


