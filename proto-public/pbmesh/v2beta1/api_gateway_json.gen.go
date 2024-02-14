// Code generated by protoc-json-shim. DO NOT EDIT.
package meshv2beta1

import (
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON is a custom marshaler for APIGateway
func (this *APIGateway) MarshalJSON() ([]byte, error) {
	str, err := ApiGatewayMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for APIGateway
func (this *APIGateway) UnmarshalJSON(b []byte) error {
	return ApiGatewayUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for APIGatewayListener
func (this *APIGatewayListener) MarshalJSON() ([]byte, error) {
	str, err := ApiGatewayMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for APIGatewayListener
func (this *APIGatewayListener) UnmarshalJSON(b []byte) error {
	return ApiGatewayUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for APIGatewayTLSConfiguration
func (this *APIGatewayTLSConfiguration) MarshalJSON() ([]byte, error) {
	str, err := ApiGatewayMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for APIGatewayTLSConfiguration
func (this *APIGatewayTLSConfiguration) UnmarshalJSON(b []byte) error {
	return ApiGatewayUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for InlineCertificate
func (this *InlineCertificate) MarshalJSON() ([]byte, error) {
	str, err := ApiGatewayMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for InlineCertificate
func (this *InlineCertificate) UnmarshalJSON(b []byte) error {
	return ApiGatewayUnmarshaler.Unmarshal(b, this)
}

var (
	ApiGatewayMarshaler   = &protojson.MarshalOptions{}
	ApiGatewayUnmarshaler = &protojson.UnmarshalOptions{DiscardUnknown: false}
)
