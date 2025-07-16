# This is not generated Dockerfile, yet!
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17

FROM $GO_BUILDER as builder

WORKDIR /go/src/github.com/knative/client
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime
# Pass to cross compile target
ENV GOFLAGS="-tags=strictfipsruntime"
ENV KN_PLUGIN_FUNC_UTIL_IMAGE=registry.redhat.io/openshift-serverless-1/kn-plugin-func-func-util-rhel8@sha256:fc5e614c54408c41f13c0bbb2467665df50f7ef7fd025652395198eb35d91e2e
ENV KN_PLUGIN_EVENT_SENDER_IMAGE=registry.redhat.io/openshift-serverless-1/kn-plugin-event-sender-rhel8@sha256:58e3ccea5c558575fe4adc6f5f6d570ebec0934de1ba1110bc32fe76fb379c7e

RUN make build-cross-package

FROM $GO_RUNTIME

ARG VERSION=knative-v1.17

RUN mkdir -p /usr/share/kn

COPY --from=builder /go/src/github.com/knative/client/kn-linux-amd64.tar.gz /usr/share/kn/
COPY --from=builder /go/src/github.com/knative/client/kn-linux-arm64.tar.gz /usr/share/kn/
COPY --from=builder /go/src/github.com/knative/client/kn-linux-ppc64le.tar.gz /usr/share/kn/
COPY --from=builder /go/src/github.com/knative/client/kn-linux-s390x.tar.gz /usr/share/kn/
COPY --from=builder /go/src/github.com/knative/client/kn-macos-amd64.tar.gz /usr/share/kn/
COPY --from=builder /go/src/github.com/knative/client/kn-macos-arm64.tar.gz /usr/share/kn/
COPY --from=builder /go/src/github.com/knative/client/kn-windows-amd64.zip  /usr/share/
COPY --from=builder /go/src/github.com/knative/client/LICENSE  /usr/share/kn/LICENSE
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-client-cli-artifacts-rhel8-container" \
      name="openshift-serverless-1/client-cli-artifacts-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Client Cli Artifacts" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Client Cli Artifacts" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Client Cli Artifacts" \
      io.k8s.description="Red Hat OpenShift Serverless Client Cli Artifacts" \
      io.openshift.tags="cli-artifacts"
