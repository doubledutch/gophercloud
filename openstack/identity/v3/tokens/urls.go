package tokens

import "github.com/doubledutch/gophercloud"

func tokenURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
