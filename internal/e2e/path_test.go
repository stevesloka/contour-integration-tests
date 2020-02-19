package e2e

import (
	"testing"

	projectcontour "github.com/projectcontour/contour/apis/projectcontour/v1"
	"github.com/stevesloka/contour-integration-tests/internal/helper"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func beforeEach() {
	helper.RemoveAllProxies()
}

func TestGetDefaultPath(t *testing.T) {
	beforeEach()
	err := helper.CreateProxy(proxy())
	if err != nil {
		t.Fatalf("Failed to create proxy: %v", err)
	}

	assert := assert.New(t)
	url, err := helper.GetUrl("/")
	if err != nil {
		t.Fatalf("Failed to parse url: %v", err)
	}
	rs, body := helper.GetRequest(url.String())

	assert.Equal(200, rs.StatusCode)
	assert.Contains(body, "This is the default app site!")
}

func proxy() *projectcontour.HTTPProxy {
	return &projectcontour.HTTPProxy{
		ObjectMeta: v1.ObjectMeta{
			Name: "test01",
		},
		Spec: projectcontour.HTTPProxySpec{
			VirtualHost: &projectcontour.VirtualHost{
				Fqdn: helper.FQDN,
			},
			Routes: []projectcontour.Route{{
				Conditions: []projectcontour.Condition{{
					Prefix: "/",
				}},
				Services: []projectcontour.Service{{
					Name: "test01",
					Port: 80,
				}},
			}},
		},
	}
}
