// SPDX-FileCopyrightText: 2019 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package genericactuator

import (
	"context"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	"github.com/gardener/gardener/extensions/pkg/controller/backupentry"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	kutil "github.com/gardener/gardener/pkg/utils/kubernetes"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
)

type actuator struct {
	backupEntryDelegate BackupEntryDelegate
	client              client.Client
	logger              logr.Logger
}

// InjectClient injects the given client into the valuesProvider.
func (a *actuator) InjectClient(client client.Client) error {
	a.client = client
	return nil
}

// InjectFunc enables injecting Kubernetes dependencies into actuator's dependencies.
func (a *actuator) InjectFunc(f inject.Func) error {
	return f(a.backupEntryDelegate)
}

// NewActuator creates a new Actuator that updates the status of the handled BackupEntry resources.
func NewActuator(backupEntryDelegate BackupEntryDelegate, logger logr.Logger) backupentry.Actuator {
	return &actuator{
		logger:              logger,
		backupEntryDelegate: backupEntryDelegate,
	}
}

// Reconcile reconciles the update of a BackupEntry
func (a *actuator) Reconcile(ctx context.Context, be *extensionsv1alpha1.BackupEntry) error {
	return a.deployEtcdBackupSecret(ctx, be)
}

func (a *actuator) deployEtcdBackupSecret(ctx context.Context, be *extensionsv1alpha1.BackupEntry) error {
	shootTechnicalID, _ := backupentry.ExtractShootDetailsFromBackupEntryName(be.Name)
	namespace := &corev1.Namespace{}
	if err := a.client.Get(ctx, kutil.Key(shootTechnicalID), namespace); err != nil {
		if apierrors.IsNotFound(err) {
			a.logger.Info("SeedNamespace for shoot not found. Avoiding etcd backup secret deployment")
			return nil
		}
		a.logger.Error(err, "failed to get seed namespace")
		return err
	}
	if namespace.DeletionTimestamp != nil {
		a.logger.Info("SeedNamespace for shoot is being terminated. Avoiding etcd backup secret deployment")
		return nil
	}

	backupSecret, err := extensionscontroller.GetSecretByReference(ctx, a.client, &be.Spec.SecretRef)
	if err != nil {
		a.logger.Error(err, "failed to read backup extension secret")
		return err
	}

	backupSecretData := backupSecret.DeepCopy().Data
	backupSecretData[backupBucketName] = []byte(be.Spec.BucketName)
	etcdSecretData, err := a.backupEntryDelegate.GetETCDSecretData(ctx, be, backupSecretData)
	if err != nil {
		return err
	}

	etcdSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      EtcdBackupSecretName,
			Namespace: shootTechnicalID,
		},
	}

	_, err = controllerutil.CreateOrUpdate(ctx, a.client, etcdSecret, func() error {
		etcdSecret.Data = etcdSecretData
		return nil
	})
	return err
}

// Delete deletes the BackupEntry
func (a *actuator) Delete(ctx context.Context, be *extensionsv1alpha1.BackupEntry) error {
	return a.backupEntryDelegate.Delete(ctx, be)
}

// Restore restores the update of a BackupEntry
func (a *actuator) Restore(ctx context.Context, be *extensionsv1alpha1.BackupEntry) error {
	return a.Reconcile(ctx, be)
}

// Migrate migrates the BackupEntry
func (a *actuator) Migrate(ctx context.Context, be *extensionsv1alpha1.BackupEntry) error {
	return nil
}
