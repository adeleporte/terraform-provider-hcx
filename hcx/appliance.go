package hcx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type InsertvCenterBody struct {
	Data InsertvCenterData `json:"data"`
}

type InsertvCenterData struct {
	Items []InsertvCenterDataItem `json:"items"`
}

type InsertvCenterDataItem struct {
	Config InsertvCenterDataItemConfig `json:"config"`
}

type InsertvCenterDataItemConfig struct {
	URL      string `json:"url"`
	Username string `json:"userName"`
	Password string `json:"password"`
}

type InsertvCenterResult struct {
	InsertvCenterData InsertvCenterData `json:"data"`
}

// InsertvCenter ...
func InsertvCenter(c *Client, body InsertvCenterBody) (InsertvCenterResult, error) {

	resp := InsertvCenterResult{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s:9443/api/admin/global/config/vcenter", c.HostURL), buf)
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
