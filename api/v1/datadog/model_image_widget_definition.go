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

// ImageWidgetDefinition The image widget allows you to embed an image on your dashboard. An image can be a PNG, JPG, or animated GIF. Only available on FREE layout dashboards
type ImageWidgetDefinition struct {
	Margin *WidgetMargin      `json:"margin,omitempty"`
	Sizing *WidgetImageSizing `json:"sizing,omitempty"`
	// Type of the widget
	Type string `json:"type"`
	// URL of the image
	Url string `json:"url"`
}

// NewImageWidgetDefinition instantiates a new ImageWidgetDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageWidgetDefinition(type_ string, url string) *ImageWidgetDefinition {
	this := ImageWidgetDefinition{}
	this.Type = type_
	this.Url = url
	return &this
}

// NewImageWidgetDefinitionWithDefaults instantiates a new ImageWidgetDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWidgetDefinitionWithDefaults() *ImageWidgetDefinition {
	this := ImageWidgetDefinition{}
	var type_ string = "image"
	this.Type = type_
	return &this
}

// GetMargin returns the Margin field value if set, zero value otherwise.
func (o *ImageWidgetDefinition) GetMargin() WidgetMargin {
	if o == nil || o.Margin == nil {
		var ret WidgetMargin
		return ret
	}
	return *o.Margin
}

// GetMarginOk returns a tuple with the Margin field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImageWidgetDefinition) GetMarginOk() (WidgetMargin, bool) {
	if o == nil || o.Margin == nil {
		var ret WidgetMargin
		return ret, false
	}
	return *o.Margin, true
}

// HasMargin returns a boolean if a field has been set.
func (o *ImageWidgetDefinition) HasMargin() bool {
	if o != nil && o.Margin != nil {
		return true
	}

	return false
}

// SetMargin gets a reference to the given WidgetMargin and assigns it to the Margin field.
func (o *ImageWidgetDefinition) SetMargin(v WidgetMargin) {
	o.Margin = &v
}

// GetSizing returns the Sizing field value if set, zero value otherwise.
func (o *ImageWidgetDefinition) GetSizing() WidgetImageSizing {
	if o == nil || o.Sizing == nil {
		var ret WidgetImageSizing
		return ret
	}
	return *o.Sizing
}

// GetSizingOk returns a tuple with the Sizing field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImageWidgetDefinition) GetSizingOk() (WidgetImageSizing, bool) {
	if o == nil || o.Sizing == nil {
		var ret WidgetImageSizing
		return ret, false
	}
	return *o.Sizing, true
}

// HasSizing returns a boolean if a field has been set.
func (o *ImageWidgetDefinition) HasSizing() bool {
	if o != nil && o.Sizing != nil {
		return true
	}

	return false
}

// SetSizing gets a reference to the given WidgetImageSizing and assigns it to the Sizing field.
func (o *ImageWidgetDefinition) SetSizing(v WidgetImageSizing) {
	o.Sizing = &v
}

// GetType returns the Type field value
func (o *ImageWidgetDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *ImageWidgetDefinition) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *ImageWidgetDefinition) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// SetUrl sets field value
func (o *ImageWidgetDefinition) SetUrl(v string) {
	o.Url = v
}

func (o ImageWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Margin != nil {
		toSerialize["margin"] = o.Margin
	}
	if o.Sizing != nil {
		toSerialize["sizing"] = o.Sizing
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

// AsWidgetDefinition wraps this instance of ImageWidgetDefinition in WidgetDefinition
func (s *ImageWidgetDefinition) AsWidgetDefinition() WidgetDefinition {
	return WidgetDefinition{WidgetDefinitionInterface: s}
}

type NullableImageWidgetDefinition struct {
	value *ImageWidgetDefinition
	isSet bool
}

func (v NullableImageWidgetDefinition) Get() *ImageWidgetDefinition {
	return v.value
}

func (v NullableImageWidgetDefinition) Set(val *ImageWidgetDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableImageWidgetDefinition) IsSet() bool {
	return v.isSet
}

func (v NullableImageWidgetDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageWidgetDefinition(val *ImageWidgetDefinition) *NullableImageWidgetDefinition {
	return &NullableImageWidgetDefinition{value: val, isSet: true}
}

func (v NullableImageWidgetDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageWidgetDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
