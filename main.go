// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-helm/helm"
	"log"
	"testing"
)

// Generate docs for website
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func TestMain(t *testing.T) {
	fmt.Println("Hello, World!")
	// Mock schema.ResourceData
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"values": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"set": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"value": {
						Type:     schema.TypeString,
						Required: true,
					},
					"type": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateFunc: validation.StringInSlice([]string{
							"auto", "string",
						}, false),
					},
				},
			},
		},
		"set_list": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"value": {
						Type:     schema.TypeList,
						Required: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
		"set_sensitive": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"value": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	}, map[string]interface{}{
		"set": []interface{}{
			map[string]interface{}{
				"name":  "extraArgs.scale-down-unneeded-time",
				"value": "2m",
			},
			map[string]interface{}{
				"name":  "extraArgs.scale-down-unneeded-time",
				"value": "110m",
			},
		},
	})

	// Call getValues
	values, err := helm.GetValues(d)
	if err != nil {
		log.Fatalf("Error getting values: %v", err)
	}

	// Inspect the output
	fmt.Printf("Values: %v\n", values)
}

func main() {
	// Run the test
	TestMain(&testing.T{})
}
