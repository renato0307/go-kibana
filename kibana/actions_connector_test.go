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

func TestCreateConnector(t *testing.T) {
	userName := "testUser"
	password := "testPassword"
	space := "testSpace"

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

	connector := CreateConnector{
		Name:            "my-connector",
		ConnectorTypeId: ".index",
		Config: ConnectorConfig{
			Index:   "test-index",
			Refresh: true,
		},
	}
	newConnector, err := c.CreateConnector(connector)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connector created is: %s", newConnector.ID)

	assert.Equal(t, newConnector.ID, "c55b6eb0-6bad-11eb-9f3b-611eebc6c3ad")
}

func TestUpdateConnector(t *testing.T) {
	userName := "testUser"
	password := "testPassword"
	space := "testSpace"

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

	connector := UpdateConnector{
		Name: "my-connector",
		Config: ConnectorConfig{
			Index: "test-index",
		},
	}
	updatedConnector, err := c.UpdateConnector("c55b6eb0-6bad-11eb-9f3b-611eebc6c3ad", connector)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connector name update is : %s", updatedConnector.Name)

	assert.Equal(t, updatedConnector.Name, "my-connector")
}
