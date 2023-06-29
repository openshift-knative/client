#!/usr/bin/env bash

readonly quarkus_version="2.13.7.SP1-redhat-00001"
readonly quarkus_platform_group_id="com.redhat.quarkus.platform"
readonly dev_mode_image="quay.io/kiegroup/kogito-swf-devmode:1.40"
export EXTERNAL_LD_FLAGS="${EXTERNAL_LD_FLAGS:-} \
-X github.com/kiegroup/kie-tools/packages/kn-plugin-workflow/pkg/metadata.DevModeImage=${dev_mode_image} \
-X github.com/kiegroup/kie-tools/packages/kn-plugin-workflow/pkg/metadata.QuarkusVersion=${quarkus_version} \
-X github.com/kiegroup/kie-tools/packages/kn-plugin-workflow/pkg/metadata.QuarkusPlatformGroupId=${quarkus_platform_group_id}"