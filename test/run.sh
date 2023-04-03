#!/usr/bin/env bash
set -e

# Usage: ./test/run.sh [-i]
#   -i: interactive mode; restart on changes
#       requires godemon: https://github.com/bduffany/godemon

PLUGIN="$PWD/protoc-gen-protobufjs"
REPO_ROOT="$PWD"

if [[ "$1" == "-i" ]]; then
  exec godemon bash -c 'sleep 0.05 && clear && '"$0"
fi

if ! [[ -e "$REPO_ROOT/protoc" ]]; then
  echo "Building protoc"
  bazel build @com_google_protobuf//:protoc &>/dev/null
  cp ./bazel-bin/external/com_google_protobuf/protoc ./protoc
fi

echo "Building protoc-gen-protobufjs"
go mod tidy
go build

cd test

# Clean
rm -rf ./out/ ./gen
mkdir -p ./out/{pbjs,protoc-gen-protobufjs} ./gen

_pbjs() {
  local out="$1"
  if [[ "$OUT" ]]; then
    out="$OUT"
  fi
  mkdir -p "$(dirname "./gen/$out")"
  pnpm pbjs \
    --target=static-module --wrap=commonjs \
    --force-message --strict-long \
    --no-delimited --no-verify \
    "$@" --out ./gen/"${out/.proto/.pbjs.js}"
  pnpm pbts ./gen/"${out/.proto/.pbjs.js}" --out ./gen/"${out/.proto/.pbjs.d.ts}"
  echo "Wrote" ./gen/"${out/.proto/.pbjs.js}" ./gen/"${out/.proto/.d.ts}"
}

echo "Generating pbjs outputs under ./gen/"

_pbjs proto/trivial.proto &
_pbjs proto/types.proto &
_pbjs proto/nesting.proto &
_pbjs proto/service.proto &
OUT=proto/multifile/multifile.proto _pbjs proto/multi/*.proto &
wait

echo "Generating protoc-gen-protobufjs outputs"

# TODO: test imports
# TODO: test well-known types

_protoc() {
  "$REPO_ROOT/protoc" --plugin="$PLUGIN" --protobufjs_out=./gen/ -I "$REPO_ROOT" -I . "$@"
}

_protoc ./proto/trivial.proto
_protoc ./proto/types.proto
_protoc ./proto/service.proto
_protoc ./proto/nesting.proto
_protoc ./proto/multifile/a.proto ./proto/multifile/b.proto --protobufjs_opt=-out=proto/multifile/multifile.ts

echo "Compiling encode.ts"
pnpm tsc
pnpm node ./out/encode.js

echo "Diffing protoscope representations"

find ./out/ -name '*.bin' | while read -r f; do
  protoscope -explicit-length-prefixes -explicit-wire-types "$f" >"$f.protoscope.txt"
done

find ./out/pbjs -name '*.protoscope.txt' | while read -r f; do
  diff -Pdpru --report-identical-files "$f" "${f/pbjs/protoc-gen-protobufjs}"
done
