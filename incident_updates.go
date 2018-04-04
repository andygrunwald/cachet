package cachet

import (
	"fmt"
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
	ComponentStatus int    `json:"component_status,omitempty"`
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
func (s *IncidentUpdatesService) GetAll(incidentID int) (*IncidentUpdateResponse, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates", incidentID)
	v := new(IncidentUpdateResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns a single incident update.
//
// Docs: https://docs.cachethq.io/reference#incidentsidupdatesid
func (s *IncidentUpdatesService) Get(incidentID int, updateID int) (*IncidentUpdate, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates/%d", incidentID, updateID)
	v := new(incidentUpdatesAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// Create creates a new incident update.
//
// Docs: https://docs.cachethq.io/reference#incidentsincidentupdates
func (s *IncidentUpdatesService) Create(incidentID int, i *IncidentUpdate) (*IncidentUpdate, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates", incidentID)
	v := new(incidentUpdatesAPIResponse)

	resp, err := s.client.Call("POST", u, i, v)
	return v.Data, resp, err
}

// Update updates an incident update.
//
// Docs: https://docs.cachethq.io/reference#incidentsincidentupdatesupdate-1
func (s *IncidentUpdatesService) Update(incidentID int, updateID int, i *IncidentUpdate) (*IncidentUpdate, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates/%d", incidentID, updateID)
	v := new(incidentUpdatesAPIResponse)

	resp, err := s.client.Call("PUT", u, i, v)
	return v.Data, resp, err
}

// Delete deletes an incident update.
//
// Docs: https://docs.cachethq.io/reference#incidentsincidentupdatesupdate
func (s *IncidentUpdatesService) Delete(incidentID int, updateID int) (*Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates/%d", incidentID, updateID)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}
