# People

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Surname** | **string** |  | 
**Name** | **string** |  | 
**Patronymic** | Pointer to **string** |  | [optional] 
**Address** | **string** |  | 

## Methods

### NewPeople

`func NewPeople(surname string, name string, address string, ) *People`

NewPeople instantiates a new People object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeopleWithDefaults

`func NewPeopleWithDefaults() *People`

NewPeopleWithDefaults instantiates a new People object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSurname

`func (o *People) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *People) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *People) SetSurname(v string)`

SetSurname sets Surname field to given value.


### GetName

`func (o *People) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *People) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *People) SetName(v string)`

SetName sets Name field to given value.


### GetPatronymic

`func (o *People) GetPatronymic() string`

GetPatronymic returns the Patronymic field if non-nil, zero value otherwise.

### GetPatronymicOk

`func (o *People) GetPatronymicOk() (*string, bool)`

GetPatronymicOk returns a tuple with the Patronymic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronymic

`func (o *People) SetPatronymic(v string)`

SetPatronymic sets Patronymic field to given value.

### HasPatronymic

`func (o *People) HasPatronymic() bool`

HasPatronymic returns a boolean if a field has been set.

### GetAddress

`func (o *People) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *People) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *People) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


