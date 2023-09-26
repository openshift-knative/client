#!/usr/bin/env bash

readonly quarkus_version="2.13.8.SP2-redhat-00001"
readonly quarkus_platform_group_id="com.redhat.quarkus.platform"
readonly dev_mode_image="registry.redhat.io/openshift-serverless-1-tech-preview/logic-swf-devmode-rhel8:1.30"
export EXTERNAL_LD_FLAGS="${EXTERNAL_LD_FLAGS:-} \
-X github.com/kiegroup/kie-tools/packages/kn-plugin-workflow/pkg/metadata.DevModeImage=${dev_mode_image} \
-X github.com/kiegroup/kie-tools/packages/kn-plugin-workflow/pkg/metadata.QuarkusVersion=${quarkus_version} \
-X github.com/kiegroup/kie-tools/packages/kn-plugin-workflow/pkg/metadata.QuarkusPlatformGroupId=${quarkus_platform_group_id}"