load("@aspect_rules_ts//ts:defs.bzl", "ts_config")
load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@npm//:defs.bzl", "npm_link_all_packages")

# gazelle:prefix github.com/buildbuddy-io/protoc-gen-protobufjs
gazelle(name = "gazelle")

go_library(
    name = "protoc-gen-protobufjs_lib",
    srcs = [
        "codegen.go",
        "descriptor.go",
        "fieldtypes.go",
        "flags.go",
        "index.go",
        "logging.go",
        "main.go",
        "ts.go",
        "util.go",
        "wire.go",
    ],
    importpath = "github.com/buildbuddy-io/protoc-gen-protobufjs",
    visibility = ["//visibility:private"],
    deps = [
        "@org_golang_google_protobuf//encoding/protojson",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//types/descriptorpb",
        "@org_golang_google_protobuf//types/pluginpb",
    ],
)

go_binary(
    name = "protoc-gen-protobufjs",
    embed = [":protoc-gen-protobufjs_lib"],
    visibility = ["//visibility:public"],
)

npm_link_all_packages(name = "node_modules")

ts_config(
    name = "tsconfig",
    src = "tsconfig.json",
    visibility = ["//visibility:public"],
)
