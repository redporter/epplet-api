package epplet

/*
#define HAVE_GLX 1
#include <X11/Xlib.h>
#include <GL/glx.h>
#include <dlfcn.h>
#include <stdlib.h>

#ifndef _EPPLET_H_GUARD
#define _EPPLET_H_GUARD
#include <epplet.h>
#endif

typedef GLXContext (*fn_bind_double_GL)(
    Epplet_gadget da, int red, int blue, int green, int alpha,
    int aux_buffers, int depth, int stencil, int accum_red,
    int accum_green, int accum_blue, int accum_alpha);
typedef GLXContext (*fn_bind_single_GL)(
    Epplet_gadget da, int red, int blue, int green, int alpha,
    int aux_buffers, int depth, int stencil, int accum_red,
    int accum_green, int accum_blue, int accum_alpha);
typedef GLXContext (*fn_default_bind_GL)(Epplet_gadget da);
typedef void (*fn_unbind_GL)(GLXContext cx);

static inline int has_glx_support(void) {
    void *h = dlopen("libepplet_glx.so", RTLD_LAZY | RTLD_GLOBAL);
    if (!h) {
        h = RTLD_DEFAULT;
    }
    return dlsym(h, "Epplet_default_bind_GL") != NULL;
}

static inline GLXContext bridge_bind_double_GL(
    Epplet_gadget da, int red, int blue, int green, int alpha,
    int aux_buffers, int depth, int stencil, int accum_red,
    int accum_green, int accum_blue, int accum_alpha)
{
    void *h = dlopen("libepplet_glx.so", RTLD_LAZY | RTLD_GLOBAL);
    if (!h) {
        h = RTLD_DEFAULT;
    }
    fn_bind_double_GL fn = (fn_bind_double_GL)dlsym(h, "Epplet_bind_double_GL");
    if (fn) {
        return fn(da, red, blue, green, alpha, aux_buffers, depth, stencil, accum_red, accum_green, accum_blue, accum_alpha);
    }
    return NULL;
}

static inline GLXContext bridge_bind_single_GL(
    Epplet_gadget da, int red, int blue, int green, int alpha,
    int aux_buffers, int depth, int stencil, int accum_red,
    int accum_green, int accum_blue, int accum_alpha)
{
    void *h = dlopen("libepplet_glx.so", RTLD_LAZY | RTLD_GLOBAL);
    if (!h) {
        h = RTLD_DEFAULT;
    }
    fn_bind_single_GL fn = (fn_bind_single_GL)dlsym(h, "Epplet_bind_single_GL");
    if (fn) {
        return fn(da, red, blue, green, alpha, aux_buffers, depth, stencil, accum_red, accum_green, accum_blue, accum_alpha);
    }
    return NULL;
}

static inline GLXContext bridge_default_bind_GL(Epplet_gadget da) {
    void *h = dlopen("libepplet_glx.so", RTLD_LAZY | RTLD_GLOBAL);
    if (!h) {
        h = RTLD_DEFAULT;
    }
    fn_default_bind_GL fn = (fn_default_bind_GL)dlsym(h, "Epplet_default_bind_GL");
    if (fn) {
        return fn(da);
    }
    return NULL;
}

static inline void bridge_unbind_GL(GLXContext cx) {
    void *h = dlopen("libepplet_glx.so", RTLD_LAZY | RTLD_GLOBAL);
    if (!h) {
        h = RTLD_DEFAULT;
    }
    fn_unbind_GL fn = (fn_unbind_GL)dlsym(h, "Epplet_unbind_GL");
    if (fn) {
        fn(cx);
    }
}

static inline void glx_make_current(Window win, GLXContext cx) {
    Display *dpy = Epplet_get_display();
    if (dpy && win && cx) {
        glXMakeCurrent(dpy, (GLXDrawable)win, cx);
    }
}

static inline void glx_swap_buffers(Window win) {
    Display *dpy = Epplet_get_display();
    if (dpy && win) {
        glXSwapBuffers(dpy, (GLXDrawable)win);
    }
}
*/
import "C"

import "unsafe"

// GLXContext represents an OpenGL GLX context handle.
type GLXContext uintptr

// HasGLXSupport checks whether the system libEpplet library was compiled with OpenGL/GLX support.
func HasGLXSupport() bool {
	return C.has_glx_support() != 0
}

// BindDoubleGL creates a double-buffered GLX context for this DrawingArea with specific bit-depth requirements.
func (da *DrawingArea) BindDoubleGL(red, blue, green, alpha, auxBuffers, depth, stencil, accumRed, accumGreen, accumBlue, accumAlpha int) GLXContext {
	if da == nil || da.handle == nil {
		return 0
	}
	cx := C.bridge_bind_double_GL(
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
	cx := C.bridge_bind_single_GL(
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
	cx := C.bridge_default_bind_GL(da.handle)
	return GLXContext(uintptr(unsafe.Pointer(cx)))
}

// MakeCurrent sets the GLX context as active for this DrawingArea.
func (da *DrawingArea) MakeCurrent(cx GLXContext) {
	if da == nil || da.handle == nil || cx == 0 {
		return
	}
	win := da.DrawingAreaWindow()
	if win == 0 {
		return
	}
	C.glx_make_current(C.Window(win), (C.GLXContext)(unsafe.Pointer(cx)))
}

// MakeCurrent sets the GLX context as active for the given window.
func (win Window) MakeCurrent(cx GLXContext) {
	if win == 0 || cx == 0 {
		return
	}
	C.glx_make_current(C.Window(win), (C.GLXContext)(unsafe.Pointer(cx)))
}

// Unbind destroys (unbinds) the GLX context.
func (cx GLXContext) Unbind() {
	if cx == 0 {
		return
	}
	C.bridge_unbind_GL((C.GLXContext)(unsafe.Pointer(cx)))
}

// UnbindGL destroys (unbinds) a GLX context.
func UnbindGL(cx GLXContext) {
	cx.Unbind()
}

// SwapBuffers swaps the front and back GL buffers for this DrawingArea's window.
func (da *DrawingArea) SwapBuffers() {
	if da == nil || da.handle == nil {
		return
	}
	win := da.DrawingAreaWindow()
	if win == 0 {
		return
	}
	C.glx_swap_buffers(C.Window(win))
}

// SwapBuffers swaps the front and back GL buffers for the given window.
func (win Window) SwapBuffers() {
	C.glx_swap_buffers(C.Window(win))
}
