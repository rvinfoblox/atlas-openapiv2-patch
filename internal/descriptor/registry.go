package descriptor

import (
	"fmt"
	genopenapi "github.com/infobloxopen/atlas-openapiv2-patch/pkg/svc"
)

type Registry struct {
	// privateOperations if true, respect private annotations and generate separate openapi files for private ops
	privateOperations bool

	// customAnnotations if true, allow custom annotations in the openapi spec
	customAnnotations bool

	// getResponse if specified, overrides the default response code for GET API
	getResponse int

	// postResponse if specified, overrides the default response code for POST API
	postResponse int

	// putResponse if specified, overrides the default response code for PUT API
	putResponse int

	// patchResponse if specified, overrides the default response code for PATCH API
	patchResponse int

	// deleteResponse if specified, overrides the default response code for DELETE API
	deleteResponse int
}

func (r *Registry) GetResponse() int {
	return r.getResponse
}

func (r *Registry) SetGetResponse(getResponse int) {
	r.getResponse = getResponse
}

func (r *Registry) PutResponse() int {
	return r.putResponse
}

func (r *Registry) SetPutResponse(putResponse int) {
	r.putResponse = putResponse
}

func (r *Registry) PatchResponse() int {
	return r.patchResponse
}

func (r *Registry) SetPatchResponse(patchResponse int) {
	r.patchResponse = patchResponse
}

func (r *Registry) PostResponse() int {
	return r.postResponse
}

func (r *Registry) SetPostResponse(postResponse int) {
	r.postResponse = postResponse
}

func (r *Registry) DeleteResponse() int {
	return r.deleteResponse
}

func (r *Registry) SetDeleteResponse(deleteResponse int) {
	r.deleteResponse = deleteResponse
}

func NewRegistry() *Registry {
	return &Registry{
		privateOperations: false,
		customAnnotations: false,
		getResponse:       genopenapi.DefaultGetResponse,
		postResponse:      genopenapi.DefaultPostResponse,
		putResponse:       genopenapi.DefaultPutResponse,
		patchResponse:     genopenapi.DefaultPatchResponse,
		deleteResponse:    genopenapi.DefaultDeleteResponse,
	}
}

// IsWithPrivateOperations whether private operations are enabled
func (r *Registry) IsWithPrivateOperations() bool {
	return r.privateOperations
}

// SetPrivateOperations if true, respect private operation annotation
func (r *Registry) SetPrivateOperations(private bool) {
	r.privateOperations = private
}

// IsWithCustomAnnotations whether custom annotations are used
func (r *Registry) IsWithCustomAnnotations() bool {
	return r.customAnnotations
}

// SetCustomAnnotations if true, use atlas custom annotations
func (r *Registry) SetCustomAnnotations(custom bool) {
	r.customAnnotations = custom
}

func (r *Registry) String() string {
	return `

		privateOperations          ` + fmt.Sprintf("%t", r.privateOperations) + `
		customAnnotations          ` + fmt.Sprintf("%t", r.customAnnotations) + `
		postResponse               ` + fmt.Sprintf("%d", r.postResponse) + `
		putResponse                ` + fmt.Sprintf("%d", r.putResponse) + `
		patchResponse              ` + fmt.Sprintf("%d", r.patchResponse) + `
		deleteResponse             ` + fmt.Sprintf("%d", r.deleteResponse) + `
`
}
