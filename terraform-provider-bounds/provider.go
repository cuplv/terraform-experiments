package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"strconv"
)

type BoundsProvider struct {
	schema.Provider
}

func (p *BoundsProvider) Validate(c *terraform.ResourceConfig) ([]string, []error) {
	s,e := p.Provider.Validate(c)
	return append(s, "boundsinfo:budget;" + strconv.Itoa(c.Raw["budget"].(int))), e
}

func (p *BoundsProvider) ValidateResource(t string, c *terraform.ResourceConfig) ([]string, []error) {
	s,e := p.Provider.ValidateResource(t,c)
	return append(s, "boundsinfo:cost;1"), e
}

func Provider() terraform.ResourceProvider {
	return &BoundsProvider{schema.Provider{
		Schema: map[string]*schema.Schema{
			"budget": &schema.Schema{
				Type: schema.TypeInt,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"bounds_thing": boundsThing(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return nil, nil
		},
	}}
}

func boundsThing() *schema.Resource {
	return &schema.Resource{
		Create: boundsCreate,
		Read: boundsRead,
		Delete: boundsDelete,
		Exists: boundsExists,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func boundsCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId(d.Get("name").(string))
	return nil
}

func boundsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func boundsDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func boundsExists(d *schema.ResourceData, m interface{}) (bool, error) {
	return true, nil
}
