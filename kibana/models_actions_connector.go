package kibana

// From https://www.elastic.co/guide/en/kibana/7.13/get-connector-api.html

type ConnectorConfig struct {
	Index              string `json:"index"`
	Refresh            bool   `json:"refresh"`
	ExecutionTimeField string `json:"executionTimeField"`
}

type Connector struct {
	ID               string          `json:"id"`
	Name             string          `json:"name"`
	ConnectorTypeId  string          `json:"connector_type_id"`
	IsPreconfigured  bool            `json:"is_preconfigured"`
	IsMissingSecrets bool            `json:"is_missing_secrets"`
	Config           ConnectorConfig `json:"config"`
}

type CreateConnector struct {
	Name            string          `json:"name"`
	ConnectorTypeId string          `json:"connector_type_id"`
	Config          ConnectorConfig `json:"config"`
}
