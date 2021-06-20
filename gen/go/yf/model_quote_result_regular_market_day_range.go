/*
 * Yahoo Finance
 *
 * Yahoo Finance API specification
 *
 * API version: 1.0.7
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yf

import (
	"encoding/json"
)

// QuoteResultRegularMarketDayRange struct for QuoteResultRegularMarketDayRange
type QuoteResultRegularMarketDayRange struct {
	Raw *string `json:"raw,omitempty"`
	Fmt *string `json:"fmt,omitempty"`
}

// NewQuoteResultRegularMarketDayRange instantiates a new QuoteResultRegularMarketDayRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteResultRegularMarketDayRange() *QuoteResultRegularMarketDayRange {
	this := QuoteResultRegularMarketDayRange{}
	return &this
}

// NewQuoteResultRegularMarketDayRangeWithDefaults instantiates a new QuoteResultRegularMarketDayRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteResultRegularMarketDayRangeWithDefaults() *QuoteResultRegularMarketDayRange {
	this := QuoteResultRegularMarketDayRange{}
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *QuoteResultRegularMarketDayRange) GetRaw() string {
	if o == nil || o.Raw == nil {
		var ret string
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultRegularMarketDayRange) GetRawOk() (*string, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *QuoteResultRegularMarketDayRange) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given string and assigns it to the Raw field.
func (o *QuoteResultRegularMarketDayRange) SetRaw(v string) {
	o.Raw = &v
}

// GetFmt returns the Fmt field value if set, zero value otherwise.
func (o *QuoteResultRegularMarketDayRange) GetFmt() string {
	if o == nil || o.Fmt == nil {
		var ret string
		return ret
	}
	return *o.Fmt
}

// GetFmtOk returns a tuple with the Fmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultRegularMarketDayRange) GetFmtOk() (*string, bool) {
	if o == nil || o.Fmt == nil {
		return nil, false
	}
	return o.Fmt, true
}

// HasFmt returns a boolean if a field has been set.
func (o *QuoteResultRegularMarketDayRange) HasFmt() bool {
	if o != nil && o.Fmt != nil {
		return true
	}

	return false
}

// SetFmt gets a reference to the given string and assigns it to the Fmt field.
func (o *QuoteResultRegularMarketDayRange) SetFmt(v string) {
	o.Fmt = &v
}

func (o QuoteResultRegularMarketDayRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Fmt != nil {
		toSerialize["fmt"] = o.Fmt
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteResultRegularMarketDayRange struct {
	value *QuoteResultRegularMarketDayRange
	isSet bool
}

func (v NullableQuoteResultRegularMarketDayRange) Get() *QuoteResultRegularMarketDayRange {
	return v.value
}

func (v *NullableQuoteResultRegularMarketDayRange) Set(val *QuoteResultRegularMarketDayRange) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteResultRegularMarketDayRange) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteResultRegularMarketDayRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteResultRegularMarketDayRange(val *QuoteResultRegularMarketDayRange) *NullableQuoteResultRegularMarketDayRange {
	return &NullableQuoteResultRegularMarketDayRange{value: val, isSet: true}
}

func (v NullableQuoteResultRegularMarketDayRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteResultRegularMarketDayRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


