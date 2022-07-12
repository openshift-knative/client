module knative.dev/client

go 1.16

require (
	github.com/google/go-cmp v0.5.7
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.10.1
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
	gotest.tools/v3 v3.1.0
	k8s.io/api v0.23.5
	k8s.io/apiextensions-apiserver v0.22.5
	k8s.io/apimachinery v0.23.5
	k8s.io/cli-runtime v0.22.5
	k8s.io/client-go v1.5.2
	k8s.io/code-generator v0.22.5
	knative.dev/eventing v0.29.2-0.20220420140829-ce4fe0990d23
	knative.dev/hack v0.0.0-20220128200847-51a42b2eb63e
	knative.dev/kn-plugin-event v1.2.0
	knative.dev/kn-plugin-func v0.24.0
	knative.dev/kn-plugin-source-kafka v0.29.0
	knative.dev/networking v0.0.0-20220120045035-ec849677a316
	knative.dev/pkg v0.0.0-20220222214439-083dd97300e1
	knative.dev/serving v0.29.5
	sigs.k8s.io/yaml v1.3.0
)

replace (
	// Tekton Triggers imports old google/cel-go, should be fixed with tektoncd/cli >=v0.24.x transitively
	github.com/google/cel-go => github.com/google/cel-go v0.11.2
	// update docker to be compatible with version used by pack and removes invalid pseudo-version
	github.com/openshift/source-to-image => github.com/boson-project/source-to-image v1.3.2

	// Enforce a version for CVE-2022-21698
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.11.1
	// Pin k8s.io dependencies to align with Knative and Tekton needs
	k8s.io/api => k8s.io/api v0.22.5
	k8s.io/apimachinery => k8s.io/apimachinery v0.22.5
	k8s.io/client-go => k8s.io/client-go v0.22.5

	// Points at: https://github.com/openshift-knative/kn-plugin-event/commit/1868b0441a3a2768724ffe5508ce0dcd6ded97c0
	knative.dev/kn-plugin-event => github.com/openshift-knative/kn-plugin-event v0.29.1-0.20220412122141-1868b0441a3a

	// Points at: https://github.com/openshift-knative/kn-plugin-func/commit/380e79884705dcc42ed849b1bd86102af14f886b
	knative.dev/kn-plugin-func => github.com/openshift-knative/kn-plugin-func v1.1.3-0.20220615075520-380e79884705
)
