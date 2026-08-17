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

import "unsafe"
import "runtime/cgo"
import "time"
import "sync"

func Init(name, version, info string, w, h int, vertical bool){
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

	//XXX pass argc and argv to e16 here	
	C.Epplet_Init(cName, cVersion, cInfo, C.int(w), C.int(h), C.int(0), nil, C.char(cVert))
}

func Cleanup(){
	C.Epplet_cleanup()
}

func Show(){
	C.Epplet_show()
}

func Remember(){
	C.Epplet_remember()
}

func Unremember(){
	C.Epplet_unremember()
}

func Loop(){
	C.Epplet_Loop()
}



// -----------------------------------------------------------------------------
type Display struct{
	ptr *C.Display
	//here goes pointer to C display structure
	//probably requires investigation which type it is
}

func GetDisplay() *Display {
	cDisp := C.Epplet_get_display()
	if cDisp == nil {
		return nil
	}
	return &Display{ptr: cDisp}
}



// ----------------------- IPC -------------------------------------------------
func SendIPC(msg string){
	cStr := C.CString(msg)
	defer C.free(unsafe.Pointer(cStr))
	C.Epplet_send_ipc(cStr)
}

// BlockForIPC blocks until an IPC message is received and returns the message string.
// If the call fails or returns NULL, an empty string is returned.
func BlockForIPC() string {
	cMsg := C.Epplet_wait_for_ipc()
	if cMsg == nil {
		return ""
	}
	return C.GoString(cMsg)
}

// WaitForIPCAsync waits for the next IPC message in a background goroutine 
// and delivers the result over a Go string channel.
func WaitForIPCAsync() <-chan string {
	ch := make(chan string, 1)
	go func() {
		cMsg := C.Epplet_wait_for_ipc()
		if cMsg != nil {
			ch <- C.GoString(cMsg)
		} else {
			close(ch)
		}
	}()
	return ch
}




// -----------------------------------------------------------------------------

type Pixmap uintptr

type ImageclassPixmaps struct {
	Pixmap Pixmap
	Mask   Pixmap
}

type Size struct {
	Width  int
	Height int
}

func ImageclassGetPixmaps(iclass string, state string, w, h int) ImageclassPixmaps {
	cIclass := C.CString(iclass)
	defer C.free(unsafe.Pointer(cIclass))
	cState := C.CString(state)
	defer C.free(unsafe.Pointer(cState))

	var cPmap C.Pixmap
	var cMask C.Pixmap

	C.Epplet_imageclass_get_pixmaps(
		cIclass,
		cState,
		&cPmap,
		&cMask,
		C.int(w),
		C.int(h),
	)

	return ImageclassPixmaps{
		Pixmap: Pixmap(cPmap),
		Mask:   Pixmap(cMask),
	}
}

func TextclassGetSize(tclass string, x, h int, txt string) Size{
	cTclass := C.CString(tclass)
	defer C.free(unsafe.Pointer(cTclass))
	cTxt := C.CString(txt)
	defer C.free(unsafe.Pointer(cTxt))
	
	var cW C.int
	var cH C.int	

	C.Epplet_textclass_get_size(
		cTclass,
		&cW,
		&cH,
		cTxt,
	)
	
	return Size{
		Height: int(cH),
		Width: int(cW),
	}
}


// -----------------------------------------------------------------------------

type TimerCallback func()

var (
	timerMu     sync.Mutex
	timerHandles = make(map[string]cgo.Handle)
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

func Timer(cb TimerCallback, d time.Duration, name string){
    cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	timerMu.Lock()
	// If a timer with this name already exists, delete its old handle
	if oldHandle, exists := timerHandles[name]; exists {
		oldHandle.Delete()
	}

	handle := cgo.NewHandle(cb)
	timerHandles[name] = handle
	timerMu.Unlock()

	seconds := d.Seconds()
	C.create_timer_bridge(unsafe.Pointer(handle), C.double(seconds), cName)
}


func RemoveTimer(name string){
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

