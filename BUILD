load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/buildbuddy-io/protoc-gen-protobufjs
gazelle(name = "gazelle")

go_library(
    name = "protoc-gen-protobufjs_lib",
    srcs = [
        "flags.go",
        "index.go",
        "main.go",
        "ts.go",
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
