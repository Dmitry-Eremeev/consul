// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package meshv2beta1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using HTTPRouteRetries within kubernetes types, where deepcopy-gen is used.
func (in *HTTPRouteRetries) DeepCopyInto(out *HTTPRouteRetries) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteRetries. Required by controller-gen.
func (in *HTTPRouteRetries) DeepCopy() *HTTPRouteRetries {
	if in == nil {
		return nil
	}
	out := new(HTTPRouteRetries)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteRetries. Required by controller-gen.
func (in *HTTPRouteRetries) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}