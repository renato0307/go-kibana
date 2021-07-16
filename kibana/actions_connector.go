package kibana

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// GetConnector - Returns a single connector
// Check https://www.elastic.co/guide/en/kibana/7.13/get-connector-api.html
func (c *Client) GetConnector(connectorID string) (*Connector, error) {
	url := fmt.Sprintf("%s/s/%s/api/actions/connector/%s", c.HostURL, c.Space, connectorID)
	log.Printf("Calling %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	connector := Connector{}
	err = json.Unmarshal(body, &connector)
	if err != nil {
		return nil, err
	}

	return &connector, nil
}
