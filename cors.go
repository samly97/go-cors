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

//
type corsOption func(*CORS)

func AllowOrigins(origins []string) corsOption {
	return func(cors *CORS) {
		s := strings.Join(origins, ", ")
		cors.Options["Access-Control-Allow-Origin"] = s
	}
}

// 	AllowCredentials: true,
// 	AllowedMethods:   []string{"GET", "POST"},
// 	AllowedHeaders: []string{"Origin", "Content-Type",
