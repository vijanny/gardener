// SPDX-FileCopyrightText: 2019 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	componentbaseconfigv1alpha1 "k8s.io/component-base/config/v1alpha1"
	"k8s.io/klog"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerManagerConfiguration defines the configuration for the Gardener controller manager.
type ControllerManagerConfiguration struct {
	metav1.TypeMeta `json:",inline"`
	// GardenClientConnection specifies the kubeconfig file and the client connection settings
	// for the proxy server to use when communicating with the garden apiserver.
	GardenClientConnection componentbaseconfigv1alpha1.ClientConnectionConfiguration `json:"gardenClientConnection"`
	// Controllers defines the configuration of the controllers.
	Controllers ControllerManagerControllerConfiguration `json:"controllers"`
	// LeaderElection defines the configuration of leader election client.
	LeaderElection LeaderElectionConfiguration `json:"leaderElection"`
	// LogLevel is the level/severity for the logs. Must be one of [info,debug,error].
	LogLevel string `json:"logLevel"`
	// KubernetesLogLevel is the log level used for Kubernetes' k8s.io/klog functions.
	KubernetesLogLevel klog.Level `json:"kubernetesLogLevel"`
	// Server defines the configuration of the HTTP server.
	Server ServerConfiguration `json:"server"`
	// FeatureGates is a map of feature names to bools that enable or disable alpha/experimental
	// features. This field modifies piecemeal the built-in default values from
	// "github.com/gardener/gardener/pkg/controllermanager/features/features.go".
	// Default: nil
	// +optional
	FeatureGates map[string]bool `json:"featureGates,omitempty"`
}

// ControllerManagerControllerConfiguration defines the configuration of the controllers.
type ControllerManagerControllerConfiguration struct {
	// CloudProfile defines the configuration of the CloudProfile controller.
	// +optional
	CloudProfile *CloudProfileControllerConfiguration `json:"cloudProfile,omitempty"`
	// ControllerRegistration defines the configuration of the ControllerRegistration controller.
	// +optional
	ControllerRegistration *ControllerRegistrationControllerConfiguration `json:"controllerRegistration,omitempty"`
	// Event defines the configuration of the Event controller.  If unset, the event controller will be disabled.
	// +optional
	Event *EventControllerConfiguration `json:"event,omitempty"`
	// Plant defines the configuration of the Plant controller.
	// +optional
	Plant *PlantControllerConfiguration `json:"plant,omitempty"`
	// Project defines the configuration of the Project controller.
	// +optional
	Project *ProjectControllerConfiguration `json:"project,omitempty"`
	// Quota defines the configuration of the Quota controller.
	// +optional
	Quota *QuotaControllerConfiguration `json:"quota,omitempty"`
	// SecretBinding defines the configuration of the SecretBinding controller.
	// +optional
	SecretBinding *SecretBindingControllerConfiguration `json:"secretBinding,omitempty"`
	// Seed defines the configuration of the Seed lifecycle controller.
	// +optional
	Seed *SeedControllerConfiguration `json:"seed,omitempty"`
	// ShootMaintenance defines the configuration of the ShootMaintenance controller.
	ShootMaintenance ShootMaintenanceControllerConfiguration `json:"shootMaintenance"`
	// ShootQuota defines the configuration of the ShootQuota controller.
	ShootQuota ShootQuotaControllerConfiguration `json:"shootQuota"`
	// ShootHibernation defines the configuration of the ShootHibernation controller.
	ShootHibernation ShootHibernationControllerConfiguration `json:"shootHibernation"`
}

// CloudProfileControllerConfiguration defines the configuration of the CloudProfile
// controller.
type CloudProfileControllerConfiguration struct {
	// ConcurrentSyncs is the number of workers used for the controller to work on
	// events.
	ConcurrentSyncs int `json:"concurrentSyncs"`
}

// ControllerRegistrationControllerConfiguration defines the configuration of the
// ControllerRegistration controller.
type ControllerRegistrationControllerConfiguration struct {
	// ConcurrentSyncs is the number of workers used for the controller to work on
	// events.
	ConcurrentSyncs int `json:"concurrentSyncs"`
}

// EventControllerConfiguration defines the configuration of the Event controller.
type EventControllerConfiguration struct {
	// ConcurrentSyncs is the number of workers used for the controller to work on
	// events.
	ConcurrentSyncs int `json:"concurrentSyncs"`
	// TTLNonShootEvents is the time-to-live for all non-shoot related events (defaults to `1h`).
	// +optional
	TTLNonShootEvents *metav1.Duration `json:"ttlNonShootEvents,omitempty"`
}

// PlantControllerConfiguration defines the configuration of the
// PlantControllerConfiguration controller.
type PlantControllerConfiguration struct {
	// ConcurrentSyncs is the number of workers used for the controller to work on
	// events.
	ConcurrentSyncs int `json:"concurrentSyncs"`
	// SyncPeriod is the duration how often the existing resources are reconciled.
	SyncPeriod metav1.Duration `json:"syncPeriod"`
}

