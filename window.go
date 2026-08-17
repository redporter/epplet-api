package epplet

/*
#include <X11/Xlib.h>

#ifndef _EPPLET_H_GUARD
#define _EPPLET_H_GUARD
#include <epplet.h>
#endif

#include <stdlib.h>

extern void goGadgetGateway(void *data);

static inline Window create_window_config_bridge(
    int w, int h, char *title,
    void *ok_handle, void *apply_handle, void *cancel_handle)
{
    return Epplet_create_window_config(
        w, h, title,
        ok_handle ? (void (*)(void *))goGadgetGateway : NULL, ok_handle,
        apply_handle ? (void (*)(void *))goGadgetGateway : NULL, apply_handle,
        cancel_handle ? (void (*)(void *))goGadgetGateway : NULL, cancel_handle
    );
}
*/
import "C"

import (
	"runtime/cgo"
	"unsafe"
)

type Window uintptr

type WindowConfigCallback func()

// GetMainWindow returns the main Epplet window handle.
func GetMainWindow() Window {
	return Window(C.Epplet_get_main_window())
}

// CreateWindow creates an auxiliary window of size (w, h) in 16-pixel tile multiples.
func CreateWindow(w, h int, title string, vertical bool) Window {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	vert := C.char(0)
	if vertical {
		vert = C.char(1)
	}

	win := C.Epplet_create_window(C.int(w), C.int(h), cTitle, vert)
	return Window(win)
}

// CreateWindowBorderless creates a borderless window of size (w, h) in tile multiples.
func CreateWindowBorderless(w, h int, title string, vertical bool) Window {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	vert := C.char(0)
	if vertical {
		vert = C.char(1)
	}

	win := C.Epplet_create_window_borderless(C.int(w), C.int(h), cTitle, vert)
	return Window(win)
}

// CreateWindowConfig creates a configuration window with optional OK, Apply, and Cancel callbacks.
func CreateWindowConfig(w, h int, title string, okCb, applyCb, cancelCb WindowConfigCallback) Window {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	var okHandlePtr, applyHandlePtr, cancelHandlePtr unsafe.Pointer

	if okCb != nil {
		h := cgo.NewHandle(okCb)
		okHandlePtr = unsafe.Pointer(h)
	}
	if applyCb != nil {
		h := cgo.NewHandle(applyCb)
		applyHandlePtr = unsafe.Pointer(h)
	}
	if cancelCb != nil {
		h := cgo.NewHandle(cancelCb)
		cancelHandlePtr = unsafe.Pointer(h)
	}

	win := C.create_window_config_bridge(
		C.int(w), C.int(h), cTitle,
		okHandlePtr, applyHandlePtr, cancelHandlePtr,
	)
	return Window(win)
}

// WindowPushContext sets newwin as the active window context for creating subsequent gadgets.
func WindowPushContext(newwin Window) {
	C.Epplet_window_push_context(C.Window(newwin))
}

// WindowPopContext pops the active window context and returns the previously popped Window.
func WindowPopContext() Window {
	return Window(C.Epplet_window_pop_context())
}

// PushContext pushes this window onto the gadget creation context stack.
func (win Window) PushContext() {
	C.Epplet_window_push_context(C.Window(win))
}

// Show displays the window.
func (win Window) Show() {
	C.Epplet_window_show(C.Window(win))
}

// Hide hides the window.
func (win Window) Hide() {
	C.Epplet_window_hide(C.Window(win))
}

// Destroy destroys the window and all gadgets attached to it.
func (win Window) Destroy() {
	C.Epplet_window_destroy(C.Window(win))
}

// Clear clears the content of the window.
func (win Window) Clear() {
	C.Epplet_clear_window(C.Window(win))
}

// ImageclassApply sets the imageclass background on the window.
func (win Window) ImageclassApply(iclass, state string) {
	cIclass := C.CString(iclass)
	defer C.free(unsafe.Pointer(cIclass))
	cState := C.CString(state)
	defer C.free(unsafe.Pointer(cState))

	C.Epplet_imageclass_apply(cIclass, cState, C.Window(win))
}

// ImageclassPaste pastes an imageclass onto the window at (x, y) with size (w, h).
func (win Window) ImageclassPaste(iclass, state string, x, y, h, w int) {
	cIclass := C.CString(iclass)
	defer C.free(unsafe.Pointer(cIclass))
	cState := C.CString(state)
	defer C.free(unsafe.Pointer(cState))

	C.Epplet_imageclass_paste(cIclass, cState, C.Window(win),
		C.int(x), C.int(y), C.int(h), C.int(w))
}

// TextclassDraw draws text using textclass onto the window at (x, y).
func (win Window) TextclassDraw(tclass, state string, x, y int, txt string) {
	cTclass := C.CString(tclass)
	defer C.free(unsafe.Pointer(cTclass))
	cState := C.CString(state)
	defer C.free(unsafe.Pointer(cState))
	cTxt := C.CString(txt)
	defer C.free(unsafe.Pointer(cTxt))

	C.Epplet_textclass_draw(
		cTclass,
		cState,
		C.Window(win),
		C.int(x),
		C.int(y),
		cTxt,
	)
}
