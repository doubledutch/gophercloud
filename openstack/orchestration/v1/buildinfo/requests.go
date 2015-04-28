package buildinfo

import "github.com/doubledutch/gophercloud"

// Get retreives data for the given stack template.
func Get(c *gophercloud.ServiceClient) GetResult {
	var res GetResult
	_, res.Err = c.Get(getURL(c), &res.Body, nil)
	return res
}
