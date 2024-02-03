# Fibonacci-sequence


Made with :blue_heart: by rnov.

Fibonacci-sequence generates the fibonacci sequence and exposes a REST API to query the sequence, as described [here](https://gist.github.com/stepanbujnak/7fa18e2e97de2fd3f593c00b09c445c2).

### Quick Start

To build and start the service locally:
```sh
make all
```

To stop the service:
```sh
make stop
```

To remove bin directory:
```sh
make clean
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

The project's design largely adheres to
the [golang-standards/project-layout](https://github.com/golang-standards/project-layout):

- `cmd`: Contains the main applications for the project.
- `internal`: Houses all the logic intended for internal use. Notably:
    - `handler`: Contains the REST API handlers. 
    - `service`: Contains fibonacci service logic.

### Notes

### //todo
