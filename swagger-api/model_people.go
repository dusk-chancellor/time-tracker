/*
People info

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the People type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &People{}

// People struct for People
type People struct {
	Surname string `json:"surname"`
	Name string `json:"name"`
	Patronymic *string `json:"patronymic,omitempty"`
	Address string `json:"address"`
}

type _People People

// NewPeople instantiates a new People object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeople(surname string, name string, address string) *People {
	this := People{}
	this.Surname = surname
	this.Name = name
	this.Address = address
	return &this
}

// NewPeopleWithDefaults instantiates a new People object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeopleWithDefaults() *People {
	this := People{}
	return &this
}

// GetSurname returns the Surname field value
func (o *People) GetSurname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value
// and a boolean to check if the value has been set.
func (o *People) GetSurnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Surname, true
}

// SetSurname sets field value
func (o *People) SetSurname(v string) {
	o.Surname = v
}

// GetName returns the Name field value
func (o *People) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *People) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *People) SetName(v string) {
	o.Name = v
}

// GetPatronymic returns the Patronymic field value if set, zero value otherwise.
func (o *People) GetPatronymic() string {
	if o == nil || IsNil(o.Patronymic) {
		var ret string
		return ret
	}
	return *o.Patronymic
}

// GetPatronymicOk returns a tuple with the Patronymic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *People) GetPatronymicOk() (*string, bool) {
	if o == nil || IsNil(o.Patronymic) {
		return nil, false
	}
	return o.Patronymic, true
}

// HasPatronymic returns a boolean if a field has been set.
func (o *People) HasPatronymic() bool {
	if o != nil && !IsNil(o.Patronymic) {
		return true
	}

	return false
}

// SetPatronymic gets a reference to the given string and assigns it to the Patronymic field.
func (o *People) SetPatronymic(v string) {
	o.Patronymic = &v
}

// GetAddress returns the Address field value
func (o *People) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *People) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *People) SetAddress(v string) {
	o.Address = v
}

func (o People) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o People) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["surname"] = o.Surname
	toSerialize["name"] = o.Name
	if !IsNil(o.Patronymic) {
		toSerialize["patronymic"] = o.Patronymic
	}
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

func (o *People) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"surname",
		"name",
		"address",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPeople := _People{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPeople)

	if err != nil {
		return err
	}

	*o = People(varPeople)

	return err
}

type NullablePeople struct {
	value *People
	isSet bool
}

func (v NullablePeople) Get() *People {
	return v.value
}

func (v *NullablePeople) Set(val *People) {
	v.value = val
	v.isSet = true
}

func (v NullablePeople) IsSet() bool {
	return v.isSet
}

func (v *NullablePeople) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeople(val *People) *NullablePeople {
	return &NullablePeople{value: val, isSet: true}
}

func (v NullablePeople) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeople) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


