/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// HTTPMethod the model 'HTTPMethod'
type HTTPMethod string

// List of HTTPMethod
const (
	HTTPMETHOD_GET     HTTPMethod = "GET"
	HTTPMETHOD_POST    HTTPMethod = "POST"
	HTTPMETHOD_PATCH   HTTPMethod = "PATCH"
	HTTPMETHOD_PUT     HTTPMethod = "PUT"
	HTTPMETHOD_DELETE  HTTPMethod = "DELETE"
	HTTPMETHOD_HEAD    HTTPMethod = "HEAD"
	HTTPMETHOD_OPTIONS HTTPMethod = "OPTIONS"
)

// Ptr returns reference to HTTPMethod value
func (v HTTPMethod) Ptr() *HTTPMethod {
	return &v
}

type NullableHTTPMethod struct {
	value *HTTPMethod
	isSet bool
}

func (v NullableHTTPMethod) Get() *HTTPMethod {
	return v.value
}

func (v NullableHTTPMethod) Set(val *HTTPMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableHTTPMethod) IsSet() bool {
	return v.isSet
}

func (v NullableHTTPMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHTTPMethod(val *HTTPMethod) *NullableHTTPMethod {
	return &NullableHTTPMethod{value: val, isSet: true}
}

func (v NullableHTTPMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHTTPMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
