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

// ChartResponseChartIndicatorsQuote struct for ChartResponseChartIndicatorsQuote
type ChartResponseChartIndicatorsQuote struct {
	High *[]float32 `json:"high,omitempty"`
	Close *[]float32 `json:"close,omitempty"`
	Volume *[]int32 `json:"volume,omitempty"`
	Low *[]float32 `json:"low,omitempty"`
	Open *[]float32 `json:"open,omitempty"`
}

// NewChartResponseChartIndicatorsQuote instantiates a new ChartResponseChartIndicatorsQuote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartResponseChartIndicatorsQuote() *ChartResponseChartIndicatorsQuote {
	this := ChartResponseChartIndicatorsQuote{}
	return &this
}

// NewChartResponseChartIndicatorsQuoteWithDefaults instantiates a new ChartResponseChartIndicatorsQuote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartResponseChartIndicatorsQuoteWithDefaults() *ChartResponseChartIndicatorsQuote {
	this := ChartResponseChartIndicatorsQuote{}
	return &this
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *ChartResponseChartIndicatorsQuote) GetHigh() []float32 {
	if o == nil || o.High == nil {
		var ret []float32
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartIndicatorsQuote) GetHighOk() (*[]float32, bool) {
	if o == nil || o.High == nil {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *ChartResponseChartIndicatorsQuote) HasHigh() bool {
	if o != nil && o.High != nil {
		return true
	}

	return false
}

// SetHigh gets a reference to the given []float32 and assigns it to the High field.
func (o *ChartResponseChartIndicatorsQuote) SetHigh(v []float32) {
	o.High = &v
}

// GetClose returns the Close field value if set, zero value otherwise.
func (o *ChartResponseChartIndicatorsQuote) GetClose() []float32 {
	if o == nil || o.Close == nil {
		var ret []float32
		return ret
	}
	return *o.Close
}

// GetCloseOk returns a tuple with the Close field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartIndicatorsQuote) GetCloseOk() (*[]float32, bool) {
	if o == nil || o.Close == nil {
		return nil, false
	}
	return o.Close, true
}

// HasClose returns a boolean if a field has been set.
func (o *ChartResponseChartIndicatorsQuote) HasClose() bool {
	if o != nil && o.Close != nil {
		return true
	}

	return false
}

// SetClose gets a reference to the given []float32 and assigns it to the Close field.
func (o *ChartResponseChartIndicatorsQuote) SetClose(v []float32) {
	o.Close = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *ChartResponseChartIndicatorsQuote) GetVolume() []int32 {
	if o == nil || o.Volume == nil {
		var ret []int32
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartIndicatorsQuote) GetVolumeOk() (*[]int32, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *ChartResponseChartIndicatorsQuote) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given []int32 and assigns it to the Volume field.
func (o *ChartResponseChartIndicatorsQuote) SetVolume(v []int32) {
	o.Volume = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *ChartResponseChartIndicatorsQuote) GetLow() []float32 {
	if o == nil || o.Low == nil {
		var ret []float32
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartIndicatorsQuote) GetLowOk() (*[]float32, bool) {
	if o == nil || o.Low == nil {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *ChartResponseChartIndicatorsQuote) HasLow() bool {
	if o != nil && o.Low != nil {
		return true
	}

	return false
}

// SetLow gets a reference to the given []float32 and assigns it to the Low field.
func (o *ChartResponseChartIndicatorsQuote) SetLow(v []float32) {
	o.Low = &v
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *ChartResponseChartIndicatorsQuote) GetOpen() []float32 {
	if o == nil || o.Open == nil {
		var ret []float32
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartIndicatorsQuote) GetOpenOk() (*[]float32, bool) {
	if o == nil || o.Open == nil {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *ChartResponseChartIndicatorsQuote) HasOpen() bool {
	if o != nil && o.Open != nil {
		return true
	}

	return false
}

// SetOpen gets a reference to the given []float32 and assigns it to the Open field.
func (o *ChartResponseChartIndicatorsQuote) SetOpen(v []float32) {
	o.Open = &v
}

func (o ChartResponseChartIndicatorsQuote) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.High != nil {
		toSerialize["high"] = o.High
	}
	if o.Close != nil {
		toSerialize["close"] = o.Close
	}
	if o.Volume != nil {
		toSerialize["volume"] = o.Volume
	}
	if o.Low != nil {
		toSerialize["low"] = o.Low
	}
	if o.Open != nil {
		toSerialize["open"] = o.Open
	}
	return json.Marshal(toSerialize)
}

type NullableChartResponseChartIndicatorsQuote struct {
	value *ChartResponseChartIndicatorsQuote
	isSet bool
}

func (v NullableChartResponseChartIndicatorsQuote) Get() *ChartResponseChartIndicatorsQuote {
	return v.value
}

func (v *NullableChartResponseChartIndicatorsQuote) Set(val *ChartResponseChartIndicatorsQuote) {
	v.value = val
	v.isSet = true
}

func (v NullableChartResponseChartIndicatorsQuote) IsSet() bool {
	return v.isSet
}

func (v *NullableChartResponseChartIndicatorsQuote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartResponseChartIndicatorsQuote(val *ChartResponseChartIndicatorsQuote) *NullableChartResponseChartIndicatorsQuote {
	return &NullableChartResponseChartIndicatorsQuote{value: val, isSet: true}
}

func (v NullableChartResponseChartIndicatorsQuote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartResponseChartIndicatorsQuote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


