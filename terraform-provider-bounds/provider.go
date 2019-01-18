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

var globalCounter *counter = initCounter(1)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"foo": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"bounds_thing": boundsThing(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return initCounter(1), nil
		},
	}
}

func boundsThing() *schema.Resource {
	return &schema.Resource{
		Create: boundsCreate,
		Delete: boundsDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
		},
		Update: boundsUpdate,
		Read: boundsRead,
	}
}

func boundsCreate(d *schema.ResourceData, m interface{}) error {
	return globalCounter.decOrDie()
}

func boundsDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func boundsUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func boundsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
