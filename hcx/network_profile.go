package hcx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type InsertNetworkProfileBody struct {
	Backings        []Backing `json:"backings"`
	Description     string    `json:"description"`
	Enterprise      string    `json:"enterprise"`
	IPScopes        []IPScope `json:"ipScopes"`
	MTU             int       `json:"mtu"`
	Name            string    `json:"name"`
	L3TenantManaged bool      `json:"l3TenantManaged"`
	OwnedBySystem   bool      `json:"ownedBySystem"`
}

type Backing struct {
	BackingID           string `json:"backingId"`
	BackingName         string `json:"backingName"`
	Type                string `json:"type"`
	VCenterInstanceUuid string `json:"vCenterInstanceUuid"`
}

type IPScope struct {
	DnsSuffix       string           `json:"dnsSuffix,omitempty"`
	Gateway         string           `json:"gateway,omitempty"`
	PrefixLength    int              `json:"prefixLength"`
	PrimaryDns      string           `json:"primaryDns,omitempty"`
	SecondaryDns    string           `json:"secondaryDns,omitempty"`
	NetworkIpRanges []NetworkIpRange `json:"networkIpRanges"`
}

type NetworkIpRange struct {
	EndAddress   string `json:"endAddress"`
	StartAddress string `json:"startAddress"`
}

type GetNetworkProfileResult struct {
	Backings        []Backing `json:"backings"`
	Description     string    `json:"description"`
	Enterprise      string    `json:"enterprise"`
	IPScopes        []IPScope `json:"ipScopes"`
	MTU             int       `json:"mtu"`
	Name            string    `json:"name"`
	L3TenantManaged bool      `json:"l3TenantManaged"`
	OwnedBySystem   bool      `json:"ownedBySystem"`
	ObjectId        string    `json:"objectId"`
}

type NetworkProfileResult struct {
	Success   bool               `json:"success"`
	Completed bool               `json:"completed"`
	Time      int                `json:"time"`
	Data      NetworkProfileData `json:"data"`
}

type NetworkProfileData struct {
	JobID    string `json:"jobId"`
	ObjectId string `json:"objectId"`
}

type UpdateNetworkProfileBody struct {
	Backings        []Backing `json:"backings"`
	Description     string    `json:"description"`
	Enterprise      string    `json:"enterprise"`
	IPScopes        []IPScope `json:"ipScopes"`
	MTU             int       `json:"mtu"`
	Name            string    `json:"name"`
	L3TenantManaged bool      `json:"l3TenantManaged"`
	OwnedBySystem   bool      `json:"ownedBySystem"`
	ObjectId        string    `json:"objectId"`
}

// InsertNetworkProfile ...
func InsertNetworkProfile(c *Client, body InsertNetworkProfileBody) (NetworkProfileResult, error) {

	resp := NetworkProfileResult{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/admin/hybridity/api/networks", c.HostURL), buf)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	// Send the request
	_, r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	return resp, nil
}

// GetNetworkProfile ...
func GetNetworkProfile(c *Client, name string) (string, error) {

	resp := []GetNetworkProfileResult{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/hybridity/api/networks", c.HostURL), nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Send the request
	_, r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	for _, j := range resp {
		if j.Name == name {
			return j.ObjectId, nil
		}
	}

	return "", errors.New("cannot find network profile")
}

// DeleteNetworkProfile ...
func DeleteNetworkProfile(c *Client, networkID string) (NetworkProfileResult, error) {

	resp := NetworkProfileResult{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/hybridity/api/networks/%s", c.HostURL, networkID), nil)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	// Send the request
	_, r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	return resp, nil
}

// UpdateNetworkProfile ...
func UpdateNetworkProfile(c *Client, body UpdateNetworkProfileBody) (NetworkProfileResult, error) {

	resp := NetworkProfileResult{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/hybridity/api/networks/%s", c.HostURL, body.ObjectId), buf)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	// Send the request
	_, r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	return resp, nil
}
