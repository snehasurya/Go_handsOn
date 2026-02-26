package main

import (
	"fmt"
)

/*
PROBLEM: Priority Task Scheduler with Cooldown

You are building a task scheduler.

Each task has:
- ID (string)
- Priority (int)       // higher value = higher priority
- Timestamp (int)      // arrival time in seconds

Rules:
1. Only ONE task can execute at a time
2. At any time, execute the highest-priority AVAILABLE task
3. Same task ID cannot be executed again until `cooldown` seconds have passed
4. If no task is eligible, the scheduler is idle
5. Tasks arrive in non-decreasing timestamp order

TASK:
Implement the function `ScheduleTasks` that returns
the order of executed task IDs.

You are expected to use:
- Priority Queue (heap)
- Time-based scheduling
- Cooldown handling

Example:
Cooldown = 3
Tasks:
A(3,0), A(3,1), B(2,1), C(1,2), A(3,4)

Output:
[A, B, C, A]
*/

type Task struct {
	ID        string
	Priority  int
	Timestamp int
}

// ScheduleTasks schedules tasks based on priority and cooldown.
// Return the execution order of task IDs.
func ScheduleTasks(tasks []Task, cooldown int) []string {
	// TODO:
	// 1. Maintain a max-heap for ready tasks (by priority)
	// 2. Maintain a min-heap or queue for cooling tasks (by next available time)
	// 3. Track current time
	// 4. Move tasks between heaps as time progresses
	// 5. Append executed task IDs to result

	return nil
}

func main() {
	tasks := []Task{
		{ID: "A", Priority: 3, Timestamp: 0},
		{ID: "A", Priority: 3, Timestamp: 1},
		{ID: "B", Priority: 2, Timestamp: 1},
		{ID: "C", Priority: 1, Timestamp: 2},
		{ID: "A", Priority: 3, Timestamp: 4},
	}

	result := ScheduleTasks(tasks, 3)
	fmt.Println(result)
}
