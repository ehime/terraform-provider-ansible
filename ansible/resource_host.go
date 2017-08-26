package ansible

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHost() *schema.Resource {
	return &schema.Resource{
		Create: resourceHostCreate,
		Read:   resourceHostRead,
		Delete: resourceHostDelete,

		Schema: map[string]*schema.Schema{
			"inventory_hostname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				ForceNew: true,
			},
			"vars": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceHostCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(d.Get("inventory_hostname").(string))

	return nil
}

func resourceHostRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHostDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
