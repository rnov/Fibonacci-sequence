package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rnov/fibonacci-sequence/internal/handler"
	"github.com/rnov/fibonacci-sequence/internal/service"
)

func main() {
	fSrv := service.NewFibonacci()

	h := handler.NewHTTPHandler(fSrv)
	r := handler.RegisterNewRouter(h)

	fmt.Println("starting server")

	// Fire up the server ":8080"
	log.Fatal(http.ListenAndServe(":8080", r))
}
