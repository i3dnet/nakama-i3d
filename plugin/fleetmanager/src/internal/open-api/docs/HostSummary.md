# HostSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of this host | [readonly] 
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
**Brand** | **string** | The server manufacturer brand name | [readonly] 
**Model** | **string** | The model of the server | [readonly] 
**IsReserve** | **int32** | If the host is reserved for a fleet isReserve set to 1 | [readonly] 
**ServiceTag** | **string** | The service tag of the host | [readonly] 
**IsODP** | **int32** | If the host is available for game hosting platform isODP set to 1 | [readonly] 
**FmOrderId** | **NullableString** | ID of Flex Metal order if this is a flex metal server | [readonly] 
**InstallStatus** | **NullableString** | Status of server auto install if it was started by one. (created / installing / finished / failed) | [readonly] 
**Status** | **string** | Whether the host is running or not | [readonly] 
**Uptime** | Pointer to **int32** | The uptime of the host. Read-only. Only provided if the server checked-in recently | [optional] 
**TrafficSum** | **int32** | Monthly traffic of all incoming and outgoing bandwidth in GB or Mbit depending on host configuration | [readonly] 
**FreeIncomingTraffic** | **bool** | Whether the host has free incoming traffic or not | [readonly] 
**OutgoingTrafficSum** | **int32** | The outgoing traffic (Mbit or GB) | [readonly] 
**IncomingTrafficSum** | **int32** | The incoming traffic (Mbit or GB) | [readonly] 

## Methods

### NewHostSummary

`func NewHostSummary(id int32, serverId int32, serverName string, serverType int32, liveHostName string, isVirtual int32, category string, osId int32, locationId int32, dcLocationId int32, instanceType string, fleetId string, newFleetId NullableString, fleetAssociatedSince int32, rackName string, dateStart string, dateEnd string, dateCancelled string, dateEndContract string, contractPeriod int32, extendPeriod int32, cancellationPeriod int32, purchaseOrder string, paymentTerm int32, pricePerMonth string, pricePerTbOveruse string, currencyId int32, bandwidthBillingType int32, bandwidthContractual int32, brand string, model string, isReserve int32, serviceTag string, isODP int32, fmOrderId NullableString, installStatus NullableString, status string, trafficSum int32, freeIncomingTraffic bool, outgoingTrafficSum int32, incomingTrafficSum int32, ) *HostSummary`

NewHostSummary instantiates a new HostSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostSummaryWithDefaults

`func NewHostSummaryWithDefaults() *HostSummary`

NewHostSummaryWithDefaults instantiates a new HostSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HostSummary) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostSummary) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostSummary) SetId(v int32)`

SetId sets Id field to given value.


### GetServerId

`func (o *HostSummary) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *HostSummary) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *HostSummary) SetServerId(v int32)`

SetServerId sets ServerId field to given value.


### GetServerName

`func (o *HostSummary) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *HostSummary) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *HostSummary) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetServerType

`func (o *HostSummary) GetServerType() int32`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *HostSummary) GetServerTypeOk() (*int32, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *HostSummary) SetServerType(v int32)`

SetServerType sets ServerType field to given value.


### GetProjectName

`func (o *HostSummary) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *HostSummary) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *HostSummary) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *HostSummary) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetClientServerName

`func (o *HostSummary) GetClientServerName() string`

GetClientServerName returns the ClientServerName field if non-nil, zero value otherwise.

### GetClientServerNameOk

`func (o *HostSummary) GetClientServerNameOk() (*string, bool)`

GetClientServerNameOk returns a tuple with the ClientServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientServerName

`func (o *HostSummary) SetClientServerName(v string)`

SetClientServerName sets ClientServerName field to given value.

### HasClientServerName

`func (o *HostSummary) HasClientServerName() bool`

HasClientServerName returns a boolean if a field has been set.

### GetClientState

`func (o *HostSummary) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *HostSummary) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *HostSummary) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *HostSummary) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetLiveHostName

`func (o *HostSummary) GetLiveHostName() string`

GetLiveHostName returns the LiveHostName field if non-nil, zero value otherwise.

### GetLiveHostNameOk

`func (o *HostSummary) GetLiveHostNameOk() (*string, bool)`

GetLiveHostNameOk returns a tuple with the LiveHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveHostName

`func (o *HostSummary) SetLiveHostName(v string)`

SetLiveHostName sets LiveHostName field to given value.


### GetIsVirtual

`func (o *HostSummary) GetIsVirtual() int32`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *HostSummary) GetIsVirtualOk() (*int32, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *HostSummary) SetIsVirtual(v int32)`

SetIsVirtual sets IsVirtual field to given value.


### GetCategory

