// SPDX-FileCopyrightText: 2018 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package validation_test

import (
	"fmt"

	"github.com/gardener/gardener/pkg/apis/core"
	. "github.com/gardener/gardener/pkg/apis/core/validation"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	gomegatypes "github.com/onsi/gomega/types"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/pointer"
)

var _ = Describe("Project Validation Tests", func() {
	Describe("#ValidateProject, #ValidateProjectUpdate", func() {
		var project *core.Project

		BeforeEach(func() {
			project = &core.Project{
				ObjectMeta: metav1.ObjectMeta{
					Name: "project-1",
				},
				Spec: core.ProjectSpec{
					CreatedBy: &rbacv1.Subject{
						APIGroup: "rbac.authorization.k8s.io",
						Kind:     rbacv1.UserKind,
						Name:     "john.doe@example.com",
					},
					Owner: &rbacv1.Subject{
						APIGroup: "rbac.authorization.k8s.io",
						Kind:     rbacv1.UserKind,
						Name:     "john.doe@example.com",
					},
					Members: []core.ProjectMember{
						{
							Subject: rbacv1.Subject{
								APIGroup: "rbac.authorization.k8s.io",
								Kind:     rbacv1.UserKind,
								Name:     "alice.doe@example.com",
							},
							Roles: []string{core.ProjectMemberAdmin},
						},
						{
							Subject: rbacv1.Subject{
								APIGroup: "rbac.authorization.k8s.io",
								Kind:     rbacv1.UserKind,
								Name:     "bob.doe@example.com",
							},
							Roles: []string{core.ProjectMemberViewer, core.ProjectMemberUserAccessManager},
						},
					},
				},
			}
		})

		It("should not return any errors", func() {
			errorList := ValidateProject(project)

			Expect(errorList).To(BeEmpty())
		})

		It("should forbid Project resources with empty metadata", func() {
			project.ObjectMeta = metav1.ObjectMeta{}

			errorList := ValidateProject(project)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeRequired),
				"Field": Equal("metadata.name"),
			}))))
		})

		It("should forbid Projects having too long names", func() {
			project.ObjectMeta.Name = "project-name-too-long"

			errorList := ValidateProject(project)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeTooLong),
				"Field": Equal("metadata.name"),
			}))))
		})

		It("should forbid Projects with namespace gardener-system-seed-lease", func() {
			project.ObjectMeta.Namespace = "gardener-system-seed-lease"

			errorList := ValidateProject(project)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeForbidden),
				"Field": Equal("metadata.namespace"),
			}))))
		})

		It("should forbid Projects having two consecutive hyphens", func() {
			project.ObjectMeta.Name = "in--valid"

			errorList := ValidateProject(project)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("metadata.name"),
			}))))
		})

		It("should forbid Project specification with empty or invalid key for description", func() {
			project.Spec.Description = pointer.StringPtr("")

			errorList := ValidateProject(project)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeRequired),
				"Field": Equal("spec.description"),
			}))))
		})

		It("should forbid Project specification with empty or invalid key for purpose", func() {
			project.Spec.Purpose = pointer.StringPtr("")

			errorList := ValidateProject(project)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeRequired),
				"Field": Equal("spec.purpose"),
			}))))
		})

		It("should not allow duplicate in roles", func() {
			project.Spec.Members[0].Roles = []string{"admin", "admin"}

			errorList := ValidateProject(project)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeDuplicate),
				"Field": Equal("spec.members[0].roles[1]"),
			}))))
		})

		It("should not allow to use unknown roles without extension prefix", func() {
			project.Spec.Members[0].Roles = []string{"unknown-role"}

			errorList := ValidateProject(project)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeNotSupported),
				"Field": Equal("spec.members[0].roles[0]"),
			}))))
		})

		It("should prevent extension roles from being too long", func() {
			project.Spec.Members[0].Roles = []string{"extension:astringthatislongerthan15chars"}

			errorList := ValidateProject(project)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeTooLong),
				"Field": Equal("spec.members[0].roles[0]"),
			}))))
		})

		It("should prevent extension roles from containing invalid characters", func() {
			project.Spec.Members[0].Roles = []string{"extension:/?as"}

			errorList := ValidateProject(project)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("spec.members[0].roles[0]"),
			}))))
		})

		It("should allow to use unknown roles with extension prefix", func() {
			project.Spec.Members[0].Roles = []string{"extension:unknown-role"}

			errorList := ValidateProject(project)

			Expect(errorList).To(BeEmpty())
		})

		It("should not allow using the owner role more than once", func() {
			project.Spec.Members[0].Roles = append(project.Spec.Members[0].Roles, core.ProjectMemberOwner)
			project.Spec.Members[1].Roles = append(project.Spec.Members[1].Roles, core.ProjectMemberOwner)

			errorList := ValidateProject(project)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeForbidden),
				"Field": Equal("spec.members[1].roles[2]"),
			}))))
		})

		DescribeTable("subject validation",
			func(apiGroup, kind, name, namespace string, expectType field.ErrorType, field string) {
				subject := rbacv1.Subject{
					APIGroup:  apiGroup,
					Kind:      kind,
					Name:      name,
					Namespace: namespace,
				}

				project.Spec.Owner = &subject
				project.Spec.CreatedBy = &subject
				project.Spec.Members = []core.ProjectMember{
					{
						Subject: subject,
						Roles:   []string{core.ProjectMemberAdmin},
					},
				}

				errList := ValidateProject(project)

				Expect(errList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
					"Type":  Equal(expectType),
					"Field": Equal(fmt.Sprintf("spec.owner.%s", field)),
				})), PointTo(MatchFields(IgnoreExtras, Fields{
					"Type":  Equal(expectType),
					"Field": Equal(fmt.Sprintf("spec.createdBy.%s", field)),
				})), PointTo(MatchFields(IgnoreExtras, Fields{
					"Type":  Equal(expectType),
					"Field": Equal(fmt.Sprintf("spec.members[0].%s", field)),
				}))))
			},

			// general
			Entry("empty name", "rbac.authorization.k8s.io", rbacv1.UserKind, "", "", field.ErrorTypeRequired, "name"),
			Entry("unknown kind", "rbac.authorization.k8s.io", "unknown", "foo", "", field.ErrorTypeNotSupported, "kind"),

			// serviceaccounts
			Entry("invalid api group name", "apps/v1beta1", rbacv1.ServiceAccountKind, "foo", "default", field.ErrorTypeNotSupported, "apiGroup"),
			Entry("invalid name", "", rbacv1.ServiceAccountKind, "foo-", "default", field.ErrorTypeInvalid, "name"),
			Entry("no namespace", "", rbacv1.ServiceAccountKind, "foo", "", field.ErrorTypeRequired, "namespace"),

			// users
			Entry("invalid api group name", "rbac.authorization.invalid", rbacv1.UserKind, "john.doe@example.com", "", field.ErrorTypeNotSupported, "apiGroup"),

			// groups
			Entry("invalid api group name", "rbac.authorization.invalid", rbacv1.GroupKind, "groupname", "", field.ErrorTypeNotSupported, "apiGroup"),
		)

		It("should forbid invalid tolerations", func() {
			tolerations := []core.Toleration{
				{},
				{Key: "foo"},
				{Key: "foo"},
				{Key: "bar", Value: pointer.StringPtr("baz")},
				{Key: "bar", Value: pointer.StringPtr("baz")},
				{Key: "baz"},
				{Key: "baz", Value: pointer.StringPtr("baz")},
			}
			project.Spec.Tolerations = &core.ProjectTolerations{
				Defaults:  tolerations,
				Whitelist: tolerations,
			}

			errorList := ValidateProject(project)

			Expect(errorList).To(ConsistOf(
				PointTo(MatchFields(IgnoreExtras, Fields{
					"Type":  Equal(field.ErrorTypeRequired),
					"Field": Equal("spec.tolerations.defaults[0].key"),
				})),
				PointTo(MatchFields(IgnoreExtras, Fields{
					"Type":  Equal(field.ErrorTypeDuplicate),
					"Field": Equal("spec.tolerations.defaults[2]"),
				})),
				PointTo(MatchFields(IgnoreExtras, Fields{
					"Type":  Equal(field.ErrorTypeDuplicate),
					"Field": Equal("spec.tolerations.defaults[4]"),
				})),
				PointTo(MatchFields(IgnoreExtras, Fields{
					"Type":  Equal(field.ErrorTypeDuplicate),
					"Field": Equal("spec.tolerations.defaults[6]"),
				})),
				PointTo(MatchFields(IgnoreExtras, Fields{
					"Type":  Equal(field.ErrorTypeRequired),
					"Field": Equal("spec.tolerations.whitelist[0].key"),
				})),
				PointTo(MatchFields(IgnoreExtras, Fields{
					"Type":  Equal(field.ErrorTypeDuplicate),
					"Field": Equal("spec.tolerations.whitelist[2]"),
				})),
				PointTo(MatchFields(IgnoreExtras, Fields{
					"Type":  Equal(field.ErrorTypeDuplicate),
					"Field": Equal("spec.tolerations.whitelist[4]"),
				})),
				PointTo(MatchFields(IgnoreExtras, Fields{
					"Type":  Equal(field.ErrorTypeDuplicate),
					"Field": Equal("spec.tolerations.whitelist[6]"),
				})),
			))
		})

		It("should forbid using a default toleration which is not in the whitelist", func() {
			project.Spec.Tolerations = &core.ProjectTolerations{
				Defaults: []core.Toleration{{Key: "foo"}},
			}

			errorList := ValidateProject(project)

			Expect(errorList).To(ConsistOf(
				PointTo(MatchFields(IgnoreExtras, Fields{
					"Type":  Equal(field.ErrorTypeForbidden),
					"Field": Equal("spec.tolerations.defaults[0]"),
				})),
			))
		})

		DescribeTable("namespace immutability",
			func(old, new *string, matcher gomegatypes.GomegaMatcher) {
				project.Spec.Namespace = old
				newProject := prepareProjectForUpdate(project)
				newProject.Spec.Namespace = new

				errList := ValidateProjectUpdate(newProject, project)

				Expect(errList).To(matcher)
			},

			Entry("namespace change w/ preset namespace", pointer.StringPtr("garden-dev"), pointer.StringPtr("garden-core"), ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("spec.namespace"),
			})))),
			Entry("namespace change w/o preset namespace", nil, pointer.StringPtr("garden-core"), BeEmpty()),
			Entry("no change (both unset)", nil, nil, BeEmpty()),
			Entry("no change (same value)", pointer.StringPtr("garden-dev"), pointer.StringPtr("garden-dev"), BeEmpty()),
		)

		It("should forbid Project updates trying to change the createdBy field", func() {
			newProject := prepareProjectForUpdate(project)
			newProject.Spec.CreatedBy.Name = "some-other-user"

			errorList := ValidateProjectUpdate(newProject, project)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("spec.createdBy"),
			}))))
		})

		It("should forbid Project updates trying to change the createdBy field", func() {
			newProject := prepareProjectForUpdate(project)
			newProject.Spec.CreatedBy.Name = "some-other-user"

			errorList := ValidateProjectUpdate(newProject, project)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("spec.createdBy"),
			}))))
		})

		It("should forbid Project updates trying to reset the owner field", func() {
			newProject := prepareProjectForUpdate(project)
			newProject.Spec.Owner = nil

			errorList := ValidateProjectUpdate(newProject, project)

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("spec.owner"),
			}))))
		})
	})
})

func prepareProjectForUpdate(project *core.Project) *core.Project {
	p := project.DeepCopy()
	p.ResourceVersion = "1"
	return p
}
