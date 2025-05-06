# Host

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of this host | [readonly] 
**UserId** | **int32** | The ID of the user who owns this server | [readonly] 
**ServerId** | **int32** | The ID of the physical machine | [readonly] 
**ServerName** | **string** | The name of the physical machine | [readonly] 
**ServerType** | **int32** | The type of the server: * 1: Bare metal server * 2: Flex metal server * 3: Virtual machine | [readonly] 
**ProjectName** | Pointer to **string** | The name of the project for the host | [optional] 
**ClientServerName** | Pointer to **string** | The name of the server defined by client | [optional] 
**ClientState** | Pointer to **string** | Client statement for the host | [optional] 
**LiveHostName** | **string** | The host name | [readonly] 
**IsVirtual** | **int32** | 0 if this is a bare metal server, 1 if it&#39;s a VM. Use &#x60;serverType&#x60; instead | [readonly] 
**Category** | **string** | Host category. Normally \&quot;Dedicated Game Servers\&quot; or \&quot;Dedicated Servers\&quot;, but can be \&quot;Broken\&quot; if the server is in a degraded state | [readonly] 
**OsId** | **int32** | Operating system ID, must be one of [&#x60;GET /v3/operatingsystem&#x60;](all#/OperatingSystem/getOperatingsystems) | [readonly] 
**LocationId** | **int32** | Legacy location ID (not used for ODP) | [readonly] 
**DcLocationId** | **int32** | Datacenter location ID. Points to one of [&#x60;GET /v3/cloud/dcLocation&#x60;](#/Cloud/getCloudDcLocations) | [readonly] 
**InstanceType** | **string** | The instance type of this server, which is one of [&#x60;GET /v3/host/instanceType&#x60;](#/Host/getHostInstanceTypes) for a bare metal, or one of [&#x60;GET /v3/cloud/instanceType&#x60;](#/Cloud/getCloudInstanceTypes) for a virtual machine | [readonly] 
**FleetId** | **string** | The fleet ID, if this host has been assigned to a fleet If &#x60;0&#x60;, the host has not been assigned to a fleet, otherwise, the fleet is assigned to this host and the host can serve application instances for that fleet. | [readonly] 
**NewFleetId** | **NullableString** | The ID of a new fleet that is to be assigned to this host after performing a bulkReserve operation at [&#x60;POST /v3/fleet/host/bulkReserve&#x60;](#/Fleet/createFleetHostBulkReserve) | [readonly] 
**FleetAssociatedSince** | **int32** | Unix timestamp when fleet assigned to the host | [readonly] 
**RackName** | **string** | The name of the rack the host is located | [readonly] 
**DateStart** | **string** | The date at which this host became active | [readonly] 
**DateEnd** | **string** | The date at which this host will expire | [readonly] 
**DateCancelled** | **string** | The date at which this host was cancelled | [readonly] 
**DateEndContract** | **string** | The date at which the contract ends (if applicable) | [readonly] 
**ContractPeriod** | **int32** | The contract period in months | [readonly] 
**ExtendPeriod** | **int32** | The service extend period in months | [readonly] 
**CancellationPeriod** | **int32** | The cancellation period in months | [readonly] 
**PurchaseOrder** | **string** | Purchase order, if one has been supplied via our billing department | [readonly] 
**PaymentTerm** | **int32** | The payment term in days (how many days are invoices generated before dateEnd) | [readonly] 
**PricePerMonth** | **string** | The price of this host per month in cents (see &#x60;currencyId&#x60; for currency) | [readonly] 
**PricePerTbOveruse** | **string** | The price of traffic overuse in cents (per TB) (see &#x60;currencyId&#x60; for currency) | [readonly] 
**CurrencyId** | **int32** | The currency of the &#x60;pricePerMonth&#x60; field: * 0: EURO * 1: USD * 9: YEN | [readonly] 
**BandwidthBillingType** | **int32** | The bandwidth billing method for this service: * 1: unmetered connection * 2: measured in TB per month * 3: measured in mbit 95th percentile | [readonly] 
**BandwidthContractual** | **int32** | The contractual maximum bandwidth usage value. In GB if bandwidthBillingType equals &#x60;2&#x60;, in mbit otherwise | [readonly] 
**IpAddress** | [**[]HostIP**](HostIP.md) | All IP addresses assigned to this host | [readonly] 
**Brand** | **string** | The server manufacturer brand name | [readonly] 
**Model** | **string** | The model of the server | [readonly] 
**NumCpu** | **int32** | Number of CPUs in this host | [readonly] 
**CpuInfo** | **string** |  | [readonly] 
**CpuType** | **string** |  | [readonly] 
**CpuLoad** | **float32** | Percentage of cpu used across all cores | [readonly] 
**Cpu** | [**HostCpu**](HostCpu.md) |  | [readonly] 
**MemUsed** | **int32** | The amount of memory that is used by the host (in megabytes) | [readonly] 
**MemMax** | **int32** | The amount of memory that is available on the host (in megabytes) | [readonly] 
**MemFree** | **int32** | The amount of free memory that is available on the host (in megabytes) | [readonly] 
**Disk** | [**[]HostDisk**](HostDisk.md) |  | [readonly] 
**Memory** | [**[]HostMemory**](HostMemory.md) |  | [readonly] 
**IsReserve** | **int32** | If the host is reserved for a fleet isReserve set to 1 | [readonly] 
**Labels** | Pointer to [**[]Label**](Label.md) | Custom key/value pairs that can be used for host | [optional] 
**ServiceTag** | **string** | The service tag of the host | [readonly] 
**IsODP** | **int32** | If the host is available for game hosting platform isODP set to 1 | [readonly] 
**FmOrderId** | **NullableString** | ID of Flex Metal order if this is a flex metal server | [readonly] 
**InstallStatus** | **NullableString** | Status of server auto install if it was started by one. (created / installing / finished / failed) | [readonly] 
**Status** | **string** | Whether the host is running or not | [readonly] 
**Uptime** | Pointer to **int32** | The uptime of the host. Read-only. Only provided if the server checked-in recently | [optional] 
**TrafficSum** | **int32** | Monthly traffic of all incoming and outgoing bandwidth in GB or Mbit depending on host configuration | [readonly] 
**Uplinks** | [**[]HostUplink**](HostUplink.md) | A list of uplinks attached to this host | [readonly] 
**FreeIncomingTraffic** | **bool** | Whether the host has free incoming traffic or not | [readonly] 
**OutgoingTrafficSum** | **int32** | The outgoing traffic (Mbit or GB) | [readonly] 
**IncomingTrafficSum** | **int32** | The incoming traffic (Mbit or GB) | [readonly] 

## Methods

### NewHost

`func NewHost(id int32, userId int32, serverId int32, serverName string, serverType int32, liveHostName string, isVirtual int32, category string, osId int32, locationId int32, dcLocationId int32, instanceType string, fleetId string, newFleetId NullableString, fleetAssociatedSince int32, rackName string, dateStart string, dateEnd string, dateCancelled string, dateEndContract string, contractPeriod int32, extendPeriod int32, cancellationPeriod int32, purchaseOrder string, paymentTerm int32, pricePerMonth string, pricePerTbOveruse string, currencyId int32, bandwidthBillingType int32, bandwidthContractual int32, ipAddress []HostIP, brand string, model string, numCpu int32, cpuInfo string, cpuType string, cpuLoad float32, cpu HostCpu, memUsed int32, memMax int32, memFree int32, disk []HostDisk, memory []HostMemory, isReserve int32, serviceTag string, isODP int32, fmOrderId NullableString, installStatus NullableString, status string, trafficSum int32, uplinks []HostUplink, freeIncomingTraffic bool, outgoingTrafficSum int32, incomingTrafficSum int32, ) *Host`

NewHost instantiates a new Host object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostWithDefaults

`func NewHostWithDefaults() *Host`

NewHostWithDefaults instantiates a new Host object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Host) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Host) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Host) SetId(v int32)`

SetId sets Id field to given value.


### GetUserId

`func (o *Host) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Host) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Host) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetServerId

`func (o *Host) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *Host) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *Host) SetServerId(v int32)`

SetServerId sets ServerId field to given value.


### GetServerName

`func (o *Host) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *Host) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *Host) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetServerType

`func (o *Host) GetServerType() int32`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *Host) GetServerTypeOk() (*int32, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *Host) SetServerType(v int32)`

SetServerType sets ServerType field to given value.


### GetProjectName

`func (o *Host) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *Host) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *Host) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *Host) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetClientServerName

`func (o *Host) GetClientServerName() string`

GetClientServerName returns the ClientServerName field if non-nil, zero value otherwise.

### GetClientServerNameOk

`func (o *Host) GetClientServerNameOk() (*string, bool)`

GetClientServerNameOk returns a tuple with the ClientServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientServerName

`func (o *Host) SetClientServerName(v string)`

SetClientServerName sets ClientServerName field to given value.

### HasClientServerName

`func (o *Host) HasClientServerName() bool`

HasClientServerName returns a boolean if a field has been set.

### GetClientState

`func (o *Host) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *Host) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *Host) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *Host) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetLiveHostName

