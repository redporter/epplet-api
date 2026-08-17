package epplet

/*
#include <X11/Xlib.h>

#ifndef _EPPLET_H_GUARD
#define _EPPLET_H_GUARD
#include <epplet.h>
#endif

#include <stdlib.h>


extern void goExposeGateway(void *data, Window win, int x, int y, int w, int h);
static inline void register_expose_bridge(void *handle) {
    Epplet_register_expose_handler(
        (void (*)(void *, Window, int, int, int, int))goExposeGateway,
        handle
    );
}


extern void goMoveResizeGateway(void *data, Window win, int x, int y, int w, int h);
static inline void register_move_resize_bridge(void *handle) {
    Epplet_register_move_resize_handler(
        (void (*)(void *, Window, int, int, int, int))goMoveResizeGateway,
        handle
    );
}


extern void goButtonPressGateway(void *data, Window win, int x, int y, int b);
static inline void register_button_press_bridge(void *handle) {
    Epplet_register_button_press_handler(
        (void (*)(void *, Window, int, int, int))goButtonPressGateway,
        handle
    );
}


extern void goButtonReleaseGateway(void *data, Window win, int x, int y, int b);
static inline void register_button_release_bridge(void *handle) {
    Epplet_register_button_release_handler(
        (void (*)(void *, Window, int, int, int))goButtonReleaseGateway,
        handle
    );
}


extern void goKeyPressGateway(void *data, Window win, char* key);
static inline void register_key_press_bridge(void *handle) {
    Epplet_register_key_press_handler(
        (void (*)(void *, Window, char*))goKeyPressGateway,
        handle
    );
}

extern void goKeyReleaseGateway(void *data, Window win, char* key);
static inline void register_key_release_bridge(void *handle) {
    Epplet_register_key_release_handler(
        (void (*)(void *, Window, char*))goKeyReleaseGateway,
        handle
    );
}


extern void goMouseMotionGateway(void *data, Window win, int x, int y);
static inline void register_mouse_motion_bridge(void *handle) {
    Epplet_register_mouse_motion_handler(
        (void (*)(void *, Window, int, int))goMouseMotionGateway,
        handle
    );
}

extern void goMouseEnterGateway(void *data, Window win);
static inline void register_mouse_enter_bridge(void *handle) {
    Epplet_register_mouse_enter_handler(
        (void (*)(void *, Window))goMouseEnterGateway,
        handle
    );
}

extern void goMouseLeaveGateway(void *data, Window win);
static inline void register_mouse_leave_bridge(void *handle) {
    Epplet_register_mouse_leave_handler(
        (void (*)(void *, Window))goMouseLeaveGateway,
        handle
    );
}


extern void goFocusInGateway(void *data, Window win);
static inline void register_focus_in_bridge(void *handle) {
    Epplet_register_focus_in_handler(
        (void (*)(void *, Window))goFocusInGateway,
        handle
    );
}

extern void goFocusOutGateway(void *data, Window win);
static inline void register_focus_out_bridge(void *handle) {
    Epplet_register_focus_out_handler(
        (void (*)(void *, Window))goFocusOutGateway,
        handle
    );
}

extern void goXEventHandlerGateway(void *data, XEvent *ev);
static inline void register_xevent_bridge(void *handle) {
    Epplet_register_event_handler(
        (void (*)(void *, XEvent *))goXEventHandlerGateway,
        handle
    );
}


extern int goDeleteEventGateway(void *data, Window win);
static inline void register_delete_event_bridge(void *handle) {
    Epplet_register_delete_event_handler(
        (int (*)(void *, Window))goDeleteEventGateway,
        handle
    );
}


*/
import "C"

import "unsafe"
import "runtime/cgo"


type ExposeHandler func(win Window, x, y, w, h int) 

//export goExposeGateway
func goExposeGateway(data unsafe.Pointer, win C.Window, x, y, w, h C.int) {
	if data == nil {
		return
	}

	fh := cgo.Handle(data)
	if fn, ok := fh.Value().(ExposeHandler); ok {
		fn(Window(win), int(x), int(y), int(w), int(h))
	}
}

func RegisterExposeHandler(handler ExposeHandler) func(){
	handle := cgo.NewHandle(handler)
	C.register_expose_bridge(unsafe.Pointer(handle))
	return func() {
		handle.Delete()
	}
}


