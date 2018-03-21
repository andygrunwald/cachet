package cachet

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSubscribersService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/subscribers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"meta":{"pagination":{"total":1,"count":1,"per_page":20,"current_page":1,"total_pages":1,"links":{"next_page":null,"previous_page":null}}},"data":[{"id":1,"email":"support@alt-three.com","verify_code":"1234567890","verified_at":"2015-07-24 14:42:24","created_at":"2015-07-24 14:42:24","updated_at":"2015-07-24 14:42:24"}]}`)
	})

	queryParams := &SubscribersQueryParams{}

	got, _, err := testClient.Subscribers.GetAll(queryParams)
	if err != nil {
		t.Errorf("Subscribers.GetAll returned error: %v", err)
	}

	expected := &SubscriberResponse{
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
		Subscribers: []Subscriber{
			{
				ID:         1,
				Email:      "support@alt-three.com",
				VerifyCode: "1234567890",
				VerifiedAt: "2015-07-24 14:42:24",
				CreatedAt:  "2015-07-24 14:42:24",
				UpdatedAt:  "2015-07-24 14:42:24",
			},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Subscribers.GetAll returned %+v, want %+v", got, expected)
	}
}

func TestSubscribersService_Create(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/subscribers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"data":{"id":1,"email":"support@alt-three.com","verify_code":"1234567890","verified_at":"2015-07-24 14:42:24","created_at":"2015-07-24 14:42:24","updated_at":"2015-07-24 14:42:24"}}`)
	})

	got, _, err := testClient.Subscribers.Create("support@alt-three.com", 1)
	if err != nil {
		t.Errorf("Subscribers.Create returned error: %v", err)
	}

	expected := &Subscriber{
		ID:         1,
		Email:      "support@alt-three.com",
		VerifyCode: "1234567890",
		VerifiedAt: "2015-07-24 14:42:24",
		CreatedAt:  "2015-07-24 14:42:24",
		UpdatedAt:  "2015-07-24 14:42:24",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Subscribers.Create returned %+v, want %+v", got, expected)
	}
}

func TestSubscribersService_Delete(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/subscribers/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := testClient.Subscribers.Delete(1)
	if err != nil {
		t.Errorf("Subscribers.Delete returned error: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Subscribers.Delete returned status %+v, want %+v", resp.StatusCode, http.StatusNoContent)
	}
}