`func (o *Host) GetLiveHostName() string`

GetLiveHostName returns the LiveHostName field if non-nil, zero value otherwise.

### GetLiveHostNameOk

`func (o *Host) GetLiveHostNameOk() (*string, bool)`

GetLiveHostNameOk returns a tuple with the LiveHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveHostName

`func (o *Host) SetLiveHostName(v string)`

SetLiveHostName sets LiveHostName field to given value.


### GetIsVirtual

`func (o *Host) GetIsVirtual() int32`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *Host) GetIsVirtualOk() (*int32, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *Host) SetIsVirtual(v int32)`

SetIsVirtual sets IsVirtual field to given value.


### GetCategory

`func (o *Host) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Host) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Host) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetOsId

`func (o *Host) GetOsId() int32`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *Host) GetOsIdOk() (*int32, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *Host) SetOsId(v int32)`

SetOsId sets OsId field to given value.


### GetLocationId

`func (o *Host) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Host) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Host) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.


### GetDcLocationId

`func (o *Host) GetDcLocationId() int32`

GetDcLocationId returns the DcLocationId field if non-nil, zero value otherwise.

### GetDcLocationIdOk

`func (o *Host) GetDcLocationIdOk() (*int32, bool)`

GetDcLocationIdOk returns a tuple with the DcLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocationId

