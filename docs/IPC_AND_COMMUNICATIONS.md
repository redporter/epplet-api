# IPC & Communications Documentation

This document describes the Inter-Process Communication (IPC) system in the `epplet-api` library used to communicate with the Enlightenment 16 (e16) window manager and external script utilities.

---

## 1. Overview

Enlightenment 16 communicates with running Epplets via an X11 ClientMessage IPC protocol. IPC messages allow Epplets to:

- Send status updates, title changes, or layout requests to e16.
- Receive command notifications, theme changes, or control signals from e16 or external shell scripts (e.g., `eesh`).

---

## 2. Sending IPC Messages

### `epplet.SendIPC`
```go
func SendIPC(msg string)
```
- **Description**: Transmits an IPC string message to the Enlightenment 16 window manager.
- **Example**:
  ```go
  epplet.SendIPC("title My Epplet Title")
  ```

---

## 3. Receiving IPC Messages

`epplet-api` provides three mechanisms for receiving and processing IPC messages:

### Option A: Command Dispatcher (Recommended)

The **Command Dispatcher** automatically parses incoming IPC strings of the form `"<command> [arguments...]"` and dispatches them to registered handler callbacks.

1. **`epplet.HandleCommand`**:
   ```go
   func HandleCommand(cmd string, fn func(args string))
   ```
   Registers a callback function executed when an IPC message starting with `cmd` is received.

2. **`epplet.HandleUnknownCommand`**:
   ```go
   func HandleUnknownCommand(fn func(rawMsg string))
   ```
   Registers a fallback callback for any IPC messages that do not match a registered command.

#### Example: Command Dispatcher Usage
```go
package main

import (
	"fmt"
	"github.com/redporter/epplet-api"
)

func main() {
	epplet.Init("IPCDemo", "1.0", "IPC Example", 4, 4, false)

	// Handle "theme <name>" IPC commands
	epplet.HandleCommand("theme", func(args string) {
		fmt.Printf("Theme changed to: %s\n", args)
	})

	// Handle "reload" IPC command
	epplet.HandleCommand("reload", func(args string) {
		fmt.Println("Reloading epplet state...")
	})

	// Fallback for unhandled IPC messages
	epplet.HandleUnknownCommand(func(rawMsg string) {
		fmt.Printf("Received raw IPC: %s\n", rawMsg)
	})

	epplet.Show()
	epplet.Loop()
}
```

---

### Option B: Asynchronous Channel Stream

```go
func WaitForIPCAsync() <-chan string
```
- **Description**: Returns a read-only Go channel (`<-chan string`) that continuously receives IPC message strings in a background goroutine.

#### Example: Channel Streaming Usage
```go
go func() {
    ipcChan := epplet.WaitForIPCAsync()
    for msg := range ipcChan {
        fmt.Println("Async IPC received:", msg)
    }
}()
```

---

### Option C: Synchronous Blocking

```go
func BlockForIPC() string
```
- **Description**: Blocks the calling thread until an IPC message arrives.
- **Note**: While blocking, event handling and timers on the calling thread are suspended. Use `WaitForIPCAsync` or `HandleCommand` for non-blocking operations.

---

## 4. Testing IPC Messages via `eesh`

From a terminal, you can send IPC commands to active Epplets using Enlightenment's `eesh` tool:

```bash
# Send an IPC message to e16 / Epplet
eesh -e "epplet_ipc IPCDemo theme dark"
```
