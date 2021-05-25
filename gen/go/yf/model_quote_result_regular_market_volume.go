/*
 * Yahoo Finance
 *
 * Yahoo Finance API specification
 *
 * API version: 1.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yf

import (
	"encoding/json"
)

// QuoteResultRegularMarketVolume struct for QuoteResultRegularMarketVolume
type QuoteResultRegularMarketVolume struct {
	Raw *int64 `json:"raw,omitempty"`
	Fmt *string `json:"fmt,omitempty"`
	LongFmt *string `json:"longFmt,omitempty"`
}

// NewQuoteResultRegularMarketVolume instantiates a new QuoteResultRegularMarketVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteResultRegularMarketVolume() *QuoteResultRegularMarketVolume {
	this := QuoteResultRegularMarketVolume{}
	return &this
}

// NewQuoteResultRegularMarketVolumeWithDefaults instantiates a new QuoteResultRegularMarketVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteResultRegularMarketVolumeWithDefaults() *QuoteResultRegularMarketVolume {
	this := QuoteResultRegularMarketVolume{}
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *QuoteResultRegularMarketVolume) GetRaw() int64 {
	if o == nil || o.Raw == nil {
		var ret int64
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultRegularMarketVolume) GetRawOk() (*int64, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *QuoteResultRegularMarketVolume) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given int64 and assigns it to the Raw field.
func (o *QuoteResultRegularMarketVolume) SetRaw(v int64) {
	o.Raw = &v
}

// GetFmt returns the Fmt field value if set, zero value otherwise.
func (o *QuoteResultRegularMarketVolume) GetFmt() string {
	if o == nil || o.Fmt == nil {
		var ret string
		return ret
	}
	return *o.Fmt
}

// GetFmtOk returns a tuple with the Fmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultRegularMarketVolume) GetFmtOk() (*string, bool) {
	if o == nil || o.Fmt == nil {
		return nil, false
	}
	return o.Fmt, true
}

// HasFmt returns a boolean if a field has been set.
func (o *QuoteResultRegularMarketVolume) HasFmt() bool {
	if o != nil && o.Fmt != nil {
		return true
	}

	return false
}

// SetFmt gets a reference to the given string and assigns it to the Fmt field.
func (o *QuoteResultRegularMarketVolume) SetFmt(v string) {
	o.Fmt = &v
}

// GetLongFmt returns the LongFmt field value if set, zero value otherwise.
func (o *QuoteResultRegularMarketVolume) GetLongFmt() string {
	if o == nil || o.LongFmt == nil {
		var ret string
		return ret
	}
	return *o.LongFmt
}

// GetLongFmtOk returns a tuple with the LongFmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultRegularMarketVolume) GetLongFmtOk() (*string, bool) {
	if o == nil || o.LongFmt == nil {
		return nil, false
	}
	return o.LongFmt, true
}

// HasLongFmt returns a boolean if a field has been set.
func (o *QuoteResultRegularMarketVolume) HasLongFmt() bool {
	if o != nil && o.LongFmt != nil {
		return true
	}

	return false
}

// SetLongFmt gets a reference to the given string and assigns it to the LongFmt field.
func (o *QuoteResultRegularMarketVolume) SetLongFmt(v string) {
	o.LongFmt = &v
}

func (o QuoteResultRegularMarketVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Fmt != nil {
		toSerialize["fmt"] = o.Fmt
	}
	if o.LongFmt != nil {
		toSerialize["longFmt"] = o.LongFmt
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteResultRegularMarketVolume struct {
	value *QuoteResultRegularMarketVolume
	isSet bool
}

func (v NullableQuoteResultRegularMarketVolume) Get() *QuoteResultRegularMarketVolume {
	return v.value
}

func (v *NullableQuoteResultRegularMarketVolume) Set(val *QuoteResultRegularMarketVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteResultRegularMarketVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteResultRegularMarketVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteResultRegularMarketVolume(val *QuoteResultRegularMarketVolume) *NullableQuoteResultRegularMarketVolume {
	return &NullableQuoteResultRegularMarketVolume{value: val, isSet: true}
}

func (v NullableQuoteResultRegularMarketVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteResultRegularMarketVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


