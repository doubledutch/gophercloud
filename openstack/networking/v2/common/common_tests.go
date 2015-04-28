package common

import (
	"github.com/doubledutch/gophercloud"
	"github.com/doubledutch/gophercloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *gophercloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc
}
