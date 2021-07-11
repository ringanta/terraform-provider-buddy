package provider

import (
	"net/http"
)

type Config struct {
	BuddyURL  string
	Token     string
	VerifySSL bool
}

type buddyAdapter struct {
	BuddyURL string
	Token    string
	*http.Client
}

type buddyProject struct {
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Status      string `json:"status"`
}

type buddyResponseWorkspaceVariable struct {
	Url         string `json:"url"`
	Id          int    `json:"id"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	SSHKey      bool   `json:"ssh_key"`
	Settable    bool   `json:"settable"`
	Encrypted   bool   `json:"encrypted"`
	Description string `json:"description"`
}

type buddyResponseProjectVariable struct {
	Url         string       `json:"url"`
	Id          int          `json:"id"`
	Key         string       `json:"key"`
	Value       string       `json:"value"`
	SSHKey      bool         `json:"ssh_key"`
	Settable    bool         `json:"settable"`
	Encrypted   bool         `json:"encrypted"`
	Description string       `json:"description"`
	Project     buddyProject `json:"project"`
}

type buddyResponseWorkspaceMember struct {
	Url            string `json:"url"`
	HTMLURL        string `json:"html_url"`
	Id             int    `json:"id"`
	Name           string `json:"name"`
	AvatarUrl      string `json:"avatar_url"`
	Title          string `json:"title"`
	Email          string `json:"email"`
	Admin          bool   `json:"admin"`
	WorkspaceOwner bool   `json:"workspace_owner"`
}

type buddyRequestWorkspaceVariable struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Type        string `json:"type"`
	Settable    bool   `json:"settable"`
	Encrypted   bool   `json:"encrypted"`
	Description string `json:"description"`
}

type buddyRequestProject struct {
	Name string `json:"name"`
}

type buddyRequestProjectVariable struct {
	Key         string              `json:"key"`
	Value       string              `json:"value"`
	Type        string              `json:"type"`
	Settable    bool                `json:"settable"`
	Encrypted   bool                `json:"encrypted"`
	Description string              `json:"description"`
	Project     buddyRequestProject `json:"project"`
}

type buddyClient interface {
	CreateWorkspaceVariable(variable buddyRequestWorkspaceVariable) (*buddyResponseWorkspaceVariable, error)
	ReadWorkspaceVariable(id string) (*buddyResponseWorkspaceVariable, error)
	UpdateWorkspaceVariable(id string, variable buddyRequestWorkspaceVariable) (*buddyResponseWorkspaceVariable, error)
	CreateProjectVariable(variable buddyRequestProjectVariable) (*buddyResponseProjectVariable, error)
	ReadProjectVariable(id string) (*buddyResponseProjectVariable, error)
	UpdateProjectVariable(id string, variable buddyRequestProjectVariable) (*buddyResponseProjectVariable, error)
	DeleteVariable(id string) error
	CreateWorkspaceMember(email string) (*buddyResponseWorkspaceMember, error)
	ReadWorkspaceMember(id string) (*buddyResponseWorkspaceMember, error)
	DeleteWorkspaceMember(id string) error
}
