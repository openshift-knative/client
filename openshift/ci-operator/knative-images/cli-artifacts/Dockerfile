# This is not generated Dockerfile, yet!
ARG CLI_ARTIFACTS=quay.io/redhat-user-workloads/ocp-serverless-tenant/serverless-operator-135/kn-client-cli-artifacts:latest

FROM $CLI_ARTIFACTS

ARG VERSION=knative-v1.15

COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-kn-cli-artifacts-rhel8-container" \
      name="openshift-serverless-1/kn-cli-artifacts-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 kn Client artifacts" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 kn Client artifacts" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 kn Client artifacts" \
      io.openshift.tags="kn-client-cli-artifacts"
