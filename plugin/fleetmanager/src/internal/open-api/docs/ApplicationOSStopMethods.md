# ApplicationOSStopMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Linux** | [**[]ApplicationStopMethod**](ApplicationStopMethod.md) | Stop methods for Linux. | [readonly] 
**Windows** | [**[]ApplicationStopMethod**](ApplicationStopMethod.md) | Stop methods for Windows. | [readonly] 

## Methods

### NewApplicationOSStopMethods

`func NewApplicationOSStopMethods(linux []ApplicationStopMethod, windows []ApplicationStopMethod, ) *ApplicationOSStopMethods`

NewApplicationOSStopMethods instantiates a new ApplicationOSStopMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationOSStopMethodsWithDefaults

`func NewApplicationOSStopMethodsWithDefaults() *ApplicationOSStopMethods`

NewApplicationOSStopMethodsWithDefaults instantiates a new ApplicationOSStopMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinux

`func (o *ApplicationOSStopMethods) GetLinux() []ApplicationStopMethod`

GetLinux returns the Linux field if non-nil, zero value otherwise.

### GetLinuxOk

`func (o *ApplicationOSStopMethods) GetLinuxOk() (*[]ApplicationStopMethod, bool)`

GetLinuxOk returns a tuple with the Linux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinux

`func (o *ApplicationOSStopMethods) SetLinux(v []ApplicationStopMethod)`

SetLinux sets Linux field to given value.


### GetWindows

`func (o *ApplicationOSStopMethods) GetWindows() []ApplicationStopMethod`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *ApplicationOSStopMethods) GetWindowsOk() (*[]ApplicationStopMethod, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *ApplicationOSStopMethods) SetWindows(v []ApplicationStopMethod)`

SetWindows sets Windows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


