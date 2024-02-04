.PHONY: build run stop clean all current next prev bench test

build:
	go build -o ./out/build ./cmd/app/

run:
	./out/build & echo $$! > run.pid

stop:
	kill `cat run.pid` && rm -f run.pid

clean:
	rm -rf ./out ./run.pid

all: build run

bench:
	go test -run=^$$ -bench=. ./... -v

test:
	go test ./... -v -cover

current:
	curl --location "http://127.0.0.1:8080/current" | jq

next:
	curl -X PUT --location "http://127.0.0.1:8080/next" \
	--header 'Content-Type: application/json' | jq

prev:
	curl --location "http://127.0.0.1:8080/previous" | jq
