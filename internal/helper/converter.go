package helper

import (
	projectcontour "github.com/projectcontour/contour/apis/projectcontour/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// UnstructuredConverter handles conversions between unstructured.Unstructured and Contour types
type UnstructuredConverter struct {
	// scheme holds an initializer for converting Unstructured to a type
	scheme *runtime.Scheme
}

// NewUnstructuredConverter returns a new UnstructuredConverter initialized
func NewUnstructuredConverter() *UnstructuredConverter {
	uc := &UnstructuredConverter{
		scheme: runtime.NewScheme(),
	}

	// Setup converter to understand custom CRD types
	projectcontour.AddKnownTypes(uc.scheme)

	return uc
}

// Convert converts a typed struct to an unstructured.Unstructured
func (c *UnstructuredConverter) Convert(proxy *projectcontour.HTTPProxy) (*unstructured.Unstructured, error) {
	obj := &unstructured.Unstructured{}
	err := c.scheme.Convert(proxy, obj, nil)
	return obj, err
}
