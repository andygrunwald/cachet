package cachet

import (
	"fmt"
)

// MetricsService contains REST endpoints that belongs to cachet metrics.
type MetricsService struct {
	client *Client
}

// Metric entity reflects one single metric
type Metric struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Suffix      string `json:"suffix,omitempty"`
	Description string `json:"description,omitempty"`
	// TODO: Check what the API response here ... string or int?
	DefaultValue int     `json:"default_value,omitempty"`
	CalcType     int     `json:"calc_type,omitempty"`
	DisplayChart bool    `json:"display_chart,omitempty"`
	CreatedAt    string  `json:"created_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
	Places       int     `json:"places,omitempty"`
	Points       []Point `json:"points,omitempty"`
}

// Point is a single point in a Metric
type Point struct {
	ID        int    `json:"id,omitempty"`
	MetricID  int    `json:"metric_id,omitempty"`
	Value     int    `json:"value,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// MetricResponse reflects the response of /metric call
type MetricResponse struct {
	Meta    Meta     `json:"meta,omitempty"`
	Metrics []Metric `json:"data,omitempty"`
}

// metricAPIResponse is an internal type to hide
// some the "data" nested level from the API.
// Some calls (e.g. Get or Create) return the metric in the "data" key.
type metricAPIResponse struct {
	Data *Metric `json:"data"`
}

// GetAll returns all metrics that have been setup.
//
// Docs: https://docs.cachethq.io/docs/get-metrics
func (s *MetricsService) GetAll() (*MetricResponse, *Response, error) {
	u := "api/v1/metrics"
	v := new(MetricResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns a single metric, without points.
//
// Docs: https://docs.cachethq.io/docs/get-a-metric
func (s *MetricsService) Get(id int) (*Metric, *Response, error) {
	u := fmt.Sprintf("api/v1/metrics/%d", id)
	v := new(metricAPIResponse)

	resp, err := s.client.Call("GET", u, nil, v)
	return v.Data, resp, err
}

// Create a new metric.
//
// Docs: https://docs.cachethq.io/docs/metrics
func (s *MetricsService) Create(m *Metric) (*Metric, *Response, error) {
	u := "api/v1/metrics"
	v := new(metricAPIResponse)

	resp, err := s.client.Call("POST", u, m, v)
	return v.Data, resp, err
}

// Delete a metric.
//
// Docs: https://docs.cachethq.io/docs/delete-a-metric
func (s *MetricsService) Delete(id int) (*Response, error) {
	u := fmt.Sprintf("api/v1/metrics/%d", id)

	resp, err := s.client.Call("DELETE", u, nil, nil)
	return resp, err
}
