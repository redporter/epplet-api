# 2D Primitive Drawing & RGB Buffers Documentation

This document describes how to perform 2D vector drawing operations, image pasting, and high-performance direct pixel manipulation using **RGB Buffers** in the `epplet-api` Go library.

---

## 1. 2D Primitive Drawing Operations

`epplet-api` exposes simple 2D drawing functions to render lines, filled boxes, outlines, and images directly onto any X11 `Window`.

### Drawing Functions & Window Receivers

| Function | Window Receiver Method | Description |
| :--- | :--- | :--- |
| `DrawLine(win, x1, y1, x2, y2, r, g, b)` | `win.DrawLine(x1, y1, x2, y2, r, g, b)` | Draws a line from `(x1, y1)` to `(x2, y2)` in RGB color `(0-255)`. |
| `DrawBox(win, x, y, w, h, r, g, b)` | `win.DrawBox(x, y, w, h, r, g, b)` | Draws a filled rectangle at `(x, y)` with size `(w, h)` in RGB color. |
| `DrawOutline(win, x, y, w, h, r, g, b)` | `win.DrawOutline(x, y, w, h, r, g, b)` | Draws a hollow rectangle outline at `(x, y)` with size `(w, h)` in RGB color. |
| `PasteImage(image, win, x, y)` | `win.PasteImage(image, x, y)` | Pastes an image file onto `win` at position `(x, y)` at its original size. |
| `PasteImageSize(image, win, x, y, w, h)` | `win.PasteImageSize(image, x, y, w, h)` | Pastes and scales an image file onto `win` at position `(x, y)` at size `(w, h)`. |
| `GetColor(r, g, b)` | — | Returns the X11 pixel color value for RGB tuple `(r, g, b)`. |
| `Sync()` | — | Synchronizes and flushes all pending X11 draw operations to the server. |

---

## 2. High-Performance RGB Buffer API

For custom software rendering, procedural graphics, CPU raycasters, or framebuffers, `epplet-api` provides the **RGB Buffer API**.

An `RGBBuf` manages a raw 32-bit ARGB pixel memory buffer created in C and exposes it as an idiomatic Go byte slice (`[]byte`) using zero-copy memory mapping.

### RGB Buffer Methods

1. **`MakeRGBBuf(w, h int) *RGBBuf`**:
   Allocates a new 32-bit ARGB pixel buffer of dimensions `w x h`.

2. **`buf.Data() []byte`**:
   Returns a Go byte slice pointing directly to the underlying raw ARGB pixel memory:
   - Slice length is `w * h * 4` bytes.
   - Byte memory order per pixel is **Blue, Green, Red, Alpha** (`[B, G, R, A]`).

3. **`buf.Paste(win Window, x, y int)` / `win.PasteBuf(buf, x, y)`**:
   Blits/pastes the RGB buffer onto `win` at position `(x, y)` with automatic Imlib2 dithering and color mapping.

4. **`buf.Free()`**:
   Frees the allocated C pixel buffer memory.

---

## 3. Code Example: Direct Pixel Software Animation

```go
package main

import (
	"math"
	"time"

	"github.com/redporter/epplet-api"
)

func main() {
	epplet.Init("PixelDemo", "1.0", "RGB Buffer Animation", 4, 4, false)

	// Get main window (64x64 pixels)
	win := epplet.GetMainWindow()

	// Allocate 56x56 RGB Buffer
	buf := epplet.MakeRGBBuf(56, 56)
	if buf == nil {
		return
	}
	defer buf.Free()

	w, h := buf.Width(), buf.Height()
	pix := buf.Data()

	var tick float64 = 0.0

	// Register Expose Event Handler
	epplet.RegisterExposeHandler(func(win epplet.Window, x, y, width, height int) {
		buf.Paste(win, 4, 4)
	})

	// Render loop timer (~30 FPS)
	var updateTimer epplet.TimerCallback
	updateTimer = func() {
		tick += 0.1

		// Procedural plasma/wave pixel rendering directly in Go byte slice
		for yPos := 0; yPos < h; yPos++ {
			for xPos := 0; xPos < w; xPos++ {
				offset := (yPos*w + xPos) * 4

				v := math.Sin(float64(xPos)*0.1+tick) + math.Cos(float64(yPos)*0.1+tick)
				intensity := byte((v + 2.0) / 4.0 * 255.0)

				pix[offset+0] = intensity       // Blue
				pix[offset+1] = 255 - intensity // Green
				pix[offset+2] = intensity / 2   // Red
				pix[offset+3] = 255             // Alpha
			}
		}

		// Paste updated buffer onto window
		win.PasteBuf(buf, 4, 4)
		epplet.Sync()

		// Reschedule timer
		epplet.Timer(updateTimer, 33*time.Millisecond, "plasma_timer")
	}

	epplet.Timer(updateTimer, 33*time.Millisecond, "plasma_timer")

	epplet.Show()
	epplet.Loop()
}
```
