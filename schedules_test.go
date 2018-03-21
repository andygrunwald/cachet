package cachet

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSchedulesService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/schedules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"meta":{"pagination":{"total":1,"count":1,"per_page":20,"current_page":1,"total_pages":1,"links":{"next_page":null,"previous_page":null}}},"data":[{"id":1,"name":"Schedule Name","status":2,"message":"Schedule Message","scheduled_at":"2015-08-01 12:30:00","completed_at":"2015-08-01 13:00:00","created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","components":[],"human_status":"Complete"}]}`)
	})

	queryParams := &SchedulesQueryParams{}

	got, _, err := testClient.Schedules.GetAll(queryParams)
	if err != nil {
		t.Errorf("Schedules.GetAll returned error: %v", err)
	}

	expected := &ScheduleResponse{
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
		Schedules: []Schedule{
			{
				ID:          1,
				Name:        "Schedule Name",
				Status:      2,
				Message:     "Schedule Message",
				ScheduledAt: "2015-08-01 12:30:00",
				CompletedAt: "2015-08-01 13:00:00",
				CreatedAt:   "2015-08-01 12:00:00",
				UpdatedAt:   "2015-08-01 12:00:00",
				Components:  []Component{},
				HumanStatus: "Complete",
			},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Schedules.GetAll returned %+v, want %+v", got, expected)
	}
}

func TestSchedulesService_Get(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/schedules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id":1,"name":"Schedule Name","status":2,"message":"Schedule Message","scheduled_at":"2015-08-01 12:30:00","completed_at":"2015-08-01 13:00:00","created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","components":[],"human_status":"Complete"}}`)
	})

	got, _, err := testClient.Schedules.Get(1)
	if err != nil {
		t.Errorf("Schedules.Get returned error: %v", err)
	}

	expected := &Schedule{
		ID:          1,
		Name:        "Schedule Name",
		Status:      2,
		Message:     "Schedule Message",
		ScheduledAt: "2015-08-01 12:30:00",
		CompletedAt: "2015-08-01 13:00:00",
		CreatedAt:   "2015-08-01 12:00:00",
		UpdatedAt:   "2015-08-01 12:00:00",
		Components:  []Component{},
		HumanStatus: "Complete",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Schedules.Get returned %+v, want %+v", got, expected)
	}
}

func TestSchedulesService_Create(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/schedules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"data":{"id":1,"name":"Schedule Name","status":2,"message":"Schedule Message","scheduled_at":"2015-08-01 12:30:00","completed_at":"2015-08-01 13:00:00","created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","components":[],"human_status":"Complete"}}`)
	})

	i := &Schedule{
		Name:        "Schedule Name",
		Message:     "Schedule Message",
		Status:      ScheduleComplete,
		ScheduledAt: "2015-08-01 12:30:00",
		CompletedAt: "2015-08-01 13:00:00",
		Components:  []Component{},
	}
	got, _, err := testClient.Schedules.Create(i)
	if err != nil {
		t.Errorf("Schedules.Create returned error: %v", err)
	}

	expected := &Schedule{
		ID:          1,
		Name:        "Schedule Name",
		Status:      2,
		Message:     "Schedule Message",
		ScheduledAt: "2015-08-01 12:30:00",
		CompletedAt: "2015-08-01 13:00:00",
		CreatedAt:   "2015-08-01 12:00:00",
		UpdatedAt:   "2015-08-01 12:00:00",
		Components:  []Component{},
		HumanStatus: "Complete",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Schedule.Create returned %+v, want %+v", got, expected)
	}
}

func TestSchedulesService_Update(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/schedules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		fmt.Fprint(w, `{"data":{"id":1,"name":"Schedule Name Update","status":2,"message":"Schedule Message","scheduled_at":"2015-08-01 12:30:00","completed_at":"2015-08-01 13:00:00","created_at":"2015-08-01 12:00:00","updated_at":"2015-08-01 12:00:00","components":[],"human_status":"Complete"}}`)
	})

	i := &Schedule{
		Name: "Schedule Name Update",
	}
	got, _, err := testClient.Schedules.Update(1, i)
	if err != nil {
		t.Errorf("Schedules.Update returned error: %v", err)
	}

	expected := &Schedule{
		ID:          1,
		Name:        "Schedule Name Update",
		Status:      2,
		Message:     "Schedule Message",
		ScheduledAt: "2015-08-01 12:30:00",
		CompletedAt: "2015-08-01 13:00:00",
		CreatedAt:   "2015-08-01 12:00:00",
		UpdatedAt:   "2015-08-01 12:00:00",
		Components:  []Component{},
		HumanStatus: "Complete",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Schedules.Update returned %+v, want %+v", got, expected)
	}
}

func TestSchedulesService_Delete(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/schedules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := testClient.Schedules.Delete(1)
	if err != nil {
		t.Errorf("Schedules.Delete returned error: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Schedules.Delete returned status %+v, want %+v", resp.StatusCode, http.StatusNoContent)
	}
}
