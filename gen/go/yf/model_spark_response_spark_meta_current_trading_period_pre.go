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

// SparkResponseSparkMetaCurrentTradingPeriodPre struct for SparkResponseSparkMetaCurrentTradingPeriodPre
type SparkResponseSparkMetaCurrentTradingPeriodPre struct {
	Timezone *string `json:"timezone,omitempty"`
	End *int32 `json:"end,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Gmtoffset *int32 `json:"gmtoffset,omitempty"`
}

// NewSparkResponseSparkMetaCurrentTradingPeriodPre instantiates a new SparkResponseSparkMetaCurrentTradingPeriodPre object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSparkResponseSparkMetaCurrentTradingPeriodPre() *SparkResponseSparkMetaCurrentTradingPeriodPre {
	this := SparkResponseSparkMetaCurrentTradingPeriodPre{}
	return &this
}

// NewSparkResponseSparkMetaCurrentTradingPeriodPreWithDefaults instantiates a new SparkResponseSparkMetaCurrentTradingPeriodPre object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSparkResponseSparkMetaCurrentTradingPeriodPreWithDefaults() *SparkResponseSparkMetaCurrentTradingPeriodPre {
	this := SparkResponseSparkMetaCurrentTradingPeriodPre{}
	return &this
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) SetTimezone(v string) {
	o.Timezone = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) SetEnd(v int32) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) SetStart(v int32) {
	o.Start = &v
}

// GetGmtoffset returns the Gmtoffset field value if set, zero value otherwise.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) GetGmtoffset() int32 {
	if o == nil || o.Gmtoffset == nil {
		var ret int32
		return ret
	}
	return *o.Gmtoffset
}

// GetGmtoffsetOk returns a tuple with the Gmtoffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) GetGmtoffsetOk() (*int32, bool) {
	if o == nil || o.Gmtoffset == nil {
		return nil, false
	}
	return o.Gmtoffset, true
}

// HasGmtoffset returns a boolean if a field has been set.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) HasGmtoffset() bool {
	if o != nil && o.Gmtoffset != nil {
		return true
	}

	return false
}

// SetGmtoffset gets a reference to the given int32 and assigns it to the Gmtoffset field.
func (o *SparkResponseSparkMetaCurrentTradingPeriodPre) SetGmtoffset(v int32) {
	o.Gmtoffset = &v
}

func (o SparkResponseSparkMetaCurrentTradingPeriodPre) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.Gmtoffset != nil {
		toSerialize["gmtoffset"] = o.Gmtoffset
	}
	return json.Marshal(toSerialize)
}

type NullableSparkResponseSparkMetaCurrentTradingPeriodPre struct {
	value *SparkResponseSparkMetaCurrentTradingPeriodPre
	isSet bool
}

func (v NullableSparkResponseSparkMetaCurrentTradingPeriodPre) Get() *SparkResponseSparkMetaCurrentTradingPeriodPre {
	return v.value
}

func (v *NullableSparkResponseSparkMetaCurrentTradingPeriodPre) Set(val *SparkResponseSparkMetaCurrentTradingPeriodPre) {
	v.value = val
	v.isSet = true
}

func (v NullableSparkResponseSparkMetaCurrentTradingPeriodPre) IsSet() bool {
	return v.isSet
}

func (v *NullableSparkResponseSparkMetaCurrentTradingPeriodPre) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSparkResponseSparkMetaCurrentTradingPeriodPre(val *SparkResponseSparkMetaCurrentTradingPeriodPre) *NullableSparkResponseSparkMetaCurrentTradingPeriodPre {
	return &NullableSparkResponseSparkMetaCurrentTradingPeriodPre{value: val, isSet: true}
}

func (v NullableSparkResponseSparkMetaCurrentTradingPeriodPre) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSparkResponseSparkMetaCurrentTradingPeriodPre) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


