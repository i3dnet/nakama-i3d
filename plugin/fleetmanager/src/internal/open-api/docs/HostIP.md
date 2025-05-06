# HostIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** | Host IP address | [readonly] 
**Version** | **int32** | IP version, possible values can be 4 or 6 | [readonly] 
**Type** | **int32** | IP type: 1) normal, 2) KVM, 253) VRRP1, 254) VRRP2, 255) gateway | [readonly] 
**Private** | **int32** | Private IP (1) or public IP (0) | [readonly] 
**Interface** | **int32** | ID of the interface (0-7) | [readonly] 
**MacAddress** | **string** | MAC address of the interface | [readonly] 
**RDns** | **string** | RDNS of the interface | [readonly] 
**VlanId** | **NullableInt32** | ID for the VLAN | [readonly] 
**Gateway** | **NullableString** | Interface gateway | [readonly] 
**Netmask** | **NullableString** | Interface netmask | [readonly] 
**Prefix** | **NullableInt32** | Interface prefix | [readonly] 

## Methods

### NewHostIP

`func NewHostIP(ipAddress string, version int32, type_ int32, private int32, interface_ int32, macAddress string, rDns string, vlanId NullableInt32, gateway NullableString, netmask NullableString, prefix NullableInt32, ) *HostIP`

NewHostIP instantiates a new HostIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostIPWithDefaults

`func NewHostIPWithDefaults() *HostIP`

NewHostIPWithDefaults instantiates a new HostIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *HostIP) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *HostIP) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *HostIP) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetVersion

`func (o *HostIP) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HostIP) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HostIP) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetType

`func (o *HostIP) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HostIP) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HostIP) SetType(v int32)`

SetType sets Type field to given value.


### GetPrivate

`func (o *HostIP) GetPrivate() int32`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *HostIP) GetPrivateOk() (*int32, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *HostIP) SetPrivate(v int32)`

SetPrivate sets Private field to given value.


### GetInterface

`func (o *HostIP) GetInterface() int32`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *HostIP) GetInterfaceOk() (*int32, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *HostIP) SetInterface(v int32)`

SetInterface sets Interface field to given value.


### GetMacAddress

`func (o *HostIP) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *HostIP) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *HostIP) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetRDns

`func (o *HostIP) GetRDns() string`

GetRDns returns the RDns field if non-nil, zero value otherwise.

### GetRDnsOk

`func (o *HostIP) GetRDnsOk() (*string, bool)`

GetRDnsOk returns a tuple with the RDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRDns

`func (o *HostIP) SetRDns(v string)`

SetRDns sets RDns field to given value.


### GetVlanId

`func (o *HostIP) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *HostIP) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *HostIP) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.


### SetVlanIdNil

`func (o *HostIP) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *HostIP) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
### GetGateway

`func (o *HostIP) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *HostIP) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *HostIP) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### SetGatewayNil

`func (o *HostIP) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *HostIP) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetNetmask

`func (o *HostIP) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *HostIP) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *HostIP) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.


### SetNetmaskNil

`func (o *HostIP) SetNetmaskNil(b bool)`

 SetNetmaskNil sets the value for Netmask to be an explicit nil

### UnsetNetmask
`func (o *HostIP) UnsetNetmask()`

UnsetNetmask ensures that no value is present for Netmask, not even an explicit nil
### GetPrefix

`func (o *HostIP) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *HostIP) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *HostIP) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.


### SetPrefixNil

`func (o *HostIP) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *HostIP) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


