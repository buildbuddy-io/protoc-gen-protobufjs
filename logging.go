package main

import (
	"fmt"
	"os"
)

func logf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

func fatalf(format string, args ...any) {
	logf("FATAL(protoc-gen-protobufjs): "+format, args...)
	os.Exit(1)
}

func debugf(format string, args ...any) {
	if os.Getenv("DEBUG") == "1" {
		fmt.Fprintf(os.Stderr, "DEBUG(protoc-gen-protobufjs): "+format+"\n", args...)
	}
}
