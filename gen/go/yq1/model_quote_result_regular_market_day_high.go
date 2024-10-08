/*
Yahoo Finance

Yahoo Finance API specification

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yq1

import (
	"encoding/json"
)

// checks if the QuoteResultRegularMarketDayHigh type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteResultRegularMarketDayHigh{}

// QuoteResultRegularMarketDayHigh struct for QuoteResultRegularMarketDayHigh
type QuoteResultRegularMarketDayHigh struct {
	Raw *float32 `json:"raw,omitempty"`
	Fmt *string `json:"fmt,omitempty"`
}

// NewQuoteResultRegularMarketDayHigh instantiates a new QuoteResultRegularMarketDayHigh object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteResultRegularMarketDayHigh() *QuoteResultRegularMarketDayHigh {
	this := QuoteResultRegularMarketDayHigh{}
	return &this
}

// NewQuoteResultRegularMarketDayHighWithDefaults instantiates a new QuoteResultRegularMarketDayHigh object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteResultRegularMarketDayHighWithDefaults() *QuoteResultRegularMarketDayHigh {
	this := QuoteResultRegularMarketDayHigh{}
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *QuoteResultRegularMarketDayHigh) GetRaw() float32 {
	if o == nil || IsNil(o.Raw) {
		var ret float32
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultRegularMarketDayHigh) GetRawOk() (*float32, bool) {
	if o == nil || IsNil(o.Raw) {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *QuoteResultRegularMarketDayHigh) HasRaw() bool {
	if o != nil && !IsNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given float32 and assigns it to the Raw field.
func (o *QuoteResultRegularMarketDayHigh) SetRaw(v float32) {
	o.Raw = &v
}

// GetFmt returns the Fmt field value if set, zero value otherwise.
func (o *QuoteResultRegularMarketDayHigh) GetFmt() string {
	if o == nil || IsNil(o.Fmt) {
		var ret string
		return ret
	}
	return *o.Fmt
}

// GetFmtOk returns a tuple with the Fmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultRegularMarketDayHigh) GetFmtOk() (*string, bool) {
	if o == nil || IsNil(o.Fmt) {
		return nil, false
	}
	return o.Fmt, true
}

// HasFmt returns a boolean if a field has been set.
func (o *QuoteResultRegularMarketDayHigh) HasFmt() bool {
	if o != nil && !IsNil(o.Fmt) {
		return true
	}

	return false
}

// SetFmt gets a reference to the given string and assigns it to the Fmt field.
func (o *QuoteResultRegularMarketDayHigh) SetFmt(v string) {
	o.Fmt = &v
}

func (o QuoteResultRegularMarketDayHigh) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteResultRegularMarketDayHigh) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Raw) {
		toSerialize["raw"] = o.Raw
	}
	if !IsNil(o.Fmt) {
		toSerialize["fmt"] = o.Fmt
	}
	return toSerialize, nil
}

type NullableQuoteResultRegularMarketDayHigh struct {
	value *QuoteResultRegularMarketDayHigh
	isSet bool
}

func (v NullableQuoteResultRegularMarketDayHigh) Get() *QuoteResultRegularMarketDayHigh {
	return v.value
}

func (v *NullableQuoteResultRegularMarketDayHigh) Set(val *QuoteResultRegularMarketDayHigh) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteResultRegularMarketDayHigh) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteResultRegularMarketDayHigh) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteResultRegularMarketDayHigh(val *QuoteResultRegularMarketDayHigh) *NullableQuoteResultRegularMarketDayHigh {
	return &NullableQuoteResultRegularMarketDayHigh{value: val, isSet: true}
}

func (v NullableQuoteResultRegularMarketDayHigh) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteResultRegularMarketDayHigh) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


