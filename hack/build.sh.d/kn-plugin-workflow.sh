#!/usr/bin/env bash

# Env variables can be injected to override values during build phase
readonly quarkus_version="${LOGIC_QUARKUS_VERSION:-2.13.8.SP3-redhat-00001}"
readonly quarkus_platform_group_id="${LOGIC_QUARKUS_GROUP:-com.redhat.quarkus.platform}"
readonly dev_mode_image="${LOGIC_DEV_MODE_IMAGE:-registry.redhat.io/openshift-serverless-1-tech-preview/logic-swf-devmode-rhel8:1.31}"

export EXTERNAL_LD_FLAGS="${EXTERNAL_LD_FLAGS:-} \
-X github.com/kiegroup/kie-tools/packages/kn-plugin-workflow/pkg/metadata.DevModeImage=${dev_mode_image} \
-X github.com/kiegroup/kie-tools/packages/kn-plugin-workflow/pkg/metadata.QuarkusVersion=${quarkus_version} \
-X github.com/kiegroup/kie-tools/packages/kn-plugin-workflow/pkg/metadata.QuarkusPlatformGroupId=${quarkus_platform_group_id}"