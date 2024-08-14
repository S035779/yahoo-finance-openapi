/*
Yahoo Finance

Yahoo Finance API specification

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yq1

import (
	"encoding/json"
	"fmt"
)

// Interval the model 'Interval'
type Interval string

// List of Interval
const (
	INTERVAL__1M Interval = "1m"
	INTERVAL__5M Interval = "5m"
	INTERVAL__15M Interval = "15m"
	INTERVAL__30M Interval = "30m"
	INTERVAL__1H Interval = "1h"
	INTERVAL__6H Interval = "6h"
	INTERVAL__1D Interval = "1d"
)

// All allowed values of Interval enum
var AllowedIntervalEnumValues = []Interval{
	"1m",
	"5m",
	"15m",
	"30m",
	"1h",
	"6h",
	"1d",
}

func (v *Interval) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Interval(value)
	for _, existing := range AllowedIntervalEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Interval", value)
}

// NewIntervalFromValue returns a pointer to a valid Interval
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIntervalFromValue(v string) (*Interval, error) {
	ev := Interval(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Interval: valid values are %v", v, AllowedIntervalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Interval) IsValid() bool {
	for _, existing := range AllowedIntervalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Interval value
func (v Interval) Ptr() *Interval {
	return &v
}

type NullableInterval struct {
	value *Interval
	isSet bool
}

func (v NullableInterval) Get() *Interval {
	return v.value
}

func (v *NullableInterval) Set(val *Interval) {
	v.value = val
	v.isSet = true
}

func (v NullableInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterval(val *Interval) *NullableInterval {
	return &NullableInterval{value: val, isSet: true}
}

func (v NullableInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

