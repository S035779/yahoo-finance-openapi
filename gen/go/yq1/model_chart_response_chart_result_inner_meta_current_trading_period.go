/*
Yahoo Finance

Yahoo Finance API specification

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yq1

import (
	"encoding/json"
)

// checks if the ChartResponseChartResultInnerMetaCurrentTradingPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChartResponseChartResultInnerMetaCurrentTradingPeriod{}

// ChartResponseChartResultInnerMetaCurrentTradingPeriod struct for ChartResponseChartResultInnerMetaCurrentTradingPeriod
type ChartResponseChartResultInnerMetaCurrentTradingPeriod struct {
	Pre *ChartResponseChartResultInnerMetaCurrentTradingPeriodPre `json:"pre,omitempty"`
	Regular *ChartResponseChartResultInnerMetaCurrentTradingPeriodRegular `json:"regular,omitempty"`
	Post *ChartResponseChartResultInnerMetaCurrentTradingPeriodPost `json:"post,omitempty"`
}

// NewChartResponseChartResultInnerMetaCurrentTradingPeriod instantiates a new ChartResponseChartResultInnerMetaCurrentTradingPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartResponseChartResultInnerMetaCurrentTradingPeriod() *ChartResponseChartResultInnerMetaCurrentTradingPeriod {
	this := ChartResponseChartResultInnerMetaCurrentTradingPeriod{}
	return &this
}

// NewChartResponseChartResultInnerMetaCurrentTradingPeriodWithDefaults instantiates a new ChartResponseChartResultInnerMetaCurrentTradingPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartResponseChartResultInnerMetaCurrentTradingPeriodWithDefaults() *ChartResponseChartResultInnerMetaCurrentTradingPeriod {
	this := ChartResponseChartResultInnerMetaCurrentTradingPeriod{}
	return &this
}

// GetPre returns the Pre field value if set, zero value otherwise.
func (o *ChartResponseChartResultInnerMetaCurrentTradingPeriod) GetPre() ChartResponseChartResultInnerMetaCurrentTradingPeriodPre {
	if o == nil || IsNil(o.Pre) {
		var ret ChartResponseChartResultInnerMetaCurrentTradingPeriodPre
		return ret
	}
	return *o.Pre
}

// GetPreOk returns a tuple with the Pre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartResultInnerMetaCurrentTradingPeriod) GetPreOk() (*ChartResponseChartResultInnerMetaCurrentTradingPeriodPre, bool) {
	if o == nil || IsNil(o.Pre) {
		return nil, false
	}
	return o.Pre, true
}

// HasPre returns a boolean if a field has been set.
func (o *ChartResponseChartResultInnerMetaCurrentTradingPeriod) HasPre() bool {
	if o != nil && !IsNil(o.Pre) {
		return true
	}

	return false
}

// SetPre gets a reference to the given ChartResponseChartResultInnerMetaCurrentTradingPeriodPre and assigns it to the Pre field.
func (o *ChartResponseChartResultInnerMetaCurrentTradingPeriod) SetPre(v ChartResponseChartResultInnerMetaCurrentTradingPeriodPre) {
	o.Pre = &v
}

// GetRegular returns the Regular field value if set, zero value otherwise.
func (o *ChartResponseChartResultInnerMetaCurrentTradingPeriod) GetRegular() ChartResponseChartResultInnerMetaCurrentTradingPeriodRegular {
	if o == nil || IsNil(o.Regular) {
		var ret ChartResponseChartResultInnerMetaCurrentTradingPeriodRegular
		return ret
	}
	return *o.Regular
}

// GetRegularOk returns a tuple with the Regular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartResultInnerMetaCurrentTradingPeriod) GetRegularOk() (*ChartResponseChartResultInnerMetaCurrentTradingPeriodRegular, bool) {
	if o == nil || IsNil(o.Regular) {
		return nil, false
	}
	return o.Regular, true
}

// HasRegular returns a boolean if a field has been set.
func (o *ChartResponseChartResultInnerMetaCurrentTradingPeriod) HasRegular() bool {
	if o != nil && !IsNil(o.Regular) {
		return true
	}

	return false
}

// SetRegular gets a reference to the given ChartResponseChartResultInnerMetaCurrentTradingPeriodRegular and assigns it to the Regular field.
func (o *ChartResponseChartResultInnerMetaCurrentTradingPeriod) SetRegular(v ChartResponseChartResultInnerMetaCurrentTradingPeriodRegular) {
	o.Regular = &v
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *ChartResponseChartResultInnerMetaCurrentTradingPeriod) GetPost() ChartResponseChartResultInnerMetaCurrentTradingPeriodPost {
	if o == nil || IsNil(o.Post) {
		var ret ChartResponseChartResultInnerMetaCurrentTradingPeriodPost
		return ret
	}
	return *o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartResultInnerMetaCurrentTradingPeriod) GetPostOk() (*ChartResponseChartResultInnerMetaCurrentTradingPeriodPost, bool) {
	if o == nil || IsNil(o.Post) {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *ChartResponseChartResultInnerMetaCurrentTradingPeriod) HasPost() bool {
	if o != nil && !IsNil(o.Post) {
		return true
	}

	return false
}

// SetPost gets a reference to the given ChartResponseChartResultInnerMetaCurrentTradingPeriodPost and assigns it to the Post field.
func (o *ChartResponseChartResultInnerMetaCurrentTradingPeriod) SetPost(v ChartResponseChartResultInnerMetaCurrentTradingPeriodPost) {
	o.Post = &v
}

func (o ChartResponseChartResultInnerMetaCurrentTradingPeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChartResponseChartResultInnerMetaCurrentTradingPeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pre) {
		toSerialize["pre"] = o.Pre
	}
	if !IsNil(o.Regular) {
		toSerialize["regular"] = o.Regular
	}
	if !IsNil(o.Post) {
		toSerialize["post"] = o.Post
	}
	return toSerialize, nil
}

type NullableChartResponseChartResultInnerMetaCurrentTradingPeriod struct {
	value *ChartResponseChartResultInnerMetaCurrentTradingPeriod
	isSet bool
}

func (v NullableChartResponseChartResultInnerMetaCurrentTradingPeriod) Get() *ChartResponseChartResultInnerMetaCurrentTradingPeriod {
	return v.value
}

func (v *NullableChartResponseChartResultInnerMetaCurrentTradingPeriod) Set(val *ChartResponseChartResultInnerMetaCurrentTradingPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableChartResponseChartResultInnerMetaCurrentTradingPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableChartResponseChartResultInnerMetaCurrentTradingPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartResponseChartResultInnerMetaCurrentTradingPeriod(val *ChartResponseChartResultInnerMetaCurrentTradingPeriod) *NullableChartResponseChartResultInnerMetaCurrentTradingPeriod {
	return &NullableChartResponseChartResultInnerMetaCurrentTradingPeriod{value: val, isSet: true}
}

func (v NullableChartResponseChartResultInnerMetaCurrentTradingPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartResponseChartResultInnerMetaCurrentTradingPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


