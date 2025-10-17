# Not generated yet!
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-9-release-golang-1.23-openshift-4.19
ARG GO_RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal

# ---------- BUILD STAGE ----------
FROM $GO_BUILDER AS builder
WORKDIR /workspace
# Allow copying from higher directory when context is limited
COPY . .
ENV GO_COMPLIANCE_DEBUG=1
ENV KN_PLUGIN_FUNC_UTIL_IMAGE=registry.redhat.io/openshift-serverless-1/kn-plugin-func-func-util-rhel9@sha256:a181b639a19b3c36367eec41e5ac0e030bcfcb247b96e6843ff60a9561ead85dn
ENV KN_PLUGIN_EVENT_SENDER_IMAGE=registry.redhat.io/openshift-serverless-1/kn-plugin-event-sender-rhel9@sha256:d28cd64bfc1d0978d6f95e8b70d5bf7c36d6b92c0d60460f7ac6fa2b2ba0a15b
# Build binaries
# RUN make build-cross-package
RUN ./hack/build.sh -p linux amd64
RUN ./hack/build.sh -p linux ppc64le
RUN ./hack/build.sh -p linux s390x
RUN ./hack/build.sh -p linux arm64
RUN ./hack/build.sh -p darwin amd64
RUN ./hack/build.sh -p darwin arm64
RUN ./hack/build.sh -p windows amd64

RUN chmod +x kn-linux-amd64 kn-linux-ppc64le kn-linux-s390x kn-windows-amd64.exe kn-darwin-amd64 kn-linux-arm64 kn-darwin-arm64

## ---------- PACKAGING STAGE ----------
RUN dnf install -y tar zip

RUN tar --transform='flags=r;s|kn-linux-amd64|kn|' -zcf kn-linux-amd64.tar.gz kn-linux-amd64 LICENSE \
  && tar --transform='flags=r;s|kn-linux-ppc64le|kn|' -zcf kn-linux-ppc64le.tar.gz kn-linux-ppc64le LICENSE \
  && tar --transform='flags=r;s|kn-linux-s390x|kn|' -zcf kn-linux-s390x.tar.gz kn-linux-s390x LICENSE \
  && tar --transform='flags=r;s|kn-darwin-amd64|kn|' -zcf kn-macos-amd64.tar.gz kn-darwin-amd64 LICENSE \
  && tar --transform='flags=r;s|kn-linux-arm64|kn|' -zcf kn-linux-arm64.tar.gz kn-linux-arm64 LICENSE \
  && tar --transform='flags=r;s|kn-darwin-arm64|kn|' -zcf kn-macos-arm64.tar.gz kn-darwin-arm64 LICENSE

RUN mkdir "windows" && mv kn-windows-amd64.exe ./windows/kn.exe && cp LICENSE ./windows/ && zip --quiet --junk-path - windows/* > kn-windows-amd64.zip

# ---------- FINAL RUNTIME IMAGE ----------
FROM $GO_RUNTIME
RUN  mkdir -p /usr/share/kn/{linux_amd64,linux_arm64,linux_ppc64le,linux_s390x,macos_amd64,macos_arm64,windows}

COPY --from=builder /workspace/kn-linux-amd64.tar.gz /usr/share/kn/linux_amd64/
COPY --from=builder /workspace/kn-linux-ppc64le.tar.gz /usr/share/kn/linux_ppc64le/
COPY --from=builder /workspace/kn-linux-s390x.tar.gz /usr/share/kn/linux_s390x/
COPY --from=builder /workspace/kn-macos-amd64.tar.gz /usr/share/kn/macos_amd64/
COPY --from=builder /workspace/kn-windows-amd64.zip  /usr/share/kn/windows/
COPY --from=builder /workspace/kn-linux-arm64.tar.gz /usr/share/kn/linux_arm64/
COPY --from=builder /workspace/kn-macos-arm64.tar.gz /usr/share/kn/macos_arm64/
COPY --from=builder /workspace/LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-client-cli-artifacts-rhel9-container" \
      name="openshift-serverless-1/kn-client-cli-artifacts-rhel9" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Client Cli Artifacts" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Client Cli Artifacts" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Client Cli Artifacts" \
      io.k8s.description="Red Hat OpenShift Serverless Client Cli Artifacts" \
      io.openshift.tags="cli-artifacts" \
      release=$VERSION \
      vendor="Red Hat, Inc." \
      cpe="cpe:/a:redhat:openshift_serverless:1.37::el9"
