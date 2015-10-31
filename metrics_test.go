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

func TestMetricsService_Create(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/metrics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"data":{"name":"Beer2","display_chart":true,"default_value":0,"calc_type":0,"places":1,"suffix":"Bottles","description":"How many beers we drank.","updated_at":"2015-10-31 16:56:18","created_at":"2015-10-31 16:56:18","id":5}}`)
	})

	m := &Metric{
		Name:         "Beer2",
		Suffix:       "Bottles",
		Description:  "How many beers we drank.",
		DefaultValue: 0,
		DisplayChart: true,
		Places:       1,
	}
	got, _, err := testClient.Metrics.Create(m)
	if err != nil {
		t.Errorf("Metrics.Create returned error: %v", err)
	}

	expected := &Metric{
		ID:           5,
		Name:         "Beer2",
		Suffix:       "Bottles",
		Description:  "How many beers we drank.",
		DisplayChart: true,
		DefaultValue: 0,
		CalcType:     0,
		Places:       1,
		CreatedAt:    "2015-10-31 16:56:18",
		UpdatedAt:    "2015-10-31 16:56:18",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Metrics.Create returned %+v, want %+v", got, expected)
	}
}

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

func TestMetricsService_GetPoints(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/metrics/1/points", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":[{"id":1,"metric_id":1,"value":4,"created_at":"2015-10-31 16:30:02","updated_at":"2015-10-31 16:30:02"},{"id":2,"metric_id":1,"value":6,"created_at":"2015-10-31 15:30:02","updated_at":"2015-10-31 16:30:02"}]}`)
	})

	got, _, err := testClient.Metrics.GetPoints(1)
	if err != nil {
		t.Errorf("Metrics.GetPoints returned error: %v", err)
	}

	expected := &[]Point{
		Point{
			ID:        1,
			MetricID:  1,
			Value:     4,
			CreatedAt: "2015-10-31 16:30:02",
			UpdatedAt: "2015-10-31 16:30:02",
		},
		Point{
			ID:        2,
			MetricID:  1,
			Value:     6,
			CreatedAt: "2015-10-31 15:30:02",
			UpdatedAt: "2015-10-31 16:30:02",
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Metrics.GetPoints returned %+v, want %+v", got, expected)
	}
}

func TestMetricsService_AddPoint(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/metrics/1/points", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"data":{"metric_id":1,"value":20,"updated_at":"2015-10-31 16:46:04","created_at":"2015-10-31 16:46:04","id":14}}`)
	})

	got, _, err := testClient.Metrics.AddPoint(1, 20, "")
	if err != nil {
		t.Errorf("Metrics.AddPoint returned error: %v", err)
	}

	expected := &Point{

		ID:        14,
		MetricID:  1,
		Value:     20,
		CreatedAt: "2015-10-31 16:46:04",
		UpdatedAt: "2015-10-31 16:46:04",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Metrics.AddPoint returned %+v, want %+v", got, expected)
	}
}

func TestMetricsService_DeletePoint(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/metrics/1/points/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := testClient.Metrics.DeletePoint(1, 1)
	if err != nil {
		t.Errorf("Metrics.DeletePoint returned error: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Metrics.DeletePoint returned status %+v, want %+v", resp.StatusCode, http.StatusNoContent)
	}
}
