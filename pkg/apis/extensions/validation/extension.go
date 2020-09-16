// SPDX-FileCopyrightText: 2019 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package validation

import (
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	apivalidation "k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateExtension validates a Extension object.
func ValidateExtension(ext *extensionsv1alpha1.Extension) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, apivalidation.ValidateObjectMeta(&ext.ObjectMeta, true, apivalidation.NameIsDNSSubdomain, field.NewPath("metadata"))...)
	allErrs = append(allErrs, ValidateExtensionSpec(&ext.Spec, field.NewPath("spec"))...)

	return allErrs
}

// ValidateExtensionUpdate validates a Extension object before an update.
func ValidateExtensionUpdate(new, old *extensionsv1alpha1.Extension) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, apivalidation.ValidateObjectMetaUpdate(&new.ObjectMeta, &old.ObjectMeta, field.NewPath("metadata"))...)
	allErrs = append(allErrs, ValidateExtensionSpecUpdate(&new.Spec, &old.Spec, new.DeletionTimestamp != nil, field.NewPath("spec"))...)
	allErrs = append(allErrs, ValidateExtension(new)...)

	return allErrs
}

// ValidateExtensionSpec validates the specification of a Extension object.
func ValidateExtensionSpec(spec *extensionsv1alpha1.ExtensionSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if len(spec.Type) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("type"), "field is required"))
	}

	return allErrs
}

// ValidateExtensionSpecUpdate validates the spec of a Extension object before an update.
func ValidateExtensionSpecUpdate(new, old *extensionsv1alpha1.ExtensionSpec, deletionTimestampSet bool, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if deletionTimestampSet && !apiequality.Semantic.DeepEqual(new, old) {
		allErrs = append(allErrs, apivalidation.ValidateImmutableField(new, old, fldPath)...)
		return allErrs
	}

	allErrs = append(allErrs, apivalidation.ValidateImmutableField(new.Type, old.Type, fldPath.Child("type"))...)

	return allErrs
}

// ValidateExtensionStatus validates the status of a Extension object.
func ValidateExtensionStatus(spec *extensionsv1alpha1.ExtensionStatus, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	return allErrs
}

// ValidateExtensionStatusUpdate validates the status field of a Extension object.
func ValidateExtensionStatusUpdate(newStatus, oldStatus extensionsv1alpha1.ExtensionStatus) field.ErrorList {
	allErrs := field.ErrorList{}

	return allErrs
}
