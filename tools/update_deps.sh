#!/usr/bin/env bash

go mod tidy
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%protoc_gen_protobufjs_dependencies