`func (o *Host) SetDcLocationId(v int32)`

SetDcLocationId sets DcLocationId field to given value.


### GetInstanceType

`func (o *Host) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *Host) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *Host) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetFleetId

`func (o *Host) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *Host) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *Host) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetNewFleetId

`func (o *Host) GetNewFleetId() string`

GetNewFleetId returns the NewFleetId field if non-nil, zero value otherwise.

### GetNewFleetIdOk

`func (o *Host) GetNewFleetIdOk() (*string, bool)`

GetNewFleetIdOk returns a tuple with the NewFleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFleetId

`func (o *Host) SetNewFleetId(v string)`

SetNewFleetId sets NewFleetId field to given value.


### SetNewFleetIdNil

`func (o *Host) SetNewFleetIdNil(b bool)`

 SetNewFleetIdNil sets the value for NewFleetId to be an explicit nil

### UnsetNewFleetId
`func (o *Host) UnsetNewFleetId()`

UnsetNewFleetId ensures that no value is present for NewFleetId, not even an explicit nil
### GetFleetAssociatedSince

`func (o *Host) GetFleetAssociatedSince() int32`

GetFleetAssociatedSince returns the FleetAssociatedSince field if non-nil, zero value otherwise.

### GetFleetAssociatedSinceOk

`func (o *Host) GetFleetAssociatedSinceOk() (*int32, bool)`

