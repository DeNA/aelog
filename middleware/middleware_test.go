package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAELogger(t *testing.T) {
	tests := []struct {
		name  string
		label string
		want  string
	}{
		{name: "empty label", label: "", want: "label is required"},
		{name: "set spancontext", label: "hoge", want: ""},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				rcv := recover()
				if rcv != nil && rcv != tt.want {
					t.Fatalf("want recover result is %s, but got recover result is %v", tt.want, rcv)
				}
			}()

			mux := AELogger(tt.label)(http.NewServeMux())
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)
		})
	}

}
