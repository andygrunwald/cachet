package cachet

import (
	"net/http"
	"testing"
)

func TestSubscriptionsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/subscription/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := testClient.Subscriptions.Delete(1)
	if err != nil {
		t.Errorf("Subscriptions.Delete returned error: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Subscriptions.Delete returned status %+v, want %+v", resp.StatusCode, http.StatusNoContent)
	}
}
