// SPDX-FileCopyrightText: 2019 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package backupentry

import (
	"context"

	extensionspredicate "github.com/gardener/gardener/extensions/pkg/predicate"

	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type secretToBackupEntryMapper struct {
	client     client.Client
	predicates []predicate.Predicate
}

func (m *secretToBackupEntryMapper) Map(obj handler.MapObject) []reconcile.Request {
	if obj.Object == nil {
		return nil
	}

	secret, ok := obj.Object.(*corev1.Secret)
	if !ok {
		return nil
	}

	backupEntryList := &extensionsv1alpha1.BackupEntryList{}
	if err := m.client.List(context.TODO(), backupEntryList); err != nil {
		return nil
	}

	var requests []reconcile.Request
	for _, backupEntry := range backupEntryList.Items {
		if backupEntry.Spec.SecretRef.Name == secret.Name && backupEntry.Spec.SecretRef.Namespace == secret.Namespace {
			if extensionspredicate.EvalGeneric(&backupEntry, m.predicates...) {
				requests = append(requests, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Name: backupEntry.Name,
					},
				})
			}
		}
	}

	return requests
}

// SecretToBackupEntryMapper returns a mapper that returns requests for BackupEntry whose
// referenced secrets have been modified.
func SecretToBackupEntryMapper(client client.Client, predicates []predicate.Predicate) handler.Mapper {
	return &secretToBackupEntryMapper{client, predicates}
}

type namespaceToBackupEntryMapper struct {
	client     client.Client
	predicates []predicate.Predicate
}

func (m *namespaceToBackupEntryMapper) Map(obj handler.MapObject) []reconcile.Request {
	if obj.Object == nil {
		return nil
	}

	namespace, ok := obj.Object.(*corev1.Namespace)
	if !ok {
		return nil
	}

	backupEntryList := &extensionsv1alpha1.BackupEntryList{}
	if err := m.client.List(context.TODO(), backupEntryList); err != nil {
		return nil
	}

	shootUID := namespace.Annotations[v1beta1constants.DeprecatedShootUID]
	var requests []reconcile.Request
	for _, backupEntry := range backupEntryList.Items {
		if !extensionspredicate.EvalGeneric(&backupEntry, m.predicates...) {
			continue
		}

		expectedTechnicalID, expectedUID := ExtractShootDetailsFromBackupEntryName(backupEntry.Name)
		if namespace.Name == expectedTechnicalID && shootUID == expectedUID {
			requests = append(requests, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Name: backupEntry.Name,
				},
			})
		}
	}
	return requests
}

// NamespaceToBackupEntryMapper returns a mapper that returns requests for BackupEntry whose
// associated Shoot's seed namespace have been modified.
func NamespaceToBackupEntryMapper(client client.Client, predicates []predicate.Predicate) handler.Mapper {
	return &namespaceToBackupEntryMapper{client, predicates}
}
