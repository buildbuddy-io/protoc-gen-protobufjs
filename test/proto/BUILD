# gazelle:exclude .
load("@rules_proto//proto:defs.bzl", "proto_library")
load("//test:defs.bzl", "protobufjs_cli_compile", "ts_proto_library")

package(default_visibility = ["//test:__subpackages__"])

SINGLE_FILES = [
    "trivial",
    "types",
    "nesting",
    "service",
    "abs_reference",
]

[
    proto_library(
        name = "%s_proto" % name,
        srcs = ["%s.proto" % name],
    )
    for name in SINGLE_FILES
]

proto_library(
    name = "multifile_proto",
    srcs = [
        "multifile/a.proto",
        "multifile/b.proto",
    ],
)

# Generate with protoc-gen-protobufjs for comparison

[
    ts_proto_library(
        name = name + "_ts_proto",
        out = name,
        proto = ":%s_proto" % name,
    )
    for name in SINGLE_FILES
]

ts_proto_library(
    name = "multifile_ts_proto",
    out = "multifile",
    proto = ":multifile_proto",
)

# Generate with pbjs for comparison

[
    protobufjs_cli_compile(
        name = name + "_baseline_pbjs_proto",
        out = name + ".pbjs",
        protos = [name + ".proto"],
    )
    for name in SINGLE_FILES
]

protobufjs_cli_compile(
    name = "multifile_baseline_pbjs",
    out = "multifile.pbjs.js",
    protos = [
        "multifile/a.proto",
        "multifile/b.proto",
    ],
)
