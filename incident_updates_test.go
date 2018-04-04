package cachet

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestIncidentUpdatesService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents/1/updates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"meta":{"pagination":{"total":1,"count":1,"per_page":20,"current_page":1,"total_pages":1,"links":{"next_page":null,"previous_page":null}}},"data":[{"id":1,"incident_id":1,"status":4,"message":"Incident Message","user_id":1,"created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","human_status":"Fixed","permalink":"https://dev.cachethq.io/incidents/1#update-1"}]}`)
	})

	got, _, err := testClient.IncidentUpdates.GetAll(1)
	if err != nil {
		t.Errorf("IncidentUpdates.GetAll returned error: %v", err)
	}

	expected := &IncidentUpdateResponse{
		Meta: Meta{
			Pagination: Pagination{
				Total:       1,
				Count:       1,
				PerPage:     20,
				CurrentPage: 1,
				TotalPages:  1,
				Links: Links{
					NextPage:     "",
					PreviousPage: "",
				},
			},
		},
		IncidentUpdates: []IncidentUpdate{
			{
				ID:          1,
				IncidentID:  1,
				Status:      4,
				Message:     "Incident Message",
				UserID:      1,
				CreatedAt:   "2015-08-01 12:00:00",
				UpdatedAt:   "2015-08-01 12:00:00",
				HumanStatus: "Fixed",
				Permalink:   "https://dev.cachethq.io/incidents/1#update-1",
			},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("IncidentUpdates.GetAll returned %+v, want %+v", got, expected)
	}
}

func TestIncidentUpdatesService_Get(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents/1/updates/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id":1,"incident_id":1,"status":4,"message":"Incident Message","user_id":1,"created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","human_status":"Fixed","permalink":"https://dev.cachethq.io/incidents/1/updates/1"}}`)
	})

	got, _, err := testClient.IncidentUpdates.Get(1, 1)
	if err != nil {
		t.Errorf("IncidentUpdates.Get returned error: %v", err)
	}

	expected := &IncidentUpdate{
		ID:          1,
		IncidentID:  1,
		Status:      4,
		Message:     "Incident Message",
		UserID:      1,
		CreatedAt:   "2015-08-01 12:00:00",
		UpdatedAt:   "2015-08-01 12:00:00",
		HumanStatus: "Fixed",
		Permalink:   "https://dev.cachethq.io/incidents/1/updates/1",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("IncidentUpdates.Get returned %+v, want %+v", got, expected)
	}
}

func TestIncidentUpdatesService_Create(t *testing.T) {
}

func TestIncidentUpdatesService_Update(t *testing.T) {
}

func TestIncidentUpdatesService_Delete(t *testing.T) {
}
