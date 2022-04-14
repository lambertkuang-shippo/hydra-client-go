/*
 * Ory Oathkeeper API
 *
 * Documentation for all of Ory Oathkeeper's APIs. 
 *
 * API version: v1.11.7
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// TrustedJwtGrantIssuer struct for TrustedJwtGrantIssuer
type TrustedJwtGrantIssuer struct {
	// The \"created_at\" indicates, when grant was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The \"expires_at\" indicates, when grant will expire, so we will reject assertion from \"issuer\" targeting \"subject\".
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	Id *string `json:"id,omitempty"`
	// The \"issuer\" identifies the principal that issued the JWT assertion (same as \"iss\" claim in JWT).
	Issuer *string `json:"issuer,omitempty"`
	PublicKey *TrustedJsonWebKey `json:"public_key,omitempty"`
	// The \"scope\" contains list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749])
	Scope []string `json:"scope,omitempty"`
	// The \"subject\" identifies the principal that is the subject of the JWT.
	Subject *string `json:"subject,omitempty"`
}

// NewTrustedJwtGrantIssuer instantiates a new TrustedJwtGrantIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustedJwtGrantIssuer() *TrustedJwtGrantIssuer {
	this := TrustedJwtGrantIssuer{}
	return &this
}

// NewTrustedJwtGrantIssuerWithDefaults instantiates a new TrustedJwtGrantIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustedJwtGrantIssuerWithDefaults() *TrustedJwtGrantIssuer {
	this := TrustedJwtGrantIssuer{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TrustedJwtGrantIssuer) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedJwtGrantIssuer) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TrustedJwtGrantIssuer) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TrustedJwtGrantIssuer) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *TrustedJwtGrantIssuer) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedJwtGrantIssuer) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *TrustedJwtGrantIssuer) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *TrustedJwtGrantIssuer) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TrustedJwtGrantIssuer) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedJwtGrantIssuer) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TrustedJwtGrantIssuer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TrustedJwtGrantIssuer) SetId(v string) {
	o.Id = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *TrustedJwtGrantIssuer) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedJwtGrantIssuer) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *TrustedJwtGrantIssuer) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *TrustedJwtGrantIssuer) SetIssuer(v string) {
	o.Issuer = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *TrustedJwtGrantIssuer) GetPublicKey() TrustedJsonWebKey {
	if o == nil || o.PublicKey == nil {
		var ret TrustedJsonWebKey
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedJwtGrantIssuer) GetPublicKeyOk() (*TrustedJsonWebKey, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *TrustedJwtGrantIssuer) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given TrustedJsonWebKey and assigns it to the PublicKey field.
func (o *TrustedJwtGrantIssuer) SetPublicKey(v TrustedJsonWebKey) {
	o.PublicKey = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *TrustedJwtGrantIssuer) GetScope() []string {
	if o == nil || o.Scope == nil {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedJwtGrantIssuer) GetScopeOk() ([]string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *TrustedJwtGrantIssuer) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *TrustedJwtGrantIssuer) SetScope(v []string) {
	o.Scope = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TrustedJwtGrantIssuer) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedJwtGrantIssuer) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TrustedJwtGrantIssuer) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TrustedJwtGrantIssuer) SetSubject(v string) {
	o.Subject = &v
}

func (o TrustedJwtGrantIssuer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableTrustedJwtGrantIssuer struct {
	value *TrustedJwtGrantIssuer
	isSet bool
}

func (v NullableTrustedJwtGrantIssuer) Get() *TrustedJwtGrantIssuer {
	return v.value
}

func (v *NullableTrustedJwtGrantIssuer) Set(val *TrustedJwtGrantIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustedJwtGrantIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustedJwtGrantIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustedJwtGrantIssuer(val *TrustedJwtGrantIssuer) *NullableTrustedJwtGrantIssuer {
	return &NullableTrustedJwtGrantIssuer{value: val, isSet: true}
}

func (v NullableTrustedJwtGrantIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustedJwtGrantIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