GetFleetAssociatedSinceOk returns a tuple with the FleetAssociatedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetAssociatedSince

`func (o *Host) SetFleetAssociatedSince(v int32)`

SetFleetAssociatedSince sets FleetAssociatedSince field to given value.


### GetRackName

`func (o *Host) GetRackName() string`

GetRackName returns the RackName field if non-nil, zero value otherwise.

### GetRackNameOk

`func (o *Host) GetRackNameOk() (*string, bool)`

GetRackNameOk returns a tuple with the RackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackName

`func (o *Host) SetRackName(v string)`

SetRackName sets RackName field to given value.


### GetDateStart

`func (o *Host) GetDateStart() string`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *Host) GetDateStartOk() (*string, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *Host) SetDateStart(v string)`

SetDateStart sets DateStart field to given value.


### GetDateEnd

`func (o *Host) GetDateEnd() string`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *Host) GetDateEndOk() (*string, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *Host) SetDateEnd(v string)`

SetDateEnd sets DateEnd field to given value.


### GetDateCancelled

`func (o *Host) GetDateCancelled() string`

GetDateCancelled returns the DateCancelled field if non-nil, zero value otherwise.

### GetDateCancelledOk

`func (o *Host) GetDateCancelledOk() (*string, bool)`

GetDateCancelledOk returns a tuple with the DateCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCancelled

`func (o *Host) SetDateCancelled(v string)`

SetDateCancelled sets DateCancelled field to given value.


### GetDateEndContract

`func (o *Host) GetDateEndContract() string`

GetDateEndContract returns the DateEndContract field if non-nil, zero value otherwise.

### GetDateEndContractOk

`func (o *Host) GetDateEndContractOk() (*string, bool)`

GetDateEndContractOk returns a tuple with the DateEndContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEndContract

`func (o *Host) SetDateEndContract(v string)`

SetDateEndContract sets DateEndContract field to given value.


### GetContractPeriod

`func (o *Host) GetContractPeriod() int32`

GetContractPeriod returns the ContractPeriod field if non-nil, zero value otherwise.

### GetContractPeriodOk

`func (o *Host) GetContractPeriodOk() (*int32, bool)`

GetContractPeriodOk returns a tuple with the ContractPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractPeriod

`func (o *Host) SetContractPeriod(v int32)`

SetContractPeriod sets ContractPeriod field to given value.


### GetExtendPeriod

`func (o *Host) GetExtendPeriod() int32`

GetExtendPeriod returns the ExtendPeriod field if non-nil, zero value otherwise.

### GetExtendPeriodOk

`func (o *Host) GetExtendPeriodOk() (*int32, bool)`

GetExtendPeriodOk returns a tuple with the ExtendPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendPeriod

`func (o *Host) SetExtendPeriod(v int32)`

SetExtendPeriod sets ExtendPeriod field to given value.


### GetCancellationPeriod

`func (o *Host) GetCancellationPeriod() int32`

GetCancellationPeriod returns the CancellationPeriod field if non-nil, zero value otherwise.

### GetCancellationPeriodOk

`func (o *Host) GetCancellationPeriodOk() (*int32, bool)`

GetCancellationPeriodOk returns a tuple with the CancellationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationPeriod

`func (o *Host) SetCancellationPeriod(v int32)`

SetCancellationPeriod sets CancellationPeriod field to given value.


### GetPurchaseOrder

