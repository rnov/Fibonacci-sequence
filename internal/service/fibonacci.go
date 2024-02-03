package service

import "sync"

// fib is the Fibonacci sequence
type fib struct {
	sync.RWMutex     // note: to protect the sequence from concurrent access
	prevFib, currFib uint32
}

// FibSequence is the interface for the Fibonacci sequence service
type FibSequence interface {
	GetNextFib() uint32
	GetCurrentFib() uint32
	GetPrevFib() uint32
}

// NewFibonacci returns a new Fibonacci sequence.
func NewFibonacci() FibSequence {
	// note: in order to meet the description of the task the current=0, prev=0
	return &fib{
		prevFib: 0,
		currFib: 0,
	}
}

// GetNextFib returns the next value in the Fibonacci sequence.
func (f *fib) GetNextFib() uint32 {
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

	// note: if the next value smaller than f.prevFib, the max value of uint32 has been exceeded(overflow), reset the sequence.
	// Although the func returns uint32 and in Fibonacci there's always going to have the same amount of numbers in the sequence for
	// uint32, the generic approach is better since it can be adapted to other units: int, unit64, etc. (bigInt is handled in different way).
	if nc < f.prevFib {
		f.currFib = 0
		f.prevFib = 0
	}

	return f.currFib
}

// GetCurrentFib returns the current value in the Fibonacci sequence. Blocking with read lock for concurrent access protection.
func (f *fib) GetCurrentFib() uint32 {
	f.RLock()
	defer f.RUnlock()
	return f.currFib
}

// GetPrevFib returns the previous value in the Fibonacci sequence. Blocking with read lock for concurrent access protection.
func (f *fib) GetPrevFib() uint32 {
	f.RLock()
	defer f.RUnlock()
	return f.prevFib
}
