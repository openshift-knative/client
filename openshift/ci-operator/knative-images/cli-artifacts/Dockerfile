# This is not generated Dockerfile, yet!
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal
ARG CLI_ARTIFACTS=registry.redhat.io/openshift-serverless-1/kn-cli-artifacts-rhel8:1.16

FROM $CLI_ARTIFACTS as builder

FROM $GO_RUNTIME

ARG VERSION=knative-v1.18

RUN mkdir -p /usr/share/kn

COPY --from=builder /usr/share/kn /usr/share/kn
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
