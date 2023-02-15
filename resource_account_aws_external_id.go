package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAccountAWSExternalID() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountAWSExternalIDCreate,
		Read:   resourceAccountAWSExternalIDRead,
		Delete: resourceAccountAWSExternalIDDelete,

		Schema: map[string]*schema.Schema{
			awsExternalIDResourceAccountIDAttrKey: {
				Type:        schema.TypeString,
				Description: "Account ID to link",
				ForceNew:    true,
				Required:    true,
			},
			awsExternalIDResourceAttrKeyExternalIDAttrKey: {
				Type:        schema.TypeString,
				Description: "ExternalID output",
				Computed:    true,
			},
		},
	}
}

func resourceAccountAWSExternalIDCreate(d *schema.ResourceData, m interface{}) error {
	accountsService := m.(*Meta).accountsService
	accountID := d.Get(awsExternalIDResourceAccountIDAttrKey).(string)

	externaID, err := accountsService.CreateExternalId(accountID)
	if err != nil {
		return err
	}

	d.SetId(accountID)
	d.Set(awsExternalIDResourceAttrKeyExternalIDAttrKey, externaID.ID)

	return nil
}

func resourceAccountAWSExternalIDRead(d *schema.ResourceData, m interface{}) error {
	externalId := d.State().Attributes[awsExternalIDResourceAttrKeyExternalIDAttrKey]
	if externalId == "" {
		d.SetId("")

		return nil
	}

	d.Set(awsExternalIDResourceAttrKeyExternalIDAttrKey, externalId)

	return nil
}

func resourceAccountAWSExternalIDDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
