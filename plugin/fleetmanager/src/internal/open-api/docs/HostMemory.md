# HostMemory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | **string** | The brand name | [readonly] 
**Model** | **string** | The model name | [readonly] 
**Size** | **int32** | In bytes | [readonly] 
**Speed** | **int32** | The speed this memory module runs at | [readonly] 
**Ecc** | **int32** | Set to 1 if ecc is supported and enabled | [readonly] 
**MemoryBank** | **int32** | The bank this module sits in | [readonly] 
**MemoryType** | **string** | The type of memory module | [readonly] 
**MemorySlot** | **string** | The slot this module sits in | [readonly] 
**MemorySerial** | **string** | The serial number of this module | [readonly] 

## Methods

### NewHostMemory

`func NewHostMemory(brand string, model string, size int32, speed int32, ecc int32, memoryBank int32, memoryType string, memorySlot string, memorySerial string, ) *HostMemory`

NewHostMemory instantiates a new HostMemory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostMemoryWithDefaults

`func NewHostMemoryWithDefaults() *HostMemory`

NewHostMemoryWithDefaults instantiates a new HostMemory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *HostMemory) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *HostMemory) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *HostMemory) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetModel

`func (o *HostMemory) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HostMemory) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HostMemory) SetModel(v string)`

SetModel sets Model field to given value.


### GetSize

`func (o *HostMemory) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *HostMemory) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *HostMemory) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSpeed

`func (o *HostMemory) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *HostMemory) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *HostMemory) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.


### GetEcc

`func (o *HostMemory) GetEcc() int32`

GetEcc returns the Ecc field if non-nil, zero value otherwise.

### GetEccOk

`func (o *HostMemory) GetEccOk() (*int32, bool)`

GetEccOk returns a tuple with the Ecc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcc

`func (o *HostMemory) SetEcc(v int32)`

SetEcc sets Ecc field to given value.


### GetMemoryBank

`func (o *HostMemory) GetMemoryBank() int32`

GetMemoryBank returns the MemoryBank field if non-nil, zero value otherwise.

### GetMemoryBankOk

`func (o *HostMemory) GetMemoryBankOk() (*int32, bool)`

GetMemoryBankOk returns a tuple with the MemoryBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBank

`func (o *HostMemory) SetMemoryBank(v int32)`

SetMemoryBank sets MemoryBank field to given value.


### GetMemoryType

`func (o *HostMemory) GetMemoryType() string`

GetMemoryType returns the MemoryType field if non-nil, zero value otherwise.

### GetMemoryTypeOk

`func (o *HostMemory) GetMemoryTypeOk() (*string, bool)`

GetMemoryTypeOk returns a tuple with the MemoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryType

`func (o *HostMemory) SetMemoryType(v string)`

SetMemoryType sets MemoryType field to given value.


### GetMemorySlot

`func (o *HostMemory) GetMemorySlot() string`

GetMemorySlot returns the MemorySlot field if non-nil, zero value otherwise.

### GetMemorySlotOk

`func (o *HostMemory) GetMemorySlotOk() (*string, bool)`

GetMemorySlotOk returns a tuple with the MemorySlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySlot

`func (o *HostMemory) SetMemorySlot(v string)`

SetMemorySlot sets MemorySlot field to given value.


### GetMemorySerial

`func (o *HostMemory) GetMemorySerial() string`

GetMemorySerial returns the MemorySerial field if non-nil, zero value otherwise.

### GetMemorySerialOk

`func (o *HostMemory) GetMemorySerialOk() (*string, bool)`

GetMemorySerialOk returns a tuple with the MemorySerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySerial

`func (o *HostMemory) SetMemorySerial(v string)`

SetMemorySerial sets MemorySerial field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


