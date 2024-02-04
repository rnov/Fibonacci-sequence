package service

import "sync"

// fib is the Fibonacci sequence struct that holds the business logic for the service.
type fib struct {
	sync.RWMutex     // note: to protect the sequence from concurrent access
	prevFib, currFib uint32
}

// Fibonacci is the interface for the Fibonacci sequence service.
type Fibonacci interface {
	Next() uint32
	Current() uint32
	Prev() uint32
}

// NewFibonacci returns an initialized Fibonacci sequence.
func NewFibonacci() Fibonacci {
	// note: in order to meet the task's description the start sec is: current=0, prev=0
	return &fib{
		prevFib: 0,
		currFib: 0,
	}
}

// Next returns the next value in the Fibonacci sequence, in case of overflow, resets to zero.
func (f *fib) Next() uint32 {
	f.Lock()
	defer f.Unlock()
	// note: start the sequence at 0 to 1.
	if f.currFib == 0 {
		f.currFib = 1
		return f.currFib
	}

	nc := f.prevFib + f.currFib
	f.prevFib = f.currFib
	f.currFib = nc

	// note: if the next value is smaller than f.prevFib, the max value of uint32 has been exceeded(overflow), reset the sequence.
	// Although the func returns uint32 and in Fibonacci there's always going to have the same amount of numbers in the sequence for
	// uint32, the generic approach is better since it can be adapted to other units: int, unit64, etc. (bigInt is handled in different way).
	if nc < f.prevFib {
		f.currFib = 0
		f.prevFib = 0
	}

	return f.currFib
}

// Current returns the current value in the Fibonacci sequence. Blocking with read lock for concurrent access protection.
func (f *fib) Current() uint32 {
	f.RLock()
	defer f.RUnlock()
	return f.currFib
}

// Prev returns the previous value in the Fibonacci sequence. Blocking with read lock for concurrent access protection.
func (f *fib) Prev() uint32 {
	f.RLock()
	defer f.RUnlock()
	return f.prevFib
}
