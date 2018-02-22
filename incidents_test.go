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
		fmt.Fprint(w, `{"meta":{"pagination":{"total":1,"count":1,"per_page":20,"current_page":1,"total_pages":1,"links":{"next_page":null,"previous_page":null}}},"data":[{"id":1,"component_id":1,"name":"Incident Name","status":2,"visible":1,"stickied":false,"message":"Incident Message","occurred_at":"2015-08-01 12:00:00","created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","deleted_at":null,"is_resolved":false,"updates":[{"id":1,"incident_id":1,"status":2,"message":"Incident Update #1","user_id":1,"created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","human_status":"Identified","permalink":"http://localhost/incidents/1#update-1"}],"human_status":"Fixed","latest_update_id":1,"latest_status":2,"latest_human_status":"Identified","latest_icon":"icon ion-alert yellows","permalink":"http://localhost/incidents/1","duration":45}]}`)
	})

	queryParams := &IncidentsQueryParams{}

	got, _, err := testClient.Incidents.GetAll(queryParams)
	if err != nil {
		t.Errorf("Incidents.GetAll returned error: %v", err)
	}

	expected := &IncidentResponse{
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
		Incidents: []Incident{
			{
				ID:          1,
				ComponentID: 1,
				Name:        "Incident Name",
				Status:      2,
				Visible:     1,
				Message:     "Incident Message",
				OccurredAt:  "2015-08-01 12:00:00",
				CreatedAt:   "2015-08-01 12:00:00",
				UpdatedAt:   "2015-08-01 12:00:00",
				DeletedAt:   "",
				IsResolved:  false,
				Updates: []IncidentUpdate{
					{
						ID:          1,
						IncidentID:  1,
						Status:      2,
						Message:     "Incident Update #1",
						UserID:      1,
						CreatedAt:   "2015-08-01 12:00:00",
						UpdatedAt:   "2015-08-01 12:00:00",
						HumanStatus: "Identified",
						Permalink:   "http://localhost/incidents/1#update-1",
					},
				},
				HumanStatus:       "Fixed",
				LatestUpdateID:    1,
				LatestStatus:      2,
				LatestHumanStatus: "Identified",
				LatestIcon:        "icon ion-alert yellows",
				Permalink:         "http://localhost/incidents/1",
				Duration:          45,
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
		fmt.Fprint(w, `{"data":{"id":1,"component_id":1,"name":"Incident Name","status":2,"visible":1,"stickied":false,"message":"Incident Message","occurred_at":"2015-08-01 12:00:00","created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","deleted_at":null,"is_resolved":false,"updates":[{"id":1,"incident_id":1,"status":2,"message":"Incident Update #1","user_id":1,"created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","human_status":"Identified","permalink":"http://localhost/incidents/1#update-1"}],"human_status":"Fixed","latest_update_id":1,"latest_status":2,"latest_human_status":"Identified","latest_icon":"icon ion-alert yellows","permalink":"http://localhost/incidents/1","duration":45}}`)
	})

	got, _, err := testClient.Incidents.Get(1)
	if err != nil {
		t.Errorf("Incidents.Get returned error: %v", err)
	}

	expected := &Incident{
		ID:          1,
		ComponentID: 1,
		Name:        "Incident Name",
		Status:      2,
		Visible:     1,
		Stickied:    false,
		Message:     "Incident Message",
		OccurredAt:  "2015-08-01 12:00:00",
		CreatedAt:   "2015-08-01 12:00:00",
		UpdatedAt:   "2015-08-01 12:00:00",
		DeletedAt:   "",
		IsResolved:  false,
		Updates: []IncidentUpdate{
			{
				ID:          1,
				IncidentID:  1,
				Status:      2,
				Message:     "Incident Update #1",
				UserID:      1,
				CreatedAt:   "2015-08-01 12:00:00",
				UpdatedAt:   "2015-08-01 12:00:00",
				HumanStatus: "Identified",
				Permalink:   "http://localhost/incidents/1#update-1",
			},
		},
		HumanStatus:       "Fixed",
		LatestUpdateID:    1,
		LatestStatus:      2,
		LatestHumanStatus: "Identified",
		LatestIcon:        "icon ion-alert yellows",
		Permalink:         "http://localhost/incidents/1",
		Duration:          45,
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Incidents.Get returned %+v, want %+v", got, expected)
	}
}

func TestIncidentsService_Create_No_Component(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"data":{"id":1,"name":"Incident Name","status":4,"visible":1,"stickied":false,"message":"Incident Message","occurred_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","created_at":"2015-08-01 12:00:00","is_resolved":false,"component":null,"updates":[],"human_status":"Fixed","latest_update_id":null,"latest_status":4,"latest_human_status":"Fixed","latest_icon":"icon ion-checkmark greens","permalink":"http://localhost/incidents/1","duration":0}}`)
	})

	i := &Incident{
		Name:    "Incident Name",
		Message: "Incident Message",
		Status:  IncidentStatusFixed,
		Visible: IncidentVisibilityPublic,
	}
	got, _, err := testClient.Incidents.Create(i)
	if err != nil {
		t.Errorf("Incidents.Create returned error: %v", err)
	}

	expected := &Incident{
		ID:                1,
		Name:              "Incident Name",
		Status:            4,
		Visible:           1,
		Stickied:          false,
		Message:           "Incident Message",
		OccurredAt:        "2015-08-01 12:00:00",
		CreatedAt:         "2015-08-01 12:00:00",
		UpdatedAt:         "2015-08-01 12:00:00",
		IsResolved:        false,
		Updates:           []IncidentUpdate{},
		HumanStatus:       "Fixed",
		LatestUpdateID:    0,
		LatestStatus:      4,
		LatestIcon:        "icon ion-checkmark greens",
		LatestHumanStatus: "Fixed",
		Permalink:         "http://localhost/incidents/1",
		Duration:          0,
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
		fmt.Fprint(w, `{"data":{"id":1,"name":"Incident Name v2","status":4,"visible":1,"stickied":false,"message":"Incident Message v2","occurred_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","created_at":"2015-08-01 12:00:00","is_resolved":true,"component":null,"updates":[],"human_status":"Fixed","latest_update_id":null,"latest_status":4,"latest_human_status":"Fixed","latest_icon":"icon ion-checkmark greens","permalink":"http://localhost/incidents/1","duration":0}}`)
	})

	i := &Incident{
		Name: "Incident Name",
	}
	got, _, err := testClient.Incidents.Update(1, i)
	if err != nil {
		t.Errorf("Incidents.Update returned error: %v", err)
	}

	expected := &Incident{
		ID:                1,
		Name:              "Incident Name v2",
		Status:            4,
		Visible:           1,
		Stickied:          false,
		Message:           "Incident Message v2",
		OccurredAt:        "2015-08-01 12:00:00",
		CreatedAt:         "2015-08-01 12:00:00",
		UpdatedAt:         "2015-08-01 12:00:00",
		IsResolved:        true,
		Updates:           []IncidentUpdate{},
		HumanStatus:       "Fixed",
		LatestUpdateID:    0,
		LatestStatus:      4,
		LatestIcon:        "icon ion-checkmark greens",
		LatestHumanStatus: "Fixed",
		Permalink:         "http://localhost/incidents/1",
		Duration:          0,
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
