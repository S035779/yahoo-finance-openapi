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

// QuoteResponseQuoteResponse struct for QuoteResponseQuoteResponse
type QuoteResponseQuoteResponse struct {
	Result *[]QuoteResult `json:"result,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewQuoteResponseQuoteResponse instantiates a new QuoteResponseQuoteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteResponseQuoteResponse() *QuoteResponseQuoteResponse {
	this := QuoteResponseQuoteResponse{}
	return &this
}

// NewQuoteResponseQuoteResponseWithDefaults instantiates a new QuoteResponseQuoteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteResponseQuoteResponseWithDefaults() *QuoteResponseQuoteResponse {
	this := QuoteResponseQuoteResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *QuoteResponseQuoteResponse) GetResult() []QuoteResult {
	if o == nil || o.Result == nil {
		var ret []QuoteResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResponseQuoteResponse) GetResultOk() (*[]QuoteResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *QuoteResponseQuoteResponse) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given []QuoteResult and assigns it to the Result field.
func (o *QuoteResponseQuoteResponse) SetResult(v []QuoteResult) {
	o.Result = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *QuoteResponseQuoteResponse) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResponseQuoteResponse) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *QuoteResponseQuoteResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *QuoteResponseQuoteResponse) SetError(v Error) {
	o.Error = &v
}

func (o QuoteResponseQuoteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteResponseQuoteResponse struct {
	value *QuoteResponseQuoteResponse
	isSet bool
}

func (v NullableQuoteResponseQuoteResponse) Get() *QuoteResponseQuoteResponse {
	return v.value
}

func (v *NullableQuoteResponseQuoteResponse) Set(val *QuoteResponseQuoteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteResponseQuoteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteResponseQuoteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteResponseQuoteResponse(val *QuoteResponseQuoteResponse) *NullableQuoteResponseQuoteResponse {
	return &NullableQuoteResponseQuoteResponse{value: val, isSet: true}
}

func (v NullableQuoteResponseQuoteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteResponseQuoteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


