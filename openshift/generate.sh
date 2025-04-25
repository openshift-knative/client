#!/usr/bin/env bash
#
# This script generates the productized Dockerfiles
#

set -o errexit
set -o nounset
set -o pipefail

repo_root_dir=$(dirname "$(realpath "${BASH_SOURCE[0]}")")/..

# --app-file-fmt is used to mimic ko build, it's assumed in --cmd flag tests
GOFLAGS='' go run github.com/openshift-knative/hack/cmd/generate@latest \
  --root-dir "${repo_root_dir}" \
  --generators dockerfile \
  --excludes ".*k8s\\.io.*" \
  --excludes ".*knative.dev/pkg/codegen.*" \
  --excludes ".*knative.dev/hack/cmd/script.*" \
  --app-file-fmt "/ko-app/%s"

#git apply $repo_root_dir/openshift/dockerfile.patch
FUNC_UTIL=$(skopeo inspect -n --format '{{.Digest}}' docker://quay.io/redhat-user-workloads/ocp-serverless-tenant/serverless-operator-136/kn-plugin-func-func-util:latest --override-os linux --override-arch amd64)
EVENT_SENDER=$(skopeo inspect -n --format '{{.Digest}}' docker://quay.io/redhat-user-workloads/ocp-serverless-tenant/serverless-operator-136/kn-plugin-event-sender:latest --override-os linux --override-arch amd64)

echo "func-util sha: ${FUNC_UTIL}"
echo "event-sender sha: ${EVENT_SENDER}"

sed -i "/RUN go build.*/ i \
ENV KN_PLUGIN_FUNC_UTIL_IMAGE=registry.redhat.io/openshift-serverless-1/kn-plugin-func-func-util-rhel8@${FUNC_UTIL}\n\
ENV KN_PLUGIN_EVENT_SENDER_IMAGE=registry.redhat.io/openshift-serverless-1/kn-plugin-event-sender-rhel8@${EVENT_SENDER}" openshift/ci-operator/knative-images/kn/Dockerfile
