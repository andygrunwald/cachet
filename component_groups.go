package cachet

import (
	"fmt"
)

// ComponentGroupsService contains REST endpoints that belongs to cachet components.
type ComponentGroupsService struct {
	client *Client
}

// ComponentGroup entity reflects one single component group
type ComponentGroup struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Order     int    `json:"order,omitempty"`
	Collapsed int    `json: "collapsed,omitempty"`
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

// GetAllGroups return all component groups that have been created.
//
// Docs: https://docs.cachethq.io/docs/get-componentgroups
func (s *ComponentGroupsService) GetAllGroups() (*ComponentGroupResponse, *Response, error) {
	u := "api/v1/components/groups"
	v := new(ComponentGroupResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetGroup return a single component group.
//
// Docs: https://docs.cachethq.io/docs/get-a-component-group
func (s *ComponentGroupsService) GetGroup(id int) (*ComponentGroup, *Response, error) {
	u := fmt.Sprintf("api/v1/components/groups/%d", id)
	v := new(componentGroupAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// CreateGroup creates a new component group.
//
// Docs: https://docs.cachethq.io/docs/post-componentgroups
func (s *ComponentGroupsService) CreateGroup(c *ComponentGroup) (*ComponentGroup, *Response, error) {
	u := "api/v1/components/groups"
	v := new(componentGroupAPIResponse)

	resp, err := s.client.Call("POST", u, c, v)
	return v.Data, resp, err
}

// UpdateGroup updates a component group.
//
// Docs: https://docs.cachethq.io/docs/put-component-group
func (s *ComponentGroupsService) UpdateGroup(id int, c *ComponentGroup) (*ComponentGroup, *Response, error) {
	u := fmt.Sprintf("api/v1/components/groups/%d", id)
	v := new(componentGroupAPIResponse)

	resp, err := s.client.Call("PUT", u, c, v)
	return v.Data, resp, err
}

// DeleteGroup deletes a component group.
//
// Docs: https://docs.cachethq.io/docs/delete-component-group
func (s *ComponentGroupsService) DeleteGroup(id int) (*Response, error) {
	u := fmt.Sprintf("api/v1/components/groups/%d", id)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}
