package descriptor

import "fmt"

type Registry struct {
	// privateOperations if true, respect private annotations and generate separate openapi files for private ops
	privateOperations bool

	// customAnnotations if true, allow custom annotations in the openapi spec
	customAnnotations bool
}

func NewRegistry() *Registry {
	return &Registry{
		privateOperations: false,
		customAnnotations: false,
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

`
}
