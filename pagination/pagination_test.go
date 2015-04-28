package pagination

import (
	"github.com/doubledutch/gophercloud"
	"github.com/doubledutch/gophercloud/testhelper"
)

func createClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{
		ProviderClient: &gophercloud.ProviderClient{TokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
