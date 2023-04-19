# protoc-gen-protobufjs

A `protoc` plugin to generate TypeScript code from `.proto` files.

It is intended as a faster replacement for `protobufjs-cli`, ideal for
large, evolving projects with many protos. The generated code uses
`protobufjs/minimal` and `long` as the only runtime dependencies.

## Improvements over protobufjs-cli (pbjs / pbts)

- Generates code 10X faster. One benchmark run on
  https://github.com/buildbuddy-io/buildbuddy brings the code generation
  time from ~19s to ~1.8s.

- Supports generating code for only a single `.proto` file and not its
  imported files[^1]. This makes it a good fit for build systems like
  [Bazel](https://bazel.build) which are optimized for many small build
  steps rather than a few big build steps.

- Imports are declared explicitly, rather than referencing global "roots".
  This makes it easier to navigate the generated code.

- Generated code preserves comments from the original proto source.

[^1]:
    By default, `protobufjs-cli` bundles all transitive dependencies
    into the output `.js` file, which causes an explosion of unnecessary
    dependency edges in the build graph. `protobufjs-cli` does have a
    `--sparse` option, but this only excludes _indirect_ imports from the
    output JS. For example, let's say we have `a.proto` that imports
    `b.proto`, and `b.proto` imports `c.proto`. Then
    `pbjs --sparse a.proto` generates code for both A and B in the output.
    With `protoc-gen-protobufjs`, the generated code only contains A, and
    instead generates an _import_ for B.

## protobufjs-cli compatibility

The generated code is compatible with `pbjs` + `pbts`, where `pbjs` is
run with the following flags:

```
--target=static-module
--wrap=es6
--force-message
--strict-long
--no-delimited
--no-verify
```

(PRs are welcome to add support for other options that would increase
compatibility with the protobufjs-cli.)

In addition, the generated code is incompatible in the following ways:

- Does not support the deprecated proto2
  [groups](https://protobuf.dev/programming-guides/proto2/#groups) feature
- Does not yet support proto2 `default` value annotations
- Does not generate `_optional` getters for proto3 optional scalar fields;
  instead types the field as `|null|undefined`
- Does not support reflection
- Does not support register global `protobuf.roots` entries like
  `protobuf.roots["default"].my.proto.Message`

## Usage

```
# Install the plugin
go install github.com/buildbuddy-io/protoc-gen-protobufjs@latest

# Make sure the installed binary (protoc-gen-protobufjs) is in $PATH
export PATH="$(go env GOPATH)/bin:$PATH"

# Compile your protos
protoc --protobufjs_out=./out ./src/hello.proto
```

## Plugin flags

The plugin supports the following flags:

<!-- #help:start -->

```
-import_path: PROTO_PATH=TS_PATH
  A mapping of proto imports to TS imports, as PROTO_PATH=TS_PATH pairs
  (this flag can be specified more than once). If the proto does not end
  with '.proto' then these act as directory prefixes.

-out: path
  Output file path. If this is set and multiple protos are provided as
  input, the generated code for all protos will be written to this file.
  Any '.ts' or '.js' extension will be ignored.

-strict_imports:
  If set, all directly imported protos must have an -import_path
  specified. This is useful when integrating with a build system where all
  direct dependencies must be explicitly specified.

-import_module_specifier_ending: string
  Suffix to apply to generated import module specifiers. May need to be
  set to ".js" in rare cases.

```

<!-- #help:end -->

Flags are specified using the `--protobufjs_opt` flag to `protoc`. To
specify multiple plugin flags, separate each flag with ":" (note that
passing `--protobufjs_opt` more than once will not accumulate flags).

Example:

```
protoc \
  --protobufjs_out=./out/ \
  --protobufjs_opt=-strict_imports:-import_path=src/foo.proto:src/foo.ts
```

<!-- TODO: usage with Bazel -->
