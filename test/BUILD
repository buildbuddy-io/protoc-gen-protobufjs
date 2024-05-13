load(":defs.bzl", "ts_jasmine_node_test")

ts_jasmine_node_test(
    name = "encode_test",
    entry_point = "encode_test.ts",
    deps = [
        "//:node_modules/long",
        "//:node_modules/protobufjs",
        "//test/proto:trivial_baseline_pbjs_proto",
        "//test/proto:trivial_ts_proto",
        "//test/proto:types_baseline_pbjs_proto",
        "//test/proto:types_ts_proto",
    ],
)
