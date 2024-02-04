# Fibonacci-sequence


Made with :blue_heart: by rnov.

Fibonacci-sequence generates the fibonacci sequence and exposes a REST API to query the sequence, as described [here](https://gist.github.com/stepanbujnak/7fa18e2e97de2fd3f593c00b09c445c2).

### Quick Start

To build and start the service locally:
```sh
make all
```

To start making request to the service:

Get previous number in the sequence.
```sh
make prev
```

Get next number in the sequence.
```sh
make next
```

Get current number in the sequence.
```sh
make current
```

### Description

The project initialises a http server that exposes the following endpoints:

- `GET /current`: Returns the current number in the sequence.
- `PUT /next`: Returns the next number in the sequence.
- `GET /previous`: Returns the previous number in the sequence.


### Design

The project's design adheres to the [golang-standards/project-layout](https://github.com/golang-standards/project-layout):

- `cmd`: Contains the main applications for the project.
- `internal`: Houses all the logic intended for internal use. Notably:
    - `handler`: Contains the REST API handlers. 
    - `service`: Contains fibonacci service logic.

Due to the simplicity of the project some of its logic is contained in the same package.

It makes use of dependency injection (widely used in go) pattern through composition to make the code more testable and maintainable.

### Testing

Run tests:
```sh
make test
```
Run benchmarks:
```sh
make bench
```

```sh
# In case console output is not clear
go test -run=^$ -bench=. ./... -v
```

### Notes

- Have left more comments than usual to explain the code as `notes`.
- Due to project simplicity, some decisions were made to keep the project simple and easy such as:
   - Implemented a limited amount of tests mostly around `next` operation in service and handlers, seems an overkill to 
     explicitly have tests for read operations due to their simplicity in this case.
   - Loading from config file or containerization is missing.
   - The project layout is simple yet is based on well-known patterns and is open to extension (playing safe).
- Added benchmark tests to measure the performance of the service. `1082263 req/s` for the most expensive operation `next`.
  Exceeding by the task requirements of `1000 req/s`.
- Based on the implementation the only place where the service could panic is through incoming requests,
  That's why I've added a panic recovery middleware to handle any panic that could be caused.
  (Technically internal panic could occur in marshal when building the response which is done by us and unless a chan
   or func is passed is safe).
- The choice of using int32 is mostly for simplicity since is not stated in the task requirements.

### Makefile

Build the project:
```sh
make build
```
Run the project:
```sh
make run
```
Stop the service:
```sh
make stop
```
Remove bin directory:
```sh
make clean
```
