package main

import "fmt"

// Define an interface
type Logger interface {
	LogValue()
	LogPtr()
}

type CustomInt int

// Method 1: Value Receiver
func (i CustomInt) LogValue() {
	fmt.Println("Logged by Value Receiver:", i)
}

// Method 2: Pointer Receiver
func (i *CustomInt) LogPtr() {
	fmt.Println("Logged by Pointer Receiver:", *i)
}

func main() {
	var v CustomInt = 10
	var p *CustomInt = &v

	// 1. Check method set of the VALUE type (CustomInt)
	fmt.Println("--- Testing Value Type (v) ---")
	v.LogValue() // OK: Value receiver method is in the set of CustomInt
	v.LogPtr()   // ERROR if uncommented: Pointer receiver method is NOT in the set of CustomInt

	// Can a VALUE type satisfy the interface?
	// var logger1 Logger = v // ERROR: CustomInt only has LogValue, not LogPtr.

	// 2. Check method set of the POINTER type (*CustomInt)
	fmt.Println("--- Testing Pointer Type (p) ---")
	p.LogValue() // OK: Value receiver methods are promoted for *CustomInt
	p.LogPtr()   // OK: Pointer receiver method is in the set of *CustomInt

	// Can a POINTER type satisfy the interface?
	var logger2 Logger = p // OK: *CustomInt has both LogValue (promoted) and LogPtr.
	logger2.LogPtr()
}
