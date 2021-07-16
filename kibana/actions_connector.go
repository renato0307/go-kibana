package kibana

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetConnector - Retrieves a connector by ID.
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

// CreateConnector - Creates a connector.
// Check https://www.elastic.co/guide/en/kibana/7.13/create-connector-api.html
func (c *Client) CreateConnector(connector CreateConnector) (*Connector, error) {
	rb, err := json.Marshal(connector)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/s/%s/api/actions/connector", c.HostURL, c.Space)
	log.Printf("Calling %s", url)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("kbn-xsrf", "true")
	req.Header.Set("content-type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newConnector := Connector{}
	err = json.Unmarshal(body, &newConnector)
	if err != nil {
		return nil, err
	}

	return &newConnector, nil
}

// UpdateConnector - Updates the attributes for an existing connector.
// Check https://www.elastic.co/guide/en/kibana/7.13/update-connector-api.html
func (c *Client) UpdateConnector(connectorID string, connector UpdateConnector) (*Connector, error) {
	rb, err := json.Marshal(connector)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/s/%s/api/actions/connector/%s", c.HostURL, c.Space, connectorID)
	log.Printf("Calling %s", url)
	req, err := http.NewRequest("PUT", url, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("kbn-xsrf", "true")
	req.Header.Set("content-type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updatedConnector := Connector{}
	err = json.Unmarshal(body, &updatedConnector)
	if err != nil {
		return nil, err
	}

	return &updatedConnector, nil
}
