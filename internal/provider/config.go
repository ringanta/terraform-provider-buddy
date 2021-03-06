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

type buddyId struct {
	Id int `json:"id"`
}

type buddyWorkspaceMember struct {
	Url       string `json:"url"`
	HTMLURL   string `json:"html_url"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	Title     string `json:"title"`
	Email     string `json:"email"`
}

type buddyPermissionSet struct {
	URL                   string `json:"url"`
	HTMLURL               string `json:"html_url"`
	Id                    int    `json:"id"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	Type                  string `json:"type"`
	RepositoryAccessLevel string `json:"repository_access_level"`
	PipelineAccessLevel   string `json:"pipeline_access_level"`
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

type buddyResponseProjectMember struct {
	Url            string             `json:"url"`
	HTMLURL        string             `json:"html_url"`
	Id             int                `json:"id"`
	Name           string             `json:"name"`
	AvatarUrl      string             `json:"avatar_url"`
	Title          string             `json:"title"`
	Email          string             `json:"email"`
	Admin          bool               `json:"admin"`
	WorkspaceOwner bool               `json:"workspace_owner"`
	PermissionSet  buddyPermissionSet `json:"permission_set"`
}

type buddyResponseListWorkspaceMember struct {
	Url     string                 `json:"url"`
	HTMLURL string                 `json:"html_url"`
	Members []buddyWorkspaceMember `json:"members"`
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

type buddyRequestProjectMember struct {
	Id            string  `json:"id"`
	PermissionSet buddyId `json:"permission_set"`
}

type buddyRequestPermissionSet struct {
	PermissionSet buddyId `json:"permission_set"`
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

	SetAdminRight(id string, admin bool) (*buddyResponseWorkspaceMember, error)

	CreateProjectMember(projectName string, variable buddyRequestProjectMember) (*buddyResponseProjectMember, error)
	ReadProjectMember(projectName string, memberId string) (*buddyResponseProjectMember, error)
	UpdateProjectMember(projectName string, memberId string, variable buddyRequestPermissionSet) (*buddyResponseProjectMember, error)
	DeleteProjectMember(projectName string, memberId string) error

	GetUser(email string) (*buddyWorkspaceMember, error)
}
