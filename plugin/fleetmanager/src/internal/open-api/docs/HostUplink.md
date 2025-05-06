# HostUplink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UplinkId** | **int32** | The cacti ID of the uplink | [readonly] 
**Aggregate** | **bool** | True if the uplink is an aggregated uplink, false if it is not aggregated | [readonly] 

## Methods

### NewHostUplink

`func NewHostUplink(uplinkId int32, aggregate bool, ) *HostUplink`

NewHostUplink instantiates a new HostUplink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostUplinkWithDefaults

`func NewHostUplinkWithDefaults() *HostUplink`

NewHostUplinkWithDefaults instantiates a new HostUplink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUplinkId

`func (o *HostUplink) GetUplinkId() int32`

GetUplinkId returns the UplinkId field if non-nil, zero value otherwise.

### GetUplinkIdOk

`func (o *HostUplink) GetUplinkIdOk() (*int32, bool)`

GetUplinkIdOk returns a tuple with the UplinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkId

`func (o *HostUplink) SetUplinkId(v int32)`

SetUplinkId sets UplinkId field to given value.


### GetAggregate

`func (o *HostUplink) GetAggregate() bool`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *HostUplink) GetAggregateOk() (*bool, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *HostUplink) SetAggregate(v bool)`

SetAggregate sets Aggregate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