`func (o *HostSummary) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HostSummary) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HostSummary) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetOsId

`func (o *HostSummary) GetOsId() int32`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *HostSummary) GetOsIdOk() (*int32, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *HostSummary) SetOsId(v int32)`

SetOsId sets OsId field to given value.


### GetLocationId

`func (o *HostSummary) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *HostSummary) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *HostSummary) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.


### GetDcLocationId

`func (o *HostSummary) GetDcLocationId() int32`

GetDcLocationId returns the DcLocationId field if non-nil, zero value otherwise.

### GetDcLocationIdOk

`func (o *HostSummary) GetDcLocationIdOk() (*int32, bool)`

GetDcLocationIdOk returns a tuple with the DcLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcLocationId

`func (o *HostSummary) SetDcLocationId(v int32)`

SetDcLocationId sets DcLocationId field to given value.


### GetInstanceType

`func (o *HostSummary) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *HostSummary) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *HostSummary) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetFleetId

`func (o *HostSummary) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *HostSummary) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *HostSummary) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetNewFleetId

`func (o *HostSummary) GetNewFleetId() string`

GetNewFleetId returns the NewFleetId field if non-nil, zero value otherwise.

### GetNewFleetIdOk

`func (o *HostSummary) GetNewFleetIdOk() (*string, bool)`

GetNewFleetIdOk returns a tuple with the NewFleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFleetId

`func (o *HostSummary) SetNewFleetId(v string)`

SetNewFleetId sets NewFleetId field to given value.


### SetNewFleetIdNil

`func (o *HostSummary) SetNewFleetIdNil(b bool)`

 SetNewFleetIdNil sets the value for NewFleetId to be an explicit nil

### UnsetNewFleetId
`func (o *HostSummary) UnsetNewFleetId()`

UnsetNewFleetId ensures that no value is present for NewFleetId, not even an explicit nil
### GetFleetAssociatedSince

`func (o *HostSummary) GetFleetAssociatedSince() int32`

GetFleetAssociatedSince returns the FleetAssociatedSince field if non-nil, zero value otherwise.

### GetFleetAssociatedSinceOk

`func (o *HostSummary) GetFleetAssociatedSinceOk() (*int32, bool)`

GetFleetAssociatedSinceOk returns a tuple with the FleetAssociatedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetAssociatedSince

`func (o *HostSummary) SetFleetAssociatedSince(v int32)`

SetFleetAssociatedSince sets FleetAssociatedSince field to given value.


### GetRackName

`func (o *HostSummary) GetRackName() string`

GetRackName returns the RackName field if non-nil, zero value otherwise.

### GetRackNameOk

`func (o *HostSummary) GetRackNameOk() (*string, bool)`

GetRackNameOk returns a tuple with the RackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackName

`func (o *HostSummary) SetRackName(v string)`

SetRackName sets RackName field to given value.


### GetDateStart

`func (o *HostSummary) GetDateStart() string`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *HostSummary) GetDateStartOk() (*string, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *HostSummary) SetDateStart(v string)`

SetDateStart sets DateStart field to given value.


### GetDateEnd

`func (o *HostSummary) GetDateEnd() string`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *HostSummary) GetDateEndOk() (*string, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *HostSummary) SetDateEnd(v string)`

SetDateEnd sets DateEnd field to given value.


### GetDateCancelled

`func (o *HostSummary) GetDateCancelled() string`

GetDateCancelled returns the DateCancelled field if non-nil, zero value otherwise.

### GetDateCancelledOk

`func (o *HostSummary) GetDateCancelledOk() (*string, bool)`

GetDateCancelledOk returns a tuple with the DateCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCancelled

`func (o *HostSummary) SetDateCancelled(v string)`

SetDateCancelled sets DateCancelled field to given value.


### GetDateEndContract

`func (o *HostSummary) GetDateEndContract() string`

GetDateEndContract returns the DateEndContract field if non-nil, zero value otherwise.

### GetDateEndContractOk

`func (o *HostSummary) GetDateEndContractOk() (*string, bool)`

GetDateEndContractOk returns a tuple with the DateEndContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEndContract

`func (o *HostSummary) SetDateEndContract(v string)`

SetDateEndContract sets DateEndContract field to given value.


### GetContractPeriod

`func (o *HostSummary) GetContractPeriod() int32`

GetContractPeriod returns the ContractPeriod field if non-nil, zero value otherwise.

### GetContractPeriodOk

`func (o *HostSummary) GetContractPeriodOk() (*int32, bool)`

GetContractPeriodOk returns a tuple with the ContractPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractPeriod

`func (o *HostSummary) SetContractPeriod(v int32)`

SetContractPeriod sets ContractPeriod field to given value.


### GetExtendPeriod

`func (o *HostSummary) GetExtendPeriod() int32`

GetExtendPeriod returns the ExtendPeriod field if non-nil, zero value otherwise.

### GetExtendPeriodOk

`func (o *HostSummary) GetExtendPeriodOk() (*int32, bool)`

GetExtendPeriodOk returns a tuple with the ExtendPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendPeriod

`func (o *HostSummary) SetExtendPeriod(v int32)`

SetExtendPeriod sets ExtendPeriod field to given value.


### GetCancellationPeriod

`func (o *HostSummary) GetCancellationPeriod() int32`

GetCancellationPeriod returns the CancellationPeriod field if non-nil, zero value otherwise.

### GetCancellationPeriodOk

`func (o *HostSummary) GetCancellationPeriodOk() (*int32, bool)`

GetCancellationPeriodOk returns a tuple with the CancellationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationPeriod

`func (o *HostSummary) SetCancellationPeriod(v int32)`

SetCancellationPeriod sets CancellationPeriod field to given value.


### GetPurchaseOrder

`func (o *HostSummary) GetPurchaseOrder() string`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *HostSummary) GetPurchaseOrderOk() (*string, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *HostSummary) SetPurchaseOrder(v string)`

