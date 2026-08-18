# Imageclass & Textclass Rendering Documentation

This document describes how to use Enlightenment 16 (e16) theme definitions—**Imageclasses** and **Textclasses**—in the `epplet-api` Go library.

---

## 1. Overview

Enlightenment 16 uses a powerful theming engine built on top of Imlib2. Themes define visual styles as:

- **Imageclasses**: Defines multi-state imagery, border frames, bevels, and transparency masks for UI elements.
- **Textclasses**: Defines fonts, text colors, drop shadows, outlines, and text alignments across different interaction states.

---

## 2. Interactive States

Most imageclass and textclass functions accept a `state` string representing the interaction state:

| State String | Description |
| :--- | :--- |
| `"normal"` | Default idle visual state. |
| `"hilited"` | Mouse hover / highlighted state. |
| `"clicked"` | Active pressed / clicked state. |
| `"disabled"` | Inactive / disabled state. |

---

## 3. Imageclass Functions

### `(win Window).ImageclassApply`
```go
func (win Window) ImageclassApply(iclass, state string)
```
- **Description**: Sets the specified imageclass as the background theme for the window, applying transparency masks if defined by the theme.
- **Parameters**:
  - `iclass`: Name of the e16 imageclass (e.g. `"EPPLET_BACKGROUND"`, `"EPPLET_BUTTON"`).
  - `state`: The state string (e.g. `"normal"`).

---

### `(win Window).ImageclassPaste`
```go
func (win Window) ImageclassPaste(iclass, state string, x, y, h, w int)
```
- **Description**: Renders and pastes a scaled portion of an imageclass onto the window at position `(x, y)` with dimensions `(w, h)`.

---

### `epplet.ImageclassGetPixmaps`
```go
func ImageclassGetPixmaps(iclass, state string, width, height int) ImageclassPixmaps
```
- **Description**: Renders an imageclass at the specified pixel size and returns raw X11 `Pixmap` and `Mask` handles for custom X11 rendering operations.
- **Returns**: `ImageclassPixmaps` struct:
  ```go
  type ImageclassPixmaps struct {
      Pixmap C.Pixmap
      Mask   C.Pixmap
  }
  ```

---

## 4. Textclass Functions

### `(win Window).TextclassDraw`
```go
func (win Window) TextclassDraw(tclass, state string, x, y int, txt string)
```
- **Description**: Renders themed text onto the window at position `(x, y)` formatted according to the specified textclass and state.
- **Parameters**:
  - `tclass`: Name of the e16 textclass (e.g. `"EPPLET_LABEL"`, `"EPPLET_BUTTON"`).
  - `state`: Interaction state string (`"normal"`, `"hilited"`, `"clicked"`).
  - `x`, `y`: Pixel coordinates relative to the window.
  - `txt`: The text string to render.

---

### `epplet.TextclassGetSize`
```go
func TextclassGetSize(tclass, state, text string) Size
```
- **Description**: Computes the exact rendered width and height dimensions (in pixels) for a text string under a specific textclass without drawing it.
- **Returns**: `Size` struct:
  ```go
  type Size struct {
      Width  int
      Height int
  }
  ```
- **Example**:
  ```go
  size := epplet.TextclassGetSize("EPPLET_LABEL", "normal", "Hello World")
  fmt.Printf("Rendered size: %dx%d pixels\n", size.Width, size.Height)
  ```

---

## 5. Code Example: Custom Imageclass & Textclass Drawing

```go
package main

import (
	"fmt"

	"github.com/redporter/epplet-api"
)

func main() {
	epplet.Init("ThemeDemo", "1.0", "Imageclass & Textclass Example", 4, 4, false)

	// Get Main Epplet Window
	win := epplet.GetMainWindow()

	// Apply themed epplet background to main window
	win.ImageclassApply("EPPLET_BACKGROUND", "normal")

	// Calculate text size dynamically
	textSize := epplet.TextclassGetSize("EPPLET_LABEL", "normal", "System Status")
	fmt.Printf("Label dimensions: %dx%d px\n", textSize.Width, textSize.Height)

	// Register Expose Event Handler to draw text and themed elements
	epplet.RegisterExposeHandler(func(w epplet.Window, x, y, width, height int) {
		// Draw themed button background tile
		w.ImageclassPaste("EPPLET_BUTTON", "normal", 4, 4, 56, 20)

		// Draw themed text on top of tile
		w.TextclassDraw("EPPLET_LABEL", "normal", 10, 8, "System Status")
	})

	epplet.Show()
	epplet.Loop()
}
```
