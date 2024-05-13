load("@aspect_rules_esbuild//esbuild:defs.bzl", "esbuild")
load("@aspect_rules_jasmine//jasmine:defs.bzl", "jasmine_test")
load("@aspect_rules_js//js:defs.bzl", "js_library")
load("@aspect_rules_ts//ts:defs.bzl", "ts_project")
load("@npm//:protobufjs-cli/package_json.bzl", protobufjs_cli_bin = "bin")
load("//:rules.bzl", "protoc_gen_protobufjs")

def ts_library(name, srcs, **kwargs):
    ts_project(
        name = name,
        tsconfig = "//:tsconfig",
        transpiler = "tsc",
        srcs = srcs,
        **kwargs
    )

def protobufjs_cli_compile(name, protos, out, **kwargs):
    js_out = out + ".js"
    d_ts_out = out + ".d.ts"

    # ../../../ is needed below since the binaries run within bazel-out/<config>/bin

    args = [
        "--target=static-module",
        "--wrap=commonjs",
        "--force-message",
        "--strict-long",
        "--no-delimited",
        "--no-verify",
        "--out=../../../$@",
    ]
    for proto in protos:
        args.append("$(location %s)" % proto)

    protobufjs_cli_bin.pbjs(
        name = name + "__pbjs",
        srcs = protos,
        outs = [js_out],
        args = args,
        **kwargs
    )

    protobufjs_cli_bin.pbts(
        name = name + "__pbts",
        srcs = [js_out],
        outs = [d_ts_out],
        args = [
            "../../../$(location %s)" % js_out,
            "--out=../../../$@",
        ],
        **kwargs
    )

    js_library(
        name = name,
        srcs = [
            "%s__pbjs" % name,
            "%s__pbts" % name,
        ],
        **kwargs
    )

def ts_proto_library(name, out, proto, **kwargs):
    protoc_gen_protobufjs(
        name = "%s__protoc" % name,
        out = out,
        proto = proto,
        **kwargs
    )

    js_library(
        name = name,
        srcs = ["%s__protoc" % name],
        **kwargs
    )

def ts_jasmine_node_test(name, entry_point, deps = [], size = "small", **kwargs):
    # jasmine_node_test can't run TS tests directly, so compile the TS to JS
    # first. Note, this produces an ES6 module. We use ES6 modules as our
    # transpilation target throughout the app since esbuild currently does not
    # do code-splitting properly on commonjs modules that are produced by SWC.

    deps = list(deps)
    deps.extend([
        "//:node_modules/@types/node",
        "//:node_modules/@types/jasmine",
    ])

    ts_library(
        name = "%s__lib" % name,
        testonly = 1,
        srcs = [entry_point],
        deps = deps,
        **kwargs
    )

    # Bundle with esbuild targeting node to avoid ES modules insanity.
    esbuild(
        name = "%s__bundle" % name,
        testonly = 1,
        entry_point = entry_point,
        deps = ["%s__lib" % name],
        platform = "node",
    )

    # jasmine_node_test(
    #     name = name,
    #     size = "small",
    #     srcs = [":%s_commonjs.test.js" % name],
    #     **kwargs
    # )

    # Copy the commonjs module to trick jasmine_node_test into thinking this is
    # a plain JS source. The test fails with "no specs found" if we try to pass
    # the commonjs module output as srcs directly.
    native.genrule(
        name = "%s__entrypoint" % name,
        srcs = [":%s__bundle.js" % name],
        outs = [":%s.test.js" % name],
        cmd_bash = "ln $(SRCS) $@ 2>/dev/null || cp $(SRCS) $@",
        tags = ["local"],
    )

    jasmine_test(
        name = name,
        args = ["*.test.js"],
        chdir = native.package_name(),
        data = [":%s.test.js" % name] + deps,
        node_modules = "//:node_modules",
    )