SetPurchaseOrder sets PurchaseOrder field to given value.


### GetPaymentTerm

`func (o *HostSummary) GetPaymentTerm() int32`

GetPaymentTerm returns the PaymentTerm field if non-nil, zero value otherwise.

### GetPaymentTermOk

`func (o *HostSummary) GetPaymentTermOk() (*int32, bool)`

GetPaymentTermOk returns a tuple with the PaymentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerm

`func (o *HostSummary) SetPaymentTerm(v int32)`

SetPaymentTerm sets PaymentTerm field to given value.


### GetPricePerMonth

`func (o *HostSummary) GetPricePerMonth() string`

GetPricePerMonth returns the PricePerMonth field if non-nil, zero value otherwise.

### GetPricePerMonthOk

`func (o *HostSummary) GetPricePerMonthOk() (*string, bool)`

GetPricePerMonthOk returns a tuple with the PricePerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerMonth

`func (o *HostSummary) SetPricePerMonth(v string)`

SetPricePerMonth sets PricePerMonth field to given value.


### GetPricePerTbOveruse

`func (o *HostSummary) GetPricePerTbOveruse() string`

GetPricePerTbOveruse returns the PricePerTbOveruse field if non-nil, zero value otherwise.

### GetPricePerTbOveruseOk

`func (o *HostSummary) GetPricePerTbOveruseOk() (*string, bool)`

GetPricePerTbOveruseOk returns a tuple with the PricePerTbOveruse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerTbOveruse

`func (o *HostSummary) SetPricePerTbOveruse(v string)`

SetPricePerTbOveruse sets PricePerTbOveruse field to given value.


### GetCurrencyId

`func (o *HostSummary) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *HostSummary) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *HostSummary) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.


### GetBandwidthBillingType

`func (o *HostSummary) GetBandwidthBillingType() int32`

GetBandwidthBillingType returns the BandwidthBillingType field if non-nil, zero value otherwise.

### GetBandwidthBillingTypeOk

`func (o *HostSummary) GetBandwidthBillingTypeOk() (*int32, bool)`

GetBandwidthBillingTypeOk returns a tuple with the BandwidthBillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthBillingType

`func (o *HostSummary) SetBandwidthBillingType(v int32)`

SetBandwidthBillingType sets BandwidthBillingType field to given value.


### GetBandwidthContractual

`func (o *HostSummary) GetBandwidthContractual() int32`

GetBandwidthContractual returns the BandwidthContractual field if non-nil, zero value otherwise.

### GetBandwidthContractualOk

`func (o *HostSummary) GetBandwidthContractualOk() (*int32, bool)`

GetBandwidthContractualOk returns a tuple with the BandwidthContractual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthContractual

`func (o *HostSummary) SetBandwidthContractual(v int32)`

SetBandwidthContractual sets BandwidthContractual field to given value.


### GetBrand

`func (o *HostSummary) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *HostSummary) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *HostSummary) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetModel

`func (o *HostSummary) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HostSummary) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HostSummary) SetModel(v string)`

SetModel sets Model field to given value.


### GetIsReserve

`func (o *HostSummary) GetIsReserve() int32`

GetIsReserve returns the IsReserve field if non-nil, zero value otherwise.

### GetIsReserveOk

`func (o *HostSummary) GetIsReserveOk() (*int32, bool)`

GetIsReserveOk returns a tuple with the IsReserve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserve

`func (o *HostSummary) SetIsReserve(v int32)`

SetIsReserve sets IsReserve field to given value.


### GetServiceTag

`func (o *HostSummary) GetServiceTag() string`

GetServiceTag returns the ServiceTag field if non-nil, zero value otherwise.

