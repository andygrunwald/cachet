package cachet

import (
	"fmt"
)

// SubscribersService contains REST endpoints that belongs to cachet subscribers.
type SubscribersService struct {
	client *Client
}

// Subscriber entity reflects one single subscriber
type Subscriber struct {
	ID         int    `json:"id,omitempty"`
	EMail      string `json:"email,omitempty"`
	VerifyCode string `json:"verify_code,omitempty"`
	VerifiedAt string `json:"verified_at,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}

// SubscriberResponse reflects the response of /subscribers call
type SubscriberResponse struct {
	Meta        Meta         `json:"meta,omitempty"`
	Subscribers []Subscriber `json:"data,omitempty"`
}

// subscriberApiResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the subscriber in the "data" key.
type subscriberAPIResponse struct {
	Data *Subscriber `json:"data"`
}

// GetAll returns all subscribers.
//
// Docs: https://docs.cachethq.io/docs/get-subscribers
func (s *SubscribersService) GetAll(opt *ListOptions) (*SubscriberResponse, *Response, error) {
	u := "api/v1/subscribers"
	v := new(SubscriberResponse)

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Create a new subscriber.
//
// Docs: https://docs.cachethq.io/docs/subscribers
func (s *SubscribersService) Create(email string, verify int) (*Subscriber, *Response, error) {
	u := "api/v1/subscribers"
	v := new(subscriberAPIResponse)

	c := struct {
		EMail  string `json:"email"`
		Verify int    `json:"verify"`
	}{
		EMail:  email,
		Verify: verify,
	}

	resp, err := s.client.Call("POST", u, c, v)
	return v.Data, resp, err
}

// Delete a subscriber.
//
// Docs: https://docs.cachethq.io/docs/delete-subscriber
func (s *SubscribersService) Delete(id int) (*Response, error) {
	u := fmt.Sprintf("api/v1/subscribers/%d", id)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}
