package kibana

type Connector struct {
	ID               string                 `json:"id"`
	Name             string                 `json:"name"`
	ConnectorTypeId  string                 `json:"connector_type_id"`
	IsPreconfigured  bool                   `json:"is_preconfigured"`
	IsMissingSecrets bool                   `json:"is_missing_secrets"`
	Config           map[string]interface{} `json:"config,omitempty"`
	Secrets          map[string]interface{} `json:"secrets,omitempty"`
}

type CreateConnector struct {
	Name            string                 `json:"name"`
	ConnectorTypeId string                 `json:"connector_type_id"`
	Config          map[string]interface{} `json:"config,omitempty"`
	Secrets         map[string]interface{} `json:"secrets,omitempty"`
}

type UpdateConnector struct {
	Name    string                 `json:"name"`
	Config  map[string]interface{} `json:"config,omitempty"`
	Secrets map[string]interface{} `json:"secrets,omitempty"`
}
