// SPDX-FileCopyrightText: 2018 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package core

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerRegistration represents a registration of an external controller.
type ControllerRegistration struct {
	metav1.TypeMeta
	// Standard object metadata.
	metav1.ObjectMeta
	// Spec contains the specification of this registration.
	Spec ControllerRegistrationSpec
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerRegistrationList is a collection of ControllerRegistrations.
type ControllerRegistrationList struct {
	metav1.TypeMeta
	// Standard list object metadata.
	metav1.ListMeta
	// Items is the list of ControllerRegistrations.
	Items []ControllerRegistration
}

// ControllerRegistrationSpec is the specification of a ControllerRegistration.
type ControllerRegistrationSpec struct {
	// Resources is a list of combinations of kinds (DNSProvider, Infrastructure, Generic, ...) and their actual types
	// (aws-route53, gcp, auditlog, ...).
	Resources []ControllerResource
	// Deployment contains information for how this controller is deployed.
	Deployment *ControllerDeployment
}

// ControllerResource is a combination of a kind (DNSProvider, Infrastructure, Generic, ...) and the actual type for this
// kind (aws-route53, gcp, auditlog, ...).
type ControllerResource struct {
	// Kind is the resource kind.
	Kind string
	// Type is the resource type.
	Type string
	// GloballyEnabled determines if this resource is required by all Shoot clusters.
	GloballyEnabled *bool
	// ReconcileTimeout defines how long Gardener should wait for the resource reconciliation.
	ReconcileTimeout *metav1.Duration
	// Primary determines if the controller backed by this ControllerRegistration is responsible for the extension
	// resource's lifecycle. This field defaults to true. There must be exactly one primary controller for this kind/type
	// combination.
	Primary *bool
}

// ControllerDeployment contains information for how this controller is deployed.
type ControllerDeployment struct {
	// Type is the deployment type.
	Type string
	// ProviderConfig contains type-specific configuration.
	ProviderConfig *runtime.RawExtension
	// Policy controls how the controller is deployed. It defaults to 'OnDemand'.
	Policy *ControllerDeploymentPolicy
	// SeedSelector contains an optional label selector for seeds. Only if the labels match then this controller will be
	// considered for a deployment.
	// An empty list means that all seeds are selected.
	SeedSelector *metav1.LabelSelector
}

// ControllerDeploymentPolicy is a string alias.
type ControllerDeploymentPolicy string

const (
	// ControllerDeploymentPolicyOnDemand specifies that the controller shall be only deployed if required by another
	// resource. If nothing requires it then the controller shall not be deployed.
	ControllerDeploymentPolicyOnDemand ControllerDeploymentPolicy = "OnDemand"
	// ControllerDeploymentPolicyAlways specifies that the controller shall be deployed always, independent of whether
	// another resource requires it.
	ControllerDeploymentPolicyAlways ControllerDeploymentPolicy = "Always"
)
