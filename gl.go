package epplet

/*
#define HAVE_GLX 1
#include <X11/Xlib.h>
#include <GL/glx.h>

#ifndef _EPPLET_H_GUARD
#define _EPPLET_H_GUARD
#include <epplet.h>
#endif

#include <stdlib.h>
*/
import "C"

import "unsafe"

// GLXContext represents an OpenGL GLX context handle.
type GLXContext uintptr

// BindDoubleGL creates a double-buffered GLX context for this DrawingArea with specific bit-depth requirements.
func (da *DrawingArea) BindDoubleGL(red, blue, green, alpha, auxBuffers, depth, stencil, accumRed, accumGreen, accumBlue, accumAlpha int) GLXContext {
	if da == nil || da.handle == nil {
		return 0
	}
	cx := C.Epplet_bind_double_GL(
		da.handle,
		C.int(red), C.int(blue), C.int(green), C.int(alpha),
		C.int(auxBuffers), C.int(depth), C.int(stencil),
		C.int(accumRed), C.int(accumGreen), C.int(accumBlue), C.int(accumAlpha),
	)
	return GLXContext(uintptr(unsafe.Pointer(cx)))
}

// BindSingleGL creates a single-buffered GLX context for this DrawingArea with specific bit-depth requirements.
func (da *DrawingArea) BindSingleGL(red, blue, green, alpha, auxBuffers, depth, stencil, accumRed, accumGreen, accumBlue, accumAlpha int) GLXContext {
	if da == nil || da.handle == nil {
		return 0
	}
	cx := C.Epplet_bind_single_GL(
		da.handle,
		C.int(red), C.int(blue), C.int(green), C.int(alpha),
		C.int(auxBuffers), C.int(depth), C.int(stencil),
		C.int(accumRed), C.int(accumGreen), C.int(accumBlue), C.int(accumAlpha),
	)
	return GLXContext(uintptr(unsafe.Pointer(cx)))
}

// DefaultBindGL creates a basic RGB double-buffered GLX context with minimal depth buffer for this DrawingArea.
func (da *DrawingArea) DefaultBindGL() GLXContext {
	if da == nil || da.handle == nil {
		return 0
	}
	cx := C.Epplet_default_bind_GL(da.handle)
	return GLXContext(uintptr(unsafe.Pointer(cx)))
}

// Unbind destroys (unbinds) the GLX context.
func (cx GLXContext) Unbind() {
	if cx == 0 {
		return
	}
	C.Epplet_unbind_GL(C.GLXContext(unsafe.Pointer(uintptr(cx))))
}

// UnbindGL destroys (unbinds) a GLX context.
func UnbindGL(cx GLXContext) {
	cx.Unbind()
}