type MoveResizeHandler func(win Window, x, y, w, h int)

//export goMoveResizeGateway
func goMoveResizeGateway(data unsafe.Pointer, win C.Window, x, y, w, h C.int) {
	if data == nil {
		return
	}

	fh := cgo.Handle(data)
	if fn, ok := fh.Value().(MoveResizeHandler); ok {
		fn(Window(win), int(x), int(y), int(w), int(h))
	}
}

func RegisterMoveResizeHandler(handler MoveResizeHandler) func(){
	handle := cgo.NewHandle(handler)
	C.register_move_resize_bridge(unsafe.Pointer(handle))
	return func() {
		handle.Delete()
	}
}


// ButtonPressHandler is invoked when a mouse button is pressed over an Epplet window.
// Parameters:
//   - win: The target X11 Window ID
//   - x, y: Pointer coordinates relative to the window
//   - button: The mouse button index (1 = Left, 2 = Middle, 3 = Right, 4/5 = Scroll)
type ButtonPressHandler func(win Window, x, y, button int)

//export goButtonPressGateway
func goButtonPressGateway(data unsafe.Pointer, win C.Window, x, y, b C.int) {
	if data == nil {
		return
	}

	h := cgo.Handle(data)
	if fn, ok := h.Value().(ButtonPressHandler); ok {
		fn(Window(win), int(x), int(y), int(b))
	}
}

func RegisterButtonPressHandler(handler ButtonPressHandler) func() {
	handle := cgo.NewHandle(handler)
	C.register_button_press_bridge(unsafe.Pointer(handle))
	return func() {
		handle.Delete()
	}
}


type ButtonReleaseHandler func(win Window, x, y, button int)

//export goButtonReleaseGateway
func goButtonReleaseGateway(data unsafe.Pointer, win C.Window, x, y, b C.int) {
	if data == nil {
		return
	}

	h := cgo.Handle(data)
	if fn, ok := h.Value().(ButtonReleaseHandler); ok {
		fn(Window(win), int(x), int(y), int(b))
	}
}

func RegisterButtonReleaseHandler(handler ButtonReleaseHandler) func() {
	handle := cgo.NewHandle(handler)
	C.register_button_release_bridge(unsafe.Pointer(handle))
	return func() {
		handle.Delete()
	}
}

type KeyPressHandler func(win Window, key string)

//export goKeyPressGateway
func goKeyPressGateway(data unsafe.Pointer, win C.Window, key *C.char) {
	if data == nil {
		return
	}

	h := cgo.Handle(data)
	if fn, ok := h.Value().(KeyPressHandler); ok {
		fn(Window(win), C.GoString(key))
	}
}

func RegisterKeyPressHandler(handler KeyPressHandler) func() {
	handle := cgo.NewHandle(handler)
	C.register_key_press_bridge(unsafe.Pointer(handle))
	return func() {
		handle.Delete()
	}
}


type KeyReleaseHandler func(win Window, key string)

//export goKeyReleaseGateway
func goKeyReleaseGateway(data unsafe.Pointer, win C.Window, key *C.char) {
	if data == nil {
		return
	}

	h := cgo.Handle(data)
	if fn, ok := h.Value().(KeyReleaseHandler); ok {
		fn(Window(win), C.GoString(key))
	}
}

func RegisterKeyReleaseHandler(handler KeyReleaseHandler) func() {
	handle := cgo.NewHandle(handler)
	C.register_key_release_bridge(unsafe.Pointer(handle))
	return func() {
		handle.Delete()
	}
}


type MouseMotionHandler func(win Window, x, y int)

//export goMouseMotionGateway
func goMouseMotionGateway(data unsafe.Pointer, win C.Window, x, y C.int) {
	if data == nil {
		return
	}

	h := cgo.Handle(data)
	if fn, ok := h.Value().(MouseMotionHandler); ok {
		fn(Window(win), int(x), int(y))
	}
}

func RegisterMouseMotionHandler(handler MouseMotionHandler) func() {
	handle := cgo.NewHandle(handler)
	C.register_mouse_motion_bridge(unsafe.Pointer(handle))
	return func() {
		handle.Delete()
	}
}


type MouseEnterHandler func(win Window)

