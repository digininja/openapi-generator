/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"encoding/json"
)

// Banana struct for Banana
type Banana struct {
	LengthCm *float32 `json:"lengthCm,omitempty"`
	Color *string `json:"color,omitempty"`
}

// NewBanana instantiates a new Banana object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBanana() *Banana {
	this := Banana{}
	return &this
}

// NewBananaWithDefaults instantiates a new Banana object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBananaWithDefaults() *Banana {
	this := Banana{}
	return &this
}

// GetLengthCm returns the LengthCm field value if set, zero value otherwise.
func (o *Banana) GetLengthCm() float32 {
	if o == nil || o.LengthCm == nil {
		var ret float32
		return ret
	}
	return *o.LengthCm
}

// GetLengthCmOk returns a tuple with the LengthCm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Banana) GetLengthCmOk() (*float32, bool) {
	if o == nil || o.LengthCm == nil {
		return nil, false
	}
	return o.LengthCm, true
}

// HasLengthCm returns a boolean if a field has been set.
func (o *Banana) HasLengthCm() bool {
	if o != nil && o.LengthCm != nil {
		return true
	}

	return false
}

// SetLengthCm gets a reference to the given float32 and assigns it to the LengthCm field.
func (o *Banana) SetLengthCm(v float32) {
	o.LengthCm = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *Banana) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Banana) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *Banana) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *Banana) SetColor(v string) {
	o.Color = &v
}

func (o Banana) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LengthCm != nil {
		toSerialize["lengthCm"] = o.LengthCm
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	return json.Marshal(toSerialize)
}

type NullableBanana struct {
	value *Banana
	isSet bool
}

func (v NullableBanana) Get() *Banana {
	return v.value
}

func (v *NullableBanana) Set(val *Banana) {
	v.value = val
	v.isSet = true
}

func (v NullableBanana) IsSet() bool {
	return v.isSet
}

func (v *NullableBanana) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBanana(val *Banana) *NullableBanana {
	return &NullableBanana{value: val, isSet: true}
}

func (v NullableBanana) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBanana) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

