package kibana

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetRule - Retrieve a rule by ID.
// Check https://www.elastic.co/guide/en/kibana/7.13/get-rule-api.html
func (c *Client) GetRule(ruleID string) (*Rule, error) {
	url := fmt.Sprintf("%s/s/%s/api/alerting/rule/%s", c.HostURL, c.Space, ruleID)
	log.Printf("Getting rule using %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	rule := Rule{}
	err = json.Unmarshal(body, &rule)
	if err != nil {
		return nil, err
	}

	return &rule, nil
}

// CreateRule - Create Kibana rules.
// https://www.elastic.co/guide/en/kibana/7.13/create-rule-api.html
func (c *Client) CreateRule(rule CreateRule) (*Rule, error) {
	url := fmt.Sprintf("%s/s/%s/api/alerting/rule", c.HostURL, c.Space)
	log.Printf("Creating rule using %s", url)

	rb, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}
	log.Printf("Creating rule %s", string(rb))

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

	newRule := Rule{}
	err = json.Unmarshal(body, &newRule)
	if err != nil {
		return nil, err
	}

	return &newRule, nil
}
