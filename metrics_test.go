package cachet

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMetricsService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/metrics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"meta":{"pagination":{"total":4,"count":4,"per_page":20,"current_page":1,"total_pages":1,"links":{"next_page":null,"previous_page":null}}},"data":[{"id":1,"name":"Cups of coffee","suffix":"Cups","description":"How many cups of coffee we've drank.","default_value":0,"calc_type":1,"display_chart":true,"created_at":"2015-10-31 14:30:02","updated_at":"2015-10-31 14:30:02","places":2,"points":[{"id":1,"metric_id":1,"value":7,"created_at":"2015-10-31 14:30:02","updated_at":"2015-10-31 14:30:02"}]}]}`)
	})

	got, _, err := testClient.Metrics.GetAll()
	if err != nil {
		t.Errorf("Metrics.GetAll returned error: %v", err)
	}

	expected := &MetricResponse{
		Meta: Meta{
			Pagination: Pagination{
				Total:       4,
				Count:       4,
				PerPage:     20,
				CurrentPage: 1,
				TotalPages:  1,
				Links: Links{
					NextPage:     "",
					PreviousPage: "",
				},
			},
		},
		Metrics: []Metric{
			Metric{
				ID:           1,
				Name:         "Cups of coffee",
				Suffix:       "Cups",
				Description:  "How many cups of coffee we've drank.",
				DefaultValue: 0,
				CalcType:     1,
				DisplayChart: true,
				CreatedAt:    "2015-10-31 14:30:02",
				UpdatedAt:    "2015-10-31 14:30:02",
				Places:       2,
				Points: []Point{
					Point{
						ID:        1,
						MetricID:  1,
						Value:     7,
						CreatedAt: "2015-10-31 14:30:02",
						UpdatedAt: "2015-10-31 14:30:02",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Metrics.GetAll returned %+v, want %+v", got, expected)
	}
}

func TestMetricsService_Get(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/metrics/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id":1,"name":"Cups of coffee","suffix":"Cups","description":"How many cups of coffee we've drank.","default_value":0,"calc_type":1,"display_chart":true,"created_at":"2015-10-31 14:30:02","updated_at":"2015-10-31 14:30:02","places":2,"points":[{"id":1,"metric_id":1,"value":7,"created_at":"2015-10-31 14:30:02","updated_at":"2015-10-31 14:30:02"}]}}`)
	})

	got, _, err := testClient.Metrics.Get(1)
	if err != nil {
		t.Errorf("Metrics.Get returned error: %v", err)
	}

	expected := &Metric{
		ID:           1,
		Name:         "Cups of coffee",
		Suffix:       "Cups",
		Description:  "How many cups of coffee we've drank.",
		DefaultValue: 0,
		CalcType:     1,
		DisplayChart: true,
		CreatedAt:    "2015-10-31 14:30:02",
		UpdatedAt:    "2015-10-31 14:30:02",
		Places:       2,
		Points: []Point{
			Point{
				ID:        1,
				MetricID:  1,
				Value:     7,
				CreatedAt: "2015-10-31 14:30:02",
				UpdatedAt: "2015-10-31 14:30:02",
			},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Metrics.Get returned %+v, want %+v", got, expected)
	}
}

// TODO: Unit test for TestMetricsService_Create

func TestMetricsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/metrics/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := testClient.Metrics.Delete(1)
	if err != nil {
		t.Errorf("Metrics.Delete returned error: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Metrics.Delete returned status %+v, want %+v", resp.StatusCode, http.StatusNoContent)
	}
}
