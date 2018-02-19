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
		fmt.Fprint(w, `{"meta":{"pagination":{"total":1,"count":1,"per_page":20,"current_page":1,"total_pages":1,"links":{"next_page":null,"previous_page":null}}},"data":[{"id":1,"component_id":0,"name":"Incident Name","status":4,"visible":1,"message":"Incident Message","scheduled_at":"2015-08-01 12:00:00","created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","deleted_at":null,"human_status":"Fixed"}]}`)
	})

	got, _, err := testClient.Incidents.GetAll(&ListOptions{})
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

func TestIncidentsService_GetAllWithOptions(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"meta":{"pagination":{"total":1,"count":1,"per_page":50,"current_page":1,"total_pages":1,"links":{"next_page":null,"previous_page":null}}},"data":[{"id":1,"component_id":0,"name":"Incident Name","status":4,"visible":1,"message":"Incident Message","scheduled_at":"2015-08-01 12:00:00","created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","deleted_at":null,"human_status":"Fixed"}]}`)
	})

	got, _, err := testClient.Incidents.GetAll(&ListOptions{PerPage: 50})
	if err != nil {
		t.Errorf("Incidents.GetAll returned error: %v", err)
	}

	expected := &IncidentResponse{
		Meta: Meta{
			Pagination: Pagination{
				Total:       1,
				Count:       1,
				PerPage:     50,
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

func TestIncidentsService_GetAllUpdates(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents/1/updates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"meta":{"pagination":{"total":2,"count":2,"per_page":20,"current_page":1,"total_pages":1,"links":{"next_page":null,"previous_page":null}}},"data":[{"id":1,"incident_id":1,"status":4,"message":"Incident Update Message","user_id":1,"created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","human_status":"Fixed","permalink":"http://cachet.app/incidents/1#update-1"},{"id":2,"incident_id":1,"status":3,"message":"Incident Update Message 2","user_id":1,"created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","human_status":"Watching","permalink":"http://cachet.app/incidents/1#update-2"}]}`)
	})

	got, _, err := testClient.Incidents.GetAllUpdates(&ListOptions{}, 1)
	if err != nil {
		t.Errorf("Incidents.GetAllUpdates returned error: %v", err)
	}

	expected := &IncidentUpdateResponse{
		Meta: Meta{
			Pagination: Pagination{
				Total:       2,
				Count:       2,
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
				Message:     "Incident Update Message",
				UserID:      1,
				CreatedAt:   "2015-08-01 12:00:00",
				UpdatedAt:   "2015-08-01 12:00:00",
				HumanStatus: "Fixed",
				Permalink:   "http://cachet.app/incidents/1#update-1",
			},
			{
				ID:          2,
				IncidentID:  1,
				Status:      3,
				Message:     "Incident Update Message 2",
				UserID:      1,
				CreatedAt:   "2015-08-01 12:00:00",
				UpdatedAt:   "2015-08-01 12:00:00",
				HumanStatus: "Watching",
				Permalink:   "http://cachet.app/incidents/1#update-2",
			},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Incidents.GetAllUpdates returned %+v, want %+v", got, expected)
	}
}

