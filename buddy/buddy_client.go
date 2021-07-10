package buddy

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

	fmt.Println(reqBody)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%v/%v/%v", b.BuddyURL, "variables", id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", b.Token))
	req.Header.Set("Content-Type", "application/json")
	resp, err := b.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data buddyResponseWorkspaceVariable
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

	fmt.Println(reqBody)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%v/%v/%v", b.BuddyURL, "variables", id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", b.Token))
	req.Header.Set("Content-Type", "application/json")
	resp, err := b.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data buddyResponseProjectVariable
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

func (b *buddyAdapter) DeleteVariable(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%v/%v/%v", b.BuddyURL, "variables", id), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", b.Token))
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
