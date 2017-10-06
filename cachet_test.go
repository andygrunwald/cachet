package cachet

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"
)

const (
	// testCachetInstance is a test instance url that won`t be called
	testCachetInstance = "https://demo.cachethq.io/"
)

var (
	// testMux is the HTTP request multiplexer used with the test server.
	testMux *http.ServeMux

	// testClient is the cachet client being tested.
	testClient *Client

	// testServer is a test HTTP server used to provide mock API responses.
	testServer *httptest.Server
)

//type testValues map[string]string

// setup sets up a test HTTP server along with a cachet.Client that is configured to talk to that test server.
// Tests should register handlers on mux which provide mock responses for the API method being tested.
func setup() {
	// Test server
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	// Cachet client configured to use test server
	testClient, _ = NewClient(testServer.URL, nil)
}

// teardown closes the test HTTP server.
func teardown() {
	testServer.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestNewClient_NoCachetInstance(t *testing.T) {
	mockData := []string{"", "://not-existing"}
	for _, data := range mockData {
		c, err := NewClient(data, nil)
		if c != nil {
			t.Errorf("NewClient return is not nil. Expected no client. Got %+v", c)
		}
		if err == nil {
			t.Error("No error occured by empty Cachet Instance. Expected one.")
		}
	}
}

func TestNewClient_HttpClient(t *testing.T) {
	customHTTPClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	mockData := []struct {
		HTTPClient         *http.Client
		ExpectedHTTPClient *http.Client
	}{
		{nil, http.DefaultClient},
		{customHTTPClient, customHTTPClient},
	}

	for _, mock := range mockData {
		c, err := NewClient(testCachetInstance, mock.HTTPClient)
		if err != nil {
			t.Errorf("An error occured. Expected nil. Got %+v.", err)
		}
		if reflect.DeepEqual(c.client, mock.ExpectedHTTPClient) == false {
			t.Errorf("Wrong HTTP Client. Expected %+v. Got %+v", mock.ExpectedHTTPClient, c.client)
		}
	}
}

func TestNewClient_Services(t *testing.T) {
	c, err := NewClient(testCachetInstance, nil)
	if err != nil {
		t.Errorf("An error occured. Expected nil. Got %+v.", err)
	}

	if c.Authentication == nil {
		t.Error("No AuthenticationService found.")
	}
	if c.General == nil {
		t.Error("No GeneralService found.")
	}
	if c.Components == nil {
		t.Error("No ComponentsService found.")
	}
	if c.Incidents == nil {
		t.Error("No IncidentsService found.")
	}
	if c.Metrics == nil {
		t.Error("No MetricsService found.")
	}
	if c.Subscribers == nil {
		t.Error("No SubscribersService found.")
	}

}

func TestNewRequest(t *testing.T) {
	c, err := NewClient(testCachetInstance, nil)
	if err != nil {
		t.Errorf("An error occured. Expected nil. Got %+v.", err)
	}

	inURL, outURL := "/foo", testCachetInstance+"foo"
	inBody, outBody := &Component{Name: "Go API (by Token)", Status: 1}, `{"name":"Go API (by Token)","status":1,"enabled":false}`+"\n"
	req, _ := c.NewRequest("POST", inURL, inBody)

	// Test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// Test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%+v) Body is %v, want %v", inBody, got, want)
	}
}

func TestNewRequest_BadURL(t *testing.T) {
	c, err := NewClient(testCachetInstance, nil)
	if err != nil {
		t.Errorf("An error occured. Expected nil. Got %+v.", err)
	}
	_, err = c.NewRequest("GET", ":", nil)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

// If a nil body is passed to cachet.NewRequest, make sure that nil is also passed to http.NewRequest.
// In most cases, passing an io.Reader that returns no content is fine,
// since there is no difference between an HTTP request body that is an empty string versus one that is not set at all.
// However in certain cases, intermediate systems may treat these differently resulting in subtle errors.
func TestNewRequest_EmptyBody(t *testing.T) {
	c, err := NewClient(testCachetInstance, nil)
	if err != nil {
		t.Errorf("An error occured. Expected nil. Got %+v.", err)
	}
	req, err := c.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %+v", err)
	}
	if req.Body != nil {
		t.Fatalf("Constructed request contains a non-nil Body")
	}
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
	}

	testMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := testClient.NewRequest("GET", "/", nil)
	body := new(foo)
	testClient.Do(req, body)

	want := &foo{"a"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}

func TestDo_ioWriter(t *testing.T) {
	setup()
	defer teardown()
	content := `{"A":"a"}`

	testMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, content)
	})

	req, _ := testClient.NewRequest("GET", "/", nil)
	var buf []byte
	actual := bytes.NewBuffer(buf)
	testClient.Do(req, actual)

	expected := []byte(content)
	if !reflect.DeepEqual(actual.Bytes(), expected) {
		t.Errorf("Response body = %v, want %v", actual, string(expected))
	}
}

func TestDo_HTTPError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := testClient.NewRequest("GET", "/", nil)
	_, err := testClient.Do(req, nil)

	if err == nil {
		t.Error("Expected HTTP 400 error.")
	}
}

// Test handling of an error caused by the internal http client's Do() function.
// A redirect loop is pretty unlikely to occur within the Cacheterrit API, but does allow us to exercise the right code path.
func TestDo_RedirectLoop(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})

	req, _ := testClient.NewRequest("GET", "/", nil)
	_, err := testClient.Do(req, nil)

	if err == nil {
		t.Error("Expected error to be returned.")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected a URL error; got %#v.", err)
	}
}
