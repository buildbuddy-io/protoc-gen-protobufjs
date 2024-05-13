workspace(name = "com_github_buildbuddy_io_protoc_gen_protobufjs")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "6b65cb7917b4d1709f9410ffe00ecf3e160edf674b78c54a894471320862184f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.39.0/rules_go-v0.39.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.39.0/rules_go-v0.39.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "ecba0f04f96b4960a5b250c8e8eeec42281035970aa8852dda73098274d14a1d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load(":deps.bzl", "protoc_gen_protobufjs_dependencies")

# gazelle:repository_macro deps.bzl%protoc_gen_protobufjs_dependencies
protoc_gen_protobufjs_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.19.5")

gazelle_dependencies()

# Required by com_google_protobuf
http_archive(
    name = "com_google_googletest",
    sha256 = "ffa17fbc5953900994e2deec164bb8949879ea09b411e07f215bfbb1f87f4632",
    strip_prefix = "googletest-1.13.0",
    urls = ["https://github.com/google/googletest/archive/v1.13.0.zip"],
)

http_archive(
    name = "com_google_protobuf",
    sha256 = "2118051b4fb3814d59d258533a4e35452934b1ddb41230261c9543384cbb4dfc",
    strip_prefix = "protobuf-3.22.2",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.22.2.tar.gz"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "rules_proto",
    sha256 = "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd",
    strip_prefix = "rules_proto-5.3.0-21.7",
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz",
    ],
)

# JS dependencies for testing

http_archive(
    name = "aspect_rules_js",
    sha256 = "5a00869efaeb308245f8132a671fe86524bdfc4f8bfd1976d26f862b316dc3c9",
    strip_prefix = "rules_js-1.42.0",
    url = "https://github.com/aspect-build/rules_js/releases/download/v1.42.0/rules_js-v1.42.0.tar.gz",
)

load("@aspect_rules_js//js:repositories.bzl", "rules_js_dependencies")

rules_js_dependencies()

load("@rules_nodejs//nodejs:repositories.bzl", "DEFAULT_NODE_VERSION", "nodejs_register_toolchains")

nodejs_register_toolchains(
    name = "nodejs",
    node_version = DEFAULT_NODE_VERSION,
)

load("@aspect_rules_js//npm:repositories.bzl", "npm_translate_lock")

npm_translate_lock(
    name = "npm",
    npmrc = "//:.npmrc",
    pnpm_lock = "//:pnpm-lock.yaml",
    verify_node_modules_ignored = "//:.bazelignore",
)

load("@npm//:repositories.bzl", "npm_repositories")

npm_repositories()

# TS dependencies for testing

http_archive(
    name = "aspect_rules_ts",
    sha256 = "b11f5bd59983a58826842029b99240fd0eeb6f1291d710db10f744b327701646",
    strip_prefix = "rules_ts-2.3.0",
    url = "https://github.com/aspect-build/rules_ts/releases/download/v2.3.0/rules_ts-v2.3.0.tar.gz",
)

load("@aspect_rules_ts//ts:repositories.bzl", "rules_ts_dependencies")

rules_ts_dependencies(
    ts_version_from = "//:package.json",
)

http_archive(
    name = "aspect_rules_jasmine",
    sha256 = "4c16ef202d1e53fd880e8ecc9e0796802201ea9c89fa32f52d5d633fff858cac",
    strip_prefix = "rules_jasmine-1.1.1",
    url = "https://github.com/aspect-build/rules_jasmine/releases/download/v1.1.1/rules_jasmine-v1.1.1.tar.gz",
)

load("@aspect_rules_jasmine//jasmine:dependencies.bzl", "rules_jasmine_dependencies")

rules_jasmine_dependencies()

# esbuild dependency for testing

http_archive(
    name = "aspect_rules_esbuild",
    sha256 = "ce206c03e27a702ba2a480ee0a1e4f8db124f3595460a77a3ae1e465243c7a73",
    strip_prefix = "rules_esbuild-0.19.0",
    url = "https://github.com/aspect-build/rules_esbuild/releases/download/v0.19.0/rules_esbuild-v0.19.0.tar.gz",
)

load("@aspect_rules_esbuild//esbuild:dependencies.bzl", "rules_esbuild_dependencies")

rules_esbuild_dependencies()

load("@aspect_rules_esbuild//esbuild:repositories.bzl", "LATEST_ESBUILD_VERSION", "esbuild_register_toolchains")

esbuild_register_toolchains(
    name = "esbuild",
    esbuild_version = LATEST_ESBUILD_VERSION,
)
