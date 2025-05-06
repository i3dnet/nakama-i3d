# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the invoice | [readonly] 
**CreationDate** | **int32** | The date the invoice was created | [readonly] 
**DueDate** | **int32** | The due date of the invoice | [readonly] 
**InvoiceNumber** | **string** | The number of the invoice | [readonly] 
**InvoiceHash** | **string** | A unique hash for this invoice | [readonly] 
**Description** | **string** | The description of the invoice | 
**Currency** | **string** | The currency that is being used | 
**CurrencySign** | **string** | The currency sign that is being used | 
**AmountIncVAT** | **string** | The amount of the invoice with VAT in cents | 
**AmountExclVAT** | **string** | The amount of the invoice without VAT in cents | 
**PaymentStatus** | **int32** | The payment status | 
**PaymentDate** | **int32** | The timestamp that the invoice was paid, if paid | 
**IsCredit** | **int32** | If this field is 1 it means that invoice is a credit invoice | 
**RemainingAmount** | **string** | The remaining amount of invoice | 
**TotalPaidAmount** | **string** | The total paid amount of invoice | 

## Methods

### NewInvoice

`func NewInvoice(id string, creationDate int32, dueDate int32, invoiceNumber string, invoiceHash string, description string, currency string, currencySign string, amountIncVAT string, amountExclVAT string, paymentStatus int32, paymentDate int32, isCredit int32, remainingAmount string, totalPaidAmount string, ) *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.


### GetCreationDate

`func (o *Invoice) GetCreationDate() int32`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Invoice) GetCreationDateOk() (*int32, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Invoice) SetCreationDate(v int32)`

SetCreationDate sets CreationDate field to given value.


### GetDueDate

`func (o *Invoice) GetDueDate() int32`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Invoice) GetDueDateOk() (*int32, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Invoice) SetDueDate(v int32)`

SetDueDate sets DueDate field to given value.


### GetInvoiceNumber

`func (o *Invoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *Invoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *Invoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetInvoiceHash

`func (o *Invoice) GetInvoiceHash() string`

GetInvoiceHash returns the InvoiceHash field if non-nil, zero value otherwise.

### GetInvoiceHashOk

`func (o *Invoice) GetInvoiceHashOk() (*string, bool)`

GetInvoiceHashOk returns a tuple with the InvoiceHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceHash

`func (o *Invoice) SetInvoiceHash(v string)`

SetInvoiceHash sets InvoiceHash field to given value.


### GetDescription

`func (o *Invoice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Invoice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Invoice) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCurrency

`func (o *Invoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCurrencySign

`func (o *Invoice) GetCurrencySign() string`

GetCurrencySign returns the CurrencySign field if non-nil, zero value otherwise.

### GetCurrencySignOk

`func (o *Invoice) GetCurrencySignOk() (*string, bool)`

GetCurrencySignOk returns a tuple with the CurrencySign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySign

`func (o *Invoice) SetCurrencySign(v string)`

SetCurrencySign sets CurrencySign field to given value.


### GetAmountIncVAT

`func (o *Invoice) GetAmountIncVAT() string`

GetAmountIncVAT returns the AmountIncVAT field if non-nil, zero value otherwise.

### GetAmountIncVATOk

`func (o *Invoice) GetAmountIncVATOk() (*string, bool)`

GetAmountIncVATOk returns a tuple with the AmountIncVAT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountIncVAT

`func (o *Invoice) SetAmountIncVAT(v string)`

SetAmountIncVAT sets AmountIncVAT field to given value.


### GetAmountExclVAT

`func (o *Invoice) GetAmountExclVAT() string`

GetAmountExclVAT returns the AmountExclVAT field if non-nil, zero value otherwise.

### GetAmountExclVATOk

`func (o *Invoice) GetAmountExclVATOk() (*string, bool)`

GetAmountExclVATOk returns a tuple with the AmountExclVAT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountExclVAT

`func (o *Invoice) SetAmountExclVAT(v string)`

SetAmountExclVAT sets AmountExclVAT field to given value.


### GetPaymentStatus

`func (o *Invoice) GetPaymentStatus() int32`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *Invoice) GetPaymentStatusOk() (*int32, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *Invoice) SetPaymentStatus(v int32)`

SetPaymentStatus sets PaymentStatus field to given value.


### GetPaymentDate

`func (o *Invoice) GetPaymentDate() int32`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *Invoice) GetPaymentDateOk() (*int32, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *Invoice) SetPaymentDate(v int32)`

SetPaymentDate sets PaymentDate field to given value.


### GetIsCredit

`func (o *Invoice) GetIsCredit() int32`

GetIsCredit returns the IsCredit field if non-nil, zero value otherwise.

### GetIsCreditOk

`func (o *Invoice) GetIsCreditOk() (*int32, bool)`

GetIsCreditOk returns a tuple with the IsCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCredit

`func (o *Invoice) SetIsCredit(v int32)`

SetIsCredit sets IsCredit field to given value.


### GetRemainingAmount

`func (o *Invoice) GetRemainingAmount() string`

GetRemainingAmount returns the RemainingAmount field if non-nil, zero value otherwise.

### GetRemainingAmountOk

`func (o *Invoice) GetRemainingAmountOk() (*string, bool)`

GetRemainingAmountOk returns a tuple with the RemainingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingAmount

`func (o *Invoice) SetRemainingAmount(v string)`

SetRemainingAmount sets RemainingAmount field to given value.


### GetTotalPaidAmount

`func (o *Invoice) GetTotalPaidAmount() string`

GetTotalPaidAmount returns the TotalPaidAmount field if non-nil, zero value otherwise.

### GetTotalPaidAmountOk

`func (o *Invoice) GetTotalPaidAmountOk() (*string, bool)`

GetTotalPaidAmountOk returns a tuple with the TotalPaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaidAmount

`func (o *Invoice) SetTotalPaidAmount(v string)`

SetTotalPaidAmount sets TotalPaidAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


