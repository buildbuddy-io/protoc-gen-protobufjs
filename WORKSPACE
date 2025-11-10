workspace(name = "com_github_buildbuddy_io_protoc_gen_protobufjs")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "bazel_skylib",
    sha256 = "6e78f0e57de26801f6f564fa7c4a48dc8b36873e416257a92bbb0937eeac8446",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.8.2/bazel-skylib-1.8.2.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.8.2/bazel-skylib-1.8.2.tar.gz",
    ],
)

http_archive(
    name = "io_bazel_rules_go",
    integrity = "sha256-aK9Uy5f73uXl6P6NIQ0VpRj51iq/1xYgw+r/Oyal/4Y=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.59.0/rules_go-v0.59.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.59.0/rules_go-v0.59.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    integrity = "sha256-Z1EU2LQz0Kn1TYEXGDO+luvEETEVZkt5Hm8gTVjpNEY=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.47.0/bazel-gazelle-v0.47.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.47.0/bazel-gazelle-v0.47.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load(":deps.bzl", "protoc_gen_protobufjs_dependencies")

# gazelle:repository_macro deps.bzl%protoc_gen_protobufjs_dependencies
protoc_gen_protobufjs_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.25.4")

gazelle_dependencies()

# Required by com_google_protobuf
http_archive(
    name = "com_google_googletest",
    integrity = "sha256-QNTslCIX3MhKnr4qaFhK2n1KM6julYdVdjJ46hxeGP8=",
    strip_prefix = "googletest-1.17.0",
    urls = ["https://github.com/google/googletest/archive/v1.17.0.zip"],
)

http_archive(
    name = "com_google_protobuf",
    integrity = "sha256-escvR0NBkLeowzVsjSbTzVsopPtU3f5+gESVq5puD2Y=",
    strip_prefix = "protobuf-4986a7722460e3163f37c98da1cb47f52c6406e1",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/4986a7722460e3163f37c98da1cb47f52c6406e1.tar.gz"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "rules_proto",
    integrity = "sha256-FKIlhwq06RhpZSz9ae8gKCd/wdxJENZdNTti1uCuIfQ=",
    strip_prefix = "rules_proto-7.1.0",
    urls = ["https://github.com/bazelbuild/rules_proto/releases/download/7.1.0/rules_proto-7.1.0.tar.gz"],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies")
rules_proto_dependencies()

load("@rules_proto//proto:toolchains.bzl", "rules_proto_toolchains")
rules_proto_toolchains()

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
http_archive(
    name = "rules_shell",
    sha256 = "e6b87c89bd0b27039e3af2c5da01147452f240f75d505f5b6880874f31036307",
    strip_prefix = "rules_shell-0.6.1",
    url = "https://github.com/bazelbuild/rules_shell/releases/download/v0.6.1/rules_shell-v0.6.1.tar.gz",
)

load("@rules_shell//shell:repositories.bzl", "rules_shell_dependencies", "rules_shell_toolchains")
rules_shell_dependencies()
rules_shell_toolchains()

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
    npmrc = "//test:.npmrc",
    pnpm_lock = "//test:pnpm-lock.yaml",
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
    ts_version_from = "//test:package.json",
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

http_archive(
    name = "toolchains_protoc",
    sha256 = "b440f7f7624d3c95b72640faa0800ffecd6eecba0799f1cbf739496f62e1689e",
    strip_prefix = "toolchains_protoc-0.6.0",
    url = "https://github.com/aspect-build/toolchains_protoc/releases/download/v0.6.0/toolchains_protoc-v0.6.0.tar.gz",
)

######################
# toolchains_protoc setup #
######################
# Fetches the toolchains_protoc dependencies.
# If you want to have a different version of some dependency,
# you should fetch it *before* calling this.
# Alternatively, you can skip calling this function, so long as you've
# already fetched all the dependencies.
load("@toolchains_protoc//protoc:repositories.bzl", "rules_protoc_dependencies")

rules_protoc_dependencies()

load("@bazel_features//:deps.bzl", "bazel_features_deps")

bazel_features_deps()

load("@toolchains_protoc//protoc:toolchain.bzl", "protoc_toolchains")

protoc_toolchains(
    name = "protoc_toolchains",
    version = "v32.3",
)
