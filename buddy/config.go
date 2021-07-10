package buddy

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
	buddyResponseWorkspaceVariable
	Project buddyProject `json:"project"`
}

type buddyRequestWorkspaceVariable struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Type        string `json:"type"`
	Settable    bool   `json:"settable"`
	Encrypted   bool   `json:"encrypted"`
	Description string `json:"description"`
}

type buddyProject struct {
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Status      string `json:"status"`
}

type buddyRequestProjectVariable struct {
	buddyRequestWorkspaceVariable
	Project string `json:"project"`
}

type buddyClient interface {
	CreateWorkspaceVariable(variable buddyRequestWorkspaceVariable) (*buddyResponseWorkspaceVariable, error)
	ReadWorkspaceVariable(id string) (*buddyResponseWorkspaceVariable, error)
	UpdateWorkspaceVariable(id string, variable buddyRequestWorkspaceVariable) (*buddyResponseWorkspaceVariable, error)
	DeleteWorkspaceVariable(id string) error
	CreateProjectVariable(variable buddyRequestProjectVariable) (*buddyResponseProjectVariable, error)
	ReadProjectVariable(id string) (*buddyResponseProjectVariable, error)
}
