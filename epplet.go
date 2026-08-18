package epplet

/*
#include <X11/Xlib.h>

#ifndef _EPPLET_H_GUARD
#define _EPPLET_H_GUARD
#include <epplet.h>
#endif

#include <stdlib.h>

extern void goTimerGateway(void *data);
static inline void create_timer_bridge(void *handle, double time, char* name) {
    Epplet_timer(
        (void (*)(void *))goTimerGateway,
        handle,
        time,
        name
    );
}
*/
import "C"

import (
	"runtime/cgo"
	"sync"
	"time"
	"unsafe"
)

// Init initializes the Epplet with Enlightenment 16 (e16).
//
// Parameters:
//   - name: The application identifier string (e.g., "E-Clock").
//   - version: Version string (e.g., "1.0").
//   - info: Descriptive text about the epplet shown in About dialogs.
//   - w, h: Width and height in 16-pixel tile grid units (e.g., 4, 4 -> 64x64 pixels).
//   - vertical: Set to true for vertical layout orientation, false for horizontal.
func Init(name, version, info string, w, h int, vertical bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cVersion := C.CString(version)
	defer C.free(unsafe.Pointer(cVersion))

	cInfo := C.CString(info)
	defer C.free(unsafe.Pointer(cInfo))

	cVert := 0
	if vertical {
		cVert = 1
	}

	C.Epplet_Init(cName, cVersion, cInfo, C.int(w), C.int(h), C.int(0), nil, C.char(cVert))
}

// Cleanup saves epplet configuration settings to disk, releases allocated X11 resources,
// and performs orderly shutdown of the epplet library.
func Cleanup() {
	C.Epplet_cleanup()
}

// Show maps and displays the primary Epplet window on the desktop.
func Show() {
	C.Epplet_show()
}

// Remember instructs Enlightenment 16 to remember the epplet's current screen position,
// desktop layer, and sticky attributes.
func Remember() {
	C.Epplet_remember()
}

// Unremember instructs Enlightenment 16 to forget any stored positioning or state attributes
// for this epplet.
func Unremember() {
	C.Epplet_unremember()
}

// Loop enters the main Epplet event processing loop. This is a blocking call that
// handles X11 events, timer ticks, IPC commands, and gadget interactions until the app exits.
func Loop() {
	C.Epplet_Loop()
}

// -----------------------------------------------------------------------------
// Display & X11 Connection
// -----------------------------------------------------------------------------

// Display wraps the underlying X11 Display connection pointer.
type Display struct {
	ptr *C.Display
}

// GetDisplay returns a pointer to the X11 Display connection used by the Epplet library.
func GetDisplay() *Display {
	cDisp := C.Epplet_get_display()
	if cDisp == nil {
		return nil
	}
	return &Display{ptr: cDisp}
}

// -----------------------------------------------------------------------------
// IPC Communications
// -----------------------------------------------------------------------------

// SendIPC sends an Enlightenment IPC message string to the e16 window manager.
func SendIPC(msg string) {
	cStr := C.CString(msg)
	defer C.free(unsafe.Pointer(cStr))
	C.Epplet_send_ipc(cStr)
}

// BlockForIPC blocks until an IPC message is received and returns the message text string.
// Returns an empty string if the call fails.
func BlockForIPC() string {
	cMsg := C.Epplet_wait_for_ipc()
	if cMsg == nil {
		return ""
	}
	return C.GoString(cMsg)
}

// WaitForIPCAsync returns a read-only Go channel that receives IPC messages asynchronously.
func WaitForIPCAsync() <-chan string {
	ch := make(chan string, 10)
	go func() {
		for {
			msg := BlockForIPC()
			if msg != "" {
				ch <- msg
			}
		}
	}()
	return ch
}

// -----------------------------------------------------------------------------
// Imageclass & Textclass Helpers
// -----------------------------------------------------------------------------

// ImageclassPixmaps represents pixmap handles for an imageclass state.
type ImageclassPixmaps struct {
	Pixmap C.Pixmap
	Mask   C.Pixmap
}

// ImageclassGetPixmaps calculates and returns the Pixmap and Mask handles for a given imageclass and state.
func ImageclassGetPixmaps(iclass, state string, width, height int) ImageclassPixmaps {
	cIclass := C.CString(iclass)
	defer C.free(unsafe.Pointer(cIclass))

	cState := C.CString(state)
	defer C.free(unsafe.Pointer(cState))

	var p, m C.Pixmap
	C.Epplet_imageclass_get_pixmaps(cIclass, cState, &p, &m, C.int(width), C.int(height))
	return ImageclassPixmaps{Pixmap: p, Mask: m}
}

// Size represents width and height dimensions in pixels.
type Size struct {
	Width  int
	Height int
}

// TextclassGetSize calculates and returns the rendered pixel width and height for text formatted with a textclass.
func TextclassGetSize(tclass, state, text string) Size {
	cTclass := C.CString(tclass)
	defer C.free(unsafe.Pointer(cTclass))

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	var w, h C.int
	C.Epplet_textclass_get_size(cTclass, &w, &h, cText)
	return Size{Width: int(w), Height: int(h)}
}

// -----------------------------------------------------------------------------
// Timers
// -----------------------------------------------------------------------------

// TimerCallback represents a function signature invoked by Epplet timers.
type TimerCallback func()

var (
	timerHandles = make(map[string]cgo.Handle)
	timerMu      sync.Mutex
)

//export goTimerGateway
func goTimerGateway(data unsafe.Pointer) {
	if data == nil {
		return
	}

	h := cgo.Handle(data)
	if cb, ok := h.Value().(TimerCallback); ok {
		cb()
	}

	// Epplet timers are one-shot by default in e16.
	// Find and cleanup the handle from our tracking map if not rescheduled.
	timerMu.Lock()
	for name, activeHandle := range timerHandles {
		if activeHandle == h {
			delete(timerHandles, name)
			h.Delete()
			break
		}
	}
	timerMu.Unlock()
}

// Timer registers a one-shot timer callback that fires after duration d.
// To create a recurring timer, invoke Timer again inside the callback function.
func Timer(cb TimerCallback, d time.Duration, name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	timerMu.Lock()
	if oldHandle, exists := timerHandles[name]; exists {
		oldHandle.Delete()
	}

	handle := cgo.NewHandle(cb)
	timerHandles[name] = handle
	timerMu.Unlock()

	seconds := d.Seconds()
	C.create_timer_bridge(unsafe.Pointer(handle), C.double(seconds), cName)
}

// RemoveTimer cancels and removes an active timer by name.
func RemoveTimer(name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	timerMu.Lock()
	if handle, exists := timerHandles[name]; exists {
		handle.Delete()
		delete(timerHandles, name)
	}
	timerMu.Unlock()

	C.Epplet_remove_timer(cName)
}

// TimerGetData retrieves the active TimerCallback function for a named timer, or nil if missing.
func TimerGetData(name string) TimerCallback {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	rawPtr := C.Epplet_timer_get_data(cName)
	if rawPtr == nil {
		return nil
	}

	h := cgo.Handle(rawPtr)
	if cb, ok := h.Value().(TimerCallback); ok {
		return cb
	}
	return nil
}
