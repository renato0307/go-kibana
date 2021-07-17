package kibana

type RuleParams struct {
	AggField            string   `json:"aggField,omitempty"`
	AggType             string   `json:"aggType,omitempty"`
	ESQuery             string   `json:"esQuery"`
	GroupBy             string   `json:"groupBy,omitempty"`
	Index               []string `json:"index"`
	Size                int      `json:"size"`
	TermField           string   `json:"termField,omitempty"`
	TermSize            int      `json:"termSize,omitempty"`
	Threshold           []int    `json:"threshold"`
	ThresholdComparator string   `json:"thresholdComparator"`
	TimeField           string   `json:"timeField"`
	TimeWindowSize      int      `json:"timeWindowSize"`
	TimeWindowUnit      string   `json:"timeWindowUnit"`
}

type RuleSchedule struct {
	Interval string `json:"interval"`
}

type RuleAction struct {
	ID     string                 `json:"id"`
	Group  string                 `json:"group"`
	Params map[string]interface{} `json:"params"`
}

type RuleExecutionStatus struct {
	LastExecutionDate string `json:"last_execution_date"`
	Status            string `json:"status"`
}

type CreateRule struct {
	Actions    []RuleAction `json:"actions"`
	Consumer   string       `json:"consumer"`
	Name       string       `json:"name"`
	NotifyWhen string       `json:"notify_when"`
	Params     RuleParams   `json:"params"`
	RuleTypeID string       `json:"rule_type_id"`
	Schedule   RuleSchedule `json:"schedule"`
	Tags       []string     `json:"tags"`
}

type Rule struct {
	Actions         []RuleAction        `json:"actions"`
	ApiKeyOwner     string              `json:"api_key_owner"`
	Consumer        string              `json:"consumer"`
	CreatedAt       string              `json:"created_at"`
	CreatedBy       string              `json:"created_by"`
	Enabled         bool                `json:"enabled"`
	ExecutionStatus RuleExecutionStatus `json:"execution_status"`
	ID              string              `json:"id"`
	MuteAll         bool                `json:"mute_all"`
	MutedAlertIDs   []string            `json:"muted_alert_ids"`
	Name            string              `json:"name"`
	NotifyWhen      string              `json:"notify_when"`
	Params          RuleParams          `json:"params"`
	RuleTypeID      string              `json:"rule_type_id"`
	Schedule        RuleSchedule        `json:"schedule"`
	ScheduledTaskId string              `json:"scheduled_task_id"`
	Tags            []string            `json:"tags"`
	UpdatedAt       string              `json:"updated_at"`
	UpdatedBy       string              `json:"updated_by"`
}
