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

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
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
