module knative.dev/client

go 1.16

require (
	github.com/google/go-cmp v0.5.7
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.10.1
	golang.org/x/mod v0.5.1
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
	gotest.tools/v3 v3.0.3
	k8s.io/api v0.22.5
	k8s.io/apiextensions-apiserver v0.22.5
	k8s.io/apimachinery v0.22.5
	k8s.io/cli-runtime v0.22.5
	k8s.io/client-go v0.22.5
	k8s.io/code-generator v0.22.5
	knative.dev/eventing v0.30.0
	knative.dev/hack v0.0.0-20220224013837-e1785985d364
	knative.dev/kn-plugin-event v0.30.1
	knative.dev/kn-plugin-source-kafka v0.30.0
	knative.dev/networking v0.0.0-20220302134042-e8b2eb995165
	knative.dev/pkg v0.0.0-20220310195447-38af013b30ff
	knative.dev/serving v0.30.0
	sigs.k8s.io/yaml v1.3.0
)

// Points at: https://github.com/openshift-knative/kn-plugin-event/commit/e518a460bb1ed73562dfd64aa6519f21b3cd9e2e
replace knative.dev/kn-plugin-event => github.com/openshift-knative/kn-plugin-event v0.29.1-0.20220712135411-e518a460bb1e
