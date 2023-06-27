package main

import (
	"github.com/spf13/pflag"

	svc "github.com/infobloxopen/atlas-openapiv2-patch/pkg/svc"
)

var (
	swaggerFiles = pflag.StringSlice("files", []string{}, "input swagger file location")

	withPrivate           = pflag.Bool("with_private", false, "if unset, generate swagger schema without operations 0 as 'private' work only if atlas_patch set")
	withCustomAnnotations = pflag.Bool("with_custom_annotations", false, "if set, you became available to use custom annotations")

	// Allow API response codes to be overridden for POST, PUT, PATCH and DELETE.
	// Note: Since GET always returns a 200, we don't permit overriding the response for that method alone.
	withPostResponse   = pflag.Int("with_post_response", svc.DefaultPostResponse, "if set, generate swagger doc with the given response code for the POST API")
	withPutResponse    = pflag.Int("with_put_response", svc.DefaultPutResponse, "if set, generate swagger doc with the given response code for the PUT API")
	withPatchResponse  = pflag.Int("with_patch_response", svc.DefaultPatchResponse, "if set, generate swagger doc with the given response code for the patch API")
	withDeleteResponse = pflag.Int("with_delete_response", svc.DefaultDeleteResponse, "if set, generate swagger doc with the given response code for the DELETE API")
)
