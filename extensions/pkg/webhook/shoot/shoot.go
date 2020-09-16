// SPDX-FileCopyrightText: 2019 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package shoot

import (
	"fmt"

	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"

	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

const (
	// WebhookName is the name of the shoot webhook.
	WebhookName = "shoot"
	// KindSystem is used for webhooks which should only apply to the to the kube-system namespace.
	KindSystem = "system"
)

var logger = log.Log.WithName("shoot-webhook")

// Args are arguments for creating a webhook targeting a shoot.
type Args struct {
	// Types is a list of resource types.
	Types []runtime.Object
	// Mutator is a mutator to be used by the admission handler. It doesn't need the shoot client.
	Mutator extensionswebhook.Mutator
	// MutatorWithShootClient is a mutator to be used by the admission handler. It needs the shoot client.
	MutatorWithShootClient extensionswebhook.MutatorWithShootClient
}

// Add creates a new webhook with the shoot as target cluster.
func New(mgr manager.Manager, args Args) (*extensionswebhook.Webhook, error) {
	logger.Info("Creating webhook", "name", WebhookName)

	// Build namespace selector from the webhook kind and provider
	namespaceSelector, err := buildSelector()
	if err != nil {
		return nil, err
	}

	wh := &extensionswebhook.Webhook{
		Name:     WebhookName,
		Types:    args.Types,
		Path:     WebhookName,
		Target:   extensionswebhook.TargetShoot,
		Selector: namespaceSelector,
	}

	switch {
	case args.Mutator != nil:
		handler, err := extensionswebhook.NewBuilder(mgr, logger).WithMutator(args.Mutator, args.Types...).Build()
		if err != nil {
			return nil, err
		}

		wh.Webhook = &admission.Webhook{Handler: handler}
		return wh, nil

	case args.MutatorWithShootClient != nil:
		handler, err := extensionswebhook.NewHandlerWithShootClient(mgr, args.Types, args.MutatorWithShootClient, logger)
		if err != nil {
			return nil, err
		}

		if _, err := inject.SchemeInto(mgr.GetScheme(), handler); err != nil {
			return nil, err
		}

		wh.Handler = handler
		return wh, nil
	}

	return nil, fmt.Errorf("neither mutator nor mutator with shoot client is set")
}

// buildSelector creates and returns a LabelSelector for the given webhook kind and provider.
func buildSelector() (*metav1.LabelSelector, error) {
	// Create and return LabelSelector
	return &metav1.LabelSelector{
		MatchExpressions: []metav1.LabelSelectorRequirement{
			{Key: v1beta1constants.GardenerPurpose, Operator: metav1.LabelSelectorOpIn, Values: []string{metav1.NamespaceSystem}},
		},
	}, nil
}
