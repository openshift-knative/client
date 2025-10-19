# DO NOT EDIT! Generated Dockerfile for cmd/kn.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-9-release-golang-1.23-openshift-4.19
ARG GO_RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

ENV KN_PLUGIN_FUNC_UTIL_IMAGE=registry.redhat.io/openshift-serverless-1/kn-plugin-func-func-util-rhel9@sha256:607e50dd3a9b61d57d125c47a7977c40c4f781bfab01a852a08b40fe42f2b893
ENV KN_PLUGIN_EVENT_SENDER_IMAGE=registry.redhat.io/openshift-serverless-1/kn-plugin-event-sender-rhel9@sha256:2d8cfab33f1c46b2d4510ea4b56a09e12b8bcf07c0eb055e84ebc077ec730145
RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/kn

FROM $GO_RUNTIME

ARG VERSION=knative-v1.17

COPY --from=builder /usr/bin/main /ko-app/kn
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-client-kn-rhel9-container" \
      name="openshift-serverless-1/client-kn-rhel9" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Client Kn" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Client Kn" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Client Kn" \
      io.k8s.description="Red Hat OpenShift Serverless Client Kn" \
      io.openshift.tags="kn" \
      cpe="cpe:/a:redhat:openshift_serverless:1.37::el9"

ENTRYPOINT ["/ko-app/kn"]
