// +build acceptance

package v1

import (
	"testing"

	"github.com/doubledutch/gophercloud"
	"github.com/doubledutch/gophercloud/openstack/orchestration/v1/stacks"
	"github.com/doubledutch/gophercloud/openstack/orchestration/v1/stacktemplates"
	th "github.com/doubledutch/gophercloud/testhelper"
)

func TestStackTemplates(t *testing.T) {
	// Create a provider client for making the HTTP requests.
	// See common.go in this directory for more information.
	client := newClient(t)

	stackName := "postman_stack_2"

	createOpts := stacks.CreateOpts{
		Name:     stackName,
		Template: template,
		Timeout:  5,
	}
	stack, err := stacks.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created stack: %+v\n", stack)
	defer func() {
		err := stacks.Delete(client, stackName, stack.ID).ExtractErr()
		th.AssertNoErr(t, err)
		t.Logf("Deleted stack (%s)", stackName)
	}()
	err = gophercloud.WaitFor(60, func() (bool, error) {
		getStack, err := stacks.Get(client, stackName, stack.ID).Extract()
		if err != nil {
			return false, err
		}
		if getStack.Status == "CREATE_COMPLETE" {
			return true, nil
		}
		return false, nil
	})

	tmpl, err := stacktemplates.Get(client, stackName, stack.ID).Extract()
	th.AssertNoErr(t, err)
	t.Logf("retrieved template: %+v\n", tmpl)

	validateOpts := stacktemplates.ValidateOpts{
		Template: map[string]interface{}{
			"heat_template_version": "2013-05-23",
			"description":           "Simple template to test heat commands",
			"parameters": map[string]interface{}{
				"flavor": map[string]interface{}{
					"default": "m1.tiny",
					"type":    "string",
				},
			},
			"resources": map[string]interface{}{
				"hello_world": map[string]interface{}{
					"type": "OS::Nova::Server",
					"properties": map[string]interface{}{
						"key_name": "heat_key",
						"flavor": map[string]interface{}{
							"get_param": "flavor",
						},
						"image":     "ad091b52-742f-469e-8f3c-fd81cadf0743",
						"user_data": "#!/bin/bash -xv\necho \"hello world\" &gt; /root/hello-world.txt\n",
					},
				},
			},
		},
	}
	validatedTemplate, err := stacktemplates.Validate(client, validateOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("validated template: %+v\n", validatedTemplate)
}
