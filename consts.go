package main

const (
	providerTokenAttrKey    = "token"
	providerEmailAttrKey    = "email"
	providerPasswordAttrKey = "password"
)

const (
	envSpotinstTokenKey    = "SPOTINST_TOKEN"
	envSpotinstEmailKey    = "SPOTINST_EMAIL"
	envSpotinstPasswordKey = "SPOTINST_PASSWORD"
)

const (
	providerName                 = "spotinstadmin"
	accountResourceName          = providerName + "_account"
	programmaticUserResourceName = providerName + "_programmatic_user"
	linkAccountResourceName      = providerName + "_account_aws_link"
	externalIDResourceName       = providerName + "_account_external_id"
)

const (
	accountResourceNameAttrKey           = "name"
	accountResourceorganizationIdAttrKey = "organization_id"
)

const (
	userResourceNameAttrKey        = "name"
	userResourceAccountIDAttrKey   = "account_id"
	userResourceDescriptionAttrKey = "description"
	userResourceAccessTokenAttrKey = "access_token"
)

const (
	linkAccountResourceAccountIDAttrKey          = "account_id"
	linkAccountResourceRoleArnAttrKey            = "aws_role_arn"
	linkAccountResourceProviderExternalIdAttrKey = "provider_external_id"
	linkAccountResourceOrganizationIdAttrKey     = "organization_id"
)

const (
	awsExternalIDResourceAccountIDAttrKey         = "account_id"
	awsExternalIDResourceAttrKeyExternalIDAttrKey = "external_id"
)
