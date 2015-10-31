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

	expected := &Ping{
		Data: "Pong!",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("General.Ping returned %+v, want %+v", got, expected)
	}
}
