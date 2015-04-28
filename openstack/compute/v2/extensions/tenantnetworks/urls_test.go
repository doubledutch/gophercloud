package tenantnetworks

import (
	"testing"

	th "github.com/doubledutch/gophercloud/testhelper"
	"github.com/doubledutch/gophercloud/testhelper/client"
)

func TestListURL(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	c := client.ServiceClient()

	th.CheckEquals(t, c.Endpoint+"os-tenant-networks", listURL(c))
}

func TestGetURL(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	c := client.ServiceClient()
	id := "1"

	th.CheckEquals(t, c.Endpoint+"os-tenant-networks/"+id, getURL(c, id))
}
