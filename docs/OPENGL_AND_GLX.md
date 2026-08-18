# OpenGL / GLX Contexts Documentation

This document describes how to create, configure, and render hardware-accelerated 3D graphics using **OpenGL / GLX Contexts** attached to `DrawingArea` widgets in `epplet-api`.

---

## 1. Architecture Overview

Enlightenment 16 Epplets allow hardware-accelerated 3D rendering directly inside a `DrawingArea` gadget.

```
+-----------------------------------------------------------+
| 1. epplet.CreateDrawingArea(x, y, w, h)                   |
+-----------------------------------------------------------+
                              |
                              v
+-----------------------------------------------------------+
| 2. glxCtx := da.DefaultBindGL()                           |
|    - Dynamically loads libepplet_glx.so                   |
|    - Chooses X11 visual & creates OpenGL context          |
+-----------------------------------------------------------+
                              |
                              v
+-----------------------------------------------------------+
| 3. Render Loop (~30 FPS Timer Callback)                   |
|    a. da.MakeCurrent(glxCtx)                              |
|    b. gl.Viewport(0, 0, w, h)                             |
|    c. Perform 3D Drawing (go-gl / OpenGL)                |
|    d. da.SwapBuffers()                                    |
+-----------------------------------------------------------+
                              |
                              v
+-----------------------------------------------------------+
| 4. glxCtx.Unbind()                                        |
+-----------------------------------------------------------+
```

---

## 2. Checking System GLX Availability

On Linux and FreeBSD systems, GLX functions reside in a separate shared object (`libepplet_glx.so`). `epplet-api` uses dynamic symbol resolution to prevent build/linker failures on installations compiled without GLX.

```go
if !epplet.HasGLXSupport() {
    fmt.Println("Warning: System libEpplet was built without GLX support.")
    return
}
```

---

## 3. Binding & Destroying Contexts

### Receiver Methods on `*DrawingArea`

1. **`da.DefaultBindGL() GLXContext`**:
   Creates a standard double-buffered RGB OpenGL context with minimal depth buffer. Recommended for most applications.

2. **`da.BindDoubleGL(...) GLXContext`**:
   Creates a double-buffered GLX context with explicit bit-depth parameters:
   ```go
   cx := da.BindDoubleGL(
       red, blue, green, alpha,
       auxBuffers, depth, stencil,
       accumRed, accumGreen, accumBlue, accumAlpha int,
   )
   ```

3. **`da.BindSingleGL(...) GLXContext`**:
   Creates a single-buffered GLX context with explicit bit-depth parameters.

4. **`glxCtx.Unbind()` / `epplet.UnbindGL(glxCtx)`**:
   Destroys and releases the GLX context.

---

## 4. Context Activation & Buffer Swapping

In Go, timer callbacks and event handlers may run on different OS threads managed by the Go runtime scheduler. Before invoking OpenGL commands, you must activate the GLX context on the current thread:

- **`da.MakeCurrent(glxCtx)` / `win.MakeCurrent(glxCtx)`**:
  Binds `glxCtx` as the active rendering context for the drawing area window on the calling thread.

- **`da.SwapBuffers()` / `win.SwapBuffers()`**:
  Swaps the front and back OpenGL framebuffers when double-buffering.

---

## 5. Integrating Pure Go OpenGL Bindings (`go-gl`)

`epplet-api` pairs seamlessly with official Go OpenGL bindings like `github.com/go-gl/gl/v2.1/gl` or `github.com/go-gl/gl/v3.3-core/gl`:

```go
da.MakeCurrent(glxCtx)
if err := gl.Init(); err != nil {
    log.Fatal(err)
}
```

---

## 6. Code Example: Complete OpenGL Epplet

```go
package main

import (
	"fmt"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/redporter/epplet-api"
)

var angle float32 = 0.0

func main() {
	epplet.Init("GLXDemo", "1.0", "OpenGL GLX Demo Epplet in Go", 4, 4, false)

	if !epplet.HasGLXSupport() {
		fmt.Println("Warning: System libEpplet was built without GLX support.")
		return
	}

	// Create 56x56 DrawingArea
	da := epplet.CreateDrawingArea(4, 4, 56, 56)
	da.Show()

	// Bind GLX context
	glxCtx := da.DefaultBindGL()
	if glxCtx == 0 {
		return
	}
	defer glxCtx.Unbind()

	// Activate context and initialize OpenGL bindings
	da.MakeCurrent(glxCtx)
	gl.Init()

	// 30 FPS Render Loop Callback
	var drawFrame epplet.TimerCallback
	drawFrame = func() {
		angle += 4.0

		// Make context active on calling thread
		da.MakeCurrent(glxCtx)

		// Set viewport & clear buffer
		gl.Viewport(0, 0, int32(da.GetWidth()), int32(da.GetHeight()))
		gl.ClearColor(0.1, 0.1, 0.15, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Projection matrix
		gl.MatrixMode(gl.PROJECTION)
		gl.LoadIdentity()
		gl.Ortho(-1.0, 1.0, -1.0, 1.0, -1.0, 1.0)

		// Modelview matrix & rotation
		gl.MatrixMode(gl.MODELVIEW)
		gl.LoadIdentity()
		gl.Rotatef(angle, 0.0, 0.0, 1.0)

		// Render colored triangle
		gl.Begin(gl.TRIANGLES)
		gl.Color3f(1.0, 0.2, 0.2)
		gl.Vertex2f(0.0, 0.6)

		gl.Color3f(0.2, 1.0, 0.2)
		gl.Vertex2f(-0.5, -0.4)

		gl.Color3f(0.2, 0.4, 1.0)
		gl.Vertex2f(0.5, -0.4)
		gl.End()

		// Swap front & back buffers
		da.SwapBuffers()

		// Reschedule timer
		epplet.Timer(drawFrame, 33*time.Millisecond, "gl_timer")
	}

	epplet.Timer(drawFrame, 33*time.Millisecond, "gl_timer")

	epplet.Show()
	epplet.Loop()
}
```
