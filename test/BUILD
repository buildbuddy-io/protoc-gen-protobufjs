load("@aspect_rules_ts//ts:defs.bzl", "ts_config")
load("@npm//:defs.bzl", "npm_link_all_packages")
load(":defs.bzl", "ts_jasmine_node_test")

npm_link_all_packages()

ts_jasmine_node_test(
    name = "encode_test",
    entry_point = "encode_test.ts",
    deps = [
        "//test:node_modules/long",
        "//test:node_modules/protobufjs",
        "//test/proto:trivial_baseline_pbjs_proto",
        "//test/proto:trivial_ts_proto",
        "//test/proto:types_baseline_pbjs_proto",
        "//test/proto:types_ts_proto",
    ],
)

ts_config(
    name = "tsconfig",
    src = "tsconfig.json",
    visibility = ["//visibility:public"],
)
