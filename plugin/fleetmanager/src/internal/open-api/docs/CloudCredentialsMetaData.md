# CloudCredentialsMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | User defined name for the cloud credential | 
**Status** | **int32** | Status of a cloud credential * 0: inactive * 1: active | 

## Methods

### NewCloudCredentialsMetaData

`func NewCloudCredentialsMetaData(name string, status int32, ) *CloudCredentialsMetaData`

NewCloudCredentialsMetaData instantiates a new CloudCredentialsMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsMetaDataWithDefaults

`func NewCloudCredentialsMetaDataWithDefaults() *CloudCredentialsMetaData`

NewCloudCredentialsMetaDataWithDefaults instantiates a new CloudCredentialsMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudCredentialsMetaData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCredentialsMetaData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCredentialsMetaData) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *CloudCredentialsMetaData) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudCredentialsMetaData) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudCredentialsMetaData) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


