load("@rules_proto//proto:defs.bzl", "ProtoInfo")

def _trim_prefix(text, prefix):
    if text.startswith(prefix):
        return text[len(prefix):]
    return text

def _package_path(ctx):
    return "/".join(ctx.build_file_path.split("/")[:-1])

ProtobufjsInfo = provider(
    doc = "Info about a generated protobufjs code.",
    fields = {
        "proto_info": "ProtoInfo for the proto_library used to generate the TS file.",
        "proto_import_paths": "Proto import paths of all protos that are included in the generated TS file.",
        "gen_import_path": "Import path corresponding to the generated file.",
    },
)

def _protoc_gen_protobufjs_impl(ctx):
    gen_import_path = _package_path(ctx) + "/" + ctx.attr.out

    # Use strict_imports since we expect all direct proto imports to
    # correspond to an explicit direct dep.
    plugin_args = [
        "-strict_imports",
        "-out=" + gen_import_path,
    ]

    args = [
        "--plugin",
        ctx.executable._protoc_gen_protobufjs.path,
        "--protobufjs_out",
        ctx.bin_dir.path + "/" + ctx.label.workspace_root,
    ]

    inputs = ctx.attr.proto[ProtoInfo].transitive_sources.to_list()

    proto_paths = depset(
        [ctx.attr.proto[ProtoInfo].proto_source_root] + [
            label[ProtoInfo].proto_source_root
            for label in ctx.attr._implicit_imports
        ],
        transitive = [
            dep[ProtobufjsInfo].proto_info.transitive_proto_path
            for dep in ctx.attr.deps
        ],
    )

    for dep in ctx.attr.deps:
        plugin_args.extend([
            "-import_path=%s=%s" % (proto_import_path, dep[ProtobufjsInfo].gen_import_path)
            for proto_import_path in dep[ProtobufjsInfo].proto_import_paths
        ])

    for label in ctx.attr._implicit_imports:
        inputs.extend(label[ProtoInfo].direct_sources)

    args.extend(["--proto_path=" + path for path in proto_paths.to_list()])

    plugin_args.extend(ctx.attr.plugin_args)
    args.append("--protobufjs_opt=" + ":".join(plugin_args))
    args += [f.path for f in ctx.attr.proto[ProtoInfo].direct_sources]

    outputs = [
        ctx.actions.declare_file(ctx.attr.out + ".js"),
        ctx.actions.declare_file(ctx.attr.out + ".d.ts"),
    ]
    ctx.actions.run(
        mnemonic = "ProtocGenPbjs",
        outputs = outputs,
        inputs = inputs,
        tools = [ctx.executable._protoc_gen_protobufjs],
        executable = ctx.executable._protoc,
        arguments = args,

        # Debug:

        # tools = [ctx.executable._protoc, ctx.executable._protoc_gen_protobufjs],
        # executable = "/bin/bash",
        # arguments = ["-c", """
        #     set -x
        #     (
        #         """ + ctx.executable._protoc.path + """ --version
        #         pwd
        #         find . | grep -P '\\.proto(.bin)?$'
        #     ) >&2
        #     """ + ctx.executable._protoc.path + " " + " ".join(args) + """
        #     (
        #         echo "AFTER running:"
        #         find | grep -P '\\.ts$'
        #     ) >&2
        # """],
    )
    return [
        DefaultInfo(files = depset(outputs)),
        ProtobufjsInfo(
            proto_info = ctx.attr.proto[ProtoInfo],
            proto_import_paths = [
                _trim_prefix(proto_src.path, ctx.attr.proto[ProtoInfo].proto_source_root + "/")
                for proto_src in ctx.attr.proto[ProtoInfo].direct_sources
            ],
            gen_import_path = gen_import_path,
        ),
    ]

protoc_gen_protobufjs = rule(
    doc = """
    Runs the protoc-gen-protobufjs plugin under protoc to generate
    JS code and TS typings from a proto library.
    """,
    implementation = _protoc_gen_protobufjs_impl,
    provides = [DefaultInfo, ProtobufjsInfo],
    attrs = {
        "out": attr.string(
            doc = "Output file name without the .js or .d.ts extension.",
            mandatory = True,
        ),
        "proto": attr.label(
            providers = [ProtoInfo],
            mandatory = True,
        ),
        "deps": attr.label_list(
            providers = [ProtobufjsInfo],
            default = [],
        ),
        "plugin_args": attr.string_list(
            doc = """
            Extra args to pass to protoc-gen-protobufjs.
            Args cannot contain ':', which is used as an arg separator.
            """,
            default = [],
        ),
        "_protoc": attr.label(
            default = "@com_google_protobuf//:protoc",
            executable = True,
            cfg = "exec",
            allow_single_file = True,
        ),
        "_protoc_gen_protobufjs": attr.label(
            default = "@com_github_buildbuddy_io_protoc_gen_protobufjs//:protoc-gen-protobufjs",
            executable = True,
            cfg = "exec",
            allow_single_file = True,
        ),
        "_implicit_imports": attr.label_list(
            doc = "Imports that are implicitly included in proto_library targets and may not necessarily be listed as deps.",
            default = [
                "@com_google_protobuf//:any_proto",
                "@com_google_protobuf//:api_proto",
                "@com_google_protobuf//:descriptor_proto",
                "@com_google_protobuf//:duration_proto",
                "@com_google_protobuf//:empty_proto",
                "@com_google_protobuf//:field_mask_proto",
                "@com_google_protobuf//:source_context_proto",
                "@com_google_protobuf//:struct_proto",
                "@com_google_protobuf//:timestamp_proto",
                "@com_google_protobuf//:type_proto",
                "@com_google_protobuf//:wrappers_proto",
            ],
        ),
    },
)
