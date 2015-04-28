package buildinfo

import (
	"testing"

	th "github.com/doubledutch/gophercloud/testhelper"
	fake "github.com/doubledutch/gophercloud/testhelper/client"
)

func TestGetTemplate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t, GetOutput)

	actual, err := Get(fake.ServiceClient()).Extract()
	th.AssertNoErr(t, err)

	expected := GetExpected
	th.AssertDeepEquals(t, expected, actual)
}
