// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package csimigration

import (
	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	extensionspredicate "github.com/gardener/gardener/extensions/pkg/predicate"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	// ControllerName is the name of the controller
	ControllerName = "csimigration_controller"

	// AnnotationKeyNeedsComplete is a constant for an annotation on the Cluster resource that indicates that
	// the control plane components require the CSIMigration<Provider>Complete feature gates.
	AnnotationKeyNeedsComplete = "csi-migration.extensions.gardener.cloud/needs-complete-feature-gates"
	// AnnotationKeyControllerFinished is a constant for an annotation on the Cluster resource that indicates that
	// the CSI migration has nothing more to do anymore because he completed earlier already.
	AnnotationKeyControllerFinished = "csi-migration.extensions.gardener.cloud/controller-finished"
)

// AddArgs are arguments for adding an csimigration controller to a manager.
type AddArgs struct {
	// ControllerOptions are the controller options used for creating a controller.
	// The options.Reconciler is always overridden with a reconciler created from the
	// given actuator.
	ControllerOptions controller.Options
	// Predicates are the predicates to use.
	Predicates []predicate.Predicate
	// CSIMigrationKubernetesVersion is the smallest Kubernetes version that is used for the CSI migration.
	CSIMigrationKubernetesVersion string
	// Type is the provider extension type.
	Type string
	// StorageClassNameToLegacyProvisioner is a map of storage class names to the used legacy provisioner name. As part
	// of the CSI migration they will be deleted so that new storage classes with the same name but a different CSI
	// provisioner can be created (storage classes are immutable).
	StorageClassNameToLegacyProvisioner map[string]string
}

// Add creates a new CSIMigration Controller and adds it to the Manager.
// and Start it when the Manager is Started.
func Add(mgr manager.Manager, args AddArgs) error {
	reconciler, err := NewReconciler(args.CSIMigrationKubernetesVersion, args.StorageClassNameToLegacyProvisioner)
	if err != nil {
		return err
	}
	args.ControllerOptions.Reconciler = reconciler

	ctrl, err := controller.New(ControllerName, mgr, args.ControllerOptions)
	if err != nil {
		return err
	}

	decoder, err := extensionscontroller.NewGardenDecoder()
	if err != nil {
		return err
	}

	defaultPredicates := []predicate.Predicate{
		extensionspredicate.ClusterShootProviderType(decoder, args.Type),
		extensionspredicate.ClusterShootKubernetesVersionAtLeast(decoder, args.CSIMigrationKubernetesVersion),
		ClusterCSIMigrationControllerNotFinished(),
	}

	return ctrl.Watch(&source.Kind{Type: &extensionsv1alpha1.Cluster{}}, &handler.EnqueueRequestForObject{}, append(defaultPredicates, args.Predicates...)...)
}
