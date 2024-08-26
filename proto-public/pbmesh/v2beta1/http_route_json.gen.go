// Code generated by protoc-json-shim. DO NOT EDIT.
package meshv2beta1

import (
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON is a custom marshaler for HTTPRoute
func (this *HTTPRoute) MarshalJSON() ([]byte, error) {
	str, err := HttpRouteMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPRoute
func (this *HTTPRoute) UnmarshalJSON(b []byte) error {
	return HttpRouteUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for HTTPRouteRule
func (this *HTTPRouteRule) MarshalJSON() ([]byte, error) {
	str, err := HttpRouteMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPRouteRule
func (this *HTTPRouteRule) UnmarshalJSON(b []byte) error {
	return HttpRouteUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for HTTPRouteMatch
func (this *HTTPRouteMatch) MarshalJSON() ([]byte, error) {
	str, err := HttpRouteMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPRouteMatch
func (this *HTTPRouteMatch) UnmarshalJSON(b []byte) error {
	return HttpRouteUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for HTTPPathMatch
func (this *HTTPPathMatch) MarshalJSON() ([]byte, error) {
	str, err := HttpRouteMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPPathMatch
func (this *HTTPPathMatch) UnmarshalJSON(b []byte) error {
	return HttpRouteUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for HTTPHeaderMatch
func (this *HTTPHeaderMatch) MarshalJSON() ([]byte, error) {
	str, err := HttpRouteMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPHeaderMatch
func (this *HTTPHeaderMatch) UnmarshalJSON(b []byte) error {
	return HttpRouteUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for HTTPQueryParamMatch
func (this *HTTPQueryParamMatch) MarshalJSON() ([]byte, error) {
	str, err := HttpRouteMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPQueryParamMatch
func (this *HTTPQueryParamMatch) UnmarshalJSON(b []byte) error {
	return HttpRouteUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for HTTPRouteFilter
func (this *HTTPRouteFilter) MarshalJSON() ([]byte, error) {
	str, err := HttpRouteMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPRouteFilter
func (this *HTTPRouteFilter) UnmarshalJSON(b []byte) error {
	return HttpRouteUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for HTTPHeaderFilter
func (this *HTTPHeaderFilter) MarshalJSON() ([]byte, error) {
	str, err := HttpRouteMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPHeaderFilter
func (this *HTTPHeaderFilter) UnmarshalJSON(b []byte) error {
	return HttpRouteUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for HTTPHeader
func (this *HTTPHeader) MarshalJSON() ([]byte, error) {
	str, err := HttpRouteMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPHeader
func (this *HTTPHeader) UnmarshalJSON(b []byte) error {
	return HttpRouteUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for HTTPURLRewriteFilter
func (this *HTTPURLRewriteFilter) MarshalJSON() ([]byte, error) {
	str, err := HttpRouteMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPURLRewriteFilter
func (this *HTTPURLRewriteFilter) UnmarshalJSON(b []byte) error {
	return HttpRouteUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for HTTPBackendRef
func (this *HTTPBackendRef) MarshalJSON() ([]byte, error) {
	str, err := HttpRouteMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPBackendRef
func (this *HTTPBackendRef) UnmarshalJSON(b []byte) error {
	return HttpRouteUnmarshaler.Unmarshal(b, this)
}

var (
	HttpRouteMarshaler   = &protojson.MarshalOptions{}
	HttpRouteUnmarshaler = &protojson.UnmarshalOptions{DiscardUnknown: false}
)