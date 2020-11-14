package mikrotik

import (
	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePool() *schema.Resource {
	return &schema.Resource{
		Create: resourcePoolCreate,
		Read:   resourcePoolRead,
		Update: resourcePoolUpdate,
		Delete: resourcePoolDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ranges": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nextpool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourcePoolCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	ranges := d.Get("ranges").(string)
	comment := d.Get("comment").(string)

	c := m.(client.Mikrotik)

	pool, err := c.AddPool(name, ranges, comment)
	if err != nil {
		return err
	}

	poolToData(pool, d)
	return nil
}

func resourcePoolRead(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	pool, err := c.FindPool(d.Id())

	if err != nil {
		d.SetId("")
		return nil
	}

	if pool == nil {
		d.SetId("")
		return nil
	}

	poolToData(pool, d)
	return nil
}

func resourcePoolUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	name := d.Get("name").(string)
	ranges := d.Get("ranges").(string)
	comment := d.Get("comment").(string)
	nextpool := d.Get("nextpool").(string)

	pool, err := c.UpdatePool(d.Id(), name, ranges, comment, nextpool)

	if err != nil {
		return err
	}

	poolToData(pool, d)
	return nil
}

func resourcePoolDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(client.Mikrotik)

	err := c.DeletePool(d.Id())

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func poolToData(pool *client.Pool, d *schema.ResourceData) error {
	d.SetId(pool.Id)
	d.Set("name", pool.Name)
	d.Set("ranges", pool.Ranges)
	d.Set("comment", pool.Comment)
	d.Set("nextpool", pool.NextPool)
	return nil
}
