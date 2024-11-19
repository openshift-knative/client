# DO NOT EDIT! Generated Dockerfile for cmd/kn.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal
ARG IMAGE_KN_PLUGIN_FUNC_UTIL=registry.redhat.io/openshift-serverless-1/kn-plugin-func-func-util-rhel8@sha256:134a07b2ecd8bb485d7f20bf23e863f5cec5620134c4a26b65e31d23c3076175
ARG IMAGE_KN_PLUGIN_EVENT_SENDER=registry.redhat.io/openshift-serverless-1/kn-plugin-event-sender-rhel8@sha256:e2ea8370284a4807712849a0ec8c786bbf7cb6cd78dc2309c14e87470df235e6

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime
ENV KN_PLUGIN_FUNC_UTIL_IMAGE=$IMAGE_KN_PLUGIN_FUNC_UTIL
ENV KN_PLUGIN_EVENT_SENDER_IMAGE=$IMAGE_KN_PLUGIN_EVENT_SENDER

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/kn

FROM $GO_RUNTIME

ARG VERSION=knative-v1.15

COPY --from=builder /usr/bin/main /ko-app/kn
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-client-kn-rhel8-container" \
      name="openshift-serverless-1/client-kn-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Client Kn" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Client Kn" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Client Kn" \
      io.k8s.description="Red Hat OpenShift Serverless Client Kn" \
      io.openshift.tags="kn"

ENTRYPOINT ["/ko-app/kn"]
