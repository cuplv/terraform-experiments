package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"sync"
	"errors"
)

type counter struct {
	v int
	mux sync.Mutex
}

func (c *counter) inc() error {
	c.mux.Lock()
	c.v ++
	c.mux.Unlock()
	return nil
}

func (c *counter) decOrDie() error {
	c.mux.Lock()
	if c.v <= 0 {
		c.mux.Unlock()
		return errors.New("Out of resources")
	} else {
		c.v --
		c.mux.Unlock()
		return nil
	}
}

func initCounter(v0 int) *counter {
	return &counter{ v: v0 }
}

// var globalCounter *counter

type BoundsProvider struct {
	schema.Provider
}

func (p *BoundsProvider) Validate(c *terraform.ResourceConfig) ([]string, []error) {
	s,e := p.Provider.Validate(c)
	return append(s, "boundsinfo:budget:foo"), e
}

func (p *BoundsProvider) ValidateResource(t string, c *terraform.ResourceConfig) ([]string, []error) {
	s,e := p.Provider.ValidateResource(t,c)
	return append(s, "boundsinfo:cost:bar"), e
}

func Provider() terraform.ResourceProvider {
	return &BoundsProvider{schema.Provider{
		Schema: map[string]*schema.Schema{
			"allowance": &schema.Schema{
				Type: schema.TypeInt,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"bounds_thing": boundsThing(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return initCounter(d.Get("allowance").(int)), nil
			// return initCounter(5), nil
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
	return m.(*counter).decOrDie()
}

func boundsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func boundsDelete(d *schema.ResourceData, m interface{}) error {
	return m.(*counter).inc()
	return nil
}

func boundsExists(d *schema.ResourceData, m interface{}) (bool, error) {
	return true, nil
}
