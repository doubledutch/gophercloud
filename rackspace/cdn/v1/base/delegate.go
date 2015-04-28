package base

import (
	"github.com/doubledutch/gophercloud"

	os "github.com/doubledutch/gophercloud/openstack/cdn/v1/base"
)

// Get retrieves the home document, allowing the user to discover the
// entire API.
func Get(c *gophercloud.ServiceClient) os.GetResult {
	return os.Get(c)
}

// Ping retrieves a ping to the server.
func Ping(c *gophercloud.ServiceClient) os.PingResult {
	return os.Ping(c)
}
