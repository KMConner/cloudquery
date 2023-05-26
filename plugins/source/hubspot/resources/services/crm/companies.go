package crm

import (
	"github.com/clarkmcc/go-hubspot/generated/v3/companies"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Companies() *schema.Table {
	return &schema.Table{
		Name:        "hubspot_crm_companies",
		Description: "https://developers.hubspot.com/docs/api/crm/companies",
		Resolver:    fetchCompanies,
		Transform: transformers.TransformWithStruct(
			companies.SimplePublicObjectWithAssociations{},
			transformers.WithPrimaryKeys("Id"),
			transformers.WithSkipFields("PropertiesWithHistory"),
		),
	}
}
