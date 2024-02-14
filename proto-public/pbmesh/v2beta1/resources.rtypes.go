// Code generated by protoc-gen-resource-types. DO NOT EDIT.

package meshv2beta1

import (
	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	GroupName = "mesh"
	Version   = "v2beta1"

	APIGatewayKind                   = "APIGateway"
	ComputedExplicitDestinationsKind = "ComputedExplicitDestinations"
	ComputedGatewayConfigurationKind = "ComputedGatewayConfiguration"
	ComputedImplicitDestinationsKind = "ComputedImplicitDestinations"
	ComputedProxyConfigurationKind   = "ComputedProxyConfiguration"
	ComputedRoutesKind               = "ComputedRoutes"
	DestinationPolicyKind            = "DestinationPolicy"
	DestinationsKind                 = "Destinations"
	DestinationsConfigurationKind    = "DestinationsConfiguration"
	GRPCRouteKind                    = "GRPCRoute"
	HTTPRouteKind                    = "HTTPRoute"
	InlineCertificateKind            = "InlineCertificate"
	MeshConfigurationKind            = "MeshConfiguration"
	MeshGatewayKind                  = "MeshGateway"
	ProxyConfigurationKind           = "ProxyConfiguration"
	ProxyStateTemplateKind           = "ProxyStateTemplate"
	TCPRouteKind                     = "TCPRoute"
)

var (
	APIGatewayType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         APIGatewayKind,
	}

	ComputedExplicitDestinationsType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         ComputedExplicitDestinationsKind,
	}

	ComputedGatewayConfigurationType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         ComputedGatewayConfigurationKind,
	}

	ComputedImplicitDestinationsType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         ComputedImplicitDestinationsKind,
	}

	ComputedProxyConfigurationType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         ComputedProxyConfigurationKind,
	}

	ComputedRoutesType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         ComputedRoutesKind,
	}

	DestinationPolicyType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         DestinationPolicyKind,
	}

	DestinationsType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         DestinationsKind,
	}

	DestinationsConfigurationType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         DestinationsConfigurationKind,
	}

	GRPCRouteType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         GRPCRouteKind,
	}

	HTTPRouteType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         HTTPRouteKind,
	}

	InlineCertificateType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         InlineCertificateKind,
	}

	MeshConfigurationType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         MeshConfigurationKind,
	}

	MeshGatewayType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         MeshGatewayKind,
	}

	ProxyConfigurationType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         ProxyConfigurationKind,
	}

	ProxyStateTemplateType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         ProxyStateTemplateKind,
	}

	TCPRouteType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         TCPRouteKind,
	}
)