`func (o *Host) GetPurchaseOrder() string`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *Host) GetPurchaseOrderOk() (*string, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *Host) SetPurchaseOrder(v string)`

SetPurchaseOrder sets PurchaseOrder field to given value.


### GetPaymentTerm

`func (o *Host) GetPaymentTerm() int32`

GetPaymentTerm returns the PaymentTerm field if non-nil, zero value otherwise.

### GetPaymentTermOk

`func (o *Host) GetPaymentTermOk() (*int32, bool)`

GetPaymentTermOk returns a tuple with the PaymentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerm

`func (o *Host) SetPaymentTerm(v int32)`

SetPaymentTerm sets PaymentTerm field to given value.


### GetPricePerMonth

`func (o *Host) GetPricePerMonth() string`

GetPricePerMonth returns the PricePerMonth field if non-nil, zero value otherwise.

### GetPricePerMonthOk

`func (o *Host) GetPricePerMonthOk() (*string, bool)`

GetPricePerMonthOk returns a tuple with the PricePerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerMonth

`func (o *Host) SetPricePerMonth(v string)`

SetPricePerMonth sets PricePerMonth field to given value.


### GetPricePerTbOveruse

`func (o *Host) GetPricePerTbOveruse() string`

GetPricePerTbOveruse returns the PricePerTbOveruse field if non-nil, zero value otherwise.

### GetPricePerTbOveruseOk

`func (o *Host) GetPricePerTbOveruseOk() (*string, bool)`

GetPricePerTbOveruseOk returns a tuple with the PricePerTbOveruse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerTbOveruse

`func (o *Host) SetPricePerTbOveruse(v string)`

SetPricePerTbOveruse sets PricePerTbOveruse field to given value.


### GetCurrencyId

`func (o *Host) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *Host) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *Host) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.


### GetBandwidthBillingType

`func (o *Host) GetBandwidthBillingType() int32`

GetBandwidthBillingType returns the BandwidthBillingType field if non-nil, zero value otherwise.

### GetBandwidthBillingTypeOk

`func (o *Host) GetBandwidthBillingTypeOk() (*int32, bool)`

GetBandwidthBillingTypeOk returns a tuple with the BandwidthBillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthBillingType

`func (o *Host) SetBandwidthBillingType(v int32)`

SetBandwidthBillingType sets BandwidthBillingType field to given value.


### GetBandwidthContractual

`func (o *Host) GetBandwidthContractual() int32`

GetBandwidthContractual returns the BandwidthContractual field if non-nil, zero value otherwise.

### GetBandwidthContractualOk

`func (o *Host) GetBandwidthContractualOk() (*int32, bool)`

GetBandwidthContractualOk returns a tuple with the BandwidthContractual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthContractual

`func (o *Host) SetBandwidthContractual(v int32)`

SetBandwidthContractual sets BandwidthContractual field to given value.


### GetIpAddress

`func (o *Host) GetIpAddress() []HostIP`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Host) GetIpAddressOk() (*[]HostIP, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Host) SetIpAddress(v []HostIP)`

SetIpAddress sets IpAddress field to given value.


### GetBrand

`func (o *Host) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *Host) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *Host) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetModel

`func (o *Host) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Host) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Host) SetModel(v string)`

SetModel sets Model field to given value.


### GetNumCpu

`func (o *Host) GetNumCpu() int32`

GetNumCpu returns the NumCpu field if non-nil, zero value otherwise.

### GetNumCpuOk

`func (o *Host) GetNumCpuOk() (*int32, bool)`

GetNumCpuOk returns a tuple with the NumCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpu

`func (o *Host) SetNumCpu(v int32)`

SetNumCpu sets NumCpu field to given value.


### GetCpuInfo

`func (o *Host) GetCpuInfo() string`

GetCpuInfo returns the CpuInfo field if non-nil, zero value otherwise.

### GetCpuInfoOk

`func (o *Host) GetCpuInfoOk() (*string, bool)`

GetCpuInfoOk returns a tuple with the CpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuInfo

`func (o *Host) SetCpuInfo(v string)`

SetCpuInfo sets CpuInfo field to given value.


### GetCpuType

`func (o *Host) GetCpuType() string`

GetCpuType returns the CpuType field if non-nil, zero value otherwise.

### GetCpuTypeOk

`func (o *Host) GetCpuTypeOk() (*string, bool)`

