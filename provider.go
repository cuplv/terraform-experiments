package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"fmt"
)

func Provider() terraform.ResourceProvider {
	// return &schema.Provider{
	// 	Schema: map[string]*schema.Schema{
	// 		"foo": &schema.Schema{
	// 			Type: schema.TypeString,
	// 			Required: true,
	// 		},
	// 	},
	// 	ResourcesMap: map[string]*schema.Resource{
	// 		"bounds_baz": bazThing(),
	// 	},
	// 	ConfigureFunc: 
	// }
	return &terraform.MockResourceProvider{
		ConfigureFn: FakeConfigure,
		// GetSchemaReturn: &terraform.ProviderSchema{
		// 	ResourceTypes: map[string]*configschema.Block
		// },
		GetSchemaReturn: GetSchemaS(),
	}
}

func GetSchemaS() *terraform.ProviderSchema {
	a, _ := SProvider().GetSchema(&terraform.ProviderSchemaRequest{})
	return a
}

func SProvider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"foo": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"bounds_baz": bazThing(),
		},
	}
}

func FakeConfigure(*terraform.ResourceConfig) error {
	fmt.Printf("hello, world\n")
	return nil
}

func bazThing() *schema.Resource {
	return &schema.Resource{
		Create: bazCreate,
		Delete: bazDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
		},
		Update: bazUpdate,
		Read: bazRead,
	}
}

func bazCreate(d *schema.ResourceData, m interface{}) error {
	fmt.Printf("baz create\n")
	return nil
}

func bazDelete(d *schema.ResourceData, m interface{}) error {
	fmt.Printf("baz delete\n")
	return nil
}

func bazUpdate(d *schema.ResourceData, m interface{}) error {
	fmt.Printf("baz update\n")
	return nil
}

func bazRead(d *schema.ResourceData, m interface{}) error {
	fmt.Printf("baz read\n")
	return nil
}
