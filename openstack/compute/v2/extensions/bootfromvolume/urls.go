package bootfromvolume

import "github.com/doubledutch/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-volumes_boot")
}
