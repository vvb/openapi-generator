/*
Test

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BaseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseItem{}

// BaseItem struct for BaseItem
type BaseItem struct {
	Title string `json:"title"`
	Type string `json:"type"`
}

type _BaseItem BaseItem

// NewBaseItem instantiates a new BaseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseItem(title string, type_ string) *BaseItem {
	this := BaseItem{}
	this.Title = title
	this.Type = type_
	return &this
}

// NewBaseItemWithDefaults instantiates a new BaseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseItemWithDefaults() *BaseItem {
	this := BaseItem{}
	return &this
}

// GetTitle returns the Title field value
func (o *BaseItem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *BaseItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *BaseItem) SetTitle(v string) {
	o.Title = v
}


// GetType returns the Type field value
func (o *BaseItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BaseItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BaseItem) SetType(v string) {
	o.Type = v
}


func (o BaseItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *BaseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBaseItem := _BaseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBaseItem)

	if err != nil {
		return err
	}

	*o = BaseItem(varBaseItem)

	return err
}

type NullableBaseItem struct {
	value *BaseItem
	isSet bool
}

func (v NullableBaseItem) Get() *BaseItem {
	return v.value
}

func (v *NullableBaseItem) Set(val *BaseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseItem(val *BaseItem) *NullableBaseItem {
	return &NullableBaseItem{value: val, isSet: true}
}

func (v NullableBaseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


