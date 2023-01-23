package users

import "github.com/kinvey/terraform-provider-spotinstadmin/client"

// Service ...
type Service struct {
	httpClient *client.Client
}

type User struct {
	ID       string `json:"userId"`
	UserName string `json:"username"`
	Type     string `json:"type"`
}

type UserDetails struct {
	ID          string `json:"userId"`
	AccessToken string `json:"accessToken"`
	UserName    string `json:"username"`
	Description string `json:"description"`
}

type createProgrammaticUserAccount struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

type createProgrammaticUserRequest struct {
	Accounts    []createProgrammaticUserAccount `json:"accounts"`
	Description string                          `json:"description"`
	Name        string                          `json:"name"`
}

type createProgrammaticUserResponse struct {
	Token string `json:"token"`
	Name  string `json:"name"`
	ID    string `json:"id"`
}