GetCpuTypeOk returns a tuple with the CpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuType

`func (o *Host) SetCpuType(v string)`

SetCpuType sets CpuType field to given value.


### GetCpuLoad

`func (o *Host) GetCpuLoad() float32`

GetCpuLoad returns the CpuLoad field if non-nil, zero value otherwise.

### GetCpuLoadOk

`func (o *Host) GetCpuLoadOk() (*float32, bool)`

GetCpuLoadOk returns a tuple with the CpuLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLoad

`func (o *Host) SetCpuLoad(v float32)`

SetCpuLoad sets CpuLoad field to given value.


### GetCpu

`func (o *Host) GetCpu() HostCpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Host) GetCpuOk() (*HostCpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Host) SetCpu(v HostCpu)`

SetCpu sets Cpu field to given value.


### GetMemUsed

`func (o *Host) GetMemUsed() int32`

GetMemUsed returns the MemUsed field if non-nil, zero value otherwise.

### GetMemUsedOk

`func (o *Host) GetMemUsedOk() (*int32, bool)`

GetMemUsedOk returns a tuple with the MemUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsed

`func (o *Host) SetMemUsed(v int32)`

SetMemUsed sets MemUsed field to given value.


### GetMemMax

`func (o *Host) GetMemMax() int32`

GetMemMax returns the MemMax field if non-nil, zero value otherwise.

### GetMemMaxOk

`func (o *Host) GetMemMaxOk() (*int32, bool)`

GetMemMaxOk returns a tuple with the MemMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemMax

`func (o *Host) SetMemMax(v int32)`

SetMemMax sets MemMax field to given value.


### GetMemFree

`func (o *Host) GetMemFree() int32`

GetMemFree returns the MemFree field if non-nil, zero value otherwise.

### GetMemFreeOk

`func (o *Host) GetMemFreeOk() (*int32, bool)`

GetMemFreeOk returns a tuple with the MemFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemFree

`func (o *Host) SetMemFree(v int32)`

SetMemFree sets MemFree field to given value.


### GetDisk

`func (o *Host) GetDisk() []HostDisk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *Host) GetDiskOk() (*[]HostDisk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *Host) SetDisk(v []HostDisk)`

SetDisk sets Disk field to given value.


### GetMemory

`func (o *Host) GetMemory() []HostMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Host) GetMemoryOk() (*[]HostMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Host) SetMemory(v []HostMemory)`

SetMemory sets Memory field to given value.


### GetIsReserve

`func (o *Host) GetIsReserve() int32`

GetIsReserve returns the IsReserve field if non-nil, zero value otherwise.

### GetIsReserveOk

`func (o *Host) GetIsReserveOk() (*int32, bool)`

GetIsReserveOk returns a tuple with the IsReserve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserve

`func (o *Host) SetIsReserve(v int32)`

SetIsReserve sets IsReserve field to given value.


### GetLabels

