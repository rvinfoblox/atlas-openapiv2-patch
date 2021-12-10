package main

import "github.com/spf13/pflag"

var (
	swaggerFiles          = pflag.StringSlice("files", []string{}, "input swagger file location")
	withPrivate           = pflag.Bool("with_private", false, "if unset, generate swagger schema without operations 0 as 'private' work only if atlas_patch set")
	withCustomAnnotations = pflag.Bool("with_custom_annotations", false, "if set, you became available to use custom annotations")
)
