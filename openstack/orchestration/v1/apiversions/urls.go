package apiversions

import "github.com/doubledutch/gophercloud"

func apiVersionsURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}
