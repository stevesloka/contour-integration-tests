package helper

import "k8s.io/apimachinery/pkg/runtime/schema"

const (
	FQDN      = "demo.projectcontour.io"
	Namespace = "contour-integration-tests"
)

var (
	HTTPProxyGVR = schema.GroupVersionResource{Group: "projectcontour.io", Version: "v1", Resource: "httpproxies"}
)
