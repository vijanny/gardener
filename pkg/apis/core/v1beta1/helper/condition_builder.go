// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package helper

import (
	"fmt"

	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConditionBuilder build a Condition.
type ConditionBuilder interface {
	WithOldCondition(old gardencorev1beta1.Condition) ConditionBuilder
	WithStatus(status gardencorev1beta1.ConditionStatus) ConditionBuilder
	WithReason(reason string) ConditionBuilder
	WithMessage(message string) ConditionBuilder
	WithCodes(codes ...gardencorev1beta1.ErrorCode) ConditionBuilder
	WithNowFunc(now func() metav1.Time) ConditionBuilder
	Build() (new gardencorev1beta1.Condition, updated bool)
}

// defaultConditionBuilder build a Condition.
type defaultConditionBuilder struct {
	old           gardencorev1beta1.Condition
	status        gardencorev1beta1.ConditionStatus
	conditionType gardencorev1beta1.ConditionType
	reason        string
	message       string
	codes         []gardencorev1beta1.ErrorCode
	nowFunc       func() metav1.Time
}

// NewConditionBuilder returns a ConditionBuilder for a specific condition.
func NewConditionBuilder(conditionType gardencorev1beta1.ConditionType) (ConditionBuilder, error) {
	if conditionType == "" {
		return nil, fmt.Errorf("conditionType cannot be empty")
	}

	return &defaultConditionBuilder{
		conditionType: conditionType,
		nowFunc:       metav1.Now,
	}, nil
}

// WithOldCondition sets the old condition. It can be used to prodive default values.
// The old's condition type is overridden to the one specified in the builder.
func (b *defaultConditionBuilder) WithOldCondition(old gardencorev1beta1.Condition) ConditionBuilder {
	old.Type = b.conditionType
	b.old = old

	return b
}

// WithStatus sets the status of the condition.
func (b *defaultConditionBuilder) WithStatus(status gardencorev1beta1.ConditionStatus) ConditionBuilder {
	b.status = status
	return b
}

// WithReason sets the reason of the condition.
func (b *defaultConditionBuilder) WithReason(reason string) ConditionBuilder {
	b.reason = reason
	return b
}

// WithMessage sets the message of the condition.
func (b *defaultConditionBuilder) WithMessage(message string) ConditionBuilder {
	b.message = message
	return b
}

// WithCodes sets the codes of the condition.
func (b *defaultConditionBuilder) WithCodes(codes ...gardencorev1beta1.ErrorCode) ConditionBuilder {
	b.codes = codes
	return b
}

// WithNowFunc sets the function used for getting the current time.
// Should only be used for tests.
func (b *defaultConditionBuilder) WithNowFunc(now func() metav1.Time) ConditionBuilder {
	b.nowFunc = now
	return b
}

// Build creates the condition and returns if there are modifications with the OldCondition.
// If OldCondition is provided:
// - Any changes to status set the `LastTransitionTime`
// - Any updates to the message or the reason cause set `LastUpdateTime` to the current time.
func (b *defaultConditionBuilder) Build() (new gardencorev1beta1.Condition, updated bool) {
	var (
		now       = b.nowFunc()
		emptyTime = metav1.Time{}
	)

	new = *b.old.DeepCopy()

	if new.LastTransitionTime == emptyTime {
		new.LastTransitionTime = now
	}

	if new.LastUpdateTime == emptyTime {
		new.LastUpdateTime = now
	}

	new.Type = b.conditionType

	if b.status != "" {
		new.Status = b.status
	} else if b.status == "" && b.old.Status == "" {
		new.Status = gardencorev1beta1.ConditionUnknown
	}

	if b.reason != "" {
		new.Reason = b.reason
	} else if b.reason == "" && b.old.Reason == "" {
		new.Reason = "ConditionInitialized"
	}

	if b.message != "" {
		new.Message = b.message
	} else if b.message == "" && b.old.Message == "" {
		new.Message = "The condition has been initialized but its semantic check has not been performed yet."
	}

	if b.codes != nil {
		new.Codes = b.codes
	} else if b.codes == nil && b.old.Codes == nil {
		new.Codes = nil
	}

	if new.Status != b.old.Status {
		new.LastTransitionTime = now
	}

	if new.Reason != b.old.Reason || new.Message != b.old.Message {
		new.LastUpdateTime = now
	}

	return new, !apiequality.Semantic.DeepEqual(new, b.old)
}
