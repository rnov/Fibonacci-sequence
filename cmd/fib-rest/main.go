package main

import (
	"fmt"
	"github.com/rnov/fibonacci-sequence/internal/handler"
	"github.com/rnov/fibonacci-sequence/internal/service"
	"log"
	"net/http"
)

func main() {
	fSrv := service.NewFibonacci()

	h := handler.NewHTTPHandler(fSrv)
	r := handler.RegisterNewRouter(h)

	fmt.Println("starting server")
	// Fire up the server ":8080"
	// todo panic recover
	log.Fatal(http.ListenAndServe(":8080", r))
}
