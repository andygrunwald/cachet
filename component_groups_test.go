package cachet

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestComponentsService_GetAllGroups(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/components/groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"meta":{"pagination":{"total":1,"count":1,"per_page":20,"current_page":1,"total_pages":1,"links":{"next_page":null,"previous_page":null}}},"data":[{"id":1,"name":"Websites","created_at":"2015-11-07 16:30:02","updated_at":"2015-11-07 16:30:02","order":1}]}`)
	})

	queryParams := &ComponentGroupsQueryParams{}

	got, _, err := testClient.ComponentGroups.GetAll(queryParams)
	if err != nil {
		t.Errorf("ComponentGroups.GetAll returned error: %v", err)
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

	got, _, err := testClient.ComponentGroups.Get(1)
	if err != nil {
		t.Errorf("ComponentGroups.Get returned error: %v", err)
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
	got, _, err := testClient.ComponentGroups.Create(cg)
	if err != nil {
		t.Errorf("ComponentGroups.Create returned error: %v", err)
	}

	expected := &ComponentGroup{
		ID:        2,
		Name:      "UnitTest",
		Order:     2,
		CreatedAt: "2015-11-07 17:25:16",
		UpdatedAt: "2015-11-07 17:25:16",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ComponentGroups.Create returned %+v, want %+v", got, expected)
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
	got, _, err := testClient.ComponentGroups.Update(1, cg)
	if err != nil {
		t.Errorf("ComponentGroups.Update returned error: %v", err)
	}

	expected := &ComponentGroup{
		ID:        1,
		Name:      "Updated Component",
		Order:     3,
		CreatedAt: "2015-11-07 16:30:02",
		UpdatedAt: "2015-11-07 17:27:32",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ComponentGroups.Update returned %+v, want %+v", got, expected)
	}
}

func TestComponentsService_DeleteGroup(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/components/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := testClient.ComponentGroups.Delete(1)
	if err != nil {
		t.Errorf("ComponentGroups.Delete returned error: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("ComponentGroups.Delete returned status %+v, want %+v", resp.StatusCode, http.StatusNoContent)
	}
}
