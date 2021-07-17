package kibana

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
