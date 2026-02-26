package main

import (
	"container/heap"
	"errors"
	"fmt"
	"sync"
)

// --- 1. Min Heap Implementation ---

// MinIntHeap is a Min-Heap of integers (slot numbers)
type MinIntHeap []int

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinIntHeap) Push(x any) {
	// Push and Pop are implemented for heap.Interface
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// --- 2. ParkingLot Structure ---

type ParkingLot struct {
	capacity       int
	occupied       map[int]string // Key: SlotNumber, Value: LicensePlate
	availableSlots MinIntHeap     // Min-Heap of available slot numbers
	mu             sync.Mutex
}

// NewParkingLot initializes the lot with a fixed capacity
func NewParkingLot(cap int) *ParkingLot {
	// Initialize the heap with all slot numbers [1, 2, ..., cap]
	initialSlots := make(MinIntHeap, cap)
	for i := 0; i < cap; i++ {
		initialSlots[i] = i + 1
	}
	heap.Init(&initialSlots) // Initialize the heap structure

	return &ParkingLot{
		capacity:       cap,
		occupied:       make(map[int]string),
		availableSlots: initialSlots,
	}
}

// Park assigns the smallest available slot (O(log N))
func (p *ParkingLot) Park(licensePlate string) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Check if the lot is full (heap is empty)
	if p.availableSlots.Len() == 0 {
		return 0, errors.New("parking lot is full")
	}

	// Pop the smallest available slot from the heap (O(log N))
	slotNum := heap.Pop(&p.availableSlots).(int)

	// Assign the slot and update the occupied map (O(1))
	p.occupied[slotNum] = licensePlate
	return slotNum, nil
}

// Unpark frees a slot and returns it to the Min-Heap (O(log N))
func (p *ParkingLot) Unpark(slotNumber int) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// 1. Validation check
	if _, exists := p.occupied[slotNumber]; !exists {
		return fmt.Errorf("slot %d is already empty or invalid", slotNumber)
	}

	// 2. Free the slot from the occupied map (O(1))
	delete(p.occupied, slotNumber)

	// 3. Push the slot number back into the availableSlots heap (O(log N))
	heap.Push(&p.availableSlots, slotNumber)

	return nil
}

// --- 3. Example Usage ---

func main() {
	fmt.Println("--- Parking Lot with Min-Heap ---")
	lot := NewParkingLot(5) // Capacity 5

	// 1. Park: Takes the smallest available slot (1, 2, 3)
	s1, _ := lot.Park("A1")
	s2, _ := lot.Park("B2")
	s3, _ := lot.Park("C3")
	fmt.Printf("Parked A1 in slot %d\n", s1)
	fmt.Printf("Parked B2 in slot %d\n", s2)
	fmt.Printf("Parked C3 in slot %d\n", s3)

	// 2. Unpark: Frees slot 1 and slot 3
	lot.Unpark(s1)
	lot.Unpark(s3)
	fmt.Printf("Unparked slots %d and %d. Available slots in heap: %v\n", s1, s3, lot.availableSlots)
	// The heap now contains {1, 3, 4, 5} in some order, with 1 at the root.

	// 3. Park again: The heap must return the smallest available slot, which is 1
	s4, _ := lot.Park("D4")
	fmt.Printf("Parked D4 in slot %d (reused smallest available)\n", s4) // Output: slot 1

	// 4. Park again: The heap returns the next smallest, which is 3
	s5, _ := lot.Park("E5")
	fmt.Printf("Parked E5 in slot %d\n", s5) // Output: slot 3
}
