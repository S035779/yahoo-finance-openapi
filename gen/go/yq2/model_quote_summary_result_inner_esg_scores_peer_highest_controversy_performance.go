/*
Yahoo Finance

Yahoo Finance API specification

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yq2

import (
	"encoding/json"
)

// checks if the QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance{}

// QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance struct for QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance
type QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance struct {
	Min *int32 `json:"min,omitempty"`
	Avg *float32 `json:"avg,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

// NewQuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance instantiates a new QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance() *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance {
	this := QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance{}
	return &this
}

// NewQuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformanceWithDefaults instantiates a new QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformanceWithDefaults() *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance {
	this := QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) SetMin(v int32) {
	o.Min = &v
}

// GetAvg returns the Avg field value if set, zero value otherwise.
func (o *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) GetAvg() float32 {
	if o == nil || IsNil(o.Avg) {
		var ret float32
		return ret
	}
	return *o.Avg
}

// GetAvgOk returns a tuple with the Avg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) GetAvgOk() (*float32, bool) {
	if o == nil || IsNil(o.Avg) {
		return nil, false
	}
	return o.Avg, true
}

// HasAvg returns a boolean if a field has been set.
func (o *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) HasAvg() bool {
	if o != nil && !IsNil(o.Avg) {
		return true
	}

	return false
}

// SetAvg gets a reference to the given float32 and assigns it to the Avg field.
func (o *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) SetAvg(v float32) {
	o.Avg = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) SetMax(v int32) {
	o.Max = &v
}

func (o QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Avg) {
		toSerialize["avg"] = o.Avg
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableQuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance struct {
	value *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance
	isSet bool
}

func (v NullableQuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) Get() *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance {
	return v.value
}

func (v *NullableQuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) Set(val *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance(val *QuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) *NullableQuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance {
	return &NullableQuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance{value: val, isSet: true}
}

func (v NullableQuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteSummaryResultInnerEsgScoresPeerHighestControversyPerformance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


