/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"bytes"
	"encoding/json"
)

// DogAllOf struct for DogAllOf
type DogAllOf struct {
	Breed *string `json:"breed,omitempty"`
}

// NewDogAllOf instantiates a new DogAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDogAllOf() *DogAllOf {
    this := DogAllOf{}
    return &this
}

// NewDogAllOfWithDefaults instantiates a new DogAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDogAllOfWithDefaults() *DogAllOf {
    this := DogAllOf{}
    return &this
}

// GetBreed returns the Breed field value if set, zero value otherwise.
func (o *DogAllOf) GetBreed() string {
	if o == nil || o.Breed == nil {
		var ret string
		return ret
	}
	return *o.Breed
}

// GetBreedOk returns a tuple with the Breed field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DogAllOf) GetBreedOk() (string, bool) {
	if o == nil || o.Breed == nil {
		var ret string
		return ret, false
	}
	return *o.Breed, true
}

// HasBreed returns a boolean if a field has been set.
func (o *DogAllOf) HasBreed() bool {
	if o != nil && o.Breed != nil {
		return true
	}

	return false
}

// SetBreed gets a reference to the given string and assigns it to the Breed field.
func (o *DogAllOf) SetBreed(v string) {
	o.Breed = &v
}

type NullableDogAllOf struct {
	Value DogAllOf
	ExplicitNull bool
}

func (v NullableDogAllOf) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableDogAllOf) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
