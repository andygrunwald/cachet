package cachet

import (
	"fmt"
)

const (
	// Upcoming
	ScheduleUpcoming = 0
	// In Progress
	ScheduleInProgress = 1
	// Complete
	ScheduleComplete = 2
)

// SchedulesService contains REST endpoints that belongs to cachet schedules.
type SchedulesService struct {
	client *Client
}

// Schedule entity reflects one single schedule
type Schedule struct {
	ID          int          `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Message     string       `json:"message,omitempty"`
	Status      int          `json:"status,omitempty"`
	ScheduledAt string       `json:"scheduled_at,omitempty"`
	CompletedAt string       `json:"completed_at,omitempty"`
	Components  []*Component `json:"components,omitempty"`
	HumanStatus string       `json:"human_status,omitempty"`
}

// ScheduleResponse reflects the response of schedules call
type ScheduleResponse struct {
	Meta      Meta       `json:"meta,omitempty"`
	Schedules []Schedule `json:"data,omitempty"`
}

// schedulesAPIResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the incident in the "data" key.
type incidentUpdatesAPIResponse struct {
	Data *Schedule `json:"data"`
}

// GetAll return all scheduled events.
//
// Docs: https://docs.cachethq.io/reference#incidentsidupdates
func (s *SchedulesService) GetAll() (*ScheduleResponse, *Response, error) {
	u := "api/v1/schedules"
	v := new(ScheduleResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns a single scheduled event.
//
// Docs: https://docs.cachethq.io/reference#incidentsidupdatesid
func (s *SchedulesService) Get(id int) (*Schedule, *Response, error) {
	u := fmt.Sprintf("api/v1/schedules/%d", id)
	v := new(incidentUpdatesAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// Create creates a new scheduled event.
//
// Docs: https://docs.cachethq.io/reference#incidentsincidentupdates
func (s *SchedulesService) Create(i *Schedule) (*Schedule, *Response, error) {
	u := "api/v1/schedules"
	v := new(incidentUpdatesAPIResponse)

	resp, err := s.client.Call("POST", u, i, v)
	return v.Data, resp, err
}

// Update updates a scheduled event.
//
// Docs: https://docs.cachethq.io/reference#incidentsincidentupdatesupdate-1
func (s *SchedulesService) Update(id int, i *Schedule) (*Schedule, *Response, error) {
	u := fmt.Sprintf("api/v1/schedules/%d", id)
	v := new(incidentUpdatesAPIResponse)

	resp, err := s.client.Call("PUT", u, i, v)
	return v.Data, resp, err
}

// Delete deletes a scheduled event.
//
// Docs: https://docs.cachethq.io/reference#incidentsincidentupdatesupdate
func (s *SchedulesService) Delete(id int) (*Response, error) {
	u := fmt.Sprintf("api/v1/schedules/%d", id)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}
