module knative.dev/client

go 1.21

require (
	github.com/google/go-cmp v0.6.0
	github.com/hashicorp/golang-lru v1.0.2
	github.com/hashicorp/hcl v1.0.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.15.0
	golang.org/x/mod v0.14.0
	golang.org/x/term v0.16.0
	gotest.tools/v3 v3.4.0
	k8s.io/api v0.28.5
	k8s.io/apiextensions-apiserver v0.28.5
	k8s.io/apimachinery v0.28.5
	k8s.io/cli-runtime v0.28.5
	k8s.io/client-go v0.28.5
	k8s.io/code-generator v0.27.10
	knative.dev/client-pkg v0.0.0-20240124090003-67fca0ca8681
	knative.dev/eventing v0.40.0
	knative.dev/hack v0.0.0-20240123162936-f3f03ac0ab1a
	knative.dev/kn-plugin-event v1.11.1
	knative.dev/kn-plugin-source-kafka v0.38.2
	knative.dev/networking v0.0.0-20240116081125-ce0738abf051
	knative.dev/pkg v0.0.0-20240116073220-b488e7be5902
	knative.dev/serving v0.40.0
	sigs.k8s.io/yaml v1.4.0
)

require (
	// Override version for display only purposes
	github.com/apache/incubator-kie-tools/packages/kn-plugin-workflow v0.33.0
	github.com/hashicorp/errwrap v1.1.0
	github.com/hashicorp/go-cleanhttp v0.5.2
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/AlecAivazis/survey/v2 v2.3.6 // indirect
	github.com/Azure/azure-sdk-for-go v68.0.0+incompatible // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.29 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.23 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.12 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.6 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20230828082145-3c4c8a2d2371 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230305170008-8188dc5388df // indirect
	github.com/apache/incubator-kie-kogito-serverless-operator/api v0.0.0 // indirect
	github.com/apache/incubator-kie-kogito-serverless-operator/workflowproj v0.0.0-20240126134605-2b5056d12bb9 // indirect
	github.com/apex/log v1.9.0 // indirect
	github.com/aws/aws-sdk-go-v2 v1.18.1 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.18.27 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.13.26 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.28 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.35 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecr v1.18.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.16.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.28 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.19.2 // indirect
	github.com/aws/smithy-go v1.13.5 // indirect
	github.com/awslabs/amazon-ecr-credential-helper/ecr-login v0.0.0-20230522190001-adf1bafd791a // indirect
	github.com/beevik/etree v1.2.0 // indirect
	github.com/buildpacks/imgutil v0.0.0-20230626185301-726f02e4225c // indirect
	github.com/buildpacks/lifecycle v0.17.0-rc.3 // indirect
	github.com/buildpacks/pack v0.30.0-pre3 // indirect
	github.com/chainguard-dev/git-urls v1.0.2 // indirect
	github.com/chrismellard/docker-credential-acr-env v0.0.0-20230304212654-82a0ddb27589 // indirect
	github.com/cloudflare/circl v1.3.3 // indirect
	github.com/containerd/containerd v1.7.2 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.14.3 // indirect
	github.com/containerd/typeurl v1.0.2 // indirect
	github.com/containers/image/v5 v5.26.1 // indirect
	github.com/containers/storage v1.48.0 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/docker/cli v24.0.2+incompatible // indirect
	github.com/docker/distribution v2.8.2+incompatible // indirect
	github.com/docker/docker v24.0.7+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.7.0 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dprotaso/go-yit v0.0.0-20220510233725-9ba8df137936 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/gdamore/tcell/v2 v2.6.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.5.0 // indirect
	github.com/go-git/go-git/v5 v5.11.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.15.4 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/glog v1.1.2 // indirect
	github.com/google/cel-go v0.16.1 // indirect
	github.com/google/go-github/v49 v49.1.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hako/durafmt v0.0.0-20210608085754-5c1018a4e16b // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.4 // indirect
	github.com/heroku/color v0.0.6 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/compress v1.16.6 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/manifestival/client-go-client v0.5.0 // indirect
	github.com/manifestival/manifestival v0.7.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/mitchellh/ioprogress v0.0.0-20180201004757-6a23b12fa88e // indirect
	github.com/moby/buildkit v0.11.6 // indirect
	github.com/moby/patternmatcher v0.5.0 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/moby/sys/sequential v0.5.0 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc4 // indirect
	github.com/opencontainers/runc v1.1.7 // indirect
	github.com/opencontainers/runtime-spec v1.1.0-rc.3 // indirect
	github.com/opencontainers/selinux v1.11.0 // indirect
	github.com/openshift-pipelines/pipelines-as-code v0.17.1 // indirect
	github.com/openshift/source-to-image v1.3.8 // indirect
	github.com/ory/viper v1.7.5 // indirect
	github.com/pb33f/libopenapi v0.10.1 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/relvacode/iso8601 v1.3.0 // indirect
	github.com/rivo/tview v0.0.0-20220307222120-9994674d60a8 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/sabhiram/go-gitignore v0.0.0-20210923224102-525f6e181f06 // indirect
	github.com/santhosh-tekuri/jsonschema/v5 v5.3.1 // indirect
	github.com/schollz/progressbar/v3 v3.13.1 // indirect
	github.com/senseyeio/duration v0.0.0-20180430131211-7c2a214ada46 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/serverlessworkflow/sdk-go/v2 v2.2.5 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/skeema/knownhosts v1.2.1 // indirect
	github.com/stoewer/go-strcase v1.2.1 // indirect
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635 // indirect
	github.com/tektoncd/cli v0.31.1 // indirect
	github.com/tektoncd/pipeline v0.47.0 // indirect
	github.com/tektoncd/triggers v0.23.1-0.20230420080448-bf603123cc0f // indirect
	github.com/thediveo/enumflag v0.10.0 // indirect
	github.com/vbatts/tar-split v0.11.3 // indirect
	github.com/vmware-labs/yaml-jsonpath v0.3.2 // indirect
	github.com/wavesoftware/go-commandline v1.0.0 // indirect
	github.com/wavesoftware/go-ensure v1.0.0 // indirect
	github.com/wavesoftware/go-retcode v1.0.0 // indirect
	github.com/xanzy/go-gitlab v0.83.0 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	k8s.io/component-base v0.28.1 // indirect
	sigs.k8s.io/controller-runtime v0.16.2 // indirect
)

