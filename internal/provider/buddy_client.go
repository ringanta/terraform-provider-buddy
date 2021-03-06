package provider

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func newBuddyClient(c *Config) *buddyAdapter {
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	config := &tls.Config{
		InsecureSkipVerify: !c.VerifySSL,
		RootCAs:            rootCAs,
	}
	tr := &http.Transport{TLSClientConfig: config}
	httpClient := &http.Client{Transport: tr}

	return &buddyAdapter{BuddyURL: strings.TrimSuffix(c.BuddyURL, "/"), Token: c.Token, Client: httpClient}
}

func (b *buddyAdapter) CreateWorkspaceVariable(variable buddyRequestWorkspaceVariable) (*buddyResponseWorkspaceVariable, error) {
	reqBody, err := json.Marshal(&variable)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/%v", b.BuddyURL, "variables"), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", b.Token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", user_agent)

	resp, err := b.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data buddyResponseWorkspaceVariable
	if resp.StatusCode != 201 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Expected return code is 201 but got %v. Failed to read response body with the following message: %v", resp.StatusCode, err.Error())
		}
		return nil, fmt.Errorf("Expected return code is 201 but got %v with the following response body %v", resp.StatusCode, string(body))
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (b *buddyAdapter) ReadWorkspaceVariable(id string) (*buddyResponseWorkspaceVariable, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v/%v", b.BuddyURL, "variables", id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", b.Token))
	req.Header.Set("User-Agent", user_agent)

	resp, err := b.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data buddyResponseWorkspaceVariable
	if resp.StatusCode == 404 {
		return &data, nil
	}

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Expected return code is 200 but got %v. Failed to read response body with the following message: %v", resp.StatusCode, err.Error())
		}
		return nil, fmt.Errorf("Expected return code is 200 but got %v with the following response body %v", resp.StatusCode, string(body))
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (b *buddyAdapter) UpdateWorkspaceVariable(id string, variable buddyRequestWorkspaceVariable) (*buddyResponseWorkspaceVariable, error) {
	reqBody, err := json.Marshal(&variable)
	if err != nil {
		return nil, err
	}

	urlPath := fmt.Sprintf("variables/%v", id)

	response, err := b.doPatch(urlPath, reqBody)
	if err != nil {
		return nil, err
	}

	var data buddyResponseWorkspaceVariable

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (b *buddyAdapter) CreateProjectVariable(variable buddyRequestProjectVariable) (*buddyResponseProjectVariable, error) {
	reqBody, err := json.Marshal(&variable)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/%v", b.BuddyURL, "variables"), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", b.Token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", user_agent)

	resp, err := b.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data buddyResponseProjectVariable
	if resp.StatusCode != 201 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Expected return code is 201 but got %v. Failed to read response body with the following message: %v", resp.StatusCode, err.Error())
		}
		return nil, fmt.Errorf("Expected return code is 201 but got %v with the following response body %v", resp.StatusCode, string(body))
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (b *buddyAdapter) ReadProjectVariable(id string) (*buddyResponseProjectVariable, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v/%v", b.BuddyURL, "variables", id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", b.Token))
	req.Header.Set("User-Agent", user_agent)

	resp, err := b.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data buddyResponseProjectVariable
	if resp.StatusCode == 404 {
		return &data, nil
	}

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Expected return code is 200 but got %v. Failed to read response body with the following message: %v", resp.StatusCode, err.Error())
		}
		return nil, fmt.Errorf("Expected return code is 200 but got %v with the following response body %v", resp.StatusCode, string(body))
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (b *buddyAdapter) UpdateProjectVariable(id string, variable buddyRequestProjectVariable) (*buddyResponseProjectVariable, error) {
	reqBody, err := json.Marshal(&variable)
	if err != nil {
		return nil, err
	}

	urlPath := fmt.Sprintf("variables/%v", id)

	response, err := b.doPatch(urlPath, reqBody)
	if err != nil {
		return nil, err
	}

	var data buddyResponseProjectVariable

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (b *buddyAdapter) DeleteVariable(id string) error {
	urlPath := fmt.Sprintf("variables/%v", id)

	return b.doDelete(urlPath)
}

func (b *buddyAdapter) CreateWorkspaceMember(email string) (*buddyResponseWorkspaceMember, error) {
	reqBody, err := json.Marshal(&struct {
		Email string `json:"email"`
	}{
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/%v", b.BuddyURL, "members"), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", b.Token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", user_agent)

	resp, err := b.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data buddyResponseWorkspaceMember
	if resp.StatusCode != 201 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Expected return code is 201 but got %v. Failed to read response body with the following message: %v", resp.StatusCode, err.Error())
		}
		return nil, fmt.Errorf("Expected return code is 201 but got %v with the following response body %v", resp.StatusCode, string(body))
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (b *buddyAdapter) ReadWorkspaceMember(id string) (*buddyResponseWorkspaceMember, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v/%v", b.BuddyURL, "members", id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", b.Token))
	req.Header.Set("User-Agent", user_agent)

	resp, err := b.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data buddyResponseWorkspaceMember
	if resp.StatusCode == 404 {
		return &data, nil
	}

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Expected return code is 200 but got %v. Failed to read response body with the following message: %v", resp.StatusCode, err.Error())
		}
		return nil, fmt.Errorf("Expected return code is 200 but got %v with the following response body %v", resp.StatusCode, string(body))
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (b *buddyAdapter) DeleteWorkspaceMember(id string) error {
	urlPath := fmt.Sprintf("members/%v", id)

	return b.doDelete(urlPath)
}

func (b *buddyAdapter) SetAdminRight(id string, admin bool) (*buddyResponseWorkspaceMember, error) {
	reqBody, err := json.Marshal(&struct {
		Admin bool `json:"admin"`
	}{
		Admin: admin,
	})
	if err != nil {
		return nil, err
	}

	urlPath := fmt.Sprintf("members/%v", id)

	response, err := b.doPatch(urlPath, reqBody)
	if err != nil {
		return nil, err
	}

	var data buddyResponseWorkspaceMember
	err = json.NewDecoder(bytes.NewReader(response)).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *buddyAdapter) CreateProjectMember(projectName string, variable buddyRequestProjectMember) (*buddyResponseProjectMember, error) {
	reqBody, err := json.Marshal(&variable)
	if err != nil {
		return nil, err
	}

	urlPath := fmt.Sprintf("%v/%v/%v", "projects", projectName, "members")
	response, err := b.doCreate(urlPath, reqBody)
	if err != nil {
		return nil, err
	}

	var data buddyResponseProjectMember
	err = json.NewDecoder(bytes.NewReader(response)).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *buddyAdapter) ReadProjectMember(projectName string, memberId string) (*buddyResponseProjectMember, error) {
	urlPath := fmt.Sprintf("%v/%v/%v/%v", "projects", projectName, "members", memberId)
	var data buddyResponseProjectMember

	response, err := b.doRead(urlPath)
	if err != nil {
		return nil, err
	}

	if response == nil {
		return &data, nil
	}

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *buddyAdapter) UpdateProjectMember(projectName string, memberId string, variable buddyRequestPermissionSet) (*buddyResponseProjectMember, error) {
	urlPath := fmt.Sprintf("%v/%v/%v/%v", "projects", projectName, "members", memberId)
	var data buddyResponseProjectMember

	reqBody, err := json.Marshal(&variable)
	if err != nil {
		return nil, err
	}

	response, err := b.doPatch(urlPath, reqBody)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *buddyAdapter) DeleteProjectMember(projectName string, memberId string) error {
	urlPath := fmt.Sprintf("%v/%v/%v/%v", "projects", projectName, "members", memberId)

	return b.doDelete(urlPath)
}

func (b *buddyAdapter) GetUser(email string) (*buddyWorkspaceMember, error) {
	response, err := b.listUsers(1, 1000)
	if err != nil {
		return nil, err
	}

	var data buddyWorkspaceMember
	for _, member := range response.Members {
		if member.Email == email {
			data = member
		}
	}

	return &data, nil
}

func (b *buddyAdapter) listUsers(pageNo int, userPerPage int) (*buddyResponseListWorkspaceMember, error) {
	urlPath := fmt.Sprintf("members?page=%v&per_page=%v&sort_name=name", pageNo, userPerPage)
	var data buddyResponseListWorkspaceMember

	response, err := b.doRead(urlPath)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(response)).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *buddyAdapter) doRead(urlPath string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", b.BuddyURL, urlPath), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", b.Token))
	req.Header.Set("User-Agent", user_agent)

	resp, err := b.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []byte

	if resp.StatusCode == 404 {
		return data, nil
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body while performing doRead with the following error message: %v", err.Error())
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Expected return code is 200 but got %v with the following response body %v", resp.StatusCode, string(data))
	}

	return data, nil
}

func (b *buddyAdapter) doCreate(urlPath string, reqBody []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/%v", b.BuddyURL, urlPath), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", b.Token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", user_agent)

	resp, err := b.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body while performing doCreate with the following error message: %v", err.Error())
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("Expected return code is 201 but got %v with the following response body %v", resp.StatusCode, string(data))
	}

	return data, nil
}

func (b *buddyAdapter) doPatch(urlPath string, reqBody []byte) ([]byte, error) {
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%v/%v", b.BuddyURL, urlPath), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", b.Token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", user_agent)

	resp, err := b.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body with the following error message: %v", err.Error())
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Expected return code is 200 but got %v with the following response body %v", resp.StatusCode, string(data))
	}

	return data, nil
}

func (b *buddyAdapter) doDelete(urlPath string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%v/%v", b.BuddyURL, urlPath), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", b.Token))
	req.Header.Set("User-Agent", user_agent)

	resp, err := b.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Expected return code is 204 but got %v. Failed to read response body with the following message: %v", resp.StatusCode, err.Error())
		}
		return fmt.Errorf("Expected return code is 204 but got %v with the following response body %v", resp.StatusCode, string(body))
	}

	return nil
}
