package cors

import (
	"net/http"
	"strings"
)

// c := cors.New(cors.Options{
// 	AllowCredentials: true,
// 	AllowedOrigins:   []string{"http://localhost:3000"},
// 	AllowedMethods:   []string{"GET", "POST"},
// 	AllowedHeaders: []string{"Origin", "Content-Type",
// 		"X-Auth-Token"},
// })

// CORS is
type CORS struct {
	Options map[string]string
}

func New(opts ...corsOption) CORS {
	cors := CORS{Options: make(map[string]string)}
	for _, opt := range opts {
		opt(&cors)
	}
	return cors
}

type corsOption func(*CORS)

func AllowOrigins(origins []string) corsOption {
	return func(cors *CORS) {
		s := strings.Join(origins, ", ")
		cors.Options["Access-Control-Allow-Origin"] = s
	}
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
