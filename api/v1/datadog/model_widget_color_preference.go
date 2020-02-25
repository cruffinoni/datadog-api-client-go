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

// WidgetColorPreference Which color to use on the widget. Available values are background or text
type WidgetColorPreference string

// List of WidgetColorPreference
const (
	WIDGETCOLORPREFERENCE_BACKGROUND WidgetColorPreference = "background"
	WIDGETCOLORPREFERENCE_TEXT       WidgetColorPreference = "text"
)

// Ptr returns reference to WidgetColorPreference value
func (v WidgetColorPreference) Ptr() *WidgetColorPreference {
	return &v
}

type NullableWidgetColorPreference struct {
	value *WidgetColorPreference
	isSet bool
}

func (v NullableWidgetColorPreference) Get() *WidgetColorPreference {
	return v.value
}

func (v NullableWidgetColorPreference) Set(val *WidgetColorPreference) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetColorPreference) IsSet() bool {
	return v.isSet
}

func (v NullableWidgetColorPreference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetColorPreference(val *WidgetColorPreference) *NullableWidgetColorPreference {
	return &NullableWidgetColorPreference{value: val, isSet: true}
}

func (v NullableWidgetColorPreference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetColorPreference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