func TestIncidentsService_GetAllUpdatesWithOptions(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents/1/updates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"meta":{"pagination":{"total":2,"count":2,"per_page":50,"current_page":1,"total_pages":1,"links":{"next_page":null,"previous_page":null}}},"data":[{"id":1,"incident_id":1,"status":4,"message":"Incident Update Message","user_id":1,"created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","human_status":"Fixed","permalink":"http://cachet.app/incidents/1#update-1"},{"id":2,"incident_id":1,"status":3,"message":"Incident Update Message 2","user_id":1,"created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","human_status":"Watching","permalink":"http://cachet.app/incidents/1#update-2"}]}`)
	})

	got, _, err := testClient.Incidents.GetAllUpdates(&ListOptions{PerPage: 50}, 1)
	if err != nil {
		t.Errorf("Incidents.GetAllUpdates returned error: %v", err)
	}

	expected := &IncidentUpdateResponse{
		Meta: Meta{
			Pagination: Pagination{
				Total:       2,
				Count:       2,
				PerPage:     50,
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
				Message:     "Incident Update Message",
				UserID:      1,
				CreatedAt:   "2015-08-01 12:00:00",
				UpdatedAt:   "2015-08-01 12:00:00",
				HumanStatus: "Fixed",
				Permalink:   "http://cachet.app/incidents/1#update-1",
			},
			{
				ID:          2,
				IncidentID:  1,
				Status:      3,
				Message:     "Incident Update Message 2",
				UserID:      1,
				CreatedAt:   "2015-08-01 12:00:00",
				UpdatedAt:   "2015-08-01 12:00:00",
				HumanStatus: "Watching",
				Permalink:   "http://cachet.app/incidents/1#update-2",
			},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Incidents.GetAllUpdates returned %+v, want %+v", got, expected)
	}
}

func TestIncidentsService_GetUpdate(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents/1/updates/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id":1,"incident_id":1,"status":4,"message":"Incident Update Message","user_id":1,"created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","human_status":"Fixed","permalink":"http://cachet.app/incidents/1#update-1"}}`)
	})

	got, _, err := testClient.Incidents.GetUpdate(1, 1)
	if err != nil {
		t.Errorf("Incidents.GetUpdate returned error: %v", err)
	}

	expected := &IncidentUpdate{
		ID:          1,
		IncidentID:  1,
		Status:      4,
		Message:     "Incident Update Message",
		UserID:      1,
		CreatedAt:   "2015-08-01 12:00:00",
		UpdatedAt:   "2015-08-01 12:00:00",
		HumanStatus: "Fixed",
		Permalink:   "http://cachet.app/incidents/1#update-1",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Incidents.GetUpdate returned %+v, want %+v", got, expected)
	}
}

func TestIncidentsService_CreateUpdate(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents/1/updates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"data":{"id":1,"incident_id":1,"status":4,"message":"Incident Update Message","user_id":1,"created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","human_status":"Fixed","permalink":"http://cachet.app/incidents/1#update-1"}}`)
	})

	i := &IncidentUpdate{
		Message: "Incident Update Message",
		Status:  IncidentStatusFixed,
	}

	got, _, err := testClient.Incidents.CreateUpdate(1, i)
	if err != nil {
		t.Errorf("Incidents.CreateUpdate returned error: %v", err)
	}

	expected := &IncidentUpdate{
		ID:          1,
		IncidentID:  1,
		Status:      4,
		Message:     "Incident Update Message",
		UserID:      1,
		CreatedAt:   "2015-08-01 12:00:00",
		UpdatedAt:   "2015-08-01 12:00:00",
		HumanStatus: "Fixed",
		Permalink:   "http://cachet.app/incidents/1#update-1",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Incidents.CreateUpdate returned %+v, want %+v", got, expected)
	}
}

func TestIncidentsService_UpdateUpdate(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents/1/updates/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		fmt.Fprint(w, `{"data":{"id":1,"incident_id":1,"status":4,"message":"Incident Update Message 2","user_id":1,"created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","human_status":"Fixed","permalink":"http://cachet.app/incidents/1#update-1"}}`)
	})

	i := &IncidentUpdate{
		Message: "Incident Update Message 2",
	}
	got, _, err := testClient.Incidents.UpdateUpdate(1, 1, i)
	if err != nil {
		t.Errorf("Incidents.UpdateUpdate returned error: %v", err)
	}

	expected := &IncidentUpdate{
		ID:          1,
		IncidentID:  1,
		Status:      4,
		Message:     "Incident Update Message 2",
		UserID:      1,
		CreatedAt:   "2015-08-01 12:00:00",
		UpdatedAt:   "2015-08-01 12:00:00",
		HumanStatus: "Fixed",
		Permalink:   "http://cachet.app/incidents/1#update-1",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Incidents.UpdateUpdate returned %+v, want %+v", got, expected)
	}
}

func TestIncidentsService_DeleteUpdate(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/incidents/1/updates/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := testClient.Incidents.DeleteUpdate(1, 1)
	if err != nil {
		t.Errorf("Incidents.DeleteUpdate returned error: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Incidents.DeleteUpdate returned status %+v, want %+v", resp.StatusCode, http.StatusNoContent)
	}
}
