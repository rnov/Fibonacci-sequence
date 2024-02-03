package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"

	"github.com/rnov/fibonacci-sequence/internal/service"
)

type HTTPHandler struct {
	fSrv service.FibSequence
}

func NewHTTPHandler(f service.FibSequence) *HTTPHandler {
	return &HTTPHandler{
		fSrv: f,
	}
}

// RegisterNewRouter registers the routes for the fibonacci sequence into gorilla mux router adding a panic recovery middleware.
func RegisterNewRouter(h *HTTPHandler) *mux.Router {
	r := mux.NewRouter()

	// Wrap the router itself with the recoverFromPanic middleware. Affects all routes.
	r.Use(recoverFromPanic)

	r.HandleFunc("/current", h.CurrentValue).Methods("GET")
	r.HandleFunc("/next", h.NextValue).Methods("PUT")
	r.HandleFunc("/previous", h.PreviousValue).Methods("GET")

	return r
}

func (h *HTTPHandler) CurrentValue(w http.ResponseWriter, _ *http.Request) {

	cVal := h.fSrv.GetCurrentFib()

	rBody := &Response{
		Value: cVal,
	}
	body, jsonErr := json.Marshal(rBody)
	if jsonErr != nil {
		// note should log error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// note: for simplicity, we are not handling the error here, but we should log it. On another note at this point
	// there's little to be done, since the code response is 200, and the error is in the response body.
	_, _ = w.Write(body)
}

func (h *HTTPHandler) NextValue(w http.ResponseWriter, _ *http.Request) {

	nVal := h.fSrv.GetNextFib()

	rBody := &Response{
		Value: nVal,
	}
	body, jsonErr := json.Marshal(rBody)
	if jsonErr != nil {
		// note should log error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(body)
}

func (h *HTTPHandler) PreviousValue(w http.ResponseWriter, _ *http.Request) {

	pVal := h.fSrv.GetPrevFib()

	rBody := &Response{
		Value: pVal,
	}
	body, jsonErr := json.Marshal(rBody)
	if jsonErr != nil {
		// note should log error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(body)
}

// Note: Usually the middleware would be in a separate package (e.g:/internal/http/middleware), but for simplicity, we are keeping it here.

// recoverFromPanic is a middleware that recovers from panics, logs the error and stack trace,
// and returns a 500 Internal Server Error response.
func recoverFromPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v\n", err)
				log.Println("Stack trace from panic: \n" + string(debug.Stack()))
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

//note: Due the simplicity of the calls, there is no need to have a separate struct for all the responses, neither to be placed in a separate file/package

// Response is the response struct for all the calls in fib sequence http handler.
type Response struct {
	Value uint32 `json:"value"`
}
