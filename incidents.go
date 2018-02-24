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
	ID                int              `json:"id,omitempty"`
	Name              string           `json:"name,omitempty"`
	Status            int              `json:"status,omitempty"`
	Message           string           `json:"message,omitempty"`
	Visible           int              `json:"visible,omitempty"`
	ComponentID       int              `json:"component_id,omitempty"`
	ComponentStatus   int              `json:"component_status,omitempty"`
	Notify            bool             `json:"notify,omitempty"`
	Stickied          bool             `json:"stickied,omitempty"`
	OccurredAt        string           `json:"occurred_at,omitempty"`
	Template          string           `json:"template,omitempty"`
	Vars              []string         `json:"vars,omitempty"`
	CreatedAt         string           `json:"created_at,omitempty"`
	UpdatedAt         string           `json:"updated_at,omitempty"`
	DeletedAt         string           `json:"deleted_at,omitempty"`
	IsResolved        bool             `json:"is_resolved,omitempty"`
	Updates           []IncidentUpdate `json:"updates,omitempty"`
	HumanStatus       string           `json:"human_status,omitempty"`
	LatestUpdateID    int              `json:"latest_update_id,omitempty"`
	LatestStatus      int              `json:"latest_status,omitempty"`
	LatestHumanStatus string           `json:"latest_human_status,omitempty"`
	LatestIcon        string           `json:"latest_icon,omitempty"`
	Permalink         string           `json:"permalink,omitempty"`
	Duration          int              `json:"duration,omitempty"`
}

// IncidentResponse reflects the response of /incidents call
type IncidentResponse struct {
	Meta      Meta       `json:"meta,omitempty"`
	Incidents []Incident `json:"data,omitempty"`
}

// IncidentsQueryParams contains fields to filter returned results
type IncidentsQueryParams struct {
	ID          int          `url:"id,omitempty"`
	Name        string       `url:"name,omitempty"`
	Status      int          `url:"status,omitempty"`
	Visible     int          `url:"visible,omitempty"`
	ComponentID int          `url:"component_id,omitempty"`
	Stickied    bool         `url:"stickied,omitempty"`
	ExtraFields QueryOptions `url:",omitempty"`
}

// incidentsAPIResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the incident in the "data" key.
type incidentsAPIResponse struct {
	Data *Incident `json:"data"`
}

// GetAll return all incidents.
//
// Docs: https://docs.cachethq.io/reference#get-incidents
func (s *IncidentsService) GetAll(filter *IncidentsQueryParams) (*IncidentResponse, *Response, error) {
	u := "api/v1/incidents"
	v := new(IncidentResponse)

	u, err := addOptions(u, filter)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns a single incident.
//
// Docs: https://docs.cachethq.io/reference#get-an-incident
func (s *IncidentsService) Get(id int) (*Incident, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d", id)
	v := new(incidentsAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// Create a new incident.
//
// Docs: https://docs.cachethq.io/reference#incidents
func (s *IncidentsService) Create(i *Incident) (*Incident, *Response, error) {
	u := "api/v1/incidents"
	v := new(incidentsAPIResponse)

	resp, err := s.client.Call("POST", u, i, v)
	return v.Data, resp, err
}

// Update updates an incident.
//
// Docs: https://docs.cachethq.io/reference#update-an-incident
func (s *IncidentsService) Update(id int, i *Incident) (*Incident, *Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d", id)
	v := new(incidentsAPIResponse)

	resp, err := s.client.Call("PUT", u, i, v)
	return v.Data, resp, err
}

// Delete delete an incident.
//
// Docs: https://docs.cachethq.io/reference#delete-an-incident
func (s *IncidentsService) Delete(id int) (*Response, error) {
	u := fmt.Sprintf("api/v1/incidents/%d", id)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}
