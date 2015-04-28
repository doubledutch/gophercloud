package bootfromvolume

import (
	"testing"

	th "github.com/doubledutch/gophercloud/testhelper"
	"github.com/doubledutch/gophercloud/testhelper/client"
)

func TestCreateURL(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	c := client.ServiceClient()

	th.CheckEquals(t, c.Endpoint+"os-volumes_boot", createURL(c))
}