require (
	contrib.go.opencensus.io/exporter/ocagent v0.7.1-0.20200907061046-05415f1de66d // indirect
	contrib.go.opencensus.io/exporter/prometheus v0.4.2 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr v1.4.10 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blendle/zapdriver v1.3.1 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cloudevents/sdk-go/sql/v2 v2.13.0 // indirect
	github.com/cloudevents/sdk-go/v2 v2.14.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.8.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.20.0 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/gnostic v0.6.9 // indirect
	github.com/google/go-containerregistry v0.15.2 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2 // indirect
	github.com/imdario/mergo v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/openzipkin/zipkin-go v0.4.2 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.46.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/prometheus/statsd_exporter v0.23.0 // indirect
	github.com/rickb777/date v1.20.1 // indirect
	github.com/rickb777/plural v1.4.1 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/xlab/treeprint v1.2.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.starlark.net v0.0.0-20230525235612-a134d8f9ddca // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/oauth2 v0.16.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.4.0 // indirect
	google.golang.org/api v0.155.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20231211222908-989df2bf70f3 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231211222908-989df2bf70f3 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231212172506-995d672761c0 // indirect
	google.golang.org/grpc v1.60.1 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/gengo v0.0.0-20221011193443-fad74ee6edd9 // indirect
	k8s.io/klog/v2 v2.100.1 // indirect
	k8s.io/kube-openapi v0.0.0-20230905202853-d090da108d2f // indirect
	knative.dev/eventing-kafka v0.38.1 // indirect
	knative.dev/func v1.13.0
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/kustomize/api v0.13.5-0.20230601165947-6ce0bf390ce3 // indirect
	sigs.k8s.io/kustomize/kyaml v0.14.3-0.20230601165947-6ce0bf390ce3 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.3.0 // indirect
)

replace (
	github.com/apache/incubator-kie-kogito-serverless-operator/api v0.0.0 => github.com/apache/incubator-kie-kogito-serverless-operator/api v0.0.0-20240126134605-2b5056d12bb9
	github.com/apache/incubator-kie-tools/packages/kn-plugin-workflow => github.com/kiegroup/kie-tools/packages/kn-plugin-workflow v1.32.1
	// Workflow plugin
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.16

	// Forked manifistival to support newer controller-runtime
	github.com/manifestival/client-go-client => github.com/manifestival/client-go-client v0.5.1-0.20240212170846-db471987b889
	// Knative deps
	k8s.io/api => k8s.io/api v0.27.10
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.27.10
	k8s.io/apimachinery => k8s.io/apimachinery v0.27.10
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.27.10
	k8s.io/client-go => k8s.io/client-go v0.27.10
	k8s.io/code-generator => k8s.io/code-generator v0.27.10
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20230501164219-8b0f38b5fd1f
	knative.dev/client-pkg => knative.dev/client-pkg v0.0.0-20231012115829-bcb06af7a827
	knative.dev/eventing => knative.dev/eventing v0.38.4
	knative.dev/func => github.com/openshift-knative/kn-plugin-func v1.1.3-0.20240213172456-3bd596fda969
	knative.dev/hack => knative.dev/hack v0.0.0-20230712131415-ddae80293c43
	knative.dev/kn-plugin-event => github.com/openshift-knative/kn-plugin-event v0.38.1-0.20231017003653-5728b23e8bab
	knative.dev/networking => knative.dev/networking v0.0.0-20231012062757-a5958051caf8
	knative.dev/pkg => knative.dev/pkg v0.0.0-20231023150739-56bfe0dd9626
	knative.dev/serving => knative.dev/serving v0.38.1
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.15.3
	sigs.k8s.io/kustomize/api => sigs.k8s.io/kustomize/api v0.13.2
	sigs.k8s.io/kustomize/kyaml => sigs.k8s.io/kustomize/kyaml v0.14.1
)
