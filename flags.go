package main

import (
	"flag"
	"strings"
)

type StringSliceFlag []string

func NewStringSliceFlag(slice *[]string) *StringSliceFlag {
	return (*StringSliceFlag)(slice)
}

func StringSlice(name string, defaultValue []string, usage string) *[]string {
	value := &[]string{}
	StringSliceVar(value, name, defaultValue, usage)
	return value
}

func StringSliceVar(value *[]string, name string, defaultValue []string, usage string) {
	if defaultValue == nil && *value != nil {
		*value = nil
	} else if len(*value) != len(defaultValue) || *value == nil {
		*value = make([]string, len(defaultValue))
	}
	copy(*value, defaultValue)
	flag.Var((*StringSliceFlag)(value), name, usage)
}

func (f *StringSliceFlag) String() string {
	return strings.Join(*f, ",")
}

func (f *StringSliceFlag) Set(values string) error {
	if values == "" {
		return nil
	}
	for _, val := range strings.Split(values, ",") {
		*f = append(*f, val)
	}
	return nil
}
