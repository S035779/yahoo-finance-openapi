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

// ChartResponseChartResult struct for ChartResponseChartResult
type ChartResponseChartResult struct {
	Meta *ChartResponseChartMeta `json:"meta,omitempty"`
	Timestamp *[]int32 `json:"timestamp,omitempty"`
	Indicators *ChartResponseChartIndicators `json:"indicators,omitempty"`
}

// NewChartResponseChartResult instantiates a new ChartResponseChartResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartResponseChartResult() *ChartResponseChartResult {
	this := ChartResponseChartResult{}
	return &this
}

// NewChartResponseChartResultWithDefaults instantiates a new ChartResponseChartResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartResponseChartResultWithDefaults() *ChartResponseChartResult {
	this := ChartResponseChartResult{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ChartResponseChartResult) GetMeta() ChartResponseChartMeta {
	if o == nil || o.Meta == nil {
		var ret ChartResponseChartMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartResult) GetMetaOk() (*ChartResponseChartMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ChartResponseChartResult) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ChartResponseChartMeta and assigns it to the Meta field.
func (o *ChartResponseChartResult) SetMeta(v ChartResponseChartMeta) {
	o.Meta = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ChartResponseChartResult) GetTimestamp() []int32 {
	if o == nil || o.Timestamp == nil {
		var ret []int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartResult) GetTimestampOk() (*[]int32, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ChartResponseChartResult) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given []int32 and assigns it to the Timestamp field.
func (o *ChartResponseChartResult) SetTimestamp(v []int32) {
	o.Timestamp = &v
}

// GetIndicators returns the Indicators field value if set, zero value otherwise.
func (o *ChartResponseChartResult) GetIndicators() ChartResponseChartIndicators {
	if o == nil || o.Indicators == nil {
		var ret ChartResponseChartIndicators
		return ret
	}
	return *o.Indicators
}

// GetIndicatorsOk returns a tuple with the Indicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartResult) GetIndicatorsOk() (*ChartResponseChartIndicators, bool) {
	if o == nil || o.Indicators == nil {
		return nil, false
	}
	return o.Indicators, true
}

// HasIndicators returns a boolean if a field has been set.
func (o *ChartResponseChartResult) HasIndicators() bool {
	if o != nil && o.Indicators != nil {
		return true
	}

	return false
}

// SetIndicators gets a reference to the given ChartResponseChartIndicators and assigns it to the Indicators field.
func (o *ChartResponseChartResult) SetIndicators(v ChartResponseChartIndicators) {
	o.Indicators = &v
}

func (o ChartResponseChartResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Indicators != nil {
		toSerialize["indicators"] = o.Indicators
	}
	return json.Marshal(toSerialize)
}

type NullableChartResponseChartResult struct {
	value *ChartResponseChartResult
	isSet bool
}

func (v NullableChartResponseChartResult) Get() *ChartResponseChartResult {
	return v.value
}

func (v *NullableChartResponseChartResult) Set(val *ChartResponseChartResult) {
	v.value = val
	v.isSet = true
}

func (v NullableChartResponseChartResult) IsSet() bool {
	return v.isSet
}

func (v *NullableChartResponseChartResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartResponseChartResult(val *ChartResponseChartResult) *NullableChartResponseChartResult {
	return &NullableChartResponseChartResult{value: val, isSet: true}
}

func (v NullableChartResponseChartResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartResponseChartResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


