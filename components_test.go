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
			Component{
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
