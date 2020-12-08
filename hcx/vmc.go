package hcx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type SDDC struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	CloudName        string `json:"cloudName,omitempty"`
	CloudURL         string `json:"cloudUrl,omitempty"`
	CloudType        string `json:"cloudType,omitempty"`
	CloudID          string `json:"cloudId,omitempty"`
	ActivationKey    string `json:"activationKey,omitempty"`
	SubscriptionID   string `json:"subscriptionId,omitempty"`
	ActivationStatus string `json:"activationStatus,omitempty"`
	DeploymentStatus string `json:"deploymentStatus,omitempty"`
	State            string `json:"state"`
}

type GetSddcsResults struct {
	SDDCs []SDDC `json:"sddcs"`
}

type VmcAccessToken struct {
	AccessToken  string `json:"access_token"`
	IDToken      string `json:"id_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refreshToken"`
}

type HcxCloudAuthorizationBody struct {
	Token string `json:"token"`
}

type ActivateHcxOnSDDCResults struct {
	JobID string `json:"jobId"`
}

type DeactivateHcxOnSDDCResults struct {
	JobID string `json:"jobId"`
}

// Authenticate

// Get SDDCs

// GET to https://connect.hcx.vmware.com/provider/csp/consumer/api/sddcs

// POST to https://connect.hcx.vmware.com/provider/csp/consumer/api/sddcs/SDDCID?action=activate (or desactivate)

func VmcAuthenticate(token string) (string, error) {

	c := Client{
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
		// Default Hashicups URL
		HostURL: "https://console.cloud.vmware.com/csp/gateway/am/api",
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/api-tokens/authorize?refresh_token=%s", c.HostURL, token), nil)
	if err != nil {
		return "", err
	}

	_, r, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	resp := VmcAccessToken{}
	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// parse response header
	return resp.AccessToken, nil

}

func HcxCloudAuthenticate(access_token string) (string, error) {

	c := Client{
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
		// Default Hashicups URL
		HostURL: "https://connect.hcx.vmware.com/provider/csp",
	}

	body := HcxCloudAuthorizationBody{
		Token: access_token,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/sessions", c.HostURL), buf)
	if err != nil {
		return "", err
	}

	resp, _, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	auth := resp.Header.Get("x-hm-authorization")
	if auth == "" {
		return "", errors.New("cannot authorize hcx cloud")
	}

	// parse response header
	return auth, nil

}

func GetSddc(hcx_auth, sddc_name string) (SDDC, error) {

	c := Client{
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
		// Default Hashicups URL
		HostURL: "https://connect.hcx.vmware.com/provider/csp/consumer",
		Token:   hcx_auth,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/sddcs", c.HostURL), nil)
	if err != nil {
		return SDDC{}, err
	}

	_, r, err := c.doRequest(req)
	if err != nil {
		return SDDC{}, err
	}

	resp := GetSddcsResults{}
	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return SDDC{}, err
	}

	for _, j := range resp.SDDCs {
		if j.Name == sddc_name {
			return j, nil
		}
	}

	// parse response header
	return SDDC{}, errors.New("cant find the sddc")

}

func ActivateHcxOnSDDC(hcx_auth, sddc_id string) (ActivateHcxOnSDDCResults, error) {

	resp := ActivateHcxOnSDDCResults{}

	c := Client{
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
		// Default Hashicups URL
		HostURL: "https://connect.hcx.vmware.com/provider/csp/consumer",
		Token:   hcx_auth,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/sddcs/%s?action=activate", c.HostURL, sddc_id), nil)
	if err != nil {
		return resp, err
	}

	_, r, err := c.doRequest(req)
	if err != nil {
		return resp, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	// parse response header
	return resp, nil

}

func DeactivateHcxOnSDDC(hcx_auth, sddc_id string) (DeactivateHcxOnSDDCResults, error) {

	resp := DeactivateHcxOnSDDCResults{}

	c := Client{
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
		// Default Hashicups URL
		HostURL: "https://connect.hcx.vmware.com/provider/csp/consumer",
		Token:   hcx_auth,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/sddcs/%s?action=deactivate", c.HostURL, sddc_id), nil)
	if err != nil {
		return resp, err
	}

	_, r, err := c.doRequest(req)
	if err != nil {
		return resp, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	// parse response header
	return resp, nil

}
