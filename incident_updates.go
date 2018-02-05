package cachet

import (
	"fmt"
)

// @todo: add constants to component status

const (
	// Docs: https://docs.cachethq.io/docs/incident-statuses

	// IncidentStatusScheduled means "This status is used for a scheduled status."
	IncidentStatusScheduled = 0
	// IncidentStatusInvestigating means "You have reports of a problem and you're currently looking into them."
	IncidentStatusInvestigating = 1
	// IncidentStatusIdentified means "You've found the issue and you're working on a fix."
	IncidentStatusIdentified = 2
	// IncidentStatusWatching means "You've since deployed a fix and you're currently watching the situation."
	IncidentStatusWatching = 3
	// IncidentStatusFixed means "The fix has worked, you're happy to close the incident."
	IncidentStatusFixed = 4

	// ComponentStatusUnknown means "The component's status is not known."
	ComponentStatusUnknown = 0
	// ComponentStatusOperational means "The component is working."
	ComponentStatusOperational = 1
	// ComponentStatusPerformanceIssues means "The component is experiencing some slowness."
	ComponentStatusPerformanceIssues = 2
	// ComponentStatusPartialOutage means "The component may not be working for everybody."
	// This could be a geographical issue for example.
	ComponentStatusPartialOutage = 3
	// ComponentStatusMajorOutage means "The component is not working for anybody."
	ComponentStatusMajorOutage = 4
)

// IncidentUpdatesService contains REST endpoints that belongs to cachet incidents.
type IncidentUpdatesService struct {
	client *Client
}

// IncidentUpdate entity reflects one single incident update
type IncidentUpdate struct {
	ID              int    `json:"id,omitempty"`
	IncidentID      int    `json:"incident_id,omitempty"`
	ComponentID     int    `json:"component_id,omitempty"`
	ComponentStatus int    `json:"component_id,omitempty"`
	Status          int    `json:"status,omitempty"`
	Message         string `json:"message,omitempty"`
	UserID          int    `json:"user_id,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
	HumanStatus     string `json:"human_status,omitempty"`
	Permalink       string `json:"permalink,omitempty"`
}

// IncidentUpdateResponse reflects the response of /incident updates call
type IncidentUpdateResponse struct {
	Meta            Meta             `json:"meta,omitempty"`
	IncidentUpdates []IncidentUpdate `json:"data,omitempty"`
}

// incidentsAPIResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the incident in the "data" key.
type incidentUpdatesAPIResponse struct {
	Data *IncidentUpdate `json:"data"`
}

// GetAll return all updates by incident.
//
// Docs: https://docs.cachethq.io/reference#incidentsidupdates
func (s *IncidentUpdatesService) GetAll(incidentId int) (*IncidentUpdateResponse, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates", incidentId)
	v := new(IncidentUpdateResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns a single incident update.
//
// Docs: https://docs.cachethq.io/reference#incidentsidupdatesid
func (s *IncidentUpdatesService) Get(incidentId int, updateId int) (*IncidentUpdate, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates/%d", incidentId, updateId)
	v := new(incidentUpdatesAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// Create creates a new incident update.
//
// Docs: https://docs.cachethq.io/reference#incidentsincidentupdates
func (s *IncidentUpdatesService) Create(incidentId int, i *IncidentUpdate) (*IncidentUpdate, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates", incidentId)
	v := new(incidentUpdatesAPIResponse)

	resp, err := s.client.Call("POST", u, i, v)
	return v.Data, resp, err
}

// Update updates an incident update.
//
// Docs: https://docs.cachethq.io/reference#incidentsincidentupdatesupdate-1
func (s *IncidentUpdatesService) Update(incidentId int, updateId int, i *IncidentUpdate) (*IncidentUpdate, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates/%d", incidentId, updateId)
	v := new(incidentUpdatesAPIResponse)

	resp, err := s.client.Call("PUT", u, i, v)
	return v.Data, resp, err
}

// Delete deletes an incident update.
//
// Docs: https://docs.cachethq.io/reference#incidentsincidentupdatesupdate
func (s *IncidentUpdatesService) Delete(incidentId int, updateId int) (*Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates/%d", incidentId, updateId)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}
