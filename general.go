package cachet

// GeneralService contains REST endpoints that belongs no specific service.
type GeneralService struct {
	client *Client
}

// PingResponse entity contains the Response of a /ping call.
type PingResponse struct {
	Data string `json:"data,omitempty"`
}

// VersionResponse entity contains the Response of a /version call.
type VersionResponse struct {
	Meta MetaVersion `json:"meta,omitempty"`
	Data string      `json:"data,omitempty"`
}

// StatusResponse entity contains the Response of a /status call.
type StatusResponse struct {
	Data *Status `json:"data"`
}

// Status entity contains the contents of API Response of a /status call.
type Status struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

// Ping calls the API test endpoint.
//
// Docs: https://docs.cachethq.io/reference#ping
func (s *GeneralService) Ping() (string, *Response, error) {
	u := "api/v1/ping"
	v := new(PingResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// Version get Cachet version
//
// Docs: https://docs.cachethq.io/reference#version
func (s *GeneralService) Version() (*VersionResponse, *Response, error) {
	u := "api/v1/version"
	v := new(VersionResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Status get Cachet status
//
// Docs: <none>
func (s *GeneralService) Status() (*Status, *Response, error) {
	u := "api/v1/status"
	v := new(StatusResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}
