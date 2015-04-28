package volumetypes

import (
	"github.com/doubledutch/gophercloud"
	os "github.com/doubledutch/gophercloud/openstack/blockstorage/v1/volumetypes"
	"github.com/doubledutch/gophercloud/pagination"
)

// List returns all volume types.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return os.List(client)
}

// Get will retrieve the volume type with the provided ID. To extract the volume
// type from the result, call the Extract method on the GetResult.
func Get(client *gophercloud.ServiceClient, id string) GetResult {
	return GetResult{os.Get(client, id)}
}
