package epplet

/*
#include <X11/Xlib.h>

#ifndef _EPPLET_H_GUARD
#define _EPPLET_H_GUARD
#include <epplet.h>
#endif

#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

// -----------------------------------------------------------------------------
// 2D Primitive Drawing Functions & Window Methods
// -----------------------------------------------------------------------------

// DrawLine draws a line from (x1, y1) to (x2, y2) on window win in (r, g, b) color.
func DrawLine(win Window, x1, y1, x2, y2, r, g, b int) {
	C.Epplet_draw_line(
		C.Window(win),
		C.int(x1), C.int(y1),
		C.int(x2), C.int(y2),
		C.int(r), C.int(g), C.int(b),
	)
}

// DrawBox draws a filled box at (x, y) of size (w, h) on window win in (r, g, b) color.
func DrawBox(win Window, x, y, w, h, r, g, b int) {
	C.Epplet_draw_box(
		C.Window(win),
		C.int(x), C.int(y),
		C.int(w), C.int(h),
		C.int(r), C.int(g), C.int(b),
	)
}

// DrawOutline draws a box outline at (x, y) of size (w, h) on window win in (r, g, b) color.
func DrawOutline(win Window, x, y, w, h, r, g, b int) {
	C.Epplet_draw_outline(
		C.Window(win),
		C.int(x), C.int(y),
		C.int(w), C.int(h),
		C.int(r), C.int(g), C.int(b),
	)
}

// GetColor returns the X11 pixel color value for RGB color (r, g, b).
func GetColor(r, g, b int) int {
	return int(C.Epplet_get_color(C.int(r), C.int(g), C.int(b)))
}

// PasteImage pastes an image file onto window win at (x, y) at its original size.
func PasteImage(image string, win Window, x, y int) {
	cImage := C.CString(image)
	defer C.free(unsafe.Pointer(cImage))
	C.Epplet_paste_image(cImage, C.Window(win), C.int(x), C.int(y))
}

// PasteImageSize pastes an image file onto window win at (x, y) with size (w, h).
func PasteImageSize(image string, win Window, x, y, w, h int) {
	cImage := C.CString(image)
	defer C.free(unsafe.Pointer(cImage))
	C.Epplet_paste_image_size(cImage, C.Window(win), C.int(x), C.int(y), C.int(w), C.int(h))
}

// Sync flushes and synchronizes all pending X11 drawing operations.
func Sync() {
	C.Esync()
}

// Window receiver methods for drawing primitives
func (win Window) DrawLine(x1, y1, x2, y2, r, g, b int) {
	DrawLine(win, x1, y1, x2, y2, r, g, b)
}

func (win Window) DrawBox(x, y, w, h, r, g, b int) {
	DrawBox(win, x, y, w, h, r, g, b)
}

func (win Window) DrawOutline(x, y, w, h, r, g, b int) {
	DrawOutline(win, x, y, w, h, r, g, b)
}

func (win Window) PasteImage(image string, x, y int) {
	PasteImage(image, win, x, y)
}

func (win Window) PasteImageSize(image string, x, y, w, h int) {
	PasteImageSize(image, win, x, y, w, h)
}

// -----------------------------------------------------------------------------
// RGB Buffer API
// -----------------------------------------------------------------------------

// RGBBuf represents a raw 32-bit RGBA drawing buffer.
type RGBBuf struct {
	ptr C.RGB_buf
	w   int
	h   int
}

// MakeRGBBuf allocates a new RGB drawing buffer of width w and height h.
func MakeRGBBuf(w, h int) *RGBBuf {
	cBuf := C.Epplet_make_rgb_buf(C.int(w), C.int(h))
	if cBuf == nil {
		return nil
	}
	return &RGBBuf{ptr: cBuf, w: w, h: h}
}

// Width returns the width of the RGB buffer.
func (buf *RGBBuf) Width() int {
	if buf == nil {
		return 0
	}
	return buf.w
}

// Height returns the height of the RGB buffer.
func (buf *RGBBuf) Height() int {
	if buf == nil {
		return 0
	}
	return buf.h
}

// Data returns a Go byte slice pointing to the raw ARGB pixel memory of the buffer.
func (buf *RGBBuf) Data() []byte {
	if buf == nil || buf.ptr == nil {
		return nil
	}
	rawPtr := C.Epplet_get_rgb_pointer(buf.ptr)
	if rawPtr == nil {
		return nil
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(rawPtr)), buf.w*buf.h*4)
}

// Paste renders and pastes the RGB buffer onto window win at (x, y).
func (buf *RGBBuf) Paste(win Window, x, y int) {
	if buf == nil || buf.ptr == nil {
		return
	}
	C.Epplet_paste_buf(buf.ptr, C.Window(win), C.int(x), C.int(y))
}

// Free releases the underlying C RGB buffer memory.
func (buf *RGBBuf) Free() {
	if buf == nil || buf.ptr == nil {
		return
	}
	C.Epplet_free_rgb_buf(buf.ptr)
	buf.ptr = nil
}

func (win Window) PasteBuf(buf *RGBBuf, x, y int) {
	if buf == nil {
		return
	}
	buf.Paste(win, x, y)
}
