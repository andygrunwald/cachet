package cachet

import (
	"fmt"
)

const (
	// Docs: https://docs.cachethq.io/docs/component-statuses

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

// ComponentsService contains REST endpoints that belongs to cachet components.
type ComponentsService struct {
	client *Client
}

// Tag ...
type Tag struct {
	Tag string `json:"tag,omitempty"`
}

// Component entity reflects one single component
type Component struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	Status      int    `json:"status,omitempty"`
	Order       int    `json:"order,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	GroupID     int    `json:"group_id,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	DeletedAt   string `json:"deleted_at,omitempty"`
	StatusName  string `json:"status_name,omitempty"`
	//Tags        Tag    `json:"tags,omitempty"`
}

// ComponentResponse reflects the response of /components call
type ComponentResponse struct {
	Meta       Meta        `json:"meta,omitempty"`
	Components []Component `json:"data,omitempty"`
}

// ComponentsQueryParams contains fields to filter returned results
type ComponentsQueryParams struct {
	ID      int    `url:"id,omitempty"`
	Name    string `url:"name,omitempty"`
	Status  int    `url:"status,omitempty"`
	Order   int    `url:"order,omitempty"`
	Enabled bool   `url:"enabled,omitempty"`
	GroupID int    `url:"group_id,omitempty"`
	QueryOptions
}

// componentApiResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the component in the "data" key.
type componentAPIResponse struct {
	Data *Component `json:"data"`
}

// GetAll return all components that have been created.
//
// Docs: https://docs.cachethq.io/reference#get-components
func (s *ComponentsService) GetAll(filter *ComponentsQueryParams) (*ComponentResponse, *Response, error) {
	u := "api/v1/components"
	v := new(ComponentResponse)

	u, err := addOptions(u, filter)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get return a single component.
//
// Docs: https://docs.cachethq.io/reference#get-a-component
func (s *ComponentsService) Get(id int) (*Component, *Response, error) {
	u := fmt.Sprintf("api/v1/components/%d", id)
	v := new(componentAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// Create a new component.
//
// Docs: https://docs.cachethq.io/reference#components
func (s *ComponentsService) Create(c *Component) (*Component, *Response, error) {
	u := "api/v1/components"
	v := new(componentAPIResponse)

	resp, err := s.client.Call("POST", u, c, v)
	return v.Data, resp, err
}

// Update updates a component.
//
// Docs: https://docs.cachethq.io/docs/update-a-component
func (s *ComponentsService) Update(id int, c *Component) (*Component, *Response, error) {
	u := fmt.Sprintf("api/v1/components/%d", id)
	v := new(componentAPIResponse)

	resp, err := s.client.Call("PUT", u, c, v)
	return v.Data, resp, err
}

// Delete deletes a component.
//
// Docs: https://docs.cachethq.io/docs/delete-a-component
func (s *ComponentsService) Delete(id int) (*Response, error) {
	u := fmt.Sprintf("api/v1/components/%d", id)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}
