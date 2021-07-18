# go-kibana

SDK for Kibana API

## Versions supported

Tests done with Kibana 7.13 from Elastic Cloud.

## APIs supported

### Alerting APIs

- [x] Create rule API to create a rule
- [x] Update rule API to update the attributes for existing rules
- [x] Get rule API to retrieve a single rule by ID
- [x] Delete rule API to permanently remove a rule
- [ ] Find rules API to retrieve a paginated set of rules by condition
- [ ] List rule types API to retrieve a list of rule types
- [ ] Enable rule API to enable a single rule by ID
- [ ] Disable rule API to disable a single rule by ID
- [ ] Mute alert API to mute alert for a single rule by ID
- [ ] Unmute alert API to unmute alert for a single rule by ID
- [ ] Mute all alerts API to mute all alerts for a single rule by ID
- [ ] Unmute all alerts API to unmute all alerts for a single rule by ID
- [ ] Get Alerting framework health API to retrieve the health of the Alerting framework

### Action and Connector APIs

- [x] Get connector API to retrieve a single connector by ID
- [ ] Get all connectors API to retrieve all connectors
- [ ] List all connector types API to retrieve a list of all connector types
- [x] Create connector API to create connectors
- [x] Update connector API to update the attributes for an existing connector
- [ ] Execute connector API to execute a connector by ID
- [x] Delete connector API to delete a connector by ID
