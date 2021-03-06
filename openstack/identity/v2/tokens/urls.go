package tokens

import "github.com/doubledutch/gophercloud"

// CreateURL generates the URL used to create new Tokens.
func CreateURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("tokens")
}
