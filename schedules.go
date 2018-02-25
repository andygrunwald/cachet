package cachet

import (
	"fmt"
)

const (
	// ScheduleUpcoming means "This scheduled event is going to happen somewhere in the future."
	ScheduleUpcoming = 0
	// ScheduleInProgress means "This scheduled event is happening at the moment."
	ScheduleInProgress = 1
	// ScheduleComplete means "This scheduled event has already finished."
	ScheduleComplete = 2
)

// SchedulesService contains REST endpoints that belongs to cachet schedules.
type SchedulesService struct {
	client *Client
}

// Schedule entity reflects one single schedule
type Schedule struct {
	ID          int         `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Message     string      `json:"message,omitempty"`
	Status      int         `json:"status,omitempty"`
	ScheduledAt string      `json:"scheduled_at,omitempty"`
	CompletedAt string      `json:"completed_at,omitempty"`
	CreatedAt   string      `json:"created_at,omitempty"`
	UpdatedAt   string      `json:"updated_at,omitempty"`
	Components  []Component `json:"components,omitempty"`
	HumanStatus string      `json:"human_status,omitempty"`
}

// ScheduleResponse reflects the response of schedules call
type ScheduleResponse struct {
	Meta      Meta       `json:"meta,omitempty"`
	Schedules []Schedule `json:"data,omitempty"`
}

// SchedulesQueryParams contains fields to filter returned results
type SchedulesQueryParams struct {
	ID          int          `url:"id,omitempty"`
	Name        string       `url:"name,omitempty"`
	Status      int          `url:"order,omitempty"`
	ExtraFields QueryOptions `url:" ,omitempty"`
}

// schedulesAPIResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the incident in the "data" key.
type schedulesAPIResponse struct {
	Data *Schedule `json:"data"`
}

// GetAll return all scheduled events.
//
// Docs: https://docs.cachethq.io/reference#incidentsidupdates
func (s *SchedulesService) GetAll(filter *SchedulesQueryParams) (*ScheduleResponse, *Response, error) {
	u := "api/v1/schedules"
	v := new(ScheduleResponse)

	u, err := addOptions(u, filter)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns a single scheduled event.
//
// Docs: https://docs.cachethq.io/reference#incidentsidupdatesid
func (s *SchedulesService) Get(id int) (*Schedule, *Response, error) {
	u := fmt.Sprintf("api/v1/schedules/%d", id)
	v := new(schedulesAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// Create creates a new scheduled event.
//
// Docs: https://docs.cachethq.io/reference#incidentsincidentupdates
func (s *SchedulesService) Create(i *Schedule) (*Schedule, *Response, error) {
	u := "api/v1/schedules"
	v := new(schedulesAPIResponse)

	resp, err := s.client.Call("POST", u, i, v)
	return v.Data, resp, err
}

// Update updates a scheduled event.
//
// Docs: https://docs.cachethq.io/reference#incidentsincidentupdatesupdate-1
func (s *SchedulesService) Update(id int, i *Schedule) (*Schedule, *Response, error) {
	u := fmt.Sprintf("api/v1/schedules/%d", id)
	v := new(schedulesAPIResponse)

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
