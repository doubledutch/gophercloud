package stacktemplates

import (
	"github.com/doubledutch/gophercloud"
	os "github.com/doubledutch/gophercloud/openstack/orchestration/v1/stacktemplates"
)

// Get retreives data for the given stack template.
func Get(c *gophercloud.ServiceClient, stackName, stackID string) os.GetResult {
	return os.Get(c, stackName, stackID)
}

// Validate validates the given stack template.
func Validate(c *gophercloud.ServiceClient, opts os.ValidateOptsBuilder) os.ValidateResult {
	return os.Validate(c, opts)
}
