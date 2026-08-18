# Auxiliary Windows & Context Stack Documentation

This document describes how to create auxiliary windows (popups, dialogs, configuration panels) and use the window context stack in the `epplet-api` Go library.

---

## 1. Overview & Window Types

By default, all UI gadgets (`CreateButton`, `CreateTextbox`, `CreateHBar`, etc.) are attached to the primary Epplet main window. However, Epplets can create auxiliary windows for settings, popups, or popout dialogs.

`epplet-api` provides three window constructor functions:

### 1. Standard Auxiliary Window
```go
win := epplet.CreateWindow(w, h int, title string, vertical bool)
```
- **Description**: Creates a standard auxiliary window with border decorations.
- **Parameters**:
  - `w`, `h`: Tile size in 16-pixel multiples (e.g., `4, 6` -> 64x96 pixels).
  - `title`: Window title string shown in window manager titlebars.
  - `vertical`: Orientation layout flag.

---

### 2. Borderless Auxiliary Window
```go
win := epplet.CreateWindowBorderless(w, h int, title string, vertical bool)
```
- **Description**: Creates an auxiliary window without window decorations, ideal for unbordered popups or tooltips.

---

### 3. Configuration Sub-Window
```go
win := epplet.CreateWindowConfig(w, h int, title string, okCb, applyCb, cancelCb WindowConfigCallback)
```
- **Description**: Creates a dedicated configuration window equipped with standard **OK**, **Apply**, and **Cancel** buttons at the bottom.
- **Parameters**:
  - `okCb`: Callback function executed when the user clicks OK (saves settings & closes window).
  - `applyCb`: Callback function executed when the user clicks Apply (saves settings without closing).
  - `cancelCb`: Callback function executed when the user clicks Cancel (reverts changes & closes window).

---

## 2. Window Methods

Each `Window` value is a lightweight type (`type Window uintptr`) supporting receiver methods:

| Method | Description |
| :--- | :--- |
| `(win Window).Show()` | Maps and displays the window on screen. |
| `(win Window).Hide()` | Hides (unmaps) the window. |
| `(win Window).Destroy()` | Destroys the window and automatically frees all attached gadgets. |
| `(win Window).Clear()` | Clears window graphics content. |
| `(win Window).ImageclassApply(iclass, state string)` | Sets an e16 imageclass as the background theme for the window. |
| `(win Window).ImageclassPaste(iclass, state string, x, y, h, w int)` | Pastes a themed imageclass element at specific coordinates. |
| `(win Window).TextclassDraw(tclass, state string, x, y int, txt string)` | Renders themed text using a textclass at (x, y). |
| `(win Window).PushContext()` | Pushes the window onto the active gadget creation stack. |

---

## 3. Window Context Stack Architecture

The **Window Context Stack** determines which window receives newly constructed gadgets.

```
Default State: Active Context -> Main Epplet Window

1. win.PushContext()
   +---------------------------------------+
   | Active Context -> Auxiliary Window    |
   +---------------------------------------+
   | - epplet.CreateTextButton(...)        |  <-- Attached to Auxiliary Window
   | - epplet.CreateTextbox(...)           |  <-- Attached to Auxiliary Window
   +---------------------------------------+

2. epplet.WindowPopContext()
   +---------------------------------------+
   | Active Context -> Main Epplet Window  |
   +---------------------------------------+
```

### Context Stack Workflow Rules

1. Call `win.PushContext()` or `epplet.WindowPushContext(win)` before creating gadgets destined for `win`.
2. Construct all desired gadgets (`CreateButton`, `CreateTextbox`, etc.).
3. Call `epplet.WindowPopContext()` to pop `win` and restore the previous window context.

---

## 4. Code Example: Configuration Sub-Window

```go
package main

import (
	"fmt"

	"github.com/redporter/epplet-api"
)

var configWin epplet.Window

func openConfigDialog() {
	if configWin != 0 {
		configWin.Show()
		return
	}

	// 1. Create a Configuration Sub-Window (4x6 tiles)
	configWin = epplet.CreateWindowConfig(
		4, 6, "Preferences",
		func() { // OK callback
			fmt.Println("Configuration Saved!")
			epplet.SaveConfig()
			configWin.Hide()
		},
		func() { // Apply callback
			fmt.Println("Configuration Applied!")
			epplet.SaveConfig()
		},
		func() { // Cancel callback
			fmt.Println("Configuration Cancelled")
			configWin.Hide()
		},
	)

	// 2. Push Window Context so gadgets attach to configWin
	configWin.PushContext()

	// 3. Create gadgets on configWin
	lbl := epplet.CreateLabel(4, 4, "Settings:", 10)
	lbl.Show()

	tb := epplet.CreateTextbox("", epplet.QueryConfigDef("user_name", "Guest"), 4, 20, 56, 18, 0, nil)
	tb.Show()

	// 4. Pop Window Context to restore Main Window as creation target
	epplet.WindowPopContext()

	// 5. Display the Configuration Window
	configWin.Show()
}

func main() {
	epplet.Init("ConfigDemo", "1.0", "Config Dialog Example", 4, 4, false)

	// Create Configure Button on Main Window
	btn := epplet.CreateTextButton("Config", 4, 4, 56, 20, openConfigDialog)
	btn.Show()

	epplet.Show()
	epplet.Loop()
}
```
