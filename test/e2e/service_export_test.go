// Copyright 2020 The Knative Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build e2e && !eventing
// +build e2e,!eventing

package e2e

import (
	"encoding/json"
	"knative.dev/client/lib/test/e2e"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"

	"sigs.k8s.io/yaml"

	"knative.dev/client/lib/test"

	corev1 "k8s.io/api/core/v1"
	clientv1alpha1 "knative.dev/client/pkg/apis/client/v1alpha1"
	pkgtest "knative.dev/pkg/test"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"
	servingtest "knative.dev/serving/pkg/testing/v1"
)

func TestServiceExport(t *testing.T) {
	//FIXME: enable once 0.19 is available
	// see: https://github.com/knative/serving/pull/9685
	if strings.HasPrefix(os.Getenv("KNATIVE_SERVING_VERSION"), "0.18") {
		t.Skip("The test is skipped on Serving version 0.18")
	}
	t.Parallel()
	it, err := test.NewKnTest()
	assert.NilError(t, err)
	defer func() {
		assert.NilError(t, it.Teardown())
	}()

	r := test.NewKnRunResultCollector(t, it)
	defer r.DumpIfFailed()

	t.Log("create service with byo revision")
	serviceCreateWithOptions(r, "hello", "--revision-name", "rev1")

	userImage := pkgtest.ImagePath("helloworld")
	if strings.Contains(userImage, "@") {
		// Images specified by digest are not set in user-image annotation.
		userImage = ""
	}
	t.Log("export service-revision1 and compare")
	serviceExport(r, "hello", e2e.BuildServiceWithOptions("hello",
		servingtest.WithConfigSpec(e2e.BuildConfigurationSpec()),
		servingtest.WithBYORevisionName("hello-rev1"),
		e2e.WithRevisionAnnotations(map[string]string{"client.knative.dev/user-image": userImage}),
	), "--mode", "replay", "-o", "json")

	t.Log("update service - add env variable")
	e2e.ServiceUpdate(r, "hello", "--env", "a=mouse", "--revision-name", "rev2", "--no-lock-to-digest")
	serviceExport(r, "hello", e2e.BuildServiceWithOptions("hello",
		servingtest.WithConfigSpec(e2e.BuildConfigurationSpec()),
		servingtest.WithBYORevisionName("hello-rev2"),
		servingtest.WithEnv(corev1.EnvVar{Name: "a", Value: "mouse"}),
	), "--mode", "replay", "-o", "json")

	t.Log("export service-revision2 with kubernetes-resources")
	serviceExportWithServiceList(r, "hello", e2e.BuildServiceListWithOptions(
		e2e.WithService(e2e.BuildServiceWithOptions("hello",
			servingtest.WithConfigSpec(e2e.BuildConfigurationSpec()),
			servingtest.WithBYORevisionName("hello-rev2"),
			e2e.WithTrafficSpec([]string{"latest"}, []int{100}, []string{""}),
			servingtest.WithEnv(corev1.EnvVar{Name: "a", Value: "mouse"}),
		)),
	), "--with-revisions", "--mode", "replay", "-o", "yaml")

	t.Log("export service-revision2 with revisions-only")
	serviceExportWithRevisionList(r, "hello", e2e.BuildServiceWithOptions("hello",
		servingtest.WithConfigSpec(e2e.BuildConfigurationSpec()),
		servingtest.WithBYORevisionName("hello-rev2"),
		e2e.WithTrafficSpec([]string{"latest"}, []int{100}, []string{""}),
		servingtest.WithEnv(corev1.EnvVar{Name: "a", Value: "mouse"}),
	), e2e.BuildKNExportWithOptions(), "--with-revisions", "--mode", "export", "-o", "yaml")

	t.Log("update service with tag and split traffic")
	e2e.ServiceUpdate(r, "hello", "--tag", "hello-rev1=candidate", "--traffic", "candidate=2%,@latest=98%")

	t.Log("export service-revision2 after tagging kubernetes-resources")
	serviceExportWithServiceList(r, "hello", e2e.BuildServiceListWithOptions(
		e2e.WithService(e2e.BuildServiceWithOptions("hello",
			servingtest.WithConfigSpec(e2e.BuildConfigurationSpec()),
			servingtest.WithBYORevisionName("hello-rev1"),
			e2e.WithRevisionAnnotations(map[string]string{
				"client.knative.dev/user-image": userImage,
				"serving.knative.dev/routes":    "hello",
			}),
		)),
		e2e.WithService(e2e.BuildServiceWithOptions("hello",
			servingtest.WithConfigSpec(e2e.BuildConfigurationSpec()),
			servingtest.WithBYORevisionName("hello-rev2"),
			e2e.WithTrafficSpec([]string{"latest", "hello-rev1"}, []int{98, 2}, []string{"", "candidate"}),
			servingtest.WithEnv(corev1.EnvVar{Name: "a", Value: "mouse"}),
		)),
	), "--with-revisions", "--mode", "replay", "-o", "yaml")

	t.Log("export service-revision2 after tagging with revisions-only")
	serviceExportWithRevisionList(r, "hello", e2e.BuildServiceWithOptions("hello",
		servingtest.WithConfigSpec(e2e.BuildConfigurationSpec()),
		servingtest.WithBYORevisionName("hello-rev2"),
		e2e.WithTrafficSpec([]string{"latest", "hello-rev1"}, []int{98, 2}, []string{"", "candidate"}),
		servingtest.WithEnv(corev1.EnvVar{Name: "a", Value: "mouse"}),
	), e2e.BuildKNExportWithOptions(
		e2e.WithKNRevision(*(e2e.BuildRevision("hello-rev1",
			servingtest.WithRevisionAnn("client.knative.dev/user-image", userImage),
			servingtest.WithRevisionAnn("serving.knative.dev/routes", "hello"),
			servingtest.WithRevisionLabel("serving.knative.dev/configuration", "hello"),
			servingtest.WithRevisionLabel("serving.knative.dev/configurationGeneration", "1"),
			servingtest.WithRevisionLabel("serving.knative.dev/routingState", "active"),
			servingtest.WithRevisionLabel("serving.knative.dev/service", "hello"),
			e2e.WithRevisionImage(pkgtest.ImagePath("helloworld")),
		))),
	), "--with-revisions", "--mode", "export", "-o", "yaml")

	t.Log("update service - untag, add env variable, traffic split and system revision name")
	e2e.ServiceUpdate(r, "hello", "--untag", "candidate")
	e2e.ServiceUpdate(r, "hello", "--env", "b=cat", "--revision-name", "hello-rev3", "--traffic", "hello-rev1=30,hello-rev2=30,hello-rev3=40")

	t.Log("export service-revision3 with kubernetes-resources")
	serviceExportWithServiceList(r, "hello", e2e.BuildServiceListWithOptions(
		e2e.WithService(e2e.BuildServiceWithOptions("hello",
			servingtest.WithConfigSpec(e2e.BuildConfigurationSpec()),
			e2e.WithRevisionAnnotations(map[string]string{
				"client.knative.dev/user-image": userImage,
				"serving.knative.dev/routes":    "hello",
			}),
			servingtest.WithBYORevisionName("hello-rev1"),
		),
		),
		e2e.WithService(e2e.BuildServiceWithOptions("hello",
			servingtest.WithConfigSpec(e2e.BuildConfigurationSpec()),
			servingtest.WithBYORevisionName("hello-rev2"),
			e2e.WithRevisionAnnotations(map[string]string{
				"serving.knative.dev/routes": "hello",
			}),
			servingtest.WithEnv(corev1.EnvVar{Name: "a", Value: "mouse"}),
		),
		),
		e2e.WithService(e2e.BuildServiceWithOptions("hello",
			servingtest.WithConfigSpec(e2e.BuildConfigurationSpec()),
			servingtest.WithBYORevisionName("hello-rev3"),
			e2e.WithTrafficSpec([]string{"hello-rev1", "hello-rev2", "hello-rev3"}, []int{30, 30, 40}, []string{"", "", ""}),
			servingtest.WithEnv(corev1.EnvVar{Name: "a", Value: "mouse"}, corev1.EnvVar{Name: "b", Value: "cat"}),
		),
		)), "--with-revisions", "--mode", "replay", "-o", "yaml")

	t.Log("export service-revision3 with revisions-only")
	serviceExportWithRevisionList(r, "hello", e2e.BuildServiceWithOptions("hello",
		servingtest.WithConfigSpec(e2e.BuildConfigurationSpec()),
		e2e.WithTrafficSpec([]string{"hello-rev1", "hello-rev2", "hello-rev3"}, []int{30, 30, 40}, []string{"", "", ""}),
		servingtest.WithBYORevisionName("hello-rev3"),
		servingtest.WithEnv(corev1.EnvVar{Name: "a", Value: "mouse"}, corev1.EnvVar{Name: "b", Value: "cat"}),
	), e2e.BuildKNExportWithOptions(
		e2e.WithKNRevision(*(e2e.BuildRevision("hello-rev1",
			servingtest.WithRevisionAnn("client.knative.dev/user-image", userImage),
			servingtest.WithRevisionAnn("serving.knative.dev/routes", "hello"),
			servingtest.WithRevisionLabel("serving.knative.dev/configuration", "hello"),
			servingtest.WithRevisionLabel("serving.knative.dev/configurationGeneration", "1"),
			servingtest.WithRevisionLabel("serving.knative.dev/routingState", "active"),
			servingtest.WithRevisionLabel("serving.knative.dev/service", "hello"),
			e2e.WithRevisionImage(pkgtest.ImagePath("helloworld")),
		))),
		e2e.WithKNRevision(*(e2e.BuildRevision("hello-rev2",
			servingtest.WithRevisionAnn("serving.knative.dev/routes", "hello"),
			servingtest.WithRevisionLabel("serving.knative.dev/configuration", "hello"),
			servingtest.WithRevisionLabel("serving.knative.dev/configurationGeneration", "2"),
			servingtest.WithRevisionLabel("serving.knative.dev/routingState", "active"),
			servingtest.WithRevisionLabel("serving.knative.dev/service", "hello"),
			e2e.WithRevisionImage(pkgtest.ImagePath("helloworld")),
			e2e.WithRevisionEnv(corev1.EnvVar{Name: "a", Value: "mouse"}),
		))),
	), "--with-revisions", "--mode", "export", "-o", "yaml")

	t.Log("send all traffic to revision 2")
	e2e.ServiceUpdate(r, "hello", "--traffic", "hello-rev2=100")

	t.Log("export kubernetes-resources - all traffic to revision 2")
	serviceExportWithServiceList(r, "hello", e2e.BuildServiceListWithOptions(
		e2e.WithService(e2e.BuildServiceWithOptions("hello",
			servingtest.WithConfigSpec(e2e.BuildConfigurationSpec()),
			servingtest.WithBYORevisionName("hello-rev2"),
			e2e.WithRevisionAnnotations(map[string]string{
				"serving.knative.dev/routes": "hello",
			}),
			servingtest.WithEnv(corev1.EnvVar{Name: "a", Value: "mouse"}),
		),
		),
		e2e.WithService(e2e.BuildServiceWithOptions("hello",
			servingtest.WithConfigSpec(e2e.BuildConfigurationSpec()),
			e2e.WithTrafficSpec([]string{"hello-rev2"}, []int{100}, []string{""}),
			servingtest.WithBYORevisionName("hello-rev3"),
			servingtest.WithEnv(corev1.EnvVar{Name: "a", Value: "mouse"}, corev1.EnvVar{Name: "b", Value: "cat"}),
		),
		),
	), "--with-revisions", "--mode", "replay", "-o", "yaml")

	t.Log("export revisions-only - all traffic to revision 2")
	serviceExportWithRevisionList(r, "hello", e2e.BuildServiceWithOptions("hello",
		servingtest.WithConfigSpec(e2e.BuildConfigurationSpec()),
		servingtest.WithBYORevisionName("hello-rev3"),
		e2e.WithTrafficSpec([]string{"hello-rev2"}, []int{100}, []string{""}),
		servingtest.WithEnv(corev1.EnvVar{Name: "a", Value: "mouse"}, corev1.EnvVar{Name: "b", Value: "cat"}),
	), e2e.BuildKNExportWithOptions(
		e2e.WithKNRevision(*(e2e.BuildRevision("hello-rev2",
			servingtest.WithRevisionAnn("serving.knative.dev/routes", "hello"),
			servingtest.WithRevisionLabel("serving.knative.dev/configuration", "hello"),
			servingtest.WithRevisionLabel("serving.knative.dev/configurationGeneration", "2"),
			servingtest.WithRevisionLabel("serving.knative.dev/routingState", "active"),
			servingtest.WithRevisionLabel("serving.knative.dev/service", "hello"),
			e2e.WithRevisionImage(pkgtest.ImagePath("helloworld")),
			e2e.WithRevisionEnv(corev1.EnvVar{Name: "a", Value: "mouse"}),
		))),
	), "--with-revisions", "--mode", "export", "-o", "yaml")

	t.Log("create and export service 'foo' and verify that serviceUID and configurationUID labels are absent")
	serviceCreateWithOptions(r, "foo")
	output := serviceExportOutput(r, "foo", "-o", "json")
	actSvc := servingv1.Service{}
	err = json.Unmarshal([]byte(output), &actSvc)
	assert.NilError(t, err)
	_, ok := actSvc.Labels["serving.knative.dev/serviceUID"]
	assert.Equal(t, ok, false)
	_, ok = actSvc.Labels["serving.knative.dev/configurationUID"]
	assert.Equal(t, ok, false)
	_, ok = actSvc.Spec.ConfigurationSpec.Template.Labels["serving.knative.dev/servingUID"]
	assert.Equal(t, ok, false)
	_, ok = actSvc.Spec.ConfigurationSpec.Template.Labels["serving.knative.dev/configurationUID"]
	assert.Equal(t, ok, false)
}

