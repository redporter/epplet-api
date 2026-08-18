# Event & Callback Handlers Documentation

This document describes how to handle native X11 window events, mouse interactions, keyboard navigation, and window lifecycle events using `epplet-api`.

---

## 1. Event Handling Architecture

Enlightenment 16 dispatches X11 events through `epplet.Loop()`. `epplet-api` bridges these events to Go callbacks using `cgo.Handle` gateways, ensuring memory safety and type safety.

- **Unregister Callbacks**: Each `Register*Handler` function returns a `func()` cleanup handle. Invoking this closure frees the underlying `cgo.Handle`.

---

## 2. Event Handler Inventory

### Rendering & Layout Events

1. **`epplet.RegisterExposeHandler`**:
   ```go
   type ExposeHandler func(win Window, x, y, w, h int)
   ```
   Invoked when a region of `win` becomes exposed or damaged and needs to be redrawn.

2. **`epplet.RegisterMoveResizeHandler`**:
   ```go
   type MoveResizeHandler func(win Window, x, y, w, h int)
   ```
   Invoked when `win` is moved or resized on the desktop.

---

### Mouse Events

1. **`epplet.RegisterButtonPressHandler`**:
   ```go
   type ButtonPressHandler func(win Window, x, y, button int)
   ```
   Invoked when a mouse button is pressed over `win`.
   - `button` values: `1` = Left, `2` = Middle, `3` = Right, `4` = Scroll Up, `5` = Scroll Down.

2. **`epplet.RegisterButtonReleaseHandler`**:
   ```go
   type ButtonReleaseHandler func(win Window, x, y, button int)
   ```
   Invoked when a mouse button is released.

3. **`epplet.RegisterMouseMotionHandler`**:
   ```go
   type MouseMotionHandler func(win Window, x, y int)
   ```
   Invoked continuously as the mouse pointer moves over `win`.

4. **`epplet.RegisterMouseEnterHandler`**:
   ```go
   type MouseEnterHandler func(win Window)
   ```
   Invoked when the mouse cursor enters the window boundary.

5. **`epplet.RegisterMouseLeaveHandler`**:
   ```go
   type MouseLeaveHandler func(win Window)
   ```
   Invoked when the mouse cursor leaves the window boundary.

---

### Keyboard Events

1. **`epplet.RegisterKeyPressHandler`**:
   ```go
   type KeyPressHandler func(win Window, key string)
   ```
   Invoked when a keyboard key is pressed while `win` has focus.
   - `key`: The key symbol string (e.g., `"Return"`, `"Escape"`, `"a"`).

2. **`epplet.RegisterKeyReleaseHandler`**:
   ```go
   type KeyReleaseHandler func(win Window, key string)
   ```
   Invoked when a keyboard key is released.

---

### Focus & Window Close Events

1. **`epplet.RegisterFocusInHandler`** / **`epplet.RegisterFocusOutHandler`**:
   ```go
   type FocusInHandler func(win Window)
   type FocusOutHandler func(win Window)
   ```
   Invoked when `win` receives or loses keyboard input focus.

2. **`epplet.RegisterDeleteEventHandler`**:
   ```go
   type DeleteEventHandler func(win Window) int
   ```
   Invoked when a window close request is sent to `win`. Return `1` to allow the window to be destroyed, or `0` to prevent destruction.

---

### Low-Level Raw X11 Events

```go
type XEventHandler func(ev *XEvent)
```

For advanced X11 event processing, `RegisterEventHandler` provides access to raw `XEvent` objects:

```go
epplet.RegisterEventHandler(func(ev *epplet.XEvent) {
    eventType := ev.Type()   // e.g. Expose, KeyPress, ClientMessage
    targetWin := ev.Window() // Target X11 Window ID
    rawPtr := ev.RawPointer() // *C.XEvent pointer for low-level Cgo
})
```

---

## 3. Code Example: Mouse Hover & Click Tracker

```go
package main

import (
	"fmt"

	"github.com/redporter/epplet-api"
)

func main() {
	epplet.Init("EventDemo", "1.0", "Event Handling Example", 4, 4, false)

	// Register Mouse Enter/Leave handlers
	epplet.RegisterMouseEnterHandler(func(win epplet.Window) {
		fmt.Println("Mouse entered epplet window!")
	})

	epplet.RegisterMouseLeaveHandler(func(win epplet.Window) {
		fmt.Println("Mouse left epplet window!")
	})

	// Register Button Press handler
	epplet.RegisterButtonPressHandler(func(win epplet.Window, x, y, btn int) {
		fmt.Printf("Mouse button %d pressed at (%d, %d)\n", btn, x, y)
	})

	// Register Key Press handler
	epplet.RegisterKeyPressHandler(func(win epplet.Window, key string) {
		fmt.Printf("Key pressed: %s\n", key)
	})

	epplet.Show()
	epplet.Loop()
}
```
