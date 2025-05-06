# OperatingSystemPartial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | For an overview of available Operating Systems, see /v3/operatingsystem. | 
**KernelParams** | Pointer to [**[]KernelParamPartial**](KernelParamPartial.md) |  | [optional] 

## Methods

### NewOperatingSystemPartial

`func NewOperatingSystemPartial(id int32, ) *OperatingSystemPartial`

NewOperatingSystemPartial instantiates a new OperatingSystemPartial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatingSystemPartialWithDefaults

`func NewOperatingSystemPartialWithDefaults() *OperatingSystemPartial`

NewOperatingSystemPartialWithDefaults instantiates a new OperatingSystemPartial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OperatingSystemPartial) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperatingSystemPartial) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperatingSystemPartial) SetId(v int32)`

SetId sets Id field to given value.


### GetKernelParams

`func (o *OperatingSystemPartial) GetKernelParams() []KernelParamPartial`

GetKernelParams returns the KernelParams field if non-nil, zero value otherwise.

### GetKernelParamsOk

`func (o *OperatingSystemPartial) GetKernelParamsOk() (*[]KernelParamPartial, bool)`

GetKernelParamsOk returns a tuple with the KernelParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelParams

`func (o *OperatingSystemPartial) SetKernelParams(v []KernelParamPartial)`

SetKernelParams sets KernelParams field to given value.

### HasKernelParams

`func (o *OperatingSystemPartial) HasKernelParams() bool`

HasKernelParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


