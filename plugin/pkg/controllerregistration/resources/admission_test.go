// SPDX-FileCopyrightText: 2019 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package resources_test

import (
	"context"

	"github.com/gardener/gardener/pkg/apis/core"
	"github.com/gardener/gardener/pkg/client/core/clientset/internalversion/fake"
	. "github.com/gardener/gardener/plugin/pkg/controllerregistration/resources"
	. "github.com/gardener/gardener/test/gomega"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/client-go/testing"
	"k8s.io/utils/pointer"
)

var _ = Describe("resources", func() {
	Describe("#Admit", func() {
		var (
			controllerRegistration core.ControllerRegistration

			attrs            admission.Attributes
			admissionHandler *Resources

			coreClient *fake.Clientset

			resourceKind = "Foo"
			resourceType = "bar"
		)

		BeforeEach(func() {
			admissionHandler, _ = New()
			admissionHandler.AssignReadyFunc(func() bool { return true })

			coreClient = &fake.Clientset{}
			admissionHandler.SetInternalCoreClientset(coreClient)

			controllerRegistration = core.ControllerRegistration{
				ObjectMeta: metav1.ObjectMeta{
					Name: "dummy",
				},
				Spec: core.ControllerRegistrationSpec{
					Resources: []core.ControllerResource{
						{
							Kind:    resourceKind,
							Type:    resourceType,
							Primary: pointer.BoolPtr(true),
						},
					},
				},
			}
		})

		It("should do nothing because the resource is not ControllerRegistration", func() {
			attrs = admission.NewAttributesRecord(nil, nil, core.Kind("SomeOtherResource").WithVersion("version"), "", controllerRegistration.Name, core.Resource("some-other-resource").WithVersion("version"), "", admission.Create, &metav1.CreateOptions{}, false, nil)

			err := admissionHandler.Validate(context.TODO(), attrs, nil)

			Expect(err).NotTo(HaveOccurred())
		})

		It("should allow the object because no other resource in the system uses the kind/type combination", func() {
			attrs = admission.NewAttributesRecord(&controllerRegistration, nil, core.Kind("ControllerRegistration").WithVersion("version"), "", controllerRegistration.Name, core.Resource("controllerregistrations").WithVersion("version"), "", admission.Create, &metav1.CreateOptions{}, false, nil)

			err := admissionHandler.Validate(context.TODO(), attrs, nil)

			Expect(err).NotTo(HaveOccurred())
		})

		It("should not deny the object because it is updated", func() {
			attrs = admission.NewAttributesRecord(&controllerRegistration, &controllerRegistration, core.Kind("ControllerRegistration").WithVersion("version"), "", controllerRegistration.Name, core.Resource("controllerregistrations").WithVersion("version"), "", admission.Update, &metav1.UpdateOptions{}, false, nil)

			coreClient.AddReactor("list", "controllerregistrations", func(action testing.Action) (bool, runtime.Object, error) {
				return true, &core.ControllerRegistrationList{
					Items: []core.ControllerRegistration{controllerRegistration},
				}, nil
			})

			err := admissionHandler.Validate(context.TODO(), attrs, nil)

			Expect(err).NotTo(HaveOccurred())
		})

		It("should deny the object because another resource in the system uses the kind/type combination", func() {
			attrs = admission.NewAttributesRecord(&controllerRegistration, nil, core.Kind("ControllerRegistration").WithVersion("version"), "", controllerRegistration.Name, core.Resource("controllerregistrations").WithVersion("version"), "", admission.Create, &metav1.CreateOptions{}, false, nil)

			controllerRegistration2 := controllerRegistration.DeepCopy()
			controllerRegistration2.Name = "another-name"

			coreClient.AddReactor("list", "controllerregistrations", func(action testing.Action) (bool, runtime.Object, error) {
				return true, &core.ControllerRegistrationList{
					Items: []core.ControllerRegistration{*controllerRegistration2},
				}, nil
			})

			err := admissionHandler.Validate(context.TODO(), attrs, nil)

			Expect(err).To(BeForbiddenError())
		})

		It("should allow the object because another resource in the system  declared the kind/type combination as secondary only", func() {
			attrs = admission.NewAttributesRecord(&controllerRegistration, nil, core.Kind("ControllerRegistration").WithVersion("version"), "", controllerRegistration.Name, core.Resource("controllerregistrations").WithVersion("version"), "", admission.Create, &metav1.CreateOptions{}, false, nil)

			controllerRegistration2 := controllerRegistration.DeepCopy()
			controllerRegistration2.Name = "another-name"
			controllerRegistration2.Spec.Resources[0].Primary = pointer.BoolPtr(false)

			coreClient.AddReactor("list", "controllerregistrations", func(action testing.Action) (bool, runtime.Object, error) {
				return true, &core.ControllerRegistrationList{
					Items: []core.ControllerRegistration{*controllerRegistration2},
				}, nil
			})

			err := admissionHandler.Validate(context.TODO(), attrs, nil)

			Expect(err).To(Succeed())
		})
	})

	Describe("#Register", func() {
		It("should register the plugin", func() {
			plugins := admission.NewPlugins()
			Register(plugins)

			registered := plugins.Registered()
			Expect(registered).To(HaveLen(1))
			Expect(registered).To(ContainElement(PluginName))
		})
	})

	Describe("#NewFactory", func() {
		It("should create a new PluginFactory", func() {
			f, err := NewFactory(nil)

			Expect(f).NotTo(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("#New", func() {
		It("should only handle CREATE or UPDATE operations", func() {
			dr, err := New()

			Expect(err).ToNot(HaveOccurred())
			Expect(dr.Handles(admission.Create)).To(BeTrue())
			Expect(dr.Handles(admission.Update)).To(BeTrue())
			Expect(dr.Handles(admission.Connect)).NotTo(BeTrue())
			Expect(dr.Handles(admission.Delete)).NotTo(BeTrue())
		})
	})

	Describe("#ValidateInitialization", func() {
		It("should return no error", func() {
			dr, _ := New()

			err := dr.ValidateInitialization()

			Expect(err).NotTo(HaveOccurred())
		})
	})
})
