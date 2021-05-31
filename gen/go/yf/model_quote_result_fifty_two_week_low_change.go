/*
 * Yahoo Finance
 *
 * Yahoo Finance API specification
 *
 * API version: 1.0.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yf

import (
	"encoding/json"
)

// QuoteResultFiftyTwoWeekLowChange struct for QuoteResultFiftyTwoWeekLowChange
type QuoteResultFiftyTwoWeekLowChange struct {
	Raw *float32 `json:"raw,omitempty"`
	Fmt *string `json:"fmt,omitempty"`
}

// NewQuoteResultFiftyTwoWeekLowChange instantiates a new QuoteResultFiftyTwoWeekLowChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteResultFiftyTwoWeekLowChange() *QuoteResultFiftyTwoWeekLowChange {
	this := QuoteResultFiftyTwoWeekLowChange{}
	return &this
}

// NewQuoteResultFiftyTwoWeekLowChangeWithDefaults instantiates a new QuoteResultFiftyTwoWeekLowChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteResultFiftyTwoWeekLowChangeWithDefaults() *QuoteResultFiftyTwoWeekLowChange {
	this := QuoteResultFiftyTwoWeekLowChange{}
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *QuoteResultFiftyTwoWeekLowChange) GetRaw() float32 {
	if o == nil || o.Raw == nil {
		var ret float32
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultFiftyTwoWeekLowChange) GetRawOk() (*float32, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *QuoteResultFiftyTwoWeekLowChange) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given float32 and assigns it to the Raw field.
func (o *QuoteResultFiftyTwoWeekLowChange) SetRaw(v float32) {
	o.Raw = &v
}

// GetFmt returns the Fmt field value if set, zero value otherwise.
func (o *QuoteResultFiftyTwoWeekLowChange) GetFmt() string {
	if o == nil || o.Fmt == nil {
		var ret string
		return ret
	}
	return *o.Fmt
}

// GetFmtOk returns a tuple with the Fmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultFiftyTwoWeekLowChange) GetFmtOk() (*string, bool) {
	if o == nil || o.Fmt == nil {
		return nil, false
	}
	return o.Fmt, true
}

// HasFmt returns a boolean if a field has been set.
func (o *QuoteResultFiftyTwoWeekLowChange) HasFmt() bool {
	if o != nil && o.Fmt != nil {
		return true
	}

	return false
}

// SetFmt gets a reference to the given string and assigns it to the Fmt field.
func (o *QuoteResultFiftyTwoWeekLowChange) SetFmt(v string) {
	o.Fmt = &v
}

func (o QuoteResultFiftyTwoWeekLowChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Fmt != nil {
		toSerialize["fmt"] = o.Fmt
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteResultFiftyTwoWeekLowChange struct {
	value *QuoteResultFiftyTwoWeekLowChange
	isSet bool
}

func (v NullableQuoteResultFiftyTwoWeekLowChange) Get() *QuoteResultFiftyTwoWeekLowChange {
	return v.value
}

func (v *NullableQuoteResultFiftyTwoWeekLowChange) Set(val *QuoteResultFiftyTwoWeekLowChange) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteResultFiftyTwoWeekLowChange) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteResultFiftyTwoWeekLowChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteResultFiftyTwoWeekLowChange(val *QuoteResultFiftyTwoWeekLowChange) *NullableQuoteResultFiftyTwoWeekLowChange {
	return &NullableQuoteResultFiftyTwoWeekLowChange{value: val, isSet: true}
}

func (v NullableQuoteResultFiftyTwoWeekLowChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteResultFiftyTwoWeekLowChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


