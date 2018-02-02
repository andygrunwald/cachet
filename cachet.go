package cachet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// A Client manages communication with the Cachet API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	// BaseURL should always be specified with a trailing slash.
	baseURL *url.URL

	// Cachet service for authentication
	Authentication *AuthenticationService

	// Services used for talking to different parts of the Cachet API.
	General         *GeneralService
	Components      *ComponentsService
	Incidents       *IncidentsService
	IncidentUpdates *IncidentUpdatesService
	Metrics         *MetricsService
	Subscribers     *SubscribersService
}

// Response is a Cachet API response.
// This wraps the standard http.Response returned from Cachet.
type Response struct {
	*http.Response
}

// NewClient returns a new Cachet API client.
// instance has to be the HTTP endpoint of the Cachet instance.
// If a nil httpClient is provided, http.DefaultClient will be used.
func NewClient(instance string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if len(instance) == 0 {
		return nil, fmt.Errorf("No Cachet instance given.")
	}
	baseURL, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:  httpClient,
		baseURL: baseURL,
	}
	c.Authentication = &AuthenticationService{client: c}
	c.General = &GeneralService{client: c}
	c.Components = &ComponentsService{client: c}
	c.Incidents = &IncidentsService{client: c}
	c.IncidentUpdates = &IncidentUpdatesService{client: c}
	c.Metrics = &MetricsService{client: c}
	c.Subscribers = &SubscribersService{client: c}

	return c, nil
}

// NewRequest creates an API request.
// A relative URL can be provided in urlStr, in which case it is resolved relative to the baseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.buildURLForRequest(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}

	// Apply Authentication
	// Docs: https://docs.cachethq.io/docs/api-authentication
	if c.Authentication.HasAuth() {
		c.addAuthentication(req)

		// If we fire requests that requires an authentication
		// we (mostly) pass data with the request.
		// We can do this by POST (see https://docs.cachethq.io/docs/post-parameters)
		// but we do this by JSON body content.
		// So we add the correct content type.
		req.Header.Add("Content-Type", "application/json")
	}

	// Just to be sure.
	// Maybe the API will accept other applications in future.
	// Who knows.
	req.Header.Add("Accept", "application/json")

	return req, nil
}

// addAuthentication adds necessary authentication.
//
// Docs: https://docs.cachethq.io/docs/api-authentication
func (c *Client) addAuthentication(req *http.Request) {
	// Apply HTTP Basic Authentication
	if c.Authentication.HasBasicAuth() {
		req.SetBasicAuth(c.Authentication.username, c.Authentication.secret)

		// Apply auth via Token
	} else if c.Authentication.HasTokenAuth() {
		req.Header.Add("X-Cachet-Token", c.Authentication.secret)
	}
}

// Call is a combine function for Client.NewRequest and Client.Do.
//
// Most API methods are quite the same.
// Get the URL, apply options, make a request, and get the response.
// Without adding special headers or something.
// To avoid a big amount of code duplication you can Client.Call.
//
// method is the HTTP method you want to call.
// u is the URL you want to call.
// body is the HTTP body.
// v is the HTTP response.
//
// For more information read https://github.com/google/go-github/issues/234
func (c *Client) Call(method, u string, body interface{}, v interface{}) (*Response, error) {
	req, err := c.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, v)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// buildURLForRequest will build the URL (as string) that will be called.
// It does several cleaning tasks for us.
func (c *Client) buildURLForRequest(urlStr string) (string, error) {
	u := c.baseURL.String()

	// If there is no / at the end, add one
	if strings.HasSuffix(u, "/") == false {
		u += "/"
	}

	// If there is a "/" at the start, remove it
	if strings.HasPrefix(urlStr, "/") == true {
		urlStr = urlStr[1:]
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	u += rel.String()

	return u, nil
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v,
// or returned as an error if an API error has occurred.
// If v implements the io.Writer interface, the raw response body will be written to v,
// without attempting to first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Wrap response
	response := &Response{Response: resp}

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			var body []byte
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				// even though there was an error, we still return the response
				// in case the caller wants to inspect it further
				return response, err
			}
			err = json.Unmarshal(body, v)
		}
	}
	return response, err
}

// CheckResponse checks the API response for errors, and returns them if present.
// A response is considered an error if it has a status code outside the 200 range.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	return fmt.Errorf("API call to %s failed: %s", r.Request.URL.String(), r.Status)
}
