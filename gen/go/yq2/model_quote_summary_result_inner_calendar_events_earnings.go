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

// checks if the QuoteSummaryResultInnerCalendarEventsEarnings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteSummaryResultInnerCalendarEventsEarnings{}

// QuoteSummaryResultInnerCalendarEventsEarnings struct for QuoteSummaryResultInnerCalendarEventsEarnings
type QuoteSummaryResultInnerCalendarEventsEarnings struct {
	EarningsDate []QuoteSummaryResultInnerDefaultKeyStatisticsSharesShortPreviousMonthDate `json:"earningsDate,omitempty"`
	EarningsAverage *QuoteSummaryResultInnerEsgScoresTotalEsg `json:"earningsAverage,omitempty"`
	EarningsLow *QuoteSummaryResultInnerEsgScoresTotalEsg `json:"earningsLow,omitempty"`
	EarningsHigh *QuoteSummaryResultInnerEsgScoresTotalEsg `json:"earningsHigh,omitempty"`
	RevenueAverage *QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue `json:"revenueAverage,omitempty"`
	RevenueLow *QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue `json:"revenueLow,omitempty"`
	RevenueHigh *QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue `json:"revenueHigh,omitempty"`
}

// NewQuoteSummaryResultInnerCalendarEventsEarnings instantiates a new QuoteSummaryResultInnerCalendarEventsEarnings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteSummaryResultInnerCalendarEventsEarnings() *QuoteSummaryResultInnerCalendarEventsEarnings {
	this := QuoteSummaryResultInnerCalendarEventsEarnings{}
	return &this
}

// NewQuoteSummaryResultInnerCalendarEventsEarningsWithDefaults instantiates a new QuoteSummaryResultInnerCalendarEventsEarnings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteSummaryResultInnerCalendarEventsEarningsWithDefaults() *QuoteSummaryResultInnerCalendarEventsEarnings {
	this := QuoteSummaryResultInnerCalendarEventsEarnings{}
	return &this
}

// GetEarningsDate returns the EarningsDate field value if set, zero value otherwise.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) GetEarningsDate() []QuoteSummaryResultInnerDefaultKeyStatisticsSharesShortPreviousMonthDate {
	if o == nil || IsNil(o.EarningsDate) {
		var ret []QuoteSummaryResultInnerDefaultKeyStatisticsSharesShortPreviousMonthDate
		return ret
	}
	return o.EarningsDate
}

// GetEarningsDateOk returns a tuple with the EarningsDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) GetEarningsDateOk() ([]QuoteSummaryResultInnerDefaultKeyStatisticsSharesShortPreviousMonthDate, bool) {
	if o == nil || IsNil(o.EarningsDate) {
		return nil, false
	}
	return o.EarningsDate, true
}

// HasEarningsDate returns a boolean if a field has been set.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) HasEarningsDate() bool {
	if o != nil && !IsNil(o.EarningsDate) {
		return true
	}

	return false
}

// SetEarningsDate gets a reference to the given []QuoteSummaryResultInnerDefaultKeyStatisticsSharesShortPreviousMonthDate and assigns it to the EarningsDate field.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) SetEarningsDate(v []QuoteSummaryResultInnerDefaultKeyStatisticsSharesShortPreviousMonthDate) {
	o.EarningsDate = v
}

// GetEarningsAverage returns the EarningsAverage field value if set, zero value otherwise.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) GetEarningsAverage() QuoteSummaryResultInnerEsgScoresTotalEsg {
	if o == nil || IsNil(o.EarningsAverage) {
		var ret QuoteSummaryResultInnerEsgScoresTotalEsg
		return ret
	}
	return *o.EarningsAverage
}

// GetEarningsAverageOk returns a tuple with the EarningsAverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) GetEarningsAverageOk() (*QuoteSummaryResultInnerEsgScoresTotalEsg, bool) {
	if o == nil || IsNil(o.EarningsAverage) {
		return nil, false
	}
	return o.EarningsAverage, true
}

// HasEarningsAverage returns a boolean if a field has been set.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) HasEarningsAverage() bool {
	if o != nil && !IsNil(o.EarningsAverage) {
		return true
	}

	return false
}

// SetEarningsAverage gets a reference to the given QuoteSummaryResultInnerEsgScoresTotalEsg and assigns it to the EarningsAverage field.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) SetEarningsAverage(v QuoteSummaryResultInnerEsgScoresTotalEsg) {
	o.EarningsAverage = &v
}

// GetEarningsLow returns the EarningsLow field value if set, zero value otherwise.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) GetEarningsLow() QuoteSummaryResultInnerEsgScoresTotalEsg {
	if o == nil || IsNil(o.EarningsLow) {
		var ret QuoteSummaryResultInnerEsgScoresTotalEsg
		return ret
	}
	return *o.EarningsLow
}

// GetEarningsLowOk returns a tuple with the EarningsLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) GetEarningsLowOk() (*QuoteSummaryResultInnerEsgScoresTotalEsg, bool) {
	if o == nil || IsNil(o.EarningsLow) {
		return nil, false
	}
	return o.EarningsLow, true
}

// HasEarningsLow returns a boolean if a field has been set.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) HasEarningsLow() bool {
	if o != nil && !IsNil(o.EarningsLow) {
		return true
	}

	return false
}

// SetEarningsLow gets a reference to the given QuoteSummaryResultInnerEsgScoresTotalEsg and assigns it to the EarningsLow field.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) SetEarningsLow(v QuoteSummaryResultInnerEsgScoresTotalEsg) {
	o.EarningsLow = &v
}

// GetEarningsHigh returns the EarningsHigh field value if set, zero value otherwise.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) GetEarningsHigh() QuoteSummaryResultInnerEsgScoresTotalEsg {
	if o == nil || IsNil(o.EarningsHigh) {
		var ret QuoteSummaryResultInnerEsgScoresTotalEsg
		return ret
	}
	return *o.EarningsHigh
}

