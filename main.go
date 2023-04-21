package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	if err := run(); err != nil {
		logf("error: %s", err)
		os.Exit(1)
	}
}

func run() error {
	// Note: flags will not actually be available at this point when invoked via
	// protoc. This is just to support running `./protoc-gen-protobufjs -help`
	// so that we can get the list of supported flags.
	flag.Parse()

	req, err := readRequest()
	if err != nil {
		return fmt.Errorf("failed to read CodeGenerationRequest from stdin: %s", err)
	}

	// CodeGenerationRequest only supports passing one "parameter", but we allow
	// passing multiple flags by separating them with ":" in the parameter.
	os.Args = append(os.Args, strings.Split(req.GetParameter(), ":")...)
	flag.Parse()

	res, err := generateCode(req)
	if err != nil {
		return err
	}

	return writeResponse(res)
}

func readRequest() (*pluginpb.CodeGeneratorRequest, error) {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	req := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(b, req); err != nil {
		return nil, err
	}
	return req, nil
}

func writeResponse(res *pluginpb.CodeGeneratorResponse) error {
	b, err := proto.Marshal(res)
	if err != nil {
		return err
	}
	if _, err := os.Stdout.Write(b); err != nil {
		return err
	}
	return nil
}
