package common

import (
	"github.com/doubledutch/gophercloud"
	"github.com/doubledutch/gophercloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *gophercloud.ServiceClient {
	return client.ServiceClient()
}
