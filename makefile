.PHONY: build run stop clean all current next prev

build:
	go build -o ./out/build fibonacci-sequence/cmd/fib-rest

run:
	./out/build & echo $$! > run.pid

stop:
	kill `cat run.pid` && rm -f run.pid

clean:
	rm -rf ./out

all: build run

current:
	curl --location "http://127.0.0.1:8080/current"

next:
	curl --location "http://127.0.0.1:8080/next"

prev:
	curl --location "http://127.0.0.1:8080/previous"
