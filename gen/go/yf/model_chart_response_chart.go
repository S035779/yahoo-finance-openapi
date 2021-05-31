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

// ChartResponseChart struct for ChartResponseChart
type ChartResponseChart struct {
	Result *[]ChartResponseChartResult `json:"result,omitempty"`
	Error *string `json:"error,omitempty"`
}

// NewChartResponseChart instantiates a new ChartResponseChart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartResponseChart() *ChartResponseChart {
	this := ChartResponseChart{}
	return &this
}

// NewChartResponseChartWithDefaults instantiates a new ChartResponseChart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartResponseChartWithDefaults() *ChartResponseChart {
	this := ChartResponseChart{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ChartResponseChart) GetResult() []ChartResponseChartResult {
	if o == nil || o.Result == nil {
		var ret []ChartResponseChartResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChart) GetResultOk() (*[]ChartResponseChartResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ChartResponseChart) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given []ChartResponseChartResult and assigns it to the Result field.
func (o *ChartResponseChart) SetResult(v []ChartResponseChartResult) {
	o.Result = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ChartResponseChart) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChart) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ChartResponseChart) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ChartResponseChart) SetError(v string) {
	o.Error = &v
}

func (o ChartResponseChart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableChartResponseChart struct {
	value *ChartResponseChart
	isSet bool
}

func (v NullableChartResponseChart) Get() *ChartResponseChart {
	return v.value
}

func (v *NullableChartResponseChart) Set(val *ChartResponseChart) {
	v.value = val
	v.isSet = true
}

func (v NullableChartResponseChart) IsSet() bool {
	return v.isSet
}

func (v *NullableChartResponseChart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartResponseChart(val *ChartResponseChart) *NullableChartResponseChart {
	return &NullableChartResponseChart{value: val, isSet: true}
}

func (v NullableChartResponseChart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartResponseChart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


