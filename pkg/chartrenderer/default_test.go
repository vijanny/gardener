// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package chartrenderer_test

import (
	"path/filepath"

	"github.com/gardener/gardener/pkg/chartrenderer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/version"
)

const alpinePod = `apiVersion: v1
kind: Pod
metadata:
  name: alpine
  namespace: default
  labels:
    chartName: alpine
    chartVersion: "0.1.0"
spec:
  restartPolicy: Never
  containers:
  - name: waiter
    image: alpine:3.3
    command: ["/bin/sleep", "9000"]
`

var _ = Describe("ChartRenderer", func() {
	var (
		alpineChartPath = filepath.Join("testdata", "alpine")

		renderer chartrenderer.Interface
	)

	BeforeEach(func() {
		renderer = chartrenderer.NewWithServerVersion(&version.Info{})
	})

	Describe("#Render", func() {
		It("should return err when chartPath is missing", func() {
			_, err := renderer.Render(filepath.Join("testdata", "missing"), "missing", "default", map[string]string{})
			Expect(err).To(HaveOccurred())
		})

		It("should return rendered chart", func() {
			chart, err := renderer.Render(alpineChartPath, "alpine", "default", map[string]string{})
			Expect(err).ToNot(HaveOccurred())

			files := chart.Files()
			Expect(files).To(HaveLen(1))
			Expect(files).To(HaveKeyWithValue("alpine/templates/alpine-pod.yaml", alpinePod))
		})
	})

	Describe("#FileContent", func() {
		It("should return empty string when template file is missing", func() {
			chart, err := renderer.Render(alpineChartPath, "alpine", "default", map[string]string{})
			Expect(err).ToNot(HaveOccurred())

			actual := chart.FileContent("missing.yaml")
			Expect(actual).To(BeEmpty())
		})

		It("should return the file content when template file exists", func() {
			chart, err := renderer.Render(alpineChartPath, "alpine", "default", map[string]string{})
			Expect(err).ToNot(HaveOccurred())

			actual := chart.FileContent("alpine-pod.yaml")
			Expect(actual).To(Equal(alpinePod))
		})
	})
})
