# CloudProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Cloud provider ID | [readonly] 
**Name** | **string** | Cloud provider&#39;s name | [readonly] 
**ShortName** | **string** | Cloud provider&#39;s short name (abbreviation) | [readonly] 
**Active** | **int32** | Cloud provider&#39;s status | [readonly] 

## Methods

### NewCloudProvider

`func NewCloudProvider(id int32, name string, shortName string, active int32, ) *CloudProvider`

NewCloudProvider instantiates a new CloudProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderWithDefaults

`func NewCloudProviderWithDefaults() *CloudProvider`

NewCloudProviderWithDefaults instantiates a new CloudProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudProvider) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudProvider) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudProvider) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CloudProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudProvider) SetName(v string)`

SetName sets Name field to given value.


### GetShortName

`func (o *CloudProvider) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *CloudProvider) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *CloudProvider) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetActive

`func (o *CloudProvider) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CloudProvider) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CloudProvider) SetActive(v int32)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


