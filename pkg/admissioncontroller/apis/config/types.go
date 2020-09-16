// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	componentbaseconfig "k8s.io/component-base/config"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AdmissionControllerConfiguration defines the configuration for the Gardener admission controller.
type AdmissionControllerConfiguration struct {
	metav1.TypeMeta
	// GardenClientConnection specifies the kubeconfig file and the client connection settings
	// when communicating with the garden apiserver.
	GardenClientConnection componentbaseconfig.ClientConnectionConfiguration
	// LogLevel is the level/severity for the logs. Must be one of [info,debug,error].
	LogLevel string
	// Server defines the configuration of the HTTP server.
	Server ServerConfiguration
}

// ServerConfiguration contains details for the HTTP(S) servers.
type ServerConfiguration struct {
	// HTTPS is the configuration for the HTTPS server.
	HTTPS HTTPSServer
	// ResourceAdmissionConfiguration is the configuration for the resource admission.
	ResourceAdmissionConfiguration *ResourceAdmissionConfiguration
}

// ResourceAdmissionConfiguration contains settings about arbitrary kinds and the size each resource should have at most.
type ResourceAdmissionConfiguration struct {
	// Limits contains configuration for resources which are subjected to size limitations.
	Limits []ResourceLimit
	// UnrestrictedSubjects contains references to users, groups, or service accounts which aren't subjected to any resource size limit.
	UnrestrictedSubjects []rbacv1.Subject
	// OperationMode specifies the mode the webhooks operates in. Allowed values are "block" and "log". Defaults to "block".
	OperationMode *ResourceAdmissionWebhookMode
}

// ResourceAdmissionWebhookMode is an alias type for the resource admission webhook mode.
type ResourceAdmissionWebhookMode string

// WildcardAll is a character which represents all elements in a set.
const WildcardAll = "*"

// ResourceLimit contains settings about a kind and the size each resource should have at most.
type ResourceLimit struct {
	// APIGroup is the name of the APIGroup that contains the limited resource. WildcardAll represents all groups.
	APIGroups []string
	// APIVersions is the version of the resource. WildcardAll represents all versions.
	APIVersions []string
	// Resource is the name of the resource this rule applies to. WildcardAll represents all resources.
	Resources []string
	// Size specifies the imposed limit.
	Size resource.Quantity
}

// Server contains information for HTTP(S) server configuration.
type Server struct {
	// BindAddress is the IP address on which to listen for the specified port.
	BindAddress string
	// Port is the port on which to serve requests.
	Port int
}

// HTTPSServer is the configuration for the HTTPSServer server.
type HTTPSServer struct {
	// Server is the configuration for the bind address and the port.
	Server
	// TLSServer contains information about the TLS configuration for a HTTPS server.
	TLS TLSServer
}

// TLSServer contains information about the TLS configuration for a HTTPS server.
type TLSServer struct {
	// ServerCertPath is the path to the server certificate file.
	ServerCertPath string
	// ServerKeyPath is the path to the private key file.
	ServerKeyPath string
}

const (
	// AdmissionModeBlock specifies that the webhook should block violating requests.
	AdmissionModeBlock ResourceAdmissionWebhookMode = "block"
	// AdmissionModeLog specifies that the webhook should only log violating requests.
	AdmissionModeLog ResourceAdmissionWebhookMode = "log"
)
