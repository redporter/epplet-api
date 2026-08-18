package epplet

/*
#include <X11/Xlib.h>

#ifndef _EPPLET_H_GUARD
#define _EPPLET_H_GUARD
#include <epplet.h>
#endif

#include <stdlib.h>

extern void goChildGateway(void *data, int pid, int exit_code);

static inline void register_child_bridge(void *handle) {
    Epplet_register_child_handler(
        (void (*)(void *, int, int))goChildGateway,
        handle
    );
}
*/
import "C"

import (
	"runtime/cgo"
	"unsafe"
)

// -----------------------------------------------------------------------------
// Command Execution & Process Spawning
// -----------------------------------------------------------------------------

// RunCommand executes a shell command synchronously and returns its exit code.
func RunCommand(cmd string) int {
	cCmd := C.CString(cmd)
	defer C.free(unsafe.Pointer(cCmd))
	return int(C.Epplet_run_command(cCmd))
}

// ReadRunCommand executes a shell command synchronously and returns its output text.
func ReadRunCommand(cmd string) string {
	cCmd := C.CString(cmd)
	defer C.free(unsafe.Pointer(cCmd))
	cStr := C.Epplet_read_run_command(cCmd)
	if cStr == nil {
		return ""
	}
	return C.GoString(cStr)
}

// SpawnCommand spawns a command asynchronously in the background and returns its process ID (PID).
func SpawnCommand(cmd string) int {
	cCmd := C.CString(cmd)
	defer C.free(unsafe.Pointer(cCmd))
	return int(C.Epplet_spawn_command(cCmd))
}

// PauseSpawnedCommand sends SIGSTOP to pause a spawned child process.
func PauseSpawnedCommand(pid int) {
	C.Epplet_pause_spawned_command(C.int(pid))
}

// UnpauseSpawnedCommand sends SIGCONT to resume a paused spawned child process.
func UnpauseSpawnedCommand(pid int) {
	C.Epplet_unpause_spawned_command(C.int(pid))
}

// KillSpawnedCommand sends SIGKILL to terminate a spawned child process.
func KillSpawnedCommand(pid int) {
	C.Epplet_kill_spawned_command(C.int(pid))
}

// DestroySpawnedCommand cleans up internal resources associated with a spawned child process.
func DestroySpawnedCommand(pid int) {
	C.Epplet_destroy_spawned_command(C.int(pid))
}

// ChildHandler is a callback invoked when a spawned child process terminates.
type ChildHandler func(pid, exitCode int)

//export goChildGateway
func goChildGateway(data unsafe.Pointer, pid, exitCode C.int) {
	if data == nil {
		return
	}
	h := cgo.Handle(data)
	if fn, ok := h.Value().(ChildHandler); ok {
		fn(int(pid), int(exitCode))
	}
}

// RegisterChildHandler registers a callback function to handle child process termination.
func RegisterChildHandler(handler ChildHandler) func() {
	handle := cgo.NewHandle(handler)
	C.register_child_bridge(unsafe.Pointer(handle))
	return func() {
		handle.Delete()
	}
}

// -----------------------------------------------------------------------------
// Dialog Helpers
// -----------------------------------------------------------------------------

// ShowAbout displays the standard Enlightenment Epplet "About" dialog for the epplet.
func ShowAbout(name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.Epplet_show_about(cName)
}

// DialogOk displays a standard OK message dialog popup window.
func DialogOk(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	C.Epplet_dialog_ok(cText)
}
