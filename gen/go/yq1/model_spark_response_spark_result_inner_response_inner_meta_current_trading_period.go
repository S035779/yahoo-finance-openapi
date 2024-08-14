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

// checks if the SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod{}

// SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod struct for SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod
type SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod struct {
	Pre *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre `json:"pre,omitempty"`
	Regular *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre `json:"regular,omitempty"`
	Post *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre `json:"post,omitempty"`
}

// NewSparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod instantiates a new SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod() *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod {
	this := SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod{}
	return &this
}

// NewSparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodWithDefaults instantiates a new SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodWithDefaults() *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod {
	this := SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod{}
	return &this
}

// GetPre returns the Pre field value if set, zero value otherwise.
func (o *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) GetPre() SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre {
	if o == nil || IsNil(o.Pre) {
		var ret SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre
		return ret
	}
	return *o.Pre
}

// GetPreOk returns a tuple with the Pre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) GetPreOk() (*SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre, bool) {
	if o == nil || IsNil(o.Pre) {
		return nil, false
	}
	return o.Pre, true
}

// HasPre returns a boolean if a field has been set.
func (o *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) HasPre() bool {
	if o != nil && !IsNil(o.Pre) {
		return true
	}

	return false
}

// SetPre gets a reference to the given SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre and assigns it to the Pre field.
func (o *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) SetPre(v SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre) {
	o.Pre = &v
}

// GetRegular returns the Regular field value if set, zero value otherwise.
func (o *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) GetRegular() SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre {
	if o == nil || IsNil(o.Regular) {
		var ret SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre
		return ret
	}
	return *o.Regular
}

// GetRegularOk returns a tuple with the Regular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) GetRegularOk() (*SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre, bool) {
	if o == nil || IsNil(o.Regular) {
		return nil, false
	}
	return o.Regular, true
}

// HasRegular returns a boolean if a field has been set.
func (o *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) HasRegular() bool {
	if o != nil && !IsNil(o.Regular) {
		return true
	}

	return false
}

// SetRegular gets a reference to the given SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre and assigns it to the Regular field.
func (o *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) SetRegular(v SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre) {
	o.Regular = &v
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) GetPost() SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre {
	if o == nil || IsNil(o.Post) {
		var ret SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre
		return ret
	}
	return *o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) GetPostOk() (*SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre, bool) {
	if o == nil || IsNil(o.Post) {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) HasPost() bool {
	if o != nil && !IsNil(o.Post) {
		return true
	}

	return false
}

// SetPost gets a reference to the given SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre and assigns it to the Post field.
func (o *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) SetPost(v SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriodPre) {
	o.Post = &v
}

func (o SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) ToMap() (map[string]interface{}, error) {
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

type NullableSparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod struct {
	value *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod
	isSet bool
}

func (v NullableSparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) Get() *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod {
	return v.value
}

func (v *NullableSparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) Set(val *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableSparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableSparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod(val *SparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) *NullableSparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod {
	return &NullableSparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod{value: val, isSet: true}
}

func (v NullableSparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSparkResponseSparkResultInnerResponseInnerMetaCurrentTradingPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


