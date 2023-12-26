/*
Apodeixis mailjet-svc

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateNotificationAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNotificationAllOfAttributes{}

// CreateNotificationAllOfAttributes struct for CreateNotificationAllOfAttributes
type CreateNotificationAllOfAttributes struct {
	Message Message `json:"message"`
}

type _CreateNotificationAllOfAttributes CreateNotificationAllOfAttributes

// NewCreateNotificationAllOfAttributes instantiates a new CreateNotificationAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNotificationAllOfAttributes(message Message) *CreateNotificationAllOfAttributes {
	this := CreateNotificationAllOfAttributes{}
	this.Message = message
	return &this
}

// NewCreateNotificationAllOfAttributesWithDefaults instantiates a new CreateNotificationAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNotificationAllOfAttributesWithDefaults() *CreateNotificationAllOfAttributes {
	this := CreateNotificationAllOfAttributes{}
	return &this
}

// GetMessage returns the Message field value
func (o *CreateNotificationAllOfAttributes) GetMessage() Message {
	if o == nil {
		var ret Message
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationAllOfAttributes) GetMessageOk() (*Message, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CreateNotificationAllOfAttributes) SetMessage(v Message) {
	o.Message = v
}

func (o CreateNotificationAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNotificationAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *CreateNotificationAllOfAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateNotificationAllOfAttributes := _CreateNotificationAllOfAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNotificationAllOfAttributes)

	if err != nil {
		return err
	}

	*o = CreateNotificationAllOfAttributes(varCreateNotificationAllOfAttributes)

	return err
}

type NullableCreateNotificationAllOfAttributes struct {
	value *CreateNotificationAllOfAttributes
	isSet bool
}

func (v NullableCreateNotificationAllOfAttributes) Get() *CreateNotificationAllOfAttributes {
	return v.value
}

func (v *NullableCreateNotificationAllOfAttributes) Set(val *CreateNotificationAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNotificationAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNotificationAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNotificationAllOfAttributes(val *CreateNotificationAllOfAttributes) *NullableCreateNotificationAllOfAttributes {
	return &NullableCreateNotificationAllOfAttributes{value: val, isSet: true}
}

func (v NullableCreateNotificationAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNotificationAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
