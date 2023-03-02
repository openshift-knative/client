#!/usr/bin/env bash

readonly quarkus_version="2.13.7.Final-redhat-00003"
readonly quarkus_platform_group_id="com.redhat.quarkus.platform"
export EXTERNAL_LD_FLAGS="${EXTERNAL_LD_FLAGS:-} \
-X github.com/kiegroup/kie-tools/packages/kn-plugin-workflow/pkg/metadata.QuarkusVersion=${quarkus_version} \
-X github.com/kiegroup/kie-tools/packages/kn-plugin-workflow/pkg/metadata.QuarkusPlatformGroupId=${quarkus_platform_group_id}"
