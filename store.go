//go:build go1.11
// +build go1.11

package csrf

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

// store represents the session storage used for CSRF tokens.
type store interface {
	// Get returns the real CSRF token from the store.
	Get(*http.Request) ([]byte, error)
	// Save stores the real CSRF token in the store and writes a
	// cookie to the http.ResponseWriter.
	// For non-cookie stores, the cookie should contain a unique (256 bit) ID
	// or key that references the token in the backend store.
	// csrf.GenerateRandomBytes is a helper function for generating secure IDs.
	Save(token []byte, w http.ResponseWriter) error
}

// cookieStore is a signed cookie session store for CSRF tokens.
type cookieStore struct {
	name     string
	maxAge   int
	secure   bool
	httpOnly bool
	path     string
	domain   string
	sc       *securecookie.SecureCookie
	sameSite SameSiteMode
}

// Get retrieves a CSRF token from the session cookie. It returns an empty token
// if decoding fails (e.g. HMAC validation fails or the named cookie doesn't exist).
func (cs *cookieStore) Get(r *http.Request) ([]byte, error) {
	_ = "STUB: not implemented"
	// Retrieve the cookie from the request
	return nil, nil
}

// Decode the HMAC authenticated cookie.

// Save stores the CSRF token in the session cookie.
func (cs *cookieStore) Save(token []byte, w http.ResponseWriter) error {
	_ = "STUB: not implemented"
	// Generate an encoded cookie value with the CSRF token.
	return nil
}

// Set the Expires field on the cookie based on the MaxAge
// If MaxAge <= 0, we don't set the Expires attribute, making the cookie
// session-only.

// Write the authenticated cookie to the response.
