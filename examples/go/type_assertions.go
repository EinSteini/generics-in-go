//go:build ignore

package main

import "fmt"

// Go does not allow type assertions or type switches directly on type
// parameter values.  You must first convert to `any`, adding boilerplate

// Optional interfaces that items in a pipeline might implement.
type Validatable interface {
	Validate() error
}

type Loggable interface {
	LogEntry() string
}

// Process tries to validate and log a generic item.
// A type switch on the type-parameter value is illegal:
//
//	switch v := item.(type) {   // ERROR: cannot use type assertion on type parameter value
//	case Validatable: ...
//	case Loggable:    ...
//	}
//
// WORKAROUND: cast to `any` first, then use a type switch.
func Process[T any](item T) {
	// Type switch must go through `any`
	switch v := any(item).(type) {
	case Validatable:
		if err := v.Validate(); err != nil {
			fmt.Println("Validation failed:", err)
			return
		}
		fmt.Println("Validated OK")
	default:
		fmt.Println("No validation available")
	}

	// Same limitation for plain type assertions
	// entry, ok := item.(Loggable)       // ERROR
	entry, ok := any(item).(Loggable) // WORKAROUND
	if ok {
		fmt.Println("Log:", entry.LogEntry())
	} else {
		fmt.Println("Not loggable")
	}
}

// --- Test Implementations ---

// Order implements both Validatable and Loggable.
type Order struct {
	ID    int
	Total float64
}

func (o Order) Validate() error {
	if o.Total <= 0 {
		return fmt.Errorf("order %d has non-positive total", o.ID)
	}
	return nil
}

func (o Order) LogEntry() string {
	return fmt.Sprintf("Order#%d total=%.2f", o.ID, o.Total)
}

// RawEvent implements neither interface.
type RawEvent struct {
	Name string
}

func main() {
	fmt.Println("--- Order (implements both) ---")
	Process(Order{ID: 1, Total: 49.99})

	fmt.Println()
	fmt.Println("--- RawEvent (implements neither) ---")
	Process(RawEvent{Name: "click"})
}