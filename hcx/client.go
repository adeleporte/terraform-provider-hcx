package hcx

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type updateConfigurationModule struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}

type updateConfigurationModuleBody struct {
	ID     int                       `json:"id"`
	Update updateConfigurationModule `json:"_update"`
}

type enterprise_get_object_groups struct {
	Type string `json:"type"`
}

// NewClient -
func NewClient(hcx, username *string, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
		// Default Hashicups URL
		HostURL: *hcx,
	}

	if (hcx != nil) && (username != nil) && (password != nil) {

		// form request body
		rb, err := json.Marshal(AuthStruct{
			Username: *username,
			Password: *password,
		})
		if err != nil {
			return nil, err
		}

		// authenticate
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/hybridity/api/sessions", c.HostURL), strings.NewReader(string(rb)))
		if err != nil {
			return nil, err
		}

		resp, _, err := c.doRequest(req)
		if err != nil {
			return nil, err
		}

		// parse response header
		c.Token = resp.Header.Get("x-hm-authorization")

	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) (*http.Response, []byte, error) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if c.Token != "" {
		req.Header.Set("x-hm-authorization", fmt.Sprintf("%s", c.Token))
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode != http.StatusAccepted {
			return nil, nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
		}
	}

	return res, body, err
}

// DeepCopy ...
func DeepCopy(src map[string]interface{}) (map[string]interface{}, error) {

	var dst map[string]interface{}

	if src == nil {
		fmt.Println("Error src")
		return nil, fmt.Errorf("src cannot be nil")
	}

	bytes, err := json.Marshal(src)

	if err != nil {
		return nil, fmt.Errorf("Unable to marshal src: %s", err)
	}
	err = json.Unmarshal(bytes, &dst)

	if err != nil {
		fmt.Println("Error unmarshal")
		fmt.Println(err)
		return nil, fmt.Errorf("Unable to unmarshal into dst: %s", err)
	}
	return dst, nil

}
