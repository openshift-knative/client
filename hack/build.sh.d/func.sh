#!/usr/bin/env bash

# There no midstream image currently, fallback to upstream one.
readonly func_util_image="${KN_PLUGIN_FUNC_UTIL_IMAGE:-quay.io/boson/alpine-socat:1.7.4.3-r1-non-root}"
readonly default_node_builder_image="${DEFAULT_NODE_BUILDER_IMAGE:-registry.access.redhat.com/ubi8/nodejs-16}"
readonly default_quarkus_builder_image="${DEFAULT_QUARKUS_BUILDER_IMAGE:-registry.access.redhat.com/ubi8/openjdk-17}"
readonly default_python_builder_image="${DEFAULT_PYTHON_BUILDER_IMAGE:-registry.access.redhat.com/ubi8/python-39}"

export EXTERNAL_LD_FLAGS="${EXTERNAL_LD_FLAGS:-} \
-X knative.dev/func/pkg/k8s.SocatImage=${func_util_image} \
-X knative.dev/func/pkg/k8s.TarImage=${func_util_image} \
-X knative.dev/func/pkg/builders/s2i.DefaultNodeBuilder=${default_node_builder_image} \
-X knative.dev/func/pkg/builders/s2i.DefaultQuarkusBuilder=${default_quarkus_builder_image} \
-X knative.dev/func/pkg/builders/s2i.DefaultPythonBuilder=${default_python_builder_image}"