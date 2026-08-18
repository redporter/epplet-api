# epplet-api

[![Go Reference](https://pkg.go.dev/badge/github.com/redporter/epplet-api.svg)](https://pkg.go.dev/github.com/redporter/epplet-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/redporter/epplet-api)](https://goreportcard.com/report/github.com/redporter/epplet-api)
[![License: BSD-3-Clause](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](LICENSE)

**`epplet-api`** is an idiomatic, high-performance Go binding library for building **Enlightenment 16 (e16) Epplets**.

It provides 100% C API coverage over `libepplet`, replacing untyped C pointers (`void *`) with type-safe Go interfaces, concrete widget receivers, zero-copy RGB framebuffers, hardware-accelerated OpenGL/GLX binding, and structured event dispatching.

---

## 🌟 Key Features

- **🛡️ 100% Type-Safe Widget Architecture**: Dedicated Go structs (`*Button`, `*Textbox`, `*ProgressBar`, `*Slider`, `*Popup`, `*DrawingArea`, etc.) implementing the unified `Gadget` interface.
- **⚡ Zero-Copy RGB Framebuffer**: Direct memory access to raw 32-bit ARGB pixel buffers via `unsafe.Slice` (`RGBBuf.Data() []byte`).
- **🎨 OpenGL / GLX 3D Acceleration**: Integrated GLX context creation, dynamic symbol resolution (`dlsym`), and seamless integration with [`github.com/go-gl/gl`](https://github.com/go-gl/gl).
- **🔄 Event & Callback Gateway**: Type-safe Go closures for mouse clicks, hover, keyboard navigation, exposed area redraws, and window destruction.
- **💾 Persistent Configuration System**: In-memory and disk persistence (`QueryConfig`, `ModifyConfig`, `QueryMultiConfig`).
- **📡 e16 IPC & Command Dispatcher**: Automated parsing and routing of Enlightenment IPC messages to Go callbacks.
- **⚙️ Asynchronous Process Management**: Spawn, pause, resume, and track background child processes with exit callbacks.

---

## 📦 Installation & System Requirements

### Prerequisites

Building Go applications with `epplet-api` requires CGO and the following development libraries:

- **Enlightenment 16**: `libepplet` / `libepplet_glx`
- **X11 / Graphics**: `libX11`, `libImlib2`
- **OpenGL** (optional for 3D): `libGL`, Mesa drivers

#### FreeBSD
```bash
pkg install e16-epplet-base libX11 imlib2 mesa-libs
```

#### Ubuntu / Debian
```bash
sudo apt-get install e16 epplets libx11-dev libimlib2-dev libgl1-mesa-dev
```

### Go Installation

```bash
go get github.com/redporter/epplet-api
```

---

## 🚀 Quick Start Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/redporter/epplet-api"
)

var progressVal int = 20

func main() {
	// 1. Initialize Epplet (4x4 tile grid = 64x64 pixels)
	epplet.Init("ClockEpplet", "1.0", "Sample Go Epplet", 4, 4, false)

	// 2. Create UI Label
	lbl := epplet.CreateLabel(4, 4, "00:00:00", 2)
	lbl.Show()

	// 3. Create Progress Bar
	bar := epplet.CreateHBar(4, 22, 56, 10, 0, &progressVal)
	bar.Show()

	// 4. Create Button
	btn := epplet.CreateTextButton("Click", 4, 38, 56, 18, func() {
		fmt.Println("Button clicked!")
	})
	btn.Show()

	// 5. Register Recurring Timer Ticker
	var tick epplet.TimerCallback
	tick = func() {
		lbl.Change(time.Now().Format("15:04:05"))

		progressVal = (progressVal + 5) % 100
		bar.DataChanged()

		epplet.Timer(tick, 1*time.Second, "clock_timer")
	}
	epplet.Timer(tick, 1*time.Second, "clock_timer")

	// 6. Map Window and Enter Main Event Loop
	epplet.Show()
	epplet.Loop()
}
```

---

## 📚 Comprehensive Documentation Guides

For in-depth architectural guides, API details, and runnable code examples, see the [`docs/`](docs) directory:

| Guide | Description |
| :--- | :--- |
| **[Core Lifecycle & Main Windowing](docs/LIFECYCLE_AND_WINDOWING.md)** | Application startup, event loops, tile math grid, and display handles. |
| **[Auxiliary Windows & Context Stack](docs/AUXILIARY_WINDOWS_AND_CONTEXT_STACK.md)** | Config dialogs, borderless windows, and window drawing context routing. |
| **[Gadgets & Widgets](docs/GADGETS_AND_WIDGETS.md)** | Full widget catalog (`Button`, `Textbox`, `Slider`, `Popup`, `ProgressBar`, etc.). |
| **[2D Primitive Drawing & RGB Buffers](docs/PRIMITIVES_AND_RGB_BUFFERS.md)** | Vector primitives, image blitting, and high-performance pixel framebuffers. |
| **[OpenGL / GLX Contexts](docs/OPENGL_AND_GLX.md)** | 3D hardware context binding, multi-thread activation, and `go-gl` integration. |
| **[IPC & Communications](docs/IPC_AND_COMMUNICATIONS.md)** | e16 ClientMessage IPC protocol, command tokenization, and `eesh` integration. |
| **[Event & Callback Handlers](docs/EVENT_AND_CALLBACK_HANDLERS.md)** | Mouse tracking, keyboard events, exposed area redraws, and window destruction. |
| **[Timers](docs/TIMERS.md)** | One-shot timers, recurring tickers, self-rescheduling callbacks, and cancellation. |
| **[Config Subsystem](docs/CONFIG_SUBSYSTEM.md)** | Persistent key-value settings, multi-value string arrays, and e16 config paths. |
| **[Command Execution & Dialogs](docs/PROCESSES_AND_DIALOGS.md)** | Process spawning (`SpawnCommand`), signal control, and built-in dialog popups. |

For a complete 1-to-1 C API symbol bridging matrix, see **[C_API_BRIDGING.md](C_API_BRIDGING.md)**.

---

## 📄 License

This library is licensed under the **BSD 3-Clause License**. See **[LICENSE](LICENSE)** for details.

*Includes third-party copyright notices for libepplet / Enlightenment 16 (Copyright © 1997–2000 Carsten Haitzler and Enlightenment contributors).*
