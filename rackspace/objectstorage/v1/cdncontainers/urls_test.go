package cdncontainers

import (
	"testing"

	"github.com/doubledutch/gophercloud"
	th "github.com/doubledutch/gophercloud/testhelper"
)

const endpoint = "http://localhost:57909/"

func endpointClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint}
}

func TestEnableURL(t *testing.T) {
	actual := enableURL(endpointClient(), "foo")
	expected := endpoint + "foo"
	th.CheckEquals(t, expected, actual)
}
