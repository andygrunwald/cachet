package cachet

import (
	"fmt"
)

// SubscriptionsService contains REST endpoints that belongs to cachet subscriptions.
type SubscriptionsService struct {
	client *Client
}

// Delete deletes a subscription.
//
// Docs: https://docs.cachethq.io/reference#incidentsincidentupdatesupdate
func (s *SubscriptionsService) Delete(id int) (*Response, error) {
	u := fmt.Sprintf("api/v1/subscription/%d", id)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}
