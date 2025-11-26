#!/usr/bin/env bash

# Copyright 2019 The OpenShift Knative Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

readonly ROOT_DIR=$(dirname $0)/..

source "$(go run knative.dev/hack/cmd/script library.sh)"
source "$(go run knative.dev/hack/cmd/script e2e-tests.sh)"

readonly E2E_TIMEOUT="60m"

# Determine if we're running locally or in CI.
if [ -n "$OPENSHIFT_BUILD_NAMESPACE" ]; then
  readonly TEST_IMAGE_TEMPLATE=$(cat <<-END
{{- with .Name }}
{{- if eq . "helloworld"}}$KNATIVE_CLIENT_TEST_HELLOWORLD{{end -}}
{{- if eq . "grpc-ping"}}$KNATIVE_CLIENT_TEST_GRPC_PING{{end -}}
{{- if eq . "servingcontainer"}}$KNATIVE_CLIENT_TEST_SERVINGCONTAINER{{end -}}
{{- if eq . "sidecarcontainer"}}$KNATIVE_CLIENT_TEST_SIDECARCONTAINER{{end -}}
{{end -}}
END
)
elif [ -n "$DOCKER_REPO_OVERRIDE" ]; then
  readonly TEST_IMAGE_TEMPLATE="${DOCKER_REPO_OVERRIDE}/{{.Name}}"
elif [ -n "$BRANCH" ]; then
  readonly TEST_IMAGE_TEMPLATE="registry.ci.openshift.org/openshift/${BRANCH}:knative-client-test-{{.Name}}"
elif [ -n "$TEMPLATE" ]; then
  readonly TEST_IMAGE_TEMPLATE="$TEMPLATE"
else
  readonly TEST_IMAGE_TEMPLATE="registry.ci.openshift.org/openshift/knative-nightly:knative-client-test-{{.Name}}"
fi

env

build_knative_client() {
  failed=0
  # run this cross platform build to ensure all the checks pass (as this is done while building artifacts)
  ./hack/build.sh -x || failed=1

  if [[ $failed -eq 0 ]]; then
    mv kn-linux-$(go env GOARCH) kn
  fi

  return $failed
}

run_unit_tests() {
  failed=0
  go test -v ./cmd/... ./pkg/... || failed=1
  return $failed
}

run_sobranch() {
  go run github.com/openshift-knative/hack/cmd/sobranch@latest "$@"
}

run_client_e2e_tests(){
  local tags=$1
  local test_name=${2:-}

  header "Running e2e tests"
  local failed=0
  # Add local dir to have access to built kn
  export PATH=$PATH:${REPO_ROOT_DIR}
  export GO111MODULE=on
  # In CI environment GOFLAGS is set to '-mod=vendor', unsetting it and providing explicit flag below
  # while invoking go e2e tests. Unsetting to keep using -mod=vendor irrespective of whether GOFLAGS is set or not.
  # Ideally this should be overridden but see https://github.com/golang/go/issues/35827
  unset GOFLAGS

  # Add anyuid scc to all authenticated users so e2e tests for --user flag can user any user id
  oc adm policy add-scc-to-group anyuid system:authenticated

  local run_append=""
  if [ -n "${test_name}" ]; then
    run_append="-run ^(${test_name})$"
  fi
  if [ -n "${tags}" ]; then
    run_append="${run_append} -tags e2e,${tags}"
  else
    run_append="${run_append} -tags e2e"
  fi

  go test \
    ./test/e2e \
    -v -timeout=$E2E_TIMEOUT \
    --imagetemplate "$TEST_IMAGE_TEMPLATE" \
    ${run_append} || failed=$?

  return $failed
}

install_serverless_operator() {
  local repository="https://github.com/openshift-knative/serverless-operator.git"
  local project_tag release so_branch
  project_tag=$(yq r "${ROOT_DIR}/openshift/project.yaml" project.tag)
  release=${project_tag/knative-/}
  so_branch=$(run_sobranch --upstream-version "${release}")

  if ! git ls-remote --heads --exit-code "$repository" "$so_branch" &>/dev/null; then
      echo "Release branch doesn't exist yet, using main"
      so_branch="main"
  fi

  local operator_dir=/tmp/serverless-operator
  local failed=0
  header "Installing serverless operator from openshift-knative/serverless-operator branch $so_branch"
  rm -rf $operator_dir
  git clone --branch "${so_branch}" $repository $operator_dir || failed=1
  pushd $operator_dir

  export SKIP_MESH_AUTH_POLICY_GENERATION=true
  export ON_CLUSTER_BUILDS=true
  export DOCKER_REPO_OVERRIDE=image-registry.openshift-image-registry.svc:5000/openshift-marketplace
  if [ "${project_tag}" == "knative-nightly" ]; then
    USE_IMAGE_RELEASE_TAG="${project_tag}"
    export USE_IMAGE_RELEASE_TAG
    make generated-files
  fi

  make images install-serving install-eventing || failed=1
  subheader "Successfully installed serverless operator."
  
  # Workaround default 'https' scheme
  oc patch knativeserving knative-serving \
    --namespace knative-serving --type merge \
    --patch '{"spec":{"config":{"network":{"default-external-scheme":"http"}}}}' || return 1

  popd
  return $failed
}
