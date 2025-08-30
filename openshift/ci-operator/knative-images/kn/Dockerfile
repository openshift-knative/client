# DO NOT EDIT! Generated Dockerfile for cmd/kn.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.24-openshift-4.20
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime
ENV GOFLAGS='-mod=mod'

ENV KN_PLUGIN_FUNC_UTIL_IMAGE=registry.redhat.io/openshift-serverless-1/kn-plugin-func-func-util-rhel8@sha256:598110e9b91c398e6f7ea6215a69f9562337968a024ded8e0f19736c5fb53457
ENV KN_PLUGIN_EVENT_SENDER_IMAGE=registry.redhat.io/openshift-serverless-1/kn-plugin-event-sender-rhel8@sha256:89752446cce6a89fc830e081bed6da50f8ad5ac2920fbe7a16c74a87fbf8f00e
RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/kn

FROM $GO_RUNTIME

ARG VERSION=

COPY --from=builder /usr/bin/main /ko-app/kn
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-kn-rhel8-container" \
      name="openshift-serverless-1/kn-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Kn" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Kn" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Kn" \
      io.k8s.description="Red Hat OpenShift Serverless Kn" \
      io.openshift.tags="kn"

ENTRYPOINT ["/ko-app/kn"]
