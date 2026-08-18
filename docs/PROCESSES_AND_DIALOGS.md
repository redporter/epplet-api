# Command Execution & Dialog Helpers Documentation

This document describes how to execute shell commands, manage background child processes, and display standard Enlightenment dialog boxes using `epplet-api`.

---

## 1. Synchronous Command Execution

`epplet-api` provides wrapper functions to execute shell commands synchronously:

### `epplet.RunCommand`
```go
func RunCommand(cmd string) int
```
- **Description**: Runs a shell command synchronously and returns its integer exit status code (`0` for success).

---

### `epplet.ReadRunCommand`
```go
func ReadRunCommand(cmd string) string
```
- **Description**: Runs a shell command synchronously and captures its standard output string.
- **Example**:
  ```go
  uptime := epplet.ReadRunCommand("uptime")
  fmt.Println("System Uptime:", uptime)
  ```

---

## 2. Asynchronous Process Spawning & Control

For long-running processes or background tasks, use the process spawning API:

### Process Lifecycle Functions

| Function | Description |
| :--- | :--- |
| `SpawnCommand(cmd string) int` | Spawns a command asynchronously in the background and returns its Process ID (`PID`). |
| `PauseSpawnedCommand(pid int)` | Sends `SIGSTOP` signal to pause execution of a spawned child process. |
| `UnpauseSpawnedCommand(pid int)` | Sends `SIGCONT` signal to resume execution of a paused child process. |
| `KillSpawnedCommand(pid int)` | Sends `SIGKILL` signal to immediately terminate a spawned child process. |
| `DestroySpawnedCommand(pid int)` | Cleans up epplet tracking resources associated with a spawned child process. |
| `RegisterChildHandler(handler ChildHandler) func()` | Registers a callback function executed whenever a spawned child process terminates. |

---

### Child Process Termination Callback

```go
type ChildHandler func(pid, exitCode int)
```

Register a `ChildHandler` to be notified when background commands finish:

```go
unregister := epplet.RegisterChildHandler(func(pid, exitCode int) {
    fmt.Printf("Child process %d exited with code %d\n", pid, exitCode)
})
defer unregister()
```

---

## 3. Dialog Helpers

Enlightenment 16 Epplets provide built-in themed dialog popups:

### `epplet.ShowAbout`
```go
func ShowAbout(name string)
```
- **Description**: Displays the standard Enlightenment 16 "About" dialog for the Epplet, displaying the application name, version, and description set during `epplet.Init()`.

---

### `epplet.DialogOk`
```go
func DialogOk(text string)
```
- **Description**: Displays a modal message box popup containing the specified text message and an **OK** button.

---

## 4. Code Example: Background Process & Dialog Controls

```go
package main

import (
	"fmt"
	"time"

	"github.com/redporter/epplet-api"
)

var activePID int

func main() {
	epplet.Init("ProcessDemo", "1.0", "Process Spawning & Dialog Demo", 4, 6, false)

	// Register child process termination handler
	epplet.RegisterChildHandler(func(pid, exitCode int) {
		fmt.Printf("Process %d finished (exit code %d)\n", pid, exitCode)
		epplet.DialogOk(fmt.Sprintf("Process %d Completed!", pid))
		activePID = 0
	})

	// Button to spawn a background process
	btnRun := epplet.CreateTextButton("Spawn Task", 4, 4, 56, 18, func() {
		if activePID != 0 {
			epplet.DialogOk("A task is already running!")
			return
		}
		activePID = epplet.SpawnCommand("sleep 3")
		fmt.Printf("Spawned background task PID: %d\n", activePID)
	})
	btnRun.Show()

	// Button to show About dialog
	btnAbout := epplet.CreateTextButton("About", 4, 26, 56, 18, func() {
		epplet.ShowAbout("ProcessDemo")
	})
	btnAbout.Show()

	epplet.Show()
	epplet.Loop()
}
```
