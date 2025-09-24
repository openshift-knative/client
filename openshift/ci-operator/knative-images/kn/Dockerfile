# DO NOT EDIT! Generated Dockerfile for cmd/kn.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.23-openshift-4.19
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

ENV KN_PLUGIN_FUNC_UTIL_IMAGE=registry.redhat.io/openshift-serverless-1/kn-plugin-func-func-util-rhel8@sha256:82e69a1890042f5636ac7bb645b1aaeaba0a4894851eccc1fa1c63296918e31a
ENV KN_PLUGIN_EVENT_SENDER_IMAGE=registry.redhat.io/openshift-serverless-1/kn-plugin-event-sender-rhel8@sha256:0c2cdf917bbad1f5bca8a2a2ef12a96cfcd0c77c190cc35687eb57169b54ab6f
RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/kn

FROM $GO_RUNTIME

ARG VERSION=knative-v1.16

COPY --from=builder /usr/bin/main /ko-app/kn
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-client-kn-rhel8-container" \
      cpe="cpe:/a:redhat:openshift_serverless:1.36::el8" \
      name="openshift-serverless-1/kn-client-kn-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Client Kn" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Client Kn" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Client Kn" \
      io.k8s.description="Red Hat OpenShift Serverless Client Kn" \
      io.openshift.tags="kn"

ENTRYPOINT ["/ko-app/kn"]