// ProjectControllerConfiguration defines the configuration of the
// Project controller.
type ProjectControllerConfiguration struct {
	// ConcurrentSyncs is the number of workers used for the controller to work on
	// events.
	ConcurrentSyncs int `json:"concurrentSyncs"`
	// MinimumLifetimeDays is the number of days a `Project` may exist before it is being
	// checked whether it is actively used or got stale.
	// +optional
	MinimumLifetimeDays *int `json:"minimumLifetimeDays,omitempty"`
	// StaleGracePeriodDays is the number of days a `Project` may be unused before it will
	// be considered for checks whether it is actively used or got stale.
	// +optional
	StaleGracePeriodDays *int `json:"staleGracePeriodDays,omitempty"`
	// StaleExpirationTimeDays is the number of days after a `Project` that has been marked as
	// 'stale'/'unused' and passed the 'stale grace period' will be considered for auto deletion.
	// +optional
	StaleExpirationTimeDays *int `json:"staleExpirationTimeDays,omitempty"`
	// StaleSyncPeriod is the duration how often the reconciliation loop for stale Projects is executed.
	// +optional
	StaleSyncPeriod *metav1.Duration `json:"staleSyncPeriod,omitempty"`
}

// QuotaControllerConfiguration defines the configuration of the Quota controller.
type QuotaControllerConfiguration struct {
	// ConcurrentSyncs is the number of workers used for the controller to work on
	// events.
	ConcurrentSyncs int `json:"concurrentSyncs"`
}

// SecretBindingControllerConfiguration defines the configuration of the
// SecretBinding controller.
type SecretBindingControllerConfiguration struct {
	// ConcurrentSyncs is the number of workers used for the controller to work on
	// events.
	ConcurrentSyncs int `json:"concurrentSyncs"`
}

// SeedControllerConfiguration defines the configuration of the
// Seed controller.
type SeedControllerConfiguration struct {
	// ConcurrentSyncs is the number of workers used for the controller to work on
	// events.
	ConcurrentSyncs int `json:"concurrentSyncs"`
	// MonitorPeriod is the duration after the seed controller will mark the `GardenletReady`
	// condition in `Seed` resources as `Unknown` in case the gardenlet did not send heartbeats.
	// +optional
	MonitorPeriod *metav1.Duration `json:"monitorPeriod,omitempty"`
	// ShootMonitorPeriod is the duration after the seed controller will mark Gardener's conditions
	// in `Shoot` resources as `Unknown` in case the gardenlet of the responsible seed cluster did
	// not send heartbeats.
	// +optional
	ShootMonitorPeriod *metav1.Duration `json:"shootMonitorPeriod,omitempty"`
	// SyncPeriod is the duration how often the existing resources are reconciled.
	SyncPeriod metav1.Duration `json:"syncPeriod"`
}

// ShootMaintenanceControllerConfiguration defines the configuration of the
// ShootMaintenance controller.
type ShootMaintenanceControllerConfiguration struct {
	// ConcurrentSyncs is the number of workers used for the controller to work on
	// events.
	ConcurrentSyncs int `json:"concurrentSyncs"`
	// EnableShootControlPlaneRestarter configures whether adequate pods of the shoot control plane are restarted during maintenance.
	// +optional
	EnableShootControlPlaneRestarter *bool `json:"enableShootControlPlaneRestarter"`
}

// ShootQuotaControllerConfiguration defines the configuration of the
// ShootQuota controller.
type ShootQuotaControllerConfiguration struct {
	// ConcurrentSyncs is the number of workers used for the controller to work on
	// events.
	ConcurrentSyncs int `json:"concurrentSyncs"`
	// SyncPeriod is the duration how often the existing resources are reconciled
	// (how often Shoots referenced Quota is checked).
	SyncPeriod metav1.Duration `json:"syncPeriod"`
}

// ShootHibernationControllerConfiguration defines the configuration of the
// ShootHibernation controller.
type ShootHibernationControllerConfiguration struct {
	// ConcurrentSyncs is the number of workers used for the controller to work on
	// events.
	ConcurrentSyncs int `json:"concurrentSyncs"`
}

// LeaderElectionConfiguration defines the configuration of leader election
// clients for components that can run with leader election enabled.
type LeaderElectionConfiguration struct {
	componentbaseconfigv1alpha1.LeaderElectionConfiguration `json:",inline"`
	// LockObjectNamespace defines the namespace of the lock object.
	LockObjectNamespace string `json:"lockObjectNamespace"`
	// LockObjectName defines the lock object name.
	LockObjectName string `json:"lockObjectName"`
}

// ServerConfiguration contains details for the HTTP(S) servers.
type ServerConfiguration struct {
	// HTTP is the configuration for the HTTP server.
	HTTP Server `json:"http"`
	// HTTPS is the configuration for the HTTPS server.
	HTTPS HTTPSServer `json:"https"`
}

// Server contains information for HTTP(S) server configuration.
type Server struct {
	// BindAddress is the IP address on which to listen for the specified port.
	BindAddress string `json:"bindAddress"`
	// Port is the port on which to serve requests.
	Port int `json:"port"`
}

// HTTPSServer is the configuration for the HTTPSServer server.
type HTTPSServer struct {
	// Server is the configuration for the bind address and the port.
	Server `json:",inline"`
	// TLSServer contains information about the TLS configuration for a HTTPS server.
	TLS TLSServer `json:"tls"`
}

// TLSServer contains information about the TLS configuration for a HTTPS server.
type TLSServer struct {
	// ServerCertPath is the path to the server certificate file.
	ServerCertPath string `json:"serverCertPath"`
	// ServerKeyPath is the path to the private key file.
	ServerKeyPath string `json:"serverKeyPath"`
}

const (
	// ControllerManagerDefaultLockObjectNamespace is the default lock namespace for leader election.
	ControllerManagerDefaultLockObjectNamespace = "garden"

	// ControllerManagerDefaultLockObjectName is the default lock name for leader election.
	ControllerManagerDefaultLockObjectName = "gardener-controller-manager-leader-election"

	// DefaultDiscoveryTTL is the default ttl for the cached discovery client.
	DefaultDiscoveryTTL = 10 * time.Second
)
