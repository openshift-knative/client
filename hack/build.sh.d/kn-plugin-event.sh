#!/usr/bin/env bash

knEventVersion="$(grep 'knative.dev/kn-plugin-event ' "$(basedir)/go.mod" \
  | head -n 1 \
  | sed -s 's|// indirect||' \
  | sed -sE 's/.* (v[0-9]+\.[0-9]+\.[0-9]+.*)/\1/')"
knEventRelease="${knEventVersion%.*}"
knEventRelease="${knEventRelease#v}"
readonly upstream_sender_image="registry.ci.openshift.org/knative/release-${knEventRelease}:client-plugin-event-sender"
readonly sender_image="${KN_PLUGIN_EVENT_SENDER_IMAGE:-${upstream_sender_image}}"
export EXTERNAL_LD_FLAGS="${EXTERNAL_LD_FLAGS:-} \
-X knative.dev/kn-plugin-event/pkg/metadata.Image=${sender_image} \
-X knative.dev/kn-plugin-event/pkg/metadata.Version=${knEventVersion}"
