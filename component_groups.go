package cachet

import (
	"fmt"
)

const (
	// ComponentGroupVisibilityPublic means "Viewable by public"
	ComponentGroupVisibilityPublic = 1
	// ComponentGroupVisibilityLoggedIn means "Only visible to logged in users"
	ComponentGroupVisibilityLoggedIn = 0
)

// ComponentGroupsService contains REST endpoints that belongs to cachet components.
type ComponentGroupsService struct {
	client *Client
}

// Tag ...
type Tag struct {
	Tag string `json:",omitempty"`
}

// ComponentGroup entity reflects one single component group
type ComponentGroup struct {
	ID                      int          `json:"id,omitempty"`
	Name                    string       `json:"name,omitempty"`
	Order                   int          `json:"order,omitempty"`
	Collapsed               int          `json:"collapsed,omitempty"`
	Visible                 int          `json:"visible,omitempty"`
	CreatedAt               string       `json:",omitempty"`
	UpdatedAt               string       `json:",omitempty"`
	EnabledComponents       []*Component `json:",omitempty"`
	EnabledComponentsLowest []*Component `json:",omitempty"`
	LowestHumanStatus       string       `json:",omitempty"`
}

// ComponentGroupResponse reflects the response of /components/groups call
type ComponentGroupResponse struct {
	Meta            Meta             `json:"meta,omitempty"`
	ComponentGroups []ComponentGroup `json:"data,omitempty"`
}

// componentGroupAPIResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the component group in the "data" key.
type componentGroupAPIResponse struct {
	Data *ComponentGroup `json:"data"`
}

// GetAll return all component groups that have been created.
//
// Docs: https://docs.cachethq.io/reference#get-componentgroups
func (s *ComponentGroupsService) GetAll() (*ComponentGroupResponse, *Response, error) {
	u := "api/v1/components/groups"
	v := new(ComponentGroupResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get return a single component group.
//
// Docs: https://docs.cachethq.io/reference#get-a-component-group
func (s *ComponentGroupsService) Get(id int) (*ComponentGroup, *Response, error) {
	u := fmt.Sprintf("api/v1/components/groups/%d", id)
	v := new(componentGroupAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// Create creates a new component group.
//
// Docs: https://docs.cachethq.io/reference#post-componentgroups
func (s *ComponentGroupsService) Create(c *ComponentGroup) (*ComponentGroup, *Response, error) {
	u := "api/v1/components/groups"
	v := new(componentGroupAPIResponse)

	resp, err := s.client.Call("POST", u, c, v)
	return v.Data, resp, err
}

// Update updates a component group.
//
// Docs: https://docs.cachethq.io/reference#put-component-group
func (s *ComponentGroupsService) Update(id int, c *ComponentGroup) (*ComponentGroup, *Response, error) {
	u := fmt.Sprintf("api/v1/components/groups/%d", id)
	v := new(componentGroupAPIResponse)

	resp, err := s.client.Call("PUT", u, c, v)
	return v.Data, resp, err
}

// Delete deletes a component group.
//
// Docs: https://docs.cachethq.io/reference#delete-component-group
func (s *ComponentGroupsService) Delete(id int) (*Response, error) {
	u := fmt.Sprintf("api/v1/components/groups/%d", id)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}
