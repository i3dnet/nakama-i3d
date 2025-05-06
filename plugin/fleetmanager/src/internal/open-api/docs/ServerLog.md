# ServerLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of this log | 
**IssueOpen** | **int32** | Does the server log has an open issue | 
**DateStamp** | **int32** | Time of the log | 
**IssueCategory** | **string** | Server log issue category | 
**IssueDescription** | **string** | Server log issue description | 
**IssueResolved** | **string** | Whether or not server log issue is resolved | 
**CatName** | **string** | Server log category name | 

## Methods

### NewServerLog

`func NewServerLog(id int32, issueOpen int32, dateStamp int32, issueCategory string, issueDescription string, issueResolved string, catName string, ) *ServerLog`

NewServerLog instantiates a new ServerLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerLogWithDefaults

`func NewServerLogWithDefaults() *ServerLog`

NewServerLogWithDefaults instantiates a new ServerLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerLog) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerLog) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerLog) SetId(v int32)`

SetId sets Id field to given value.


### GetIssueOpen

`func (o *ServerLog) GetIssueOpen() int32`

GetIssueOpen returns the IssueOpen field if non-nil, zero value otherwise.

### GetIssueOpenOk

`func (o *ServerLog) GetIssueOpenOk() (*int32, bool)`

GetIssueOpenOk returns a tuple with the IssueOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueOpen

`func (o *ServerLog) SetIssueOpen(v int32)`

SetIssueOpen sets IssueOpen field to given value.


### GetDateStamp

`func (o *ServerLog) GetDateStamp() int32`

GetDateStamp returns the DateStamp field if non-nil, zero value otherwise.

### GetDateStampOk

`func (o *ServerLog) GetDateStampOk() (*int32, bool)`

GetDateStampOk returns a tuple with the DateStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStamp

`func (o *ServerLog) SetDateStamp(v int32)`

SetDateStamp sets DateStamp field to given value.


### GetIssueCategory

`func (o *ServerLog) GetIssueCategory() string`

GetIssueCategory returns the IssueCategory field if non-nil, zero value otherwise.

### GetIssueCategoryOk

`func (o *ServerLog) GetIssueCategoryOk() (*string, bool)`

GetIssueCategoryOk returns a tuple with the IssueCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCategory

`func (o *ServerLog) SetIssueCategory(v string)`

SetIssueCategory sets IssueCategory field to given value.


### GetIssueDescription

`func (o *ServerLog) GetIssueDescription() string`

GetIssueDescription returns the IssueDescription field if non-nil, zero value otherwise.

### GetIssueDescriptionOk

`func (o *ServerLog) GetIssueDescriptionOk() (*string, bool)`

GetIssueDescriptionOk returns a tuple with the IssueDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDescription

`func (o *ServerLog) SetIssueDescription(v string)`

SetIssueDescription sets IssueDescription field to given value.


### GetIssueResolved

`func (o *ServerLog) GetIssueResolved() string`

GetIssueResolved returns the IssueResolved field if non-nil, zero value otherwise.

### GetIssueResolvedOk

`func (o *ServerLog) GetIssueResolvedOk() (*string, bool)`

GetIssueResolvedOk returns a tuple with the IssueResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueResolved

`func (o *ServerLog) SetIssueResolved(v string)`

SetIssueResolved sets IssueResolved field to given value.


### GetCatName

`func (o *ServerLog) GetCatName() string`

GetCatName returns the CatName field if non-nil, zero value otherwise.

### GetCatNameOk

`func (o *ServerLog) GetCatNameOk() (*string, bool)`

GetCatNameOk returns a tuple with the CatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatName

`func (o *ServerLog) SetCatName(v string)`

SetCatName sets CatName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


