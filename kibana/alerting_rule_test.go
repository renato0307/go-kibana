package kibana

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRule(t *testing.T) {
	userName := "testUser"
	password := "testPassword"
	space := "testSpace"
	ruleID := "testId"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"id": "0a037d60-6b62-11eb-9e0d-85d233e3ee35",
			"notify_when": "onActionGroupChange",
			"params": {
				"aggType": "avg",
				"timeWindowSize": 5,
				"timeWindowUnit": "m"
			},
			"consumer": "alerts",
			"rule_type_id": "test.rule.type",
			"schedule": {
			  "interval": "1m"
			},
			"actions": [
				{
					"id":"dceeb5d0-6b41-11eb-802b-85b0c1bc8ba2",
					"group":"threshold met",
					"params":{
					   "level":"info"
					}
				 }				
			],
			"tags": [ "tag1", "tag2"],
			"name": "test rule",
			"enabled": true,
			"throttle": null,
			"api_key_owner": "elastic",
			"created_by": "elastic",
			"updated_by": "elastic",
			"mute_all": false,
			"muted_alert_ids": [ "asdfgh", "qwerty" ],
			"updated_at": "2021-02-10T05:37:19.086Z",
			"created_at": "2021-02-10T05:37:19.086Z",
			"scheduled_task_id": "0b092d90-6b62-11eb-9e0d-85d233e3ee35",
			"execution_status": {
			  "last_execution_date": "2021-02-10T17:55:14.262Z",
			  "status": "ok"
			}
		  }`)
	}))
	defer ts.Close()

	c, err := NewClient(&ts.URL, &userName, &password, &space)
	if err != nil {
		log.Fatal(err)
	}

	rule, err := c.GetRule(ruleID)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "alerts", rule.Consumer)
	assert.Equal(t, "2021-02-10T17:55:14.262Z", rule.ExecutionStatus.LastExecutionDate)
	assert.Equal(t, "ok", rule.ExecutionStatus.Status)
	assert.Equal(t, "0a037d60-6b62-11eb-9e0d-85d233e3ee35", rule.ID)
	assert.Equal(t, "test rule", rule.Name)
	assert.Equal(t, "onActionGroupChange", rule.NotifyWhen)
	assert.Equal(t, 5, rule.Params.TimeWindowSize)
	assert.Equal(t, "avg", rule.Params.AggType)
}

func TestCreateRule(t *testing.T) {
	userName := "testUser"
	password := "testPassword"
	space := "testSpace"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"id": "0a037d60-6b62-11eb-9e0d-85d233e3ee35",
			"notify_when": "onActionGroupChange",
			"params": {
			  "aggType": "avg",
			  "timeWindowSize": 5,
			  "timeWindowUnit": "m"
			},
			"consumer": "alerts",
			"rule_type_id": "test.rule.type",
			"schedule": {
			  "interval": "1m"
			},
			"actions": [
				{
					"id":"dceeb5d0-6b41-11eb-802b-85b0c1bc8ba2",
					"group":"query matched",
					"params":{
						"documents": [
							{
								"level":"info",
								"message":"alert message"
							}
						]
					}
				 }				
			],
			"tags": [ "tag1", "tag2"],
			"name": "test rule",
			"enabled": true,
			"throttle": null,
			"api_key_owner": "elastic",
			"created_by": "elastic",
			"updated_by": "elastic",
			"mute_all": false,
			"muted_alert_ids": [ "asdfgh", "qwerty" ],
			"updated_at": "2021-02-10T05:37:19.086Z",
			"created_at": "2021-02-10T05:37:19.086Z",
			"scheduled_task_id": "0b092d90-6b62-11eb-9e0d-85d233e3ee35",
			"execution_status": {
			  "last_execution_date": "2021-02-10T17:55:14.262Z",
			  "status": "ok"
			}
		  }`)
	}))
	defer ts.Close()

	c, err := NewClient(&ts.URL, &userName, &password, &space)
	if err != nil {
		log.Fatal(err)
	}

	rule := CreateRule{
		Actions: []RuleAction{
			{
				ID:    "dceeb5d0-6b41-11eb-802b-85b0c1bc8ba2",
				Group: "query matched",
				Params: map[string]interface{}{
					"documents": []interface{}{
						map[string]interface{}{
							"level":   "info",
							"message": "alert message",
						},
					},
				},
			},
		},
		Consumer:   "alerts",
		Name:       "test rule",
		NotifyWhen: "onActionGroupChange",
		RuleTypeID: "test.rule.type",
		Params: RuleParams{
			AggType:        "avg",
			TimeWindowSize: 5,
			TimeWindowUnit: "m",
		},
		Schedule: RuleSchedule{
			Interval: "1m",
		},
		Tags: []string{
			"tag1", "tag2",
		},
	}

	newRule, err := c.CreateRule(rule)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, rule.Actions, newRule.Actions)
	assert.Equal(t, rule.Name, newRule.Name)
	assert.Equal(t, rule.Consumer, newRule.Consumer)
	assert.Equal(t, rule.Params, newRule.Params)
	assert.Equal(t, rule.NotifyWhen, newRule.NotifyWhen)
	assert.Equal(t, rule.RuleTypeID, newRule.RuleTypeID)
	assert.Equal(t, rule.Tags, newRule.Tags)
}
