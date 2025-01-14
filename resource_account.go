package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/kinvey/terraform-provider-spotinstadmin/services/accounts"
)

func resourceAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountCreate,
		Read:   resourceAccountRead,
		Update: resourceAccountUpdate,
		Delete: resourceAccountDelete,

		Schema: map[string]*schema.Schema{
			accountResourceNameAttrKey: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			accountResourceorganizationIdAttrKey: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAccountCreate(d *schema.ResourceData, m interface{}) error {
	accountsService := m.(*Meta).accountsService
	name := d.Get(accountResourceNameAttrKey).(string)

	out, err := accountsService.Create(name)

	if err != nil {
		return err
	}

	d.SetId(out.ID)

	d.Set("organization_id", out.OrganizationID)

	return resourceAccountRead(d, m)
}

func resourceAccountRead(d *schema.ResourceData, m interface{}) error {
	accountsService := m.(*Meta).accountsService
	obj, err := accountsService.Get(d.Id())
	if err != nil {
		if accounts.IsAccountNotFoundErr(err) {
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set(accountResourceNameAttrKey, obj.Name)
	d.Set("organization_id", obj.OrganizationID)
	d.Set("provider_external_id", obj.ProviderExternalID)

	return nil
}

func resourceAccountUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceAccountRead(d, m)
}

func resourceAccountDelete(d *schema.ResourceData, m interface{}) error {
	accountsService := m.(*Meta).accountsService
	return accountsService.Delete(d.Id())
}
