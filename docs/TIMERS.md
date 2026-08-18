# Timers Documentation

This document describes how to schedule, query, and cancel timers in the `epplet-api` Go library.

---

## 1. Overview

The `epplet-api` timer system integrates directly with the Enlightenment 16 main event loop (`epplet.Loop()`).

- **One-Shot Default**: In the e16 Epplet C API, timers are **one-shot** by default. Once a timer fires, its callback is invoked and the timer handle is destroyed.
- **Recurring Pattern**: To create a recurring timer (e.g. clock update, status ticker, animation frame), the callback function reschedules itself at the end of its execution block.

---

## 2. Timer Functions

### `epplet.Timer`
```go
func Timer(cb TimerCallback, d time.Duration, name string)
```
- **Description**: Registers a one-shot timer with a unique `name` that fires callback `cb` after duration `d`.
- **Parameters**:
  - `cb`: The `func()` signature callback.
  - `d`: Standard Go `time.Duration` (e.g., `1*time.Second`, `100*time.Millisecond`).
  - `name`: Unique identifier string for the timer. Re-registering with the same name replaces any previously active timer under that name.

---

### `epplet.RemoveTimer`
```go
func RemoveTimer(name string)
```
- **Description**: Cancels and removes an active timer matching `name`, releasing its underlying resources.

---

### `epplet.TimerGetData`
```go
func TimerGetData(name string) TimerCallback
```
- **Description**: Retrieves the active `TimerCallback` function associated with `name`, or `nil` if no active timer exists under that name.

---

## 3. Recurring Timer Design Pattern

Because Epplet timers fire once, use self-referential functions to implement recurring tickers:

```go
var clockTimer epplet.TimerCallback

clockTimer = func() {
    // 1. Perform tick operations
    updateClockDisplay()

    // 2. Reschedule timer for the next tick
    epplet.Timer(clockTimer, 1*time.Second, "clock_timer")
}

// Initial trigger
epplet.Timer(clockTimer, 1*time.Second, "clock_timer")
```

---

## 4. Code Examples

### Example A: Digital Clock Ticker
```go
package main

import (
	"fmt"
	"time"

	"github.com/redporter/epplet-api"
)

func main() {
	epplet.Init("ClockDemo", "1.0", "Digital Clock Timer Example", 4, 4, false)

	label := epplet.CreateLabel(4, 20, "00:00:00", 2)
	label.Show()

	var tick epplet.TimerCallback
	tick = func() {
		now := time.Now().Format("15:04:05")
		label.Change(now)

		// Reschedule timer
		epplet.Timer(tick, 1*time.Second, "clock_ticker")
	}

	// Start clock ticker
	epplet.Timer(tick, 1*time.Second, "clock_ticker")

	epplet.Show()
	epplet.Loop()
}
```

---

### Example B: One-Shot Delay with Cancellation

```go
package main

import (
	"fmt"
	"time"

	"github.com/redporter/epplet-api"
)

func main() {
	epplet.Init("TimerDemo", "1.0", "Timer Cancel Example", 4, 4, false)

	// Schedule a delayed message in 5 seconds
	epplet.Timer(func() {
		fmt.Println("Delayed action executed!")
	}, 5*time.Second, "delayed_task")

	// Button to cancel delayed message
	btnCancel := epplet.CreateTextButton("Cancel Task", 4, 20, 56, 20, func() {
		epplet.RemoveTimer("delayed_task")
		fmt.Println("Delayed task cancelled!")
	})
	btnCancel.Show()

	epplet.Show()
	epplet.Loop()
}
```
