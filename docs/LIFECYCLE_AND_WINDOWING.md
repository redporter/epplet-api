# Core Lifecycle & Main Windowing Documentation

This document explains the architecture, execution lifecycle, window hierarchy, and context management in the Go `epplet-api` wrapper for Enlightenment 16 (e16) Epplets.

---

## 1. Application Lifecycle Overview

An Epplet application follows a deterministic lifecycle managed by Enlightenment 16:

```
+-------------------------------------------------------------+
| 1. epplet.Init(name, version, info, w, h, vertical)         |
|    - Connects to X11 display                                |
|    - Allocates main window (grid w * h in 16px tile units)  |
+-------------------------------------------------------------+
                              |
                              v
+-------------------------------------------------------------+
| 2. Widget & Event Registration                              |
|    - epplet.CreateButton / CreateTextbox / CreateHBar       |
|    - Register event handlers & timers                       |
+-------------------------------------------------------------+
                              |
                              v
+-------------------------------------------------------------+
| 3. epplet.Show()                                            |
|    - Maps and draws the epplet main window on desktop       |
+-------------------------------------------------------------+
                              |
                              v
+-------------------------------------------------------------+
| 4. epplet.Loop()  (Blocking Event Loop)                     |
|    - Dispatches X11 events, gadget clicks, timers           |
+-------------------------------------------------------------+
                              |
                              v
+-------------------------------------------------------------+
| 5. epplet.Cleanup()                                         |
|    - Flushes config to disk, releases X11 resources         |
+-------------------------------------------------------------+
```

---

## 2. Core Lifecycle Functions

### `epplet.Init`
```go
func Init(name, version, info string, w, h int, vertical bool)
```
- **Description**: Initializes the epplet with the Enlightenment 16 desktop manager.
- **Parameters**:
  - `name`: Identifier string for the epplet (e.g. `"E-Clock"`).
  - `version`: Version string (e.g. `"1.0"`).
  - `info`: Descriptive summary string displayed in standard About dialogs.
  - `w`, `h`: Tile dimensions in grid units. Each unit is **16 pixels** (e.g. `4, 4` creates a 64x64 pixel main window).
  - `vertical`: `false` for horizontal layout alignment, `true` for vertical.

---

### `epplet.Show`
```go
func Show()
```
- **Description**: Maps the main Epplet window on screen and renders its visual state.

---

### `epplet.Loop`
```go
func Loop()
```
- **Description**: Enters the main X11 event dispatch loop. This function blocks until the application terminates.

---

### `epplet.Cleanup`
```go
func Cleanup()
```
- **Description**: Performs graceful shutdown. Flushes configuration key-value pairs to disk and frees X11 display allocations.

---

### `epplet.Remember` / `epplet.Unremember`
```go
func Remember()
func Unremember()
```
- **Description**: Instructs Enlightenment 16 to save (`Remember()`) or forget (`Unremember()`) the epplet's screen position, desktop workspace, and layer state.

---

## 3. Window Management & Context Stack

### Main Window Access
```go
func GetMainWindow() Window
```
- **Description**: Returns the X11 `Window` handle of the primary Epplet container window.

### Display Connection
```go
func GetDisplay() *Display
```
- **Description**: Returns the X11 `Display` handle connected to the X server.

---

### Auxiliary Window Creation

Epplets can create auxiliary popups, sub-windows, or dialog containers:

1. **Standard Auxiliary Window**:
   ```go
   win := epplet.CreateWindow(w, h int, title string, vertical bool)
   ```
2. **Borderless Auxiliary Window**:
   ```go
   win := epplet.CreateWindowBorderless(w, h int, title string, vertical bool)
   ```
3. **Configuration Sub-Window**:
   ```go
   win := epplet.CreateWindowConfig(w, h int, title string, okCb, applyCb, cancelCb WindowConfigCallback)
   ```

---

### Window Context Stack

To route rendering operations (`DrawLine`, `ImageclassApply`, `PasteImage`, etc.) to a specific window, use the window context stack:

```go
// Push window onto drawing context stack
win.PushContext()

// Perform rendering operations targetting win
win.DrawBox(4, 4, 32, 32, 255, 0, 0)

// Restore previous window context
epplet.WindowPopContext()
```

---

## 4. Code Example: Complete Minimal Epplet

```go
package main

import (
	"fmt"
	"time"

	"github.com/redporter/epplet-api"
)

func main() {
	// 1. Initialize (4x4 tiles = 64x64 pixels)
	epplet.Init("SimpleApp", "1.0", "Minimal Epplet Example", 4, 4, false)

	// 2. Create UI Gadgets
	btn := epplet.CreateTextButton("Click", 4, 4, 56, 20, func() {
		fmt.Println("Button clicked!")
	})
	btn.Show()

	// 3. Register Recurring Timer (reschedule in callback)
	var tick epplet.TimerCallback
	tick = func() {
		fmt.Println("Tick!")
		epplet.Timer(tick, 1*time.Second, "tick_timer")
	}
	epplet.Timer(tick, 1*time.Second, "tick_timer")

	// 4. Map and Display Window
	epplet.Show()

	// 5. Enter Event Dispatch Loop (blocking)
	epplet.Loop()
}
```
