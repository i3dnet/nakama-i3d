# ApplicationInstanceIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** | Application instance IP address | 
**IpVersion** | **int32** | Application instance IP version | [readonly] 
**Private** | **int32** | If application instance has private ip address. | [readonly] 

## Methods

### NewApplicationInstanceIP

`func NewApplicationInstanceIP(ipAddress string, ipVersion int32, private int32, ) *ApplicationInstanceIP`

NewApplicationInstanceIP instantiates a new ApplicationInstanceIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceIPWithDefaults

`func NewApplicationInstanceIPWithDefaults() *ApplicationInstanceIP`

NewApplicationInstanceIPWithDefaults instantiates a new ApplicationInstanceIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *ApplicationInstanceIP) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ApplicationInstanceIP) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ApplicationInstanceIP) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetIpVersion

`func (o *ApplicationInstanceIP) GetIpVersion() int32`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *ApplicationInstanceIP) GetIpVersionOk() (*int32, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *ApplicationInstanceIP) SetIpVersion(v int32)`

SetIpVersion sets IpVersion field to given value.


### GetPrivate

`func (o *ApplicationInstanceIP) GetPrivate() int32`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ApplicationInstanceIP) GetPrivateOk() (*int32, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ApplicationInstanceIP) SetPrivate(v int32)`

SetPrivate sets Private field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


