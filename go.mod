module knative.dev/client

go 1.16

require (
	github.com/emicklei/go-restful/v3 v3.9.0 // indirect
	github.com/google/go-cmp v0.5.7
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.10.1
	golang.org/x/mod v0.5.1
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
	gotest.tools/v3 v3.1.0
	k8s.io/api v0.23.5
	k8s.io/apiextensions-apiserver v0.22.5
	k8s.io/apimachinery v0.23.5
	k8s.io/cli-runtime v0.22.5
	k8s.io/client-go v1.5.2
	k8s.io/code-generator v0.22.5
	knative.dev/eventing v0.30.4-0.20220601140224-ec844245369e
	knative.dev/hack v0.0.0-20220629135030-82afbd367967
	knative.dev/kn-plugin-event v0.30.1
	knative.dev/kn-plugin-func v0.25.0
	knative.dev/kn-plugin-source-kafka v0.30.0
	knative.dev/networking v0.0.0-20220302134042-e8b2eb995165
	knative.dev/pkg v0.0.0-20220314165618-a637a96a1bd9
	knative.dev/serving v0.30.3-0.20220630163432-b5f775538f3d
	sigs.k8s.io/yaml v1.3.0
)

replace (
	// Enforce a version for CVE-2022-1996
	github.com/emicklei/go-restful => github.com/emicklei/go-restful/v3 v3.8.0
	// Tekton Triggers imports old google/cel-go, should be fixed with tektoncd/cli >=v0.24.x transitively
	github.com/google/cel-go => github.com/google/cel-go v0.11.2
	// update docker to be compatible with version used by pack and removes invalid pseudo-version
	github.com/openshift/source-to-image => github.com/boson-project/source-to-image v1.3.2
	// Pin k8s.io dependencies to align with Knative and Tekton needs
	k8s.io/api => k8s.io/api v0.22.5
	k8s.io/apimachinery => k8s.io/apimachinery v0.22.5
	k8s.io/client-go => k8s.io/client-go v0.22.5

	// Points at: https://github.com/openshift-knative/kn-plugin-event/commit/059d9ada55c23b1c5fa8797d22be308899020d38
	knative.dev/kn-plugin-event => github.com/openshift-knative/kn-plugin-event v0.29.1-0.20220719090345-059d9ada55c2
	// Points at: https://github.com/openshift-knative/kn-plugin-func/commit/4a318c18b944966fefa320324424f2b23b48a74d
	knative.dev/kn-plugin-func => github.com/openshift-knative/kn-plugin-func v1.1.3-0.20220714121505-4a318c18b944

)
