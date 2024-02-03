package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rnov/fibonacci-sequence/internal/service"
)

func TestFib_GetNextFib(t *testing.T) {
	f := service.NewFibonacci()
	h := NewHTTPHandler(f)

	tests := []struct {
		name        string
		expectedRes Response
	}{
		{name: "fib-sec-1", expectedRes: Response{Value: 1}},
		{name: "fib-sec-2", expectedRes: Response{Value: 1}},
		{name: "fib-sec-3", expectedRes: Response{Value: 2}},
	}

	servicesRouter := mux.NewRouter()
	servicesRouter.HandleFunc("/next", h.NextValue).Methods("PUT")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("PUT", "/next", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			servicesRouter.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Fatalf("test %s failed, expected status code 200, got %d", test.name, rr.Code)
			}
			if rr.Body.Len() > 0 {
				res := Response{}
				if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
					t.Fatal(err)
				}
				if res.Value != test.expectedRes.Value {
					t.Fatalf("test %s failed, expected %d, got %d", test.name, test.expectedRes.Value, res.Value)
				}
			}
		})
	}
}
