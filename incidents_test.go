package cachet

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestIncidentsService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"meta":{"pagination":{"total":1,"count":1,"per_page":"20","current_page":1,"total_pages":1,"links":{"next_page":null,"previous_page":null}}},"data":[{"id":1,"component_id":0,"name":"Incident Name","status":4,"visible":1,"message":"Incident Message","scheduled_at":"2015-08-01 12:00:00","created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","deleted_at":null,"human_status":"Fixed"}]}`)
	})

	got, _, err := testClient.Incidents.GetAll()
	if err != nil {
		t.Errorf("Incidents.GetAll returned error: %v", err)
	}

	expected := &IncidentResponse{
		Meta: Meta{
			Pagination: Pagination{
				Total:       1,
				Count:       1,
				// TODO Why is this 0? This should be 20. ...Maybe, because it is returned as string. But in components this is an int. Lets check this later
				PerPage:     0,
				CurrentPage: 1,
				TotalPages:  1,
				Links: Links{
					NextPage:     "",
					PreviousPage: "",
				},
			},
		},
		Incidents: []Incident{
			Incident{
				ID:          1,
				ComponentID: 0,
				Name:        "Incident Name",
				Status:      4,
				Visible:     1,
				Message:     "Incident Message",
				ScheduledAt: "2015-08-01 12:00:00",
				CreatedAt:   "2015-08-01 12:00:00",
				UpdatedAt:   "2015-08-01 12:00:00",
				DeletedAt:   "",
				HumanStatus: "Fixed",
			},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Incidents.GetAll returned %+v, want %+v", got, expected)
	}
}

func TestIncidentsService_Get(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id":1,"component_id":0,"name":"Incident Name","status":4,"visible":1,"message":"Incident Message","scheduled_at":"2015-08-01 12:00:00","created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","deleted_at":null,"human_status":"Fixed"}}`)
	})

	got, _, err := testClient.Incidents.Get(1)
	if err != nil {
		t.Errorf("Incidents.Get returned error: %v", err)
	}

	expected := &Incident{
		ID:          1,
		ComponentID: 0,
		Name:        "Incident Name",
		Status:      4,
		Visible:     1,
		Message:     "Incident Message",
		ScheduledAt: "2015-08-01 12:00:00",
		CreatedAt:   "2015-08-01 12:00:00",
		UpdatedAt:   "2015-08-01 12:00:00",
		DeletedAt:   "",
		HumanStatus: "Fixed",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Incidents.Get returned %+v, want %+v", got, expected)
	}
}

func TestIncidentsService_Create(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"data":{"id":1,"component_id":0,"name":"Incident Name","status":4,"visible":1,"message":"Incident Message","scheduled_at":"2015-08-01 12:00:00","created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","deleted_at":null,"human_status":"Fixed"}}`)
	})

	i := &Incident{
		Name:    "Incident Name",
		Message: "Incident Message",
		Status:  IncidentStatusFixed,
		Visible: 1,
	}
	got, _, err := testClient.Incidents.Create(i)
	if err != nil {
		t.Errorf("Incidents.Create returned error: %v", err)
	}

	expected := &Incident{
		ID:          1,
		ComponentID: 0,
		Name:        "Incident Name",
		Status:      4,
		Visible:     1,
		Message:     "Incident Message",
		ScheduledAt: "2015-08-01 12:00:00",
		CreatedAt:   "2015-08-01 12:00:00",
		UpdatedAt:   "2015-08-01 12:00:00",
		DeletedAt:   "",
		HumanStatus: "Fixed",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Incidents.Create returned %+v, want %+v", got, expected)
	}
}

func TestIncidentsService_Update(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		fmt.Fprint(w, `{"data":{"id":1,"component_id":0,"name":"Incident Name","status":4,"visible":1,"message":"Incident Message","scheduled_at":"2015-08-01 12:00:00","created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","deleted_at":null,"human_status":"Fixed"}}`)
	})

	i := &Incident{
		Name: "Incident Name",
	}
	got, _, err := testClient.Incidents.Update(1, i)
	if err != nil {
		t.Errorf("Incidents.Update returned error: %v", err)
	}

	expected := &Incident{
		ID:          1,
		ComponentID: 0,
		Name:        "Incident Name",
		Status:      4,
		Visible:     1,
		Message:     "Incident Message",
		ScheduledAt: "2015-08-01 12:00:00",
		CreatedAt:   "2015-08-01 12:00:00",
		UpdatedAt:   "2015-08-01 12:00:00",
		DeletedAt:   "",
		HumanStatus: "Fixed",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Incidents.Update returned %+v, want %+v", got, expected)
	}
}

func TestIncidentsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := testClient.Incidents.Delete(1)
	if err != nil {
		t.Errorf("Incidents.Delete returned error: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Incidents.Delete returned status %+v, want %+v", resp.StatusCode, http.StatusNoContent)
	}
}
