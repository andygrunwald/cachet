package cachet

const (
	// HTTP Basic Authentication
	authTypeBasic = 1
	// Auth by API token
	authTypeToken = 2
)

// AuthenticationService contains contains Authentication related functions.
type AuthenticationService struct {
	client   *Client
	username string
	secret   string
	authType int
}

// SetBasicAuth sets basic parameters for HTTP Basic auth.
// Attention! Basic Auth is no longer supported by Cachet since ~v2.1.
// You can only use this if you running an older version.
// Checkout https://github.com/CachetHQ/Cachet/issues/1658
func (s *AuthenticationService) SetBasicAuth(username, password string) {
	s.username = username
	s.secret = password
	s.authType = authTypeBasic
}

// SetTokenAuth sets the API token for "token auth"
func (s *AuthenticationService) SetTokenAuth(token string) {
	s.secret = token
	s.authType = authTypeToken
}

// HasAuth checks if an auth type is used
func (s *AuthenticationService) HasAuth() bool {
	return s.authType > 0
}

// HasBasicAuth checks if the auth type is HTTP Basic auth
func (s *AuthenticationService) HasBasicAuth() bool {
	return s.authType == authTypeBasic
}

// HasTokenAuth checks we auth with an API token against the API
func (s *AuthenticationService) HasTokenAuth() bool {
	return s.authType == authTypeToken
}
