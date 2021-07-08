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

type buddyResponseGlobalVariable struct {
	Url         string `json:"url"`
	Id          int    `json:"id"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	SSHKey      bool   `json:"ssh_key"`
	Settable    bool   `json:"settable"`
	Encrypted   bool   `json:"encrypted"`
	Description string `json:"description"`
}

type buddyRequestGlobalVariable struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Type        string `json:"type"`
	Settable    bool   `json:"settable"`
	Encrypted   bool   `json:"encrypted"`
	Description string `json:"description"`
}

type buddyClient interface {
	CreateGlobalVariable(key string, value string, varType string, description string, settable bool, encrypted bool) (*buddyResponseGlobalVariable, error)
	ReadGlobalVariable(id string) (*buddyResponseGlobalVariable, error)
	UpdateGlobalVariable(id string, key string, value string, varType string, description string, settable bool, encrypted bool) (*buddyResponseGlobalVariable, error)
}

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

func (b *buddyAdapter) CreateGlobalVariable(key string, value string, varType string, description string, settable bool, encrypted bool) (*buddyResponseGlobalVariable, error) {
	reqBody, err := json.Marshal(&buddyRequestGlobalVariable{
		Key:         key,
		Value:       value,
		Type:        varType,
		Description: description,
		Settable:    settable,
		Encrypted:   encrypted,
	})
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

	var data buddyResponseGlobalVariable
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

func (b *buddyAdapter) ReadGlobalVariable(id string) (*buddyResponseGlobalVariable, error) {
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

	var data buddyResponseGlobalVariable
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

func (b *buddyAdapter) UpdateGlobalVariable(id string, key string, value string, varType string, description string, settable bool, encrypted bool) (*buddyResponseGlobalVariable, error) {
	reqBody, err := json.Marshal(&buddyRequestGlobalVariable{
		Key:         key,
		Value:       value,
		Type:        varType,
		Description: description,
		Settable:    settable,
		Encrypted:   encrypted,
	})
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

	var data buddyResponseGlobalVariable
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
