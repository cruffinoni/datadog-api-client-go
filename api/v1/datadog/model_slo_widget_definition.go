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

// SLOWidgetDefinition Use the SLO and uptime widget to track your SLOs (Service Level Objectives) and uptime on screenboards and timeboards.
type SLOWidgetDefinition struct {
	ShowErrorBudget *bool                `json:"show_error_budget,omitempty"`
	SloId           *string              `json:"slo_id,omitempty"`
	TimeWindows     *[]WidgetTimeWindows `json:"time_windows,omitempty"`
	// Title of the widget
	Title      *string          `json:"title,omitempty"`
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title
	TitleSize *string `json:"title_size,omitempty"`
	// Type of the widget
	Type     string          `json:"type"`
	ViewMode *WidgetViewMode `json:"view_mode,omitempty"`
	ViewType string          `json:"view_type"`
}

// NewSLOWidgetDefinition instantiates a new SLOWidgetDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLOWidgetDefinition(type_ string, viewType string) *SLOWidgetDefinition {
	this := SLOWidgetDefinition{}
	this.Type = type_
	this.ViewType = viewType
	return &this
}

// NewSLOWidgetDefinitionWithDefaults instantiates a new SLOWidgetDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLOWidgetDefinitionWithDefaults() *SLOWidgetDefinition {
	this := SLOWidgetDefinition{}
	var type_ string = "slo"
	this.Type = type_
	var viewType string = "detail"
	this.ViewType = viewType
	return &this
}

// GetShowErrorBudget returns the ShowErrorBudget field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetShowErrorBudget() bool {
	if o == nil || o.ShowErrorBudget == nil {
		var ret bool
		return ret
	}
	return *o.ShowErrorBudget
}

// GetShowErrorBudgetOk returns a tuple with the ShowErrorBudget field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetShowErrorBudgetOk() (bool, bool) {
	if o == nil || o.ShowErrorBudget == nil {
		var ret bool
		return ret, false
	}
	return *o.ShowErrorBudget, true
}

// HasShowErrorBudget returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasShowErrorBudget() bool {
	if o != nil && o.ShowErrorBudget != nil {
		return true
	}

	return false
}

// SetShowErrorBudget gets a reference to the given bool and assigns it to the ShowErrorBudget field.
func (o *SLOWidgetDefinition) SetShowErrorBudget(v bool) {
	o.ShowErrorBudget = &v
}

// GetSloId returns the SloId field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetSloId() string {
	if o == nil || o.SloId == nil {
		var ret string
		return ret
	}
	return *o.SloId
}

// GetSloIdOk returns a tuple with the SloId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetSloIdOk() (string, bool) {
	if o == nil || o.SloId == nil {
		var ret string
		return ret, false
	}
	return *o.SloId, true
}

// HasSloId returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasSloId() bool {
	if o != nil && o.SloId != nil {
		return true
	}

	return false
}

// SetSloId gets a reference to the given string and assigns it to the SloId field.
func (o *SLOWidgetDefinition) SetSloId(v string) {
	o.SloId = &v
}

// GetTimeWindows returns the TimeWindows field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetTimeWindows() []WidgetTimeWindows {
	if o == nil || o.TimeWindows == nil {
		var ret []WidgetTimeWindows
		return ret
	}
	return *o.TimeWindows
}

// GetTimeWindowsOk returns a tuple with the TimeWindows field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetTimeWindowsOk() ([]WidgetTimeWindows, bool) {
	if o == nil || o.TimeWindows == nil {
		var ret []WidgetTimeWindows
		return ret, false
	}
	return *o.TimeWindows, true
}

// HasTimeWindows returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasTimeWindows() bool {
	if o != nil && o.TimeWindows != nil {
		return true
	}

	return false
}

// SetTimeWindows gets a reference to the given []WidgetTimeWindows and assigns it to the TimeWindows field.
func (o *SLOWidgetDefinition) SetTimeWindows(v []WidgetTimeWindows) {
	o.TimeWindows = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetTitleOk() (string, bool) {
	if o == nil || o.Title == nil {
		var ret string
		return ret, false
	}
	return *o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SLOWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetTitleAlignOk() (WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret, false
	}
	return *o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasTitleAlign() bool {
	if o != nil && o.TitleAlign != nil {
		return true
	}

	return false
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *SLOWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetTitleSizeOk() (string, bool) {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret, false
	}
	return *o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasTitleSize() bool {
	if o != nil && o.TitleSize != nil {
		return true
	}

	return false
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *SLOWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value
func (o *SLOWidgetDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *SLOWidgetDefinition) SetType(v string) {
	o.Type = v
}

// GetViewMode returns the ViewMode field value if set, zero value otherwise.
func (o *SLOWidgetDefinition) GetViewMode() WidgetViewMode {
	if o == nil || o.ViewMode == nil {
		var ret WidgetViewMode
		return ret
	}
	return *o.ViewMode
}

// GetViewModeOk returns a tuple with the ViewMode field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SLOWidgetDefinition) GetViewModeOk() (WidgetViewMode, bool) {
	if o == nil || o.ViewMode == nil {
		var ret WidgetViewMode
		return ret, false
	}
	return *o.ViewMode, true
}

// HasViewMode returns a boolean if a field has been set.
func (o *SLOWidgetDefinition) HasViewMode() bool {
	if o != nil && o.ViewMode != nil {
		return true
	}

	return false
}

// SetViewMode gets a reference to the given WidgetViewMode and assigns it to the ViewMode field.
func (o *SLOWidgetDefinition) SetViewMode(v WidgetViewMode) {
	o.ViewMode = &v
}

// GetViewType returns the ViewType field value
func (o *SLOWidgetDefinition) GetViewType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ViewType
}

// SetViewType sets field value
func (o *SLOWidgetDefinition) SetViewType(v string) {
	o.ViewType = v
}

func (o SLOWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ShowErrorBudget != nil {
		toSerialize["show_error_budget"] = o.ShowErrorBudget
	}
	if o.SloId != nil {
		toSerialize["slo_id"] = o.SloId
	}
	if o.TimeWindows != nil {
		toSerialize["time_windows"] = o.TimeWindows
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.TitleAlign != nil {
		toSerialize["title_align"] = o.TitleAlign
	}
	if o.TitleSize != nil {
		toSerialize["title_size"] = o.TitleSize
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.ViewMode != nil {
		toSerialize["view_mode"] = o.ViewMode
	}
	if true {
		toSerialize["view_type"] = o.ViewType
	}
	return json.Marshal(toSerialize)
}

// AsWidgetDefinition wraps this instance of SLOWidgetDefinition in WidgetDefinition
func (s *SLOWidgetDefinition) AsWidgetDefinition() WidgetDefinition {
	return WidgetDefinition{WidgetDefinitionInterface: s}
}

type NullableSLOWidgetDefinition struct {
	value *SLOWidgetDefinition
	isSet bool
}

func (v NullableSLOWidgetDefinition) Get() *SLOWidgetDefinition {
	return v.value
}

func (v NullableSLOWidgetDefinition) Set(val *SLOWidgetDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableSLOWidgetDefinition) IsSet() bool {
	return v.isSet
}

func (v NullableSLOWidgetDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLOWidgetDefinition(val *SLOWidgetDefinition) *NullableSLOWidgetDefinition {
	return &NullableSLOWidgetDefinition{value: val, isSet: true}
}

func (v NullableSLOWidgetDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLOWidgetDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
