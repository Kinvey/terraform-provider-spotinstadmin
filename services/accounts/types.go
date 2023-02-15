package accounts

import "github.com/kinvey/terraform-provider-spotinstadmin/client"

// Service is a client for creating accounts
type Service struct {
	httpClient *client.Client
}

// Account represesnts Spotinst account in API
type Account struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	OrganizationID     string `json:"organizationId"`
	ProviderExternalID string `json:"providerExternalId,omitempty"`
}

type ExternalID struct {
	ID         string `json:"externalId"`
	Expiration string `json:"maxValidUntil"`
}

// AccountNotFoundError is raised when looking up account
// fails because there's no such account in Spotinst:w
type AccountNotFoundError struct {
	AccountID string
}
