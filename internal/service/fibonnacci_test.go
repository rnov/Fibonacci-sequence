package service

import (
	"testing"
)

type testFib struct {
	name          string
	expectedValue uint32
}

func TestFib_GetNextFib(t *testing.T) {
	tests := FibonacciUint32
	f := NewFibonacci()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			fn := f.GetNextFib()
			if fn != test.expectedValue {
				t.Fatalf("test %s failed, expected %d, got %d", test.name, test.expectedValue, fn)
			}

		})
	}
}

// max value of fibonacci sequence: 2971215073 for uint32 max value: 4294967295
var FibonacciUint32 = []testFib{
	{"fib-sec-1", 1},
	{"fib-sec-2", 1},
	{"fib-sec-3", 2},
	{"fib-sec-4", 3},
	{"fib-sec-5", 5},
	{"fib-sec-6", 8},
	{"fib-sec-7", 13},
	{"fib-sec-8", 21},
	{"fib-sec-9", 34},
	{"fib-sec-10", 55},
	{"fib-sec-11", 89},
	{"fib-sec-12", 144},
	{"fib-sec-13", 233},
	{"fib-sec-14", 377},
	{"fib-sec-15", 610},
	{"fib-sec-16", 987},
	{"fib-sec-17", 1597},
	{"fib-sec-18", 2584},
	{"fib-sec-19", 4181},
	{"fib-sec-20", 6765},
	{"fib-sec-21", 10946},
	{"fib-sec-22", 17711},
	{"fib-sec-23", 28657},
	{"fib-sec-24", 46368},
	{"fib-sec-25", 75025},
	{"fib-sec-26", 121393},
	{"fib-sec-27", 196418},
	{"fib-sec-28", 317811},
	{"fib-sec-29", 514229},
	{"fib-sec-30", 832040},
	{"fib-sec-31", 1346269},
	{"fib-sec-32", 2178309},
	{"fib-sec-33", 3524578},
	{"fib-sec-34", 5702887},
	{"fib-sec-35", 9227465},
	{"fib-sec-36", 14930352},
	{"fib-sec-37", 24157817},
	{"fib-sec-38", 39088169},
	{"fib-sec-39", 63245986},
	{"fib-sec-40", 102334155},
	{"fib-sec-41", 165580141},
	{"fib-sec-42", 267914296},
	{"fib-sec-43", 433494437},
	{"fib-sec-44", 701408733},
	{"fib-sec-45", 1134903170},
	{"fib-sec-46", 1836311903},
	{"fib-sec-47", 2971215073},
	{"fib-sec-50", 0}, // overflow value of uint32
}