// GetEarningsHighOk returns a tuple with the EarningsHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) GetEarningsHighOk() (*QuoteSummaryResultInnerEsgScoresTotalEsg, bool) {
	if o == nil || IsNil(o.EarningsHigh) {
		return nil, false
	}
	return o.EarningsHigh, true
}

// HasEarningsHigh returns a boolean if a field has been set.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) HasEarningsHigh() bool {
	if o != nil && !IsNil(o.EarningsHigh) {
		return true
	}

	return false
}

// SetEarningsHigh gets a reference to the given QuoteSummaryResultInnerEsgScoresTotalEsg and assigns it to the EarningsHigh field.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) SetEarningsHigh(v QuoteSummaryResultInnerEsgScoresTotalEsg) {
	o.EarningsHigh = &v
}

// GetRevenueAverage returns the RevenueAverage field value if set, zero value otherwise.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) GetRevenueAverage() QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue {
	if o == nil || IsNil(o.RevenueAverage) {
		var ret QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue
		return ret
	}
	return *o.RevenueAverage
}

// GetRevenueAverageOk returns a tuple with the RevenueAverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) GetRevenueAverageOk() (*QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue, bool) {
	if o == nil || IsNil(o.RevenueAverage) {
		return nil, false
	}
	return o.RevenueAverage, true
}

// HasRevenueAverage returns a boolean if a field has been set.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) HasRevenueAverage() bool {
	if o != nil && !IsNil(o.RevenueAverage) {
		return true
	}

	return false
}

// SetRevenueAverage gets a reference to the given QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue and assigns it to the RevenueAverage field.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) SetRevenueAverage(v QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue) {
	o.RevenueAverage = &v
}

// GetRevenueLow returns the RevenueLow field value if set, zero value otherwise.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) GetRevenueLow() QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue {
	if o == nil || IsNil(o.RevenueLow) {
		var ret QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue
		return ret
	}
	return *o.RevenueLow
}

// GetRevenueLowOk returns a tuple with the RevenueLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) GetRevenueLowOk() (*QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue, bool) {
	if o == nil || IsNil(o.RevenueLow) {
		return nil, false
	}
	return o.RevenueLow, true
}

// HasRevenueLow returns a boolean if a field has been set.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) HasRevenueLow() bool {
	if o != nil && !IsNil(o.RevenueLow) {
		return true
	}

	return false
}

// SetRevenueLow gets a reference to the given QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue and assigns it to the RevenueLow field.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) SetRevenueLow(v QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue) {
	o.RevenueLow = &v
}

// GetRevenueHigh returns the RevenueHigh field value if set, zero value otherwise.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) GetRevenueHigh() QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue {
	if o == nil || IsNil(o.RevenueHigh) {
		var ret QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue
		return ret
	}
	return *o.RevenueHigh
}

// GetRevenueHighOk returns a tuple with the RevenueHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) GetRevenueHighOk() (*QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue, bool) {
	if o == nil || IsNil(o.RevenueHigh) {
		return nil, false
	}
	return o.RevenueHigh, true
}

// HasRevenueHigh returns a boolean if a field has been set.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) HasRevenueHigh() bool {
	if o != nil && !IsNil(o.RevenueHigh) {
		return true
	}

	return false
}

// SetRevenueHigh gets a reference to the given QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue and assigns it to the RevenueHigh field.
func (o *QuoteSummaryResultInnerCalendarEventsEarnings) SetRevenueHigh(v QuoteSummaryResultInnerDefaultKeyStatisticsEnterpriseValue) {
	o.RevenueHigh = &v
}

func (o QuoteSummaryResultInnerCalendarEventsEarnings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteSummaryResultInnerCalendarEventsEarnings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EarningsDate) {
		toSerialize["earningsDate"] = o.EarningsDate
	}
	if !IsNil(o.EarningsAverage) {
		toSerialize["earningsAverage"] = o.EarningsAverage
	}
	if !IsNil(o.EarningsLow) {
		toSerialize["earningsLow"] = o.EarningsLow
	}
	if !IsNil(o.EarningsHigh) {
		toSerialize["earningsHigh"] = o.EarningsHigh
	}
	if !IsNil(o.RevenueAverage) {
		toSerialize["revenueAverage"] = o.RevenueAverage
	}
	if !IsNil(o.RevenueLow) {
		toSerialize["revenueLow"] = o.RevenueLow
	}
	if !IsNil(o.RevenueHigh) {
		toSerialize["revenueHigh"] = o.RevenueHigh
	}
	return toSerialize, nil
}

type NullableQuoteSummaryResultInnerCalendarEventsEarnings struct {
	value *QuoteSummaryResultInnerCalendarEventsEarnings
	isSet bool
}

func (v NullableQuoteSummaryResultInnerCalendarEventsEarnings) Get() *QuoteSummaryResultInnerCalendarEventsEarnings {
	return v.value
}

func (v *NullableQuoteSummaryResultInnerCalendarEventsEarnings) Set(val *QuoteSummaryResultInnerCalendarEventsEarnings) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteSummaryResultInnerCalendarEventsEarnings) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteSummaryResultInnerCalendarEventsEarnings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteSummaryResultInnerCalendarEventsEarnings(val *QuoteSummaryResultInnerCalendarEventsEarnings) *NullableQuoteSummaryResultInnerCalendarEventsEarnings {
	return &NullableQuoteSummaryResultInnerCalendarEventsEarnings{value: val, isSet: true}
}

func (v NullableQuoteSummaryResultInnerCalendarEventsEarnings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteSummaryResultInnerCalendarEventsEarnings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


