package cachet

import (
	"fmt"
)

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

	// IncidentVisibilityPublic means "Viewable by public"
	IncidentVisibilityPublic = 1
	// IncidentVisibilityLoggedIn means "Only visible to logged in users"
	IncidentVisibilityLoggedIn = 0
)

// IncidentsService contains REST endpoints that belongs to cachet incidents.
type IncidentsService struct {
	client *Client
}

// Incident entity reflects one single incident
type Incident struct {
	ID              int    `json:"id,omitempty"`
	ComponentID     int    `json:"component_id,omitempty"`
	ComponentStatus int    `json:"component_status,omitempty"`
	Name            string `json:"name,omitempty"`
	Status          int    `json:"status,omitempty"`
	Visible         int    `json:"visible,omitempty"`
	Message         string `json:"message,omitempty"`
	ScheduledAt     string `json:"scheduled_at,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
	DeletedAt       string `json:"deleted_at,omitempty"`
	HumanStatus     string `json:"human_status,omitempty"`
	Notify          bool   `json:"notify,omitempty"`
}

// IncidentUpdate entity reflects one single incident update
type IncidentUpdate struct {
	ID          int    `json:"id,omitempty"`
	IncidentID  int    `json:"incident_id,omitempty"`
	Status      int    `json:"status,omitempty"`
	Message     string `json:"message,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	HumanStatus string `json:"human_status,omitempty"`
	Permalink   string `json:"permalink,omitempty"`
}

// IncidentResponse reflects the response of /incidents call
type IncidentResponse struct {
	Meta      Meta       `json:"meta,omitempty"`
	Incidents []Incident `json:"data,omitempty"`
}

// IncidentUpdate reflects the response of /incidents/:incident/updates call
type IncidentUpdateResponse struct {
	Meta            Meta             `json:"meta,omitempty"`
	IncidentUpdates []IncidentUpdate `json:"data,omitempty"`
}

// incidentsAPIResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the incident in the "data" key.
type incidentsAPIResponse struct {
	Data *Incident `json:"data"`
}

// incidentUpdatesAPIResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the incident update in the "data" key.
type incidentUpdatesAPIResponse struct {
	Data *IncidentUpdate `json:"data"`
}

// GetAll return all incidents.
//
// Docs: https://docs.cachethq.io/docs/get-incidents
func (s *IncidentsService) GetAll(opt *ListOptions) (*IncidentResponse, *Response, error) {
	u := "api/v1/incidents"
	v := new(IncidentResponse)

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns a single incident.
//
// Docs: https://docs.cachethq.io/docs/get-an-incident
func (s *IncidentsService) Get(id int) (*Incident, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d", id)
	v := new(incidentsAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// Create a new incident.
//
// Docs: https://docs.cachethq.io/docs/incidents
func (s *IncidentsService) Create(i *Incident) (*Incident, *Response, error) {
	u := "api/v1/incidents"
	v := new(incidentsAPIResponse)

	resp, err := s.client.Call("POST", u, i, v)
	return v.Data, resp, err
}

// Update updates an incident.
//
// Docs: https://docs.cachethq.io/docs/update-an-incident
func (s *IncidentsService) Update(id int, i *Incident) (*Incident, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d", id)
	v := new(incidentsAPIResponse)

	resp, err := s.client.Call("PUT", u, i, v)
	return v.Data, resp, err
}

// Delete delete an incident.
//
// Docs: https://docs.cachethq.io/docs/delete-an-incident
func (s *IncidentsService) Delete(id int) (*Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d", id)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}

// GetAllUpdates return all incident updates that have been created for an incident.
//
// Docs: https://docs.cachethq.io/docs/incidentsidupdates
func (s *IncidentsService) GetAllUpdates(opt *ListOptions, id int) (*IncidentUpdateResponse, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates", id)
	v := new(IncidentUpdateResponse)

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetUpdate returns a single incident update.
//
// Docs: https://docs.cachethq.io/docs/incidentsidupdatesid
func (s *IncidentsService) GetUpdate(id, uid int) (*IncidentUpdate, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates/%d", id, uid)
	v := new(incidentUpdatesAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// CreateUpdate creates a new incident update.
//
// Docs: https://docs.cachethq.io/docs/incidentsincidentupdates
func (s *IncidentsService) CreateUpdate(id int, i *IncidentUpdate) (*IncidentUpdate, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates", id)
	v := new(incidentUpdatesAPIResponse)

	resp, err := s.client.Call("POST", u, i, v)
	return v.Data, resp, err
}

// UpdateUpdate updates an incident update.
//
// Docs: https://docs.cachethq.io/docs/incidentsincidentupdatesupdate-1
func (s *IncidentsService) UpdateUpdate(id, uid int, i *IncidentUpdate) (*IncidentUpdate, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates/%d", id, uid)
	v := new(incidentUpdatesAPIResponse)

	resp, err := s.client.Call("PUT", u, i, v)
	return v.Data, resp, err
}

// DeleteUpdate delete an incident update.
//
// Docs: https://docs.cachethq.io/docs/incidentsincidentupdatesupdate
func (s *IncidentsService) DeleteUpdate(id, uid int) (*Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d/updates/%d", id, uid)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}
