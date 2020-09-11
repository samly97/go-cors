package cors

import (
	"net/http"
	"strings"
)

// The CORS middleware will apply CORS headers to requests which require
// CORS. Headers are attached to the preflight requests and to the actual
// requests as well.
//
// The user would have to specify the headers that are applied in the
// 'New' initialization function. The user could find a list of functional
// options available from functions that returns the 'corsOption' function
// type.
//
// Note: for requests that require user authentication, the CORS
// middleware may need to be wrapped around the users middlware first and
// not the other way around.
//
// This is because CORS need to send preflights before the actual request
// goes through. So if your user authentication needs to pass cookies
// for authentication, then they aren't attached during the preflight
// and depending on how your user auth middleware is implemented (early
// exit if nil cookie), then we may not even reach the CORS implementation
// at all.
//
// E.g.:
// Do this...
//	corsMw.ApplyFn(userAuthMw.ApplyFn(...))
// Not this...
//	userAuthMw.ApplyFn(corsMw.ApplyFn(...))
type CORS struct {
	Headers map[string]string
}

// New initializes a CORS object with the options specified by the
// variadic corsOptions function.
func New(opts ...corsOption) CORS {
	cors := CORS{Options: make(map[string]string)}
	for _, opt := range opts {
		opt(&cors)
	}
	return cors
}

func (c *CORS) Apply(next http.Handler) http.HandlerFunc {
	return c.ApplyFn(next.ServeHTTP)
}

func (c *CORS) ApplyFn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			// Set Prelight headers here
			w.Header().Set("Access-Control-Allow-Origin",
				"http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
			w.Header().Set("Access-Control-Allow-Headers",
				"Content-Type, Origin")
			w.Header().Set("Access-Control-Max-Age", "300")
			return
		} else {
			w.Header().Set("Access-Control-Allow-Origin",
				"http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
			w.Header().Set("Access-Control-Allow-Headers",
				"Content-Type, Origin")
			w.Header().Set("Access-Control-Max-Age", "300")
			// Forward to non-simple request
			next(w, r)
		}
	}
}

// corsOptions is a function type used in initializing a CORS struct via
// the 'New' function. Options are added inline as pointer to the object
// is passed.
//
// Order of initialization shouldn't matter.
type corsOption func(*CORS)

// AllowOrigins whitelists different origins (sites: A, B, ..., Y)
// to access resources on site Z.
func AllowOrigins(origins []string) corsOption {
	return func(cors *CORS) {
		s := strings.Join(origins, ", ")
		cors.Options["Access-Control-Allow-Origin"] = s
	}
}

// AllowCredentials will allow authentication from different sites to
// host. For instance, this includes site cookies.
func AllowCredentials(allow bool) corsOption {
	cors.Options["Allow-Control-Allow-Credentials"] = string(allow)
}

// AllowMethods will allow the HTTP methods from different sites to host.
// E.g.
//	cors.AllowMethods([]string{"GET", "POST"}))
//	GET and POST methods allowed on host
func AllowMethods(methods []string) corsOption {
	s := strings.Join(methods, ", ")
	cors.Options["Allow-Control-Allow-Methods"] = s
}

// AllowHeaders will whitelist the specified non-simple header types.
// Simple header example:
//	Content-Type: text/plain
// Non-simple header example:
//	Content-Type: application/json
// Checkout Mozilla's documentation for more information on non-simple
// headers: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
func AllowHeaders(headers []string) corsOption {
	s := strings.Join(headers, ", ")
	cors.Options["Access-Control-Allow-Headers"] = s
}
