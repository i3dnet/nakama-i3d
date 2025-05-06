# MirrorState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | **int32** | ID of the location | 
**LocationName** | **string** | Name of the location | 
**FileServerIds** | **[]int32** | An array with the file server IDs that are used for this mirroring process | 
**Transfers** | [**[]Transfer**](Transfer.md) | An array with the file transfer progress per file server | 

## Methods

### NewMirrorState

`func NewMirrorState(locationId int32, locationName string, fileServerIds []int32, transfers []Transfer, ) *MirrorState`

NewMirrorState instantiates a new MirrorState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMirrorStateWithDefaults

`func NewMirrorStateWithDefaults() *MirrorState`

NewMirrorStateWithDefaults instantiates a new MirrorState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *MirrorState) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *MirrorState) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *MirrorState) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.


### GetLocationName

`func (o *MirrorState) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *MirrorState) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *MirrorState) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.


### GetFileServerIds

`func (o *MirrorState) GetFileServerIds() []int32`

GetFileServerIds returns the FileServerIds field if non-nil, zero value otherwise.

### GetFileServerIdsOk

`func (o *MirrorState) GetFileServerIdsOk() (*[]int32, bool)`

GetFileServerIdsOk returns a tuple with the FileServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServerIds

`func (o *MirrorState) SetFileServerIds(v []int32)`

SetFileServerIds sets FileServerIds field to given value.


### GetTransfers

`func (o *MirrorState) GetTransfers() []Transfer`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *MirrorState) GetTransfersOk() (*[]Transfer, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *MirrorState) SetTransfers(v []Transfer)`

SetTransfers sets Transfers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


