// SPDX-FileCopyrightText: 2018 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package shoot

import (
	"context"
	"net"
	"time"

	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/gardener/gardener/pkg/operation/botanist/component"
	"github.com/gardener/gardener/pkg/operation/botanist/controlplane/kubescheduler"
	"github.com/gardener/gardener/pkg/operation/etcdencryption"
	"github.com/gardener/gardener/pkg/operation/garden"

	"github.com/Masterminds/semver"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Builder is an object that builds Shoot objects.
type Builder struct {
	shootObjectFunc  func() (*gardencorev1beta1.Shoot, error)
	cloudProfileFunc func(string) (*gardencorev1beta1.CloudProfile, error)
	shootSecretFunc  func(context.Context, client.Client, string, string) (*corev1.Secret, error)
	projectName      string
	internalDomain   *garden.Domain
	defaultDomains   []*garden.Domain
	disableDNS       bool
}

// Shoot is an object containing information about a Shoot cluster.
type Shoot struct {
	Info         *gardencorev1beta1.Shoot
	Secret       *corev1.Secret
	CloudProfile *gardencorev1beta1.CloudProfile

	SeedNamespace               string
	KubernetesMajorMinorVersion string
	KubernetesVersion           *semver.Version

	DisableDNS            bool
	InternalClusterDomain string
	ExternalClusterDomain *string
	ExternalDomain        *garden.Domain

	WantsClusterAutoscaler     bool
	WantsVerticalPodAutoscaler bool
	WantsAlertmanager          bool
	IgnoreAlerts               bool
	HibernationEnabled         bool
	KonnectivityTunnelEnabled  bool
	NodeLocalDNSEnabled        bool
	Networks                   *Networks

	Components *Components

	OperatingSystemConfigsMap map[string]OperatingSystemConfigs
	Extensions                map[string]Extension
	InfrastructureStatus      []byte
	ControlPlaneStatus        []byte
	MachineDeployments        []extensionsv1alpha1.MachineDeployment

	ETCDEncryption *etcdencryption.EncryptionConfig

	ResourceRefs map[string]autoscalingv1.CrossVersionObjectReference
}

// Components contains different components deployed in the Shoot cluster.
type Components struct {
	Extensions      *Extensions
	ControlPlane    *ControlPlane
	ClusterIdentity component.Deployer
}

// ControlPlane contains references to K8S control plane components.
type ControlPlane struct {
	KubeAPIServerService component.DeployWaiter
	KubeAPIServerSNI     component.DeployWaiter
	KubeScheduler        kubescheduler.KubeScheduler
}

// Extensions contains references to extension resources.
type Extensions struct {
	DNS              *DNS
	Infrastructure   Infrastructure
	Network          component.DeployMigrateWaiter
	ContainerRuntime ContainerRuntime
}

// DNS contains references to internal and external DNSProvider and DNSEntry deployers.
type DNS struct {
	ExternalOwner       component.DeployWaiter
	ExternalProvider    component.DeployWaiter
	ExternalEntry       component.DeployWaiter
	InternalOwner       component.DeployWaiter
	InternalProvider    component.DeployWaiter
	InternalEntry       component.DeployWaiter
	AdditionalProviders map[string]component.DeployWaiter
	NginxOwner          component.DeployWaiter
	NginxEntry          component.DeployWaiter
}

// Infrastructure contains references to an Infrastructure extension deployer and its generated
// provider status.
type Infrastructure interface {
	component.DeployWaiter
	SetSSHPublicKey([]byte)
	ProviderStatus() *runtime.RawExtension
	NodesCIDR() *string
}

// ContainerRuntime contains references to a ContainerRuntime extension deployer.
type ContainerRuntime interface {
	component.DeployMigrateWaiter
	DeleteStaleResources(ctx context.Context) error
}

// Networks contains pre-calculated subnets and IP address for various components.
type Networks struct {
	// Pods subnet
	Pods *net.IPNet
	// Services subnet
	Services *net.IPNet
	// APIServer is the ClusterIP of default/kubernetes Service
	APIServer net.IP
	// CoreDNS is the ClusterIP of kube-system/coredns Service
	CoreDNS net.IP
}

// OperatingSystemConfigs contains operating system configs for the downloader script as well as for the original cloud config.
type OperatingSystemConfigs struct {
	Downloader OperatingSystemConfig
	Original   OperatingSystemConfig
}

// OperatingSystemConfig contains the operating system config's name and data.
type OperatingSystemConfig struct {
	Name string
	Data OperatingSystemConfigData
}

// OperatingSystemConfigData contains the actual content, a command to load it and all units that
// shall be considered for restart on change.
type OperatingSystemConfigData struct {
	Content string
	Command *string
	Units   []string
}

// Extension contains information about the extension api resouce as well as configuration information.
type Extension struct {
	extensionsv1alpha1.Extension
	Timeout time.Duration
}

// IncompleteDNSConfigError is a custom error type.
type IncompleteDNSConfigError struct{}

// Error prints the error message of the IncompleteDNSConfigError error.
func (e *IncompleteDNSConfigError) Error() string {
	return "unable to figure out which secret should be used for dns"
}

// IsIncompleteDNSConfigError returns true if the error indicates that not the DNS config is incomplete.
func IsIncompleteDNSConfigError(err error) bool {
	_, ok := err.(*IncompleteDNSConfigError)
	return ok
}
