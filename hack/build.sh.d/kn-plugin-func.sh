#!/usr/bin/env bash

# There no midstream image currently, fallback to upstream one.
readonly func_util_image="${KN_PLUGIN_FUNC_UTIL_IMAGE:-quay.io/boson/alpine-socat:1.7.4.3-r1-non-root}"
export EXTERNAL_LD_FLAGS="${EXTERNAL_LD_FLAGS:-} \
-X knative.dev/kn-plugin-func/k8s.SocatImage=${func_util_image} \
-X knative.dev/kn-plugin-func/k8s.TarImage=${func_util_image}"

