package kibana

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Kibana URL
const HostURL string = "http://localhost:5601"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Username   string
	Password   string
	Space      string
}

// NewClient -
func NewClient(host, username, password, space *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL, // Default Kibana URL
		Username:   *username,
		Password:   *password,
		Space:      *space,
	}

	if host != nil {
		c.HostURL = *host
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.Username, c.Password)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
