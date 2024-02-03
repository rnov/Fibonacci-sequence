package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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
			} else {
				t.Fatalf("test %s failed, expected %d, got empty body", test.name, test.expectedRes.Value)
			}
		})
	}
}

// BenchmarkFib_GetNextFib benchmarks overall performance of the GetNextFib endpoint in NS/op.
func BenchmarkFib_GetNextFib(b *testing.B) {
	f := service.NewFibonacci()
	h := NewHTTPHandler(f)

	servicesRouter := RegisterNewRouter(h)

	req, err := http.NewRequest("PUT", "/next", nil)
	if err != nil {
		b.Fatal(err)
	}

	rr := httptest.NewRecorder()

	b.ResetTimer() // Reset timer to exclude setup time
	for i := 0; i < b.N; i++ {
		servicesRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			b.Fatalf("benchmark failed, expected status code 200, got %d", rr.Code)
		}
		// Reset the response recorder for the next iteration
		rr = httptest.NewRecorder()
	}
}

// BenchmarkFib_GetNextFibIn1Sec benchmarks the amount of requests done in 1 second on the heaviest resource wise endpoint.
func BenchmarkFib_GetNextFibIn1Sec(b *testing.B) {
	f := service.NewFibonacci()
	h := NewHTTPHandler(f)

	servicesRouter := RegisterNewRouter(h)

	req, err := http.NewRequest("PUT", "/next", nil)
	if err != nil {
		b.Fatal(err)
	}

	rr := httptest.NewRecorder()

	var i int
	startTime := time.Now()
	for ; time.Since(startTime) < time.Second; i++ {
		servicesRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			b.Fatalf("benchmark failed, expected status code 200, got %d", rr.Code)
		}

		// Reset the response recorder for the next iteration
		rr = httptest.NewRecorder()
	}

	b.Logf("Amount of requests done in 1 second: %d", i)
}

// TestPanicRecoveryMiddleware tests that the recoverFromPanic middleware recovers from handler panics
func TestPanicRecoveryMiddleware(t *testing.T) {
	// Handler that intentionally panics
	panicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})

	// Wrap the panic handler with the recoverFromPanic middleware
	testHandler := recoverFromPanic(panicHandler)

	// Create an httptest server using the panic recovery wrapped handler
	server := httptest.NewServer(testHandler)
	defer server.Close()

	// Make a request to the test server
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to make request to test server: %v", err)
	}
	defer resp.Body.Close()

	// Check that the response code is 500 Internal Server Error
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code 500, got: %d", resp.StatusCode)
	}

}
