package cachet

import (
	"fmt"
)

const (
	// Docs: https://docs.cachethq.io/docs/component-statuses

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
}

// ComponentGroup entity reflects one single component group
type ComponentGroup struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Order     int    `json:"order,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// ComponentResponse reflects the response of /components call
type ComponentResponse struct {
	Meta       Meta        `json:"meta,omitempty"`
	Components []Component `json:"data,omitempty"`
}

// ComponentGroupResponse reflects the response of /components/groups call
type ComponentGroupResponse struct {
	Meta            Meta             `json:"meta,omitempty"`
	ComponentGroups []ComponentGroup `json:"data,omitempty"`
}

// componentApiResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the component in the "data" key.
type componentAPIResponse struct {
	Data *Component `json:"data"`
}

// componentGroupAPIResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the component group in the "data" key.
type componentGroupAPIResponse struct {
	Data *ComponentGroup `json:"data"`
}

// GetAll return all components that have been created.
//
// Docs: https://docs.cachethq.io/docs/get-components
func (s *ComponentsService) GetAll() (*ComponentResponse, *Response, error) {
	u := "api/v1/components"
	v := new(ComponentResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get return a single component.
//
// Docs: https://docs.cachethq.io/docs/get-a-component
func (s *ComponentsService) Get(id int) (*Component, *Response, error) {
	u := fmt.Sprintf("api/v1/components/%d", id)
	v := new(componentAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// Create a new component.
//
// Docs: https://docs.cachethq.io/docs/components
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

// GetAllGroups return all component groups that have been created.
//
// Docs: https://docs.cachethq.io/docs/get-componentgroups
func (s *ComponentsService) GetAllGroups() (*ComponentGroupResponse, *Response, error) {
	u := "api/v1/components/groups"
	v := new(ComponentGroupResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetGroup return a single component group.
//
// Docs: https://docs.cachethq.io/docs/get-a-component-group
func (s *ComponentsService) GetGroup(id int) (*ComponentGroup, *Response, error) {
	u := fmt.Sprintf("api/v1/components/groups/%d", id)
	v := new(componentGroupAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// CreateGroup creates a new component group.
//
// Docs: https://docs.cachethq.io/docs/post-componentgroups
func (s *ComponentsService) CreateGroup(c *ComponentGroup) (*ComponentGroup, *Response, error) {
	u := "api/v1/components/groups"
	v := new(componentGroupAPIResponse)

	resp, err := s.client.Call("POST", u, c, v)
	return v.Data, resp, err
}

// UpdateGroup updates a component group.
//
// Docs: https://docs.cachethq.io/docs/put-component-group
func (s *ComponentsService) UpdateGroup(id int, c *ComponentGroup) (*ComponentGroup, *Response, error) {
	u := fmt.Sprintf("api/v1/components/groups/%d", id)
	v := new(componentGroupAPIResponse)

	resp, err := s.client.Call("PUT", u, c, v)
	return v.Data, resp, err
}

// DeleteGroup deletes a component group.
//
// Docs: https://docs.cachethq.io/docs/delete-component-group
func (s *ComponentsService) DeleteGroup(id int) (*Response, error) {
	u := fmt.Sprintf("api/v1/components/groups/%d", id)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}
