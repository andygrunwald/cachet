package cachet

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestComponentsService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/components", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"meta":{"pagination":{"total":1,"count":1,"per_page":20,"current_page":1,"total_pages":1,"links":{"next_page":null,"previous_page":null}}},"data":[{"id":1,"name":"API","description":"This is the Cachet API.","link":"","status":1,"order":0,"group_id":0,"created_at":"2015-07-24 14:42:10","updated_at":"2015-07-24 14:42:10","deleted_at":null,"status_name":"Operational"}]}`)
	})

	got, _, err := testClient.Components.GetAll()
	if err != nil {
		t.Errorf("Components.GetAll returned error: %v", err)
	}

	expected := &ComponentResponse{
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
		Components: []Component{
			{
				ID:          1,
				Name:        "API",
				Description: "This is the Cachet API.",
				Link:        "",
				Status:      1,
				Order:       0,
				GroupID:     0,
				CreatedAt:   "2015-07-24 14:42:10",
				UpdatedAt:   "2015-07-24 14:42:10",
				DeletedAt:   "",
				StatusName:  "Operational",
			},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Components.GetAll returned %+v, want %+v", got, expected)
	}
}

func TestComponentsService_Get(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/components/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id":1,"name":"API","description":"Used by third-parties to connect to us","link":"","status":1,"order":0,"group_id":0,"created_at":"2015-10-31 08:30:01","updated_at":"2015-10-31 08:30:01","deleted_at":null,"status_name":"Operational"}}`)
	})

	got, _, err := testClient.Components.Get(1)
	if err != nil {
		t.Errorf("Components.Get returned error: %v", err)
	}

	expected := &Component{
		ID:          1,
		Name:        "API",
		Description: "Used by third-parties to connect to us",
		Link:        "",
		Status:      1,
		Order:       0,
		GroupID:     0,
		CreatedAt:   "2015-10-31 08:30:01",
		UpdatedAt:   "2015-10-31 08:30:01",
		DeletedAt:   "",
		StatusName:  "Operational",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Components.Get returned %+v, want %+v", got, expected)
	}
}

func TestComponentsService_Create(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/components", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"data":{"id":1,"name":"Component Name","description":"Description","link":"","status":1,"order":0,"group_id":0,"created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","deleted_at":null,"status_name":"Operational"}}`)
	})

	co := &Component{
		Name:   "Go API (by Token) - Updated!",
		Status: 1,
	}
	got, _, err := testClient.Components.Create(co)
	if err != nil {
		t.Errorf("Components.Create returned error: %v", err)
	}

	expected := &Component{
		ID:          1,
		Name:        "Component Name",
		Description: "Description",
		Link:        "",
		Status:      1,
		Order:       0,
		GroupID:     0,
		CreatedAt:   "2015-08-01 12:00:00",
		UpdatedAt:   "2015-08-01 12:00:00",
		DeletedAt:   "",
		StatusName:  "Operational",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Components.Create returned %+v, want %+v", got, expected)
	}
}

func TestComponentsService_Update(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/components/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		fmt.Fprint(w, `{"data":{"id":1,"name":"Component Name","description":"Description","link":"","status":1,"order":0,"group_id":0,"created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","deleted_at":null,"status_name":"Operational"}}`)
	})

	co := &Component{
		Name:   "Component Name",
		Status: 1,
	}
	got, _, err := testClient.Components.Update(1, co)
	if err != nil {
		t.Errorf("Components.Update returned error: %v", err)
	}

	expected := &Component{
		ID:          1,
		Name:        "Component Name",
		Description: "Description",
		Link:        "",
		Status:      1,
		Order:       0,
		GroupID:     0,
		CreatedAt:   "2015-08-01 12:00:00",
		UpdatedAt:   "2015-08-01 12:00:00",
		DeletedAt:   "",
		StatusName:  "Operational",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Components.Update returned %+v, want %+v", got, expected)
	}
}

func TestComponentsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/components/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := testClient.Components.Delete(1)
	if err != nil {
		t.Errorf("Components.Delete returned error: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Components.Delete returned status %+v, want %+v", resp.StatusCode, http.StatusNoContent)
	}
}

func TestComponentsService_GetAllGroups(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/components/groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"meta":{"pagination":{"total":1,"count":1,"per_page":20,"current_page":1,"total_pages":1,"links":{"next_page":null,"previous_page":null}}},"data":[{"id":1,"name":"Websites","created_at":"2015-11-07 16:30:02","updated_at":"2015-11-07 16:30:02","order":1}]}`)
	})

	got, _, err := testClient.Components.GetAllGroups()
	if err != nil {
		t.Errorf("Components.GetAllGroups returned error: %v", err)
	}

	expected := &ComponentGroupResponse{
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
		ComponentGroups: []ComponentGroup{
			{
				ID:        1,
				Name:      "Websites",
				Order:     1,
				CreatedAt: "2015-11-07 16:30:02",
				UpdatedAt: "2015-11-07 16:30:02",
			},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Components.GetAllGroups returned %+v, want %+v", got, expected)
	}
}

func TestComponentsService_GetGroup(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/components/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id":1,"name":"Websites","created_at":"2015-11-07 16:30:02","updated_at":"2015-11-07 16:30:02","order":1}}`)
	})

	got, _, err := testClient.Components.GetGroup(1)
	if err != nil {
		t.Errorf("Components.GetGroup returned error: %v", err)
	}

	expected := &ComponentGroup{
		ID:        1,
		Name:      "Websites",
		Order:     1,
		CreatedAt: "2015-11-07 16:30:02",
		UpdatedAt: "2015-11-07 16:30:02",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Components.GetGroup returned %+v, want %+v", got, expected)
	}
}

func TestComponentsService_CreateGroup(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/components/groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"data":{"name":"UnitTest","order":2,"updated_at":"2015-11-07 17:25:16","created_at":"2015-11-07 17:25:16","id":2}}`)
	})

	cg := &ComponentGroup{
		Name:  "UnitTest",
		Order: 2,
	}
	got, _, err := testClient.Components.CreateGroup(cg)
	if err != nil {
		t.Errorf("Components.CreateGroup returned error: %v", err)
	}

	expected := &ComponentGroup{
		ID:        2,
		Name:      "UnitTest",
		Order:     2,
		CreatedAt: "2015-11-07 17:25:16",
		UpdatedAt: "2015-11-07 17:25:16",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Components.CreateGroup returned %+v, want %+v", got, expected)
	}
}

func TestComponentsService_UpdateGroup(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/components/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		fmt.Fprint(w, `{"data":{"id":1,"name":"Updated Component","created_at":"2015-11-07 16:30:02","updated_at":"2015-11-07 17:27:32","order":3}}`)
	})

	cg := &ComponentGroup{
		Name:  "Updated Component",
		Order: 3,
	}
	got, _, err := testClient.Components.UpdateGroup(1, cg)
	if err != nil {
		t.Errorf("Components.Update returned error: %v", err)
	}

	expected := &ComponentGroup{
		ID:        1,
		Name:      "Updated Component",
		Order:     3,
		CreatedAt: "2015-11-07 16:30:02",
		UpdatedAt: "2015-11-07 17:27:32",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Components.UpdateGroup returned %+v, want %+v", got, expected)
	}
}

func TestComponentsService_DeleteGroup(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/components/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := testClient.Components.DeleteGroup(1)
	if err != nil {
		t.Errorf("Components.DeleteGroup returned error: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Components.DeleteGroup returned status %+v, want %+v", resp.StatusCode, http.StatusNoContent)
	}
}
