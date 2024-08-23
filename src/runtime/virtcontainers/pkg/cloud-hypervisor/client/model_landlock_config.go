/*
Cloud Hypervisor API

Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LandlockConfig struct for LandlockConfig
type LandlockConfig struct {
	Path   string `json:"path"`
	Access string `json:"access"`
}

// NewLandlockConfig instantiates a new LandlockConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLandlockConfig(path string, access string) *LandlockConfig {
	this := LandlockConfig{}
	this.Path = path
	this.Access = access
	return &this
}

// NewLandlockConfigWithDefaults instantiates a new LandlockConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLandlockConfigWithDefaults() *LandlockConfig {
	this := LandlockConfig{}
	return &this
}

// GetPath returns the Path field value
func (o *LandlockConfig) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *LandlockConfig) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *LandlockConfig) SetPath(v string) {
	o.Path = v
}

// GetAccess returns the Access field value
func (o *LandlockConfig) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *LandlockConfig) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *LandlockConfig) SetAccess(v string) {
	o.Access = v
}

func (o LandlockConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableLandlockConfig struct {
	value *LandlockConfig
	isSet bool
}

func (v NullableLandlockConfig) Get() *LandlockConfig {
	return v.value
}

func (v *NullableLandlockConfig) Set(val *LandlockConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableLandlockConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableLandlockConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLandlockConfig(val *LandlockConfig) *NullableLandlockConfig {
	return &NullableLandlockConfig{value: val, isSet: true}
}

func (v NullableLandlockConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLandlockConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}