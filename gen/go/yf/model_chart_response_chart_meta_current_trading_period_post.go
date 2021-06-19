/*
 * Yahoo Finance
 *
 * Yahoo Finance API specification
 *
 * API version: 1.0.6
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yf

import (
	"encoding/json"
)

// ChartResponseChartMetaCurrentTradingPeriodPost struct for ChartResponseChartMetaCurrentTradingPeriodPost
type ChartResponseChartMetaCurrentTradingPeriodPost struct {
	Timezone *string `json:"timezone,omitempty"`
	Start *int32 `json:"start,omitempty"`
	End *int32 `json:"end,omitempty"`
	Gmtoffset *int32 `json:"gmtoffset,omitempty"`
}

// NewChartResponseChartMetaCurrentTradingPeriodPost instantiates a new ChartResponseChartMetaCurrentTradingPeriodPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartResponseChartMetaCurrentTradingPeriodPost() *ChartResponseChartMetaCurrentTradingPeriodPost {
	this := ChartResponseChartMetaCurrentTradingPeriodPost{}
	return &this
}

// NewChartResponseChartMetaCurrentTradingPeriodPostWithDefaults instantiates a new ChartResponseChartMetaCurrentTradingPeriodPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartResponseChartMetaCurrentTradingPeriodPostWithDefaults() *ChartResponseChartMetaCurrentTradingPeriodPost {
	this := ChartResponseChartMetaCurrentTradingPeriodPost{}
	return &this
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) SetTimezone(v string) {
	o.Timezone = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) SetStart(v int32) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) SetEnd(v int32) {
	o.End = &v
}

// GetGmtoffset returns the Gmtoffset field value if set, zero value otherwise.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) GetGmtoffset() int32 {
	if o == nil || o.Gmtoffset == nil {
		var ret int32
		return ret
	}
	return *o.Gmtoffset
}

// GetGmtoffsetOk returns a tuple with the Gmtoffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) GetGmtoffsetOk() (*int32, bool) {
	if o == nil || o.Gmtoffset == nil {
		return nil, false
	}
	return o.Gmtoffset, true
}

// HasGmtoffset returns a boolean if a field has been set.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) HasGmtoffset() bool {
	if o != nil && o.Gmtoffset != nil {
		return true
	}

	return false
}

// SetGmtoffset gets a reference to the given int32 and assigns it to the Gmtoffset field.
func (o *ChartResponseChartMetaCurrentTradingPeriodPost) SetGmtoffset(v int32) {
	o.Gmtoffset = &v
}

func (o ChartResponseChartMetaCurrentTradingPeriodPost) MarshalJSON() ([]byte, error) {
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

type NullableChartResponseChartMetaCurrentTradingPeriodPost struct {
	value *ChartResponseChartMetaCurrentTradingPeriodPost
	isSet bool
}

func (v NullableChartResponseChartMetaCurrentTradingPeriodPost) Get() *ChartResponseChartMetaCurrentTradingPeriodPost {
	return v.value
}

func (v *NullableChartResponseChartMetaCurrentTradingPeriodPost) Set(val *ChartResponseChartMetaCurrentTradingPeriodPost) {
	v.value = val
	v.isSet = true
}

func (v NullableChartResponseChartMetaCurrentTradingPeriodPost) IsSet() bool {
	return v.isSet
}

func (v *NullableChartResponseChartMetaCurrentTradingPeriodPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartResponseChartMetaCurrentTradingPeriodPost(val *ChartResponseChartMetaCurrentTradingPeriodPost) *NullableChartResponseChartMetaCurrentTradingPeriodPost {
	return &NullableChartResponseChartMetaCurrentTradingPeriodPost{value: val, isSet: true}
}

func (v NullableChartResponseChartMetaCurrentTradingPeriodPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartResponseChartMetaCurrentTradingPeriodPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


