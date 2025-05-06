# HostDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskType** | **string** | The type of disk | [readonly] 
**DiskMedium** | **string** | The medium of this disk | [readonly] 
**Model** | **string** | The model name of this disk | [readonly] 
**Product** | **string** | The product string of this disk | [readonly] 
**DiskSerial** | **string** | The serial number of this disk | [readonly] 
**FirmwareVersion** | **string** | Firmware version | [readonly] 
**RotationRate** | **int32** | Rotation rate (does not apply to SSD) | [readonly] 
**SectorSizeLogical** | **int32** | Logical sector size | [readonly] 
**SectorSizePhysical** | **int32** | Physical sector size | [readonly] 
**Size** | **int32** | In bytes | [readonly] 

## Methods

### NewHostDisk

`func NewHostDisk(diskType string, diskMedium string, model string, product string, diskSerial string, firmwareVersion string, rotationRate int32, sectorSizeLogical int32, sectorSizePhysical int32, size int32, ) *HostDisk`

NewHostDisk instantiates a new HostDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostDiskWithDefaults

`func NewHostDiskWithDefaults() *HostDisk`

NewHostDiskWithDefaults instantiates a new HostDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskType

`func (o *HostDisk) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *HostDisk) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *HostDisk) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.


### GetDiskMedium

`func (o *HostDisk) GetDiskMedium() string`

GetDiskMedium returns the DiskMedium field if non-nil, zero value otherwise.

### GetDiskMediumOk

`func (o *HostDisk) GetDiskMediumOk() (*string, bool)`

GetDiskMediumOk returns a tuple with the DiskMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMedium

`func (o *HostDisk) SetDiskMedium(v string)`

SetDiskMedium sets DiskMedium field to given value.


### GetModel

`func (o *HostDisk) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HostDisk) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HostDisk) SetModel(v string)`

SetModel sets Model field to given value.


### GetProduct

`func (o *HostDisk) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *HostDisk) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *HostDisk) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetDiskSerial

`func (o *HostDisk) GetDiskSerial() string`

GetDiskSerial returns the DiskSerial field if non-nil, zero value otherwise.

### GetDiskSerialOk

`func (o *HostDisk) GetDiskSerialOk() (*string, bool)`

GetDiskSerialOk returns a tuple with the DiskSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSerial

`func (o *HostDisk) SetDiskSerial(v string)`

SetDiskSerial sets DiskSerial field to given value.


### GetFirmwareVersion

`func (o *HostDisk) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *HostDisk) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *HostDisk) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.


### GetRotationRate

`func (o *HostDisk) GetRotationRate() int32`

GetRotationRate returns the RotationRate field if non-nil, zero value otherwise.

### GetRotationRateOk

`func (o *HostDisk) GetRotationRateOk() (*int32, bool)`

GetRotationRateOk returns a tuple with the RotationRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationRate

`func (o *HostDisk) SetRotationRate(v int32)`

SetRotationRate sets RotationRate field to given value.


### GetSectorSizeLogical

`func (o *HostDisk) GetSectorSizeLogical() int32`

GetSectorSizeLogical returns the SectorSizeLogical field if non-nil, zero value otherwise.

### GetSectorSizeLogicalOk

`func (o *HostDisk) GetSectorSizeLogicalOk() (*int32, bool)`

GetSectorSizeLogicalOk returns a tuple with the SectorSizeLogical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorSizeLogical

`func (o *HostDisk) SetSectorSizeLogical(v int32)`

SetSectorSizeLogical sets SectorSizeLogical field to given value.


### GetSectorSizePhysical

`func (o *HostDisk) GetSectorSizePhysical() int32`

GetSectorSizePhysical returns the SectorSizePhysical field if non-nil, zero value otherwise.

### GetSectorSizePhysicalOk

`func (o *HostDisk) GetSectorSizePhysicalOk() (*int32, bool)`

GetSectorSizePhysicalOk returns a tuple with the SectorSizePhysical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorSizePhysical

`func (o *HostDisk) SetSectorSizePhysical(v int32)`

SetSectorSizePhysical sets SectorSizePhysical field to given value.


### GetSize

`func (o *HostDisk) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *HostDisk) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *HostDisk) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


