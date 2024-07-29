/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
	"fmt"
)

// checks if the Pet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pet{}

// Pet struct for Pet
type Pet struct {
	Id *int64 `json:"id,omitempty"`
	Category *Category `json:"category,omitempty"`
	Name string `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags []Tag `json:"tags,omitempty"`
	// pet status in the store
	// Deprecated
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Pet Pet

// NewPet instantiates a new Pet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPet(name string, photoUrls []string) *Pet {
	this := Pet{}
	this.Name = name
	this.PhotoUrls = photoUrls
	return &this
}

// NewPetWithDefaults instantiates a new Pet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPetWithDefaults() *Pet {
	this := Pet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Pet) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Pet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Pet) SetId(v int64) {
	o.Id = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Pet) GetCategory() Category {
	if o == nil || IsNil(o.Category) {
		var ret Category
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetCategoryOk() (*Category, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Pet) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given Category and assigns it to the Category field.
func (o *Pet) SetCategory(v Category) {
	o.Category = &v
}

// GetName returns the Name field value
func (o *Pet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Pet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Pet) SetName(v string) {
	o.Name = v
}

// GetPhotoUrls returns the PhotoUrls field value
func (o *Pet) GetPhotoUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PhotoUrls
}

// GetPhotoUrlsOk returns a tuple with the PhotoUrls field value
// and a boolean to check if the value has been set.
func (o *Pet) GetPhotoUrlsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhotoUrls, true
}

// SetPhotoUrls sets field value
func (o *Pet) SetPhotoUrls(v []string) {
	o.PhotoUrls = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Pet) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Pet) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *Pet) SetTags(v []Tag) {
	o.Tags = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
// Deprecated
func (o *Pet) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Pet) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Pet) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
// Deprecated
func (o *Pet) SetStatus(v string) {
	o.Status = &v
}

func (o Pet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	toSerialize["name"] = o.Name
	toSerialize["photoUrls"] = o.PhotoUrls
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Pet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"photoUrls",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{} {
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == ""{
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil{
			return err
		}
	}
	varPet := _Pet{}

	err = json.Unmarshal(data, &varPet)

	if err != nil {
		return err
	}

	*o = Pet(varPet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "category")
		delete(additionalProperties, "name")
		delete(additionalProperties, "photoUrls")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePet struct {
	value *Pet
	isSet bool
}

func (v NullablePet) Get() *Pet {
	return v.value
}

func (v *NullablePet) Set(val *Pet) {
	v.value = val
	v.isSet = true
}

func (v NullablePet) IsSet() bool {
	return v.isSet
}

func (v *NullablePet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePet(val *Pet) *NullablePet {
	return &NullablePet{value: val, isSet: true}
}

func (v NullablePet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


