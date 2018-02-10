package cachet

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGeneralService_Ping(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/ping", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data": "Pong!"}`)
	})

	got, _, err := testClient.General.Ping()
	if err != nil {
		t.Errorf("General.Ping returned error: %v", err)
	}

	expected := "Pong!"
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("General.Ping returned %+v, want %+v", got, expected)
	}
}

func TestGeneralService_Version(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/version", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"meta":{"on_latest":true,"latest":{"tag_name":"v2.4.0","prelease":false,"draft":false}},"data":"2.4.0"}`)
	})

	got, _, err := testClient.General.Version()
	if err != nil {
		t.Errorf("General.Version returned error: %v", err)
	}

	expected := &VersionResponse{
		Meta: MetaVersion{
			OnLatest: true,
			Latest: Latest{
				TagName:  "v2.4.0",
				Prelease: false,
				Draft:    false,
			},
		},
		Data: "2.4.0",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("General.Version returned %+v, want %+v", got, expected)
	}
}

func TestGeneralService_Status(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"status":"info","message":"Some systems are experiencing issues"}}`)
	})

	got, _, err := testClient.General.Status()
	if err != nil {
		t.Errorf("General.Status returned error: %v", err)
	}

	expected := &Status{
		Status:  "info",
		Message: "Some systems are experiencing issues",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("General.Status returned %+v, want %+v", got, expected)
	}
}