// serviceExportOutput returns the export output of given service
func serviceExportOutput(r *test.KnRunResultCollector, serviceName string, options ...string) string {
	command := []string{"service", "export", serviceName}
	command = append(command, options...)
	out := r.KnTest().Kn().Run(command...)
	return out.Stdout
}

func serviceExport(r *test.KnRunResultCollector, serviceName string, expService *servingv1.Service, options ...string) {
	command := []string{"service", "export", serviceName}
	command = append(command, options...)
	out := r.KnTest().Kn().Run(command...)
	validateExportedService(r.T(), r.KnTest(), out.Stdout, expService)
	r.AssertNoError(out)
}

func serviceExportWithServiceList(r *test.KnRunResultCollector, serviceName string, expServiceList *servingv1.ServiceList, options ...string) {
	command := []string{"service", "export", serviceName}
	command = append(command, options...)
	out := r.KnTest().Kn().Run(command...)
	validateExportedServiceList(r.T(), r.KnTest(), out.Stdout, expServiceList)
	r.AssertNoError(out)
}

func serviceExportWithRevisionList(r *test.KnRunResultCollector, serviceName string, expService *servingv1.Service, knExport *clientv1alpha1.Export, options ...string) {
	command := []string{"service", "export", serviceName}
	command = append(command, options...)
	out := r.KnTest().Kn().Run(command...)
	validateExportedServiceandRevisionList(r.T(), r.KnTest(), out.Stdout, expService, knExport)
	r.AssertNoError(out)
}

func validateExportedService(t *testing.T, it *test.KnTest, out string, expService *servingv1.Service) {
	actSvc := servingv1.Service{}
	err := json.Unmarshal([]byte(out), &actSvc)
	assert.NilError(t, err)
	assert.DeepEqual(t, expService, &actSvc)
}

func validateExportedServiceList(t *testing.T, it *test.KnTest, out string, expServiceList *servingv1.ServiceList) {
	actSvcList := servingv1.ServiceList{}
	err := yaml.Unmarshal([]byte(out), &actSvcList)
	assert.NilError(t, err)
	assert.DeepEqual(t, expServiceList, &actSvcList)
}

func validateExportedServiceandRevisionList(t *testing.T, it *test.KnTest, out string, expService *servingv1.Service, knExport *clientv1alpha1.Export) {
	actSvc := clientv1alpha1.Export{}
	err := yaml.Unmarshal([]byte(out), &actSvc)
	assert.NilError(t, err)

	knExport.Spec.Service = *expService
	assert.DeepEqual(t, knExport, &actSvc)
}
