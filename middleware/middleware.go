/*
Package middleware provides the http.Handler with spancontext.
*/
package middleware

import (
	"net/http"

	"github.com/DeNA/aelog/internal/spancontext"
)

// AELogger is middleware to set spancontext in the context.
// In this middleware, label is required.
// If you don't set label, this middleware panics.
func AELogger(label string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if label == "" {
				panic("label is required")
			}

			ctx, done := spancontext.Set(r, label)
			defer done()
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
