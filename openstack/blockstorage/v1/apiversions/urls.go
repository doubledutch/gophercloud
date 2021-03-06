package apiversions

import (
	"strings"

	"github.com/doubledutch/gophercloud"
)

func getURL(c *gophercloud.ServiceClient, version string) string {
	return c.ServiceURL(strings.TrimRight(version, "/") + "/")
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("")
}