### GetServiceTagOk

`func (o *HostSummary) GetServiceTagOk() (*string, bool)`

GetServiceTagOk returns a tuple with the ServiceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTag

`func (o *HostSummary) SetServiceTag(v string)`

SetServiceTag sets ServiceTag field to given value.


### GetIsODP

`func (o *HostSummary) GetIsODP() int32`

GetIsODP returns the IsODP field if non-nil, zero value otherwise.

### GetIsODPOk

`func (o *HostSummary) GetIsODPOk() (*int32, bool)`

GetIsODPOk returns a tuple with the IsODP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsODP

`func (o *HostSummary) SetIsODP(v int32)`

SetIsODP sets IsODP field to given value.


### GetFmOrderId

`func (o *HostSummary) GetFmOrderId() string`

GetFmOrderId returns the FmOrderId field if non-nil, zero value otherwise.

### GetFmOrderIdOk

`func (o *HostSummary) GetFmOrderIdOk() (*string, bool)`

GetFmOrderIdOk returns a tuple with the FmOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFmOrderId

`func (o *HostSummary) SetFmOrderId(v string)`

SetFmOrderId sets FmOrderId field to given value.


### SetFmOrderIdNil

`func (o *HostSummary) SetFmOrderIdNil(b bool)`

 SetFmOrderIdNil sets the value for FmOrderId to be an explicit nil

### UnsetFmOrderId
`func (o *HostSummary) UnsetFmOrderId()`

UnsetFmOrderId ensures that no value is present for FmOrderId, not even an explicit nil
### GetInstallStatus

`func (o *HostSummary) GetInstallStatus() string`

GetInstallStatus returns the InstallStatus field if non-nil, zero value otherwise.

### GetInstallStatusOk

`func (o *HostSummary) GetInstallStatusOk() (*string, bool)`

GetInstallStatusOk returns a tuple with the InstallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStatus

`func (o *HostSummary) SetInstallStatus(v string)`

SetInstallStatus sets InstallStatus field to given value.


### SetInstallStatusNil

`func (o *HostSummary) SetInstallStatusNil(b bool)`

 SetInstallStatusNil sets the value for InstallStatus to be an explicit nil

### UnsetInstallStatus
`func (o *HostSummary) UnsetInstallStatus()`

UnsetInstallStatus ensures that no value is present for InstallStatus, not even an explicit nil
### GetStatus

`func (o *HostSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUptime

`func (o *HostSummary) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *HostSummary) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *HostSummary) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *HostSummary) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetTrafficSum

`func (o *HostSummary) GetTrafficSum() int32`

GetTrafficSum returns the TrafficSum field if non-nil, zero value otherwise.

### GetTrafficSumOk

`func (o *HostSummary) GetTrafficSumOk() (*int32, bool)`

GetTrafficSumOk returns a tuple with the TrafficSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSum

`func (o *HostSummary) SetTrafficSum(v int32)`

SetTrafficSum sets TrafficSum field to given value.


### GetFreeIncomingTraffic

`func (o *HostSummary) GetFreeIncomingTraffic() bool`

GetFreeIncomingTraffic returns the FreeIncomingTraffic field if non-nil, zero value otherwise.

### GetFreeIncomingTrafficOk

`func (o *HostSummary) GetFreeIncomingTrafficOk() (*bool, bool)`

GetFreeIncomingTrafficOk returns a tuple with the FreeIncomingTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeIncomingTraffic

`func (o *HostSummary) SetFreeIncomingTraffic(v bool)`

SetFreeIncomingTraffic sets FreeIncomingTraffic field to given value.


### GetOutgoingTrafficSum

`func (o *HostSummary) GetOutgoingTrafficSum() int32`

GetOutgoingTrafficSum returns the OutgoingTrafficSum field if non-nil, zero value otherwise.

### GetOutgoingTrafficSumOk

`func (o *HostSummary) GetOutgoingTrafficSumOk() (*int32, bool)`

GetOutgoingTrafficSumOk returns a tuple with the OutgoingTrafficSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTrafficSum

`func (o *HostSummary) SetOutgoingTrafficSum(v int32)`

SetOutgoingTrafficSum sets OutgoingTrafficSum field to given value.


### GetIncomingTrafficSum

`func (o *HostSummary) GetIncomingTrafficSum() int32`

GetIncomingTrafficSum returns the IncomingTrafficSum field if non-nil, zero value otherwise.

### GetIncomingTrafficSumOk

`func (o *HostSummary) GetIncomingTrafficSumOk() (*int32, bool)`

GetIncomingTrafficSumOk returns a tuple with the IncomingTrafficSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingTrafficSum

`func (o *HostSummary) SetIncomingTrafficSum(v int32)`

SetIncomingTrafficSum sets IncomingTrafficSum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


