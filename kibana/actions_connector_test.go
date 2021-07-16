package kibana

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConnector(t *testing.T) {
	userName := "testUser"
	password := "testPassword"
	space := "testSpace"
	connectorID := "testId"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"id": "c55b6eb0-6bad-11eb-9f3b-611eebc6c3ad",
			"connector_type_id": ".index",
			"name": "my-connector",
			"config": {
			  "index": "test-index",
			  "refresh": false,
			  "executionTimeField": null
			},
			"is_preconfigured": false
		  }`)
	}))
	defer ts.Close()

	c, err := NewClient(&ts.URL, &userName, &password, &space)
	if err != nil {
		log.Fatal(err)
	}

	connector, err := c.GetConnector(connectorID)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, connector.Name, "my-connector")

}
