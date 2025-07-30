# DO NOT EDIT! Generated Dockerfile for cmd/kn.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.24-openshift-4.20
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

ENV KN_PLUGIN_FUNC_UTIL_IMAGE=registry.redhat.io/openshift-serverless-1/kn-plugin-func-func-util-rhel8@sha256:ebe3203b3e883057de5d54a434252e48517cf50b5d637ef0e90e11313317d39d
ENV KN_PLUGIN_EVENT_SENDER_IMAGE=registry.redhat.io/openshift-serverless-1/kn-plugin-event-sender-rhel8@sha256:f7795088777ea84fc6180b81b6131962944e34918e2c06671033a1a572581773
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