`func (o *Host) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Host) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Host) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Host) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetServiceTag

`func (o *Host) GetServiceTag() string`

GetServiceTag returns the ServiceTag field if non-nil, zero value otherwise.

### GetServiceTagOk

`func (o *Host) GetServiceTagOk() (*string, bool)`

GetServiceTagOk returns a tuple with the ServiceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTag

`func (o *Host) SetServiceTag(v string)`

SetServiceTag sets ServiceTag field to given value.


### GetIsODP

`func (o *Host) GetIsODP() int32`

GetIsODP returns the IsODP field if non-nil, zero value otherwise.

### GetIsODPOk

`func (o *Host) GetIsODPOk() (*int32, bool)`

GetIsODPOk returns a tuple with the IsODP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsODP

`func (o *Host) SetIsODP(v int32)`

SetIsODP sets IsODP field to given value.


### GetFmOrderId

`func (o *Host) GetFmOrderId() string`

GetFmOrderId returns the FmOrderId field if non-nil, zero value otherwise.

### GetFmOrderIdOk

`func (o *Host) GetFmOrderIdOk() (*string, bool)`

GetFmOrderIdOk returns a tuple with the FmOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFmOrderId

`func (o *Host) SetFmOrderId(v string)`

SetFmOrderId sets FmOrderId field to given value.


### SetFmOrderIdNil

`func (o *Host) SetFmOrderIdNil(b bool)`

 SetFmOrderIdNil sets the value for FmOrderId to be an explicit nil

### UnsetFmOrderId
`func (o *Host) UnsetFmOrderId()`

UnsetFmOrderId ensures that no value is present for FmOrderId, not even an explicit nil
### GetInstallStatus

`func (o *Host) GetInstallStatus() string`

GetInstallStatus returns the InstallStatus field if non-nil, zero value otherwise.

### GetInstallStatusOk

`func (o *Host) GetInstallStatusOk() (*string, bool)`

GetInstallStatusOk returns a tuple with the InstallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStatus

`func (o *Host) SetInstallStatus(v string)`

SetInstallStatus sets InstallStatus field to given value.


### SetInstallStatusNil

`func (o *Host) SetInstallStatusNil(b bool)`

 SetInstallStatusNil sets the value for InstallStatus to be an explicit nil

### UnsetInstallStatus
`func (o *Host) UnsetInstallStatus()`

UnsetInstallStatus ensures that no value is present for InstallStatus, not even an explicit nil
### GetStatus

`func (o *Host) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Host) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Host) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUptime

`func (o *Host) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *Host) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *Host) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *Host) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetTrafficSum

`func (o *Host) GetTrafficSum() int32`

GetTrafficSum returns the TrafficSum field if non-nil, zero value otherwise.

### GetTrafficSumOk

`func (o *Host) GetTrafficSumOk() (*int32, bool)`

GetTrafficSumOk returns a tuple with the TrafficSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSum

`func (o *Host) SetTrafficSum(v int32)`

SetTrafficSum sets TrafficSum field to given value.


### GetUplinks

`func (o *Host) GetUplinks() []HostUplink`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *Host) GetUplinksOk() (*[]HostUplink, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *Host) SetUplinks(v []HostUplink)`

SetUplinks sets Uplinks field to given value.


### GetFreeIncomingTraffic

`func (o *Host) GetFreeIncomingTraffic() bool`

GetFreeIncomingTraffic returns the FreeIncomingTraffic field if non-nil, zero value otherwise.

### GetFreeIncomingTrafficOk

`func (o *Host) GetFreeIncomingTrafficOk() (*bool, bool)`

GetFreeIncomingTrafficOk returns a tuple with the FreeIncomingTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeIncomingTraffic

`func (o *Host) SetFreeIncomingTraffic(v bool)`

SetFreeIncomingTraffic sets FreeIncomingTraffic field to given value.


### GetOutgoingTrafficSum

`func (o *Host) GetOutgoingTrafficSum() int32`

GetOutgoingTrafficSum returns the OutgoingTrafficSum field if non-nil, zero value otherwise.

### GetOutgoingTrafficSumOk

`func (o *Host) GetOutgoingTrafficSumOk() (*int32, bool)`

GetOutgoingTrafficSumOk returns a tuple with the OutgoingTrafficSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTrafficSum

`func (o *Host) SetOutgoingTrafficSum(v int32)`

SetOutgoingTrafficSum sets OutgoingTrafficSum field to given value.


### GetIncomingTrafficSum

`func (o *Host) GetIncomingTrafficSum() int32`

GetIncomingTrafficSum returns the IncomingTrafficSum field if non-nil, zero value otherwise.

### GetIncomingTrafficSumOk

`func (o *Host) GetIncomingTrafficSumOk() (*int32, bool)`

GetIncomingTrafficSumOk returns a tuple with the IncomingTrafficSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingTrafficSum

`func (o *Host) SetIncomingTrafficSum(v int32)`

SetIncomingTrafficSum sets IncomingTrafficSum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


