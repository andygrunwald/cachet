package cachet

// GeneralService contains REST endpoints that belongs no specific service.
type GeneralService struct {
	client *Client
}

// Ping entity contains the Response of a /ping call.
type Ping struct {
	Data string `json:"data,omitempty"`
}

// Ping calls the API test endpoint.
//
// Docs: https://docs.cachethq.io/docs/ping
func (s *GeneralService) Ping() (string, *Response, error) {
	u := "api/v1/ping"
	v := new(Ping)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}
