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

// ChartMetaCurrentTradingPeriodPre struct for ChartMetaCurrentTradingPeriodPre
type ChartMetaCurrentTradingPeriodPre struct {
	Timezone *string `json:"timezone,omitempty"`
	Start *int32 `json:"start,omitempty"`
	End *int32 `json:"end,omitempty"`
	Gmtoffset *int32 `json:"gmtoffset,omitempty"`
}

// NewChartMetaCurrentTradingPeriodPre instantiates a new ChartMetaCurrentTradingPeriodPre object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartMetaCurrentTradingPeriodPre() *ChartMetaCurrentTradingPeriodPre {
	this := ChartMetaCurrentTradingPeriodPre{}
	return &this
}

// NewChartMetaCurrentTradingPeriodPreWithDefaults instantiates a new ChartMetaCurrentTradingPeriodPre object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartMetaCurrentTradingPeriodPreWithDefaults() *ChartMetaCurrentTradingPeriodPre {
	this := ChartMetaCurrentTradingPeriodPre{}
	return &this
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ChartMetaCurrentTradingPeriodPre) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMetaCurrentTradingPeriodPre) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ChartMetaCurrentTradingPeriodPre) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ChartMetaCurrentTradingPeriodPre) SetTimezone(v string) {
	o.Timezone = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ChartMetaCurrentTradingPeriodPre) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMetaCurrentTradingPeriodPre) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ChartMetaCurrentTradingPeriodPre) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ChartMetaCurrentTradingPeriodPre) SetStart(v int32) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ChartMetaCurrentTradingPeriodPre) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMetaCurrentTradingPeriodPre) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ChartMetaCurrentTradingPeriodPre) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ChartMetaCurrentTradingPeriodPre) SetEnd(v int32) {
	o.End = &v
}

// GetGmtoffset returns the Gmtoffset field value if set, zero value otherwise.
func (o *ChartMetaCurrentTradingPeriodPre) GetGmtoffset() int32 {
	if o == nil || o.Gmtoffset == nil {
		var ret int32
		return ret
	}
	return *o.Gmtoffset
}

// GetGmtoffsetOk returns a tuple with the Gmtoffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMetaCurrentTradingPeriodPre) GetGmtoffsetOk() (*int32, bool) {
	if o == nil || o.Gmtoffset == nil {
		return nil, false
	}
	return o.Gmtoffset, true
}

// HasGmtoffset returns a boolean if a field has been set.
func (o *ChartMetaCurrentTradingPeriodPre) HasGmtoffset() bool {
	if o != nil && o.Gmtoffset != nil {
		return true
	}

	return false
}

// SetGmtoffset gets a reference to the given int32 and assigns it to the Gmtoffset field.
func (o *ChartMetaCurrentTradingPeriodPre) SetGmtoffset(v int32) {
	o.Gmtoffset = &v
}

func (o ChartMetaCurrentTradingPeriodPre) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Gmtoffset != nil {
		toSerialize["gmtoffset"] = o.Gmtoffset
	}
	return json.Marshal(toSerialize)
}

type NullableChartMetaCurrentTradingPeriodPre struct {
	value *ChartMetaCurrentTradingPeriodPre
	isSet bool
}

func (v NullableChartMetaCurrentTradingPeriodPre) Get() *ChartMetaCurrentTradingPeriodPre {
	return v.value
}

func (v *NullableChartMetaCurrentTradingPeriodPre) Set(val *ChartMetaCurrentTradingPeriodPre) {
	v.value = val
	v.isSet = true
}

func (v NullableChartMetaCurrentTradingPeriodPre) IsSet() bool {
	return v.isSet
}

func (v *NullableChartMetaCurrentTradingPeriodPre) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartMetaCurrentTradingPeriodPre(val *ChartMetaCurrentTradingPeriodPre) *NullableChartMetaCurrentTradingPeriodPre {
	return &NullableChartMetaCurrentTradingPeriodPre{value: val, isSet: true}
}

func (v NullableChartMetaCurrentTradingPeriodPre) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartMetaCurrentTradingPeriodPre) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


