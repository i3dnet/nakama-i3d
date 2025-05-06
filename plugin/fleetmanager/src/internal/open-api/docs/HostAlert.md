# HostAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Host alert ID | [readonly] 
**Percentage** | **int32** | At what percentage of bandwidth usage to trigger the alert | 
**SendMail** | **int32** | Will send an email notification when alert is triggered | 
**SendTicket** | **int32** | WIll create a ticket notification when alert is triggered | 
**Triggered** | **int32** | UNIX timestamp of the last time the alert was triggered | [readonly] 
**CreatedAt** | **int32** | UNIX timestamp of the time the alert was created | [readonly] 

## Methods

### NewHostAlert

`func NewHostAlert(id int32, percentage int32, sendMail int32, sendTicket int32, triggered int32, createdAt int32, ) *HostAlert`

NewHostAlert instantiates a new HostAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostAlertWithDefaults

`func NewHostAlertWithDefaults() *HostAlert`

NewHostAlertWithDefaults instantiates a new HostAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HostAlert) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostAlert) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostAlert) SetId(v int32)`

SetId sets Id field to given value.


### GetPercentage

`func (o *HostAlert) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *HostAlert) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *HostAlert) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.


### GetSendMail

`func (o *HostAlert) GetSendMail() int32`

GetSendMail returns the SendMail field if non-nil, zero value otherwise.

### GetSendMailOk

`func (o *HostAlert) GetSendMailOk() (*int32, bool)`

GetSendMailOk returns a tuple with the SendMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendMail

`func (o *HostAlert) SetSendMail(v int32)`

SetSendMail sets SendMail field to given value.


### GetSendTicket

`func (o *HostAlert) GetSendTicket() int32`

GetSendTicket returns the SendTicket field if non-nil, zero value otherwise.

### GetSendTicketOk

`func (o *HostAlert) GetSendTicketOk() (*int32, bool)`

GetSendTicketOk returns a tuple with the SendTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTicket

`func (o *HostAlert) SetSendTicket(v int32)`

SetSendTicket sets SendTicket field to given value.


### GetTriggered

`func (o *HostAlert) GetTriggered() int32`

GetTriggered returns the Triggered field if non-nil, zero value otherwise.

### GetTriggeredOk

`func (o *HostAlert) GetTriggeredOk() (*int32, bool)`

GetTriggeredOk returns a tuple with the Triggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggered

`func (o *HostAlert) SetTriggered(v int32)`

SetTriggered sets Triggered field to given value.


### GetCreatedAt

`func (o *HostAlert) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HostAlert) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HostAlert) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