//export goMouseEnterGateway
func goMouseEnterGateway(data unsafe.Pointer, win C.Window) {
	if data == nil {
		return
	}

	h := cgo.Handle(data)
	if fn, ok := h.Value().(MouseEnterHandler); ok {
		fn(Window(win))
	}
}

func RegisterMouseEnterHandler(handler MouseEnterHandler) func() {
	handle := cgo.NewHandle(handler)
	C.register_mouse_enter_bridge(unsafe.Pointer(handle))
	return func() {
		handle.Delete()
	}
}


type MouseLeaveHandler func(win Window)

//export goMouseLeaveGateway
func goMouseLeaveGateway(data unsafe.Pointer, win C.Window) {
	if data == nil {
		return
	}

	h := cgo.Handle(data)
	if fn, ok := h.Value().(MouseLeaveHandler); ok {
		fn(Window(win))
	}
}

func RegisterMouseLeaveHandler(handler MouseEnterHandler) func() {
	handle := cgo.NewHandle(handler)
	C.register_mouse_leave_bridge(unsafe.Pointer(handle))
	return func() {
		handle.Delete()
	}
}


type FocusInHandler func(win Window)

//export goFocusInGateway
func goFocusInGateway(data unsafe.Pointer, win C.Window) {
	if data == nil {
		return
	}

	h := cgo.Handle(data)
	if fn, ok := h.Value().(FocusInHandler); ok {
		fn(Window(win))
	}
}

func RegisterFocusInHandler(handler FocusInHandler) func() {
	handle := cgo.NewHandle(handler)
	C.register_focus_in_bridge(unsafe.Pointer(handle))
	return func() {
		handle.Delete()
	}
}


type FocusOutHandler func(win Window)

//export goFocusOutGateway
func goFocusOutGateway(data unsafe.Pointer, win C.Window) {
	if data == nil {
		return
	}

	h := cgo.Handle(data)
	if fn, ok := h.Value().(FocusOutHandler); ok {
		fn(Window(win))
	}
}

func RegisterFocusOutHandler(handler FocusOutHandler) func() {
	handle := cgo.NewHandle(handler)
	C.register_focus_out_bridge(unsafe.Pointer(handle))
	return func() {
		handle.Delete()
	}
}


type XEvent struct {
	raw *C.XEvent
}

// Type returns the X11 event type constant (e.g., KeyPress, Expose, ClientMessage).
func (e *XEvent) Type() int {
	if e == nil || e.raw == nil {
		return 0
	}
	// XEvent has 'type' as its first field across all union variants
	anyEvent := (*C.XAnyEvent)(unsafe.Pointer(e.raw))
	return int(anyEvent._type)
}

// Window returns the target X11 Window ID if applicable to this event type.
func (e *XEvent) Window() Window {
	if e == nil || e.raw == nil {
		return 0
	}
	// Cast the union to XAnyEvent to safely read window ID
	anyEvent := (*C.XAnyEvent)(unsafe.Pointer(e.raw))
	return Window(anyEvent.window)
}

// RawPointer gives access to the underlying *C.XEvent for low-level Xlib calls.
func (e *XEvent) RawPointer() unsafe.Pointer {
	return unsafe.Pointer(e.raw)
}

// XEventHandler is the user-facing Go callback.
type XEventHandler func(ev *XEvent)

//export goXEventHandlerGateway
func goXEventHandlerGateway(data unsafe.Pointer, ev *C.XEvent) {
	if data == nil {
		return
	}

	h := cgo.Handle(data)
	if fn, ok := h.Value().(XEventHandler); ok {
		fn(&XEvent{raw: ev})
	}
}

func RegisterEventHandler(handler XEventHandler) func() {
	handle := cgo.NewHandle(handler)
	C.register_xevent_bridge(unsafe.Pointer(handle))
	return func() {
		handle.Delete()
	}
}


type DeleteEventHandler func(win Window) int

//export goDeleteEventGateway
func goDeleteEventGateway(data unsafe.Pointer, win C.Window) C.int{
	if data == nil {
		return 1
	}

	h := cgo.Handle(data)
	if fn, ok := h.Value().(DeleteEventHandler); ok {
		return C.int(fn(Window(win)))
	}
	
	return 1
}

func RegisterDeleteEventHandler(handler DeleteEventHandler) func() {
	handle := cgo.NewHandle(handler)
	C.register_delete_event_bridge(unsafe.Pointer(handle))
	return func() {
		handle.Delete()
	}
}
