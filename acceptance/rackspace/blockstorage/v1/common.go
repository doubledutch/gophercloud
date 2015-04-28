// +build acceptance

package v1

import (
	"os"
	"testing"

	"github.com/doubledutch/gophercloud"
	"github.com/doubledutch/gophercloud/acceptance/tools"
	"github.com/doubledutch/gophercloud/rackspace"
	th "github.com/doubledutch/gophercloud/testhelper"
)

func newClient() (*gophercloud.ServiceClient, error) {
	opts, err := rackspace.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}
	opts = tools.OnlyRS(opts)
	region := os.Getenv("RS_REGION")

	provider, err := rackspace.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}

	return rackspace.NewBlockStorageV1(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}

func setup(t *testing.T) *gophercloud.ServiceClient {
	client, err := newClient()
	th.AssertNoErr(t, err)

	return client
}
