package epplet

/*
#include <X11/Xlib.h>

#ifndef _EPPLET_H_GUARD
#define _EPPLET_H_GUARD
#include <epplet.h>
#endif

#include <stdlib.h>

extern void goGadgetGateway(void *data);

static inline Epplet_gadget create_button_bridge(
    char *label, char *image, int x, int y, int w, int h, char *std,
    Window parent, Epplet_gadget pop_parent, void *handle)
{
    return Epplet_create_button(
        label, image, x, y, w, h, std, parent, pop_parent,
        handle ? (void (*)(void *))goGadgetGateway : NULL,
        handle
    );
}

static inline Epplet_gadget create_text_button_bridge(
    char *label, int x, int y, int w, int h, void *handle)
{
    return Epplet_create_text_button(
        label, x, y, w, h,
        handle ? (void (*)(void *))goGadgetGateway : NULL,
        handle
    );
}

static inline Epplet_gadget create_std_button_bridge(
    char *std, int x, int y, void *handle)
{
    return Epplet_create_std_button(
        std, x, y,
        handle ? (void (*)(void *))goGadgetGateway : NULL,
        handle
    );
}

static inline Epplet_gadget create_image_button_bridge(
    char *image, int x, int y, int w, int h, void *handle)
{
    return Epplet_create_image_button(
        image, x, y, w, h,
        handle ? (void (*)(void *))goGadgetGateway : NULL,
        handle
    );
}

static inline Epplet_gadget create_textbox_bridge(
    char *image, char *contents, int x, int y, int w, int h, char size, void *handle)
{
    return Epplet_create_textbox(
        image, contents, x, y, w, h, size,
        handle ? (void (*)(void *))goGadgetGateway : NULL,
        handle
    );
}

static inline Epplet_gadget create_hslider_bridge(
    int x, int y, int len, int min, int max, int step, int jump, int *val, void *handle)
{
    return Epplet_create_hslider(
        x, y, len, min, max, step, jump, val,
        handle ? (void (*)(void *))goGadgetGateway : NULL,
        handle
    );
}

static inline Epplet_gadget create_vslider_bridge(
    int x, int y, int len, int min, int max, int step, int jump, int *val, void *handle)
{
    return Epplet_create_vslider(
        x, y, len, min, max, step, jump, val,
        handle ? (void (*)(void *))goGadgetGateway : NULL,
        handle
    );
}

static inline Epplet_gadget create_togglebutton_bridge(
    char *label, char *pixmap, int x, int y, int w, int h, int *val, void *handle)
{
    return Epplet_create_togglebutton(
        label, pixmap, x, y, w, h, val,
        handle ? (void (*)(void *))goGadgetGateway : NULL,
        handle
    );
}

static inline void add_popup_entry_bridge(
    Epplet_gadget gadget, char *label, char *pixmap, void *handle)
{
    Epplet_add_popup_entry(
        gadget, label, pixmap,
        handle ? (void (*)(void *))goGadgetGateway : NULL,
        handle
    );
}

static inline void add_sized_popup_entry_bridge(
    Epplet_gadget gadget, char *label, char *pixmap, int w, int h, void *handle)
{
    Epplet_add_sized_popup_entry(
        gadget, label, pixmap, w, h,
        handle ? (void (*)(void *))goGadgetGateway : NULL,
        handle
    );
}
*/
import "C"

import (
	"runtime/cgo"
	"unsafe"
)

type GadgetType int

const (
	GadgetButton       GadgetType = 0
	GadgetDrawingArea  GadgetType = 1
	GadgetTextbox      GadgetType = 2
	GadgetHSlider      GadgetType = 3
	GadgetVSlider      GadgetType = 4
	GadgetToggleButton GadgetType = 5
	GadgetPopupButton  GadgetType = 6
	GadgetPopup        GadgetType = 7
	GadgetImage        GadgetType = 8
	GadgetLabel        GadgetType = 9
	GadgetHBar         GadgetType = 10
	GadgetVBar         GadgetType = 11
)

const (
	StdButtonArrowUp     = "ARROW_UP"
	StdButtonArrowDown   = "ARROW_DOWN"
	StdButtonArrowLeft   = "ARROW_LEFT"
	StdButtonArrowRight  = "ARROW_RIGHT"
	StdButtonPlay        = "PLAY"
	StdButtonStop        = "STOP"
	StdButtonPause       = "PAUSE"
	StdButtonPrevious    = "PREVIOUS"
	StdButtonNext        = "NEXT"
	StdButtonEject       = "EJECT"
	StdButtonClose       = "CLOSE"
	StdButtonFastForward = "FAST_FORWARD"
	StdButtonRewind      = "REWIND"
	StdButtonRepeat      = "REPEAT"
	StdButtonSkip        = "SKIP"
	StdButtonHelp        = "HELP"
	StdButtonConfigure   = "CONFIGURE"
)

type GadgetCallback func()

//export goGadgetGateway
func goGadgetGateway(data unsafe.Pointer) {
	if data == nil {
		return
	}

	h := cgo.Handle(data)
	if cb, ok := h.Value().(GadgetCallback); ok && cb != nil {
		cb()
	}
}

// -----------------------------------------------------------------------------
// Global Gadget Functions
// -----------------------------------------------------------------------------

// Redraw redraws all gadgets on the Epplet window.
func Redraw() {
	C.Epplet_redraw()
}

// -----------------------------------------------------------------------------
// Gadget Interface & Base Struct
// -----------------------------------------------------------------------------

type Gadget interface {
	GetX() int
	GetY() int
	GetWidth() int
	GetHeight() int
	GetType() int
	Show()
	Hide()
	Move(x, y int)
	Destroy()
	DataChanged()
	Draw(unOnly, force bool)
	Handle() C.Epplet_gadget
}

type baseGadget struct {
	handle C.Epplet_gadget
}

func (g *baseGadget) GetX() int {
	return int(C.Epplet_gadget_get_x(g.handle))
}

func (g *baseGadget) GetY() int {
	return int(C.Epplet_gadget_get_y(g.handle))
}

func (g *baseGadget) GetWidth() int {
	return int(C.Epplet_gadget_get_width(g.handle))
}

func (g *baseGadget) GetHeight() int {
	return int(C.Epplet_gadget_get_height(g.handle))
}

func (g *baseGadget) GetType() int {
	return int(C.Epplet_gadget_get_type(g.handle))
}

func (g *baseGadget) Handle() C.Epplet_gadget {
	return g.handle
}

func (g *baseGadget) Show() {
	C.Epplet_gadget_show(g.handle)
}

func (g *baseGadget) Hide() {
	C.Epplet_gadget_hide(g.handle)
}

func (g *baseGadget) Move(x, y int) {
	C.Epplet_gadget_move(g.handle, C.int(x), C.int(y))
}

func (g *baseGadget) Destroy() {
	C.Epplet_gadget_destroy(g.handle)
}

func (g *baseGadget) DataChanged() {
	C.Epplet_gadget_data_changed(g.handle)
}

func (g *baseGadget) Draw(unOnly, force bool) {
	unOnlyInt := 0
	if unOnly {
		unOnlyInt = 1
	}
	forceInt := 0
	if force {
		forceInt = 1
	}
	C.Epplet_gadget_draw(g.handle, C.int(unOnlyInt), C.int(forceInt))
}

// -----------------------------------------------------------------------------
// Concrete Gadget Types
// -----------------------------------------------------------------------------

type Button struct{ baseGadget }
type Textbox struct{ baseGadget }
type DrawingArea struct{ baseGadget }
type Slider struct{ baseGadget }
type ToggleButton struct{ baseGadget }
type PopupButton struct{ baseGadget }
type Popup struct{ baseGadget }
type ImageGadget struct{ baseGadget }
type Label struct{ baseGadget }
type ProgressBar struct{ baseGadget }

// Helper function to extract handle safely from a Gadget interface or struct pointer
func getGadgetHandle(g Gadget) C.Epplet_gadget {
	if g == nil {
		return nil
	}
	return g.Handle()
}

// -----------------------------------------------------------------------------
// Button Constructors & Type-Safe Methods
// -----------------------------------------------------------------------------

func CreateButton(label, image string, x, y, w, h int, std string, parent Window, popParent Gadget, cb GadgetCallback) *Button {
	var cLabel, cImage, cStd *C.char
	if label != "" {
		cLabel = C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
	}
	if image != "" {
		cImage = C.CString(image)
		defer C.free(unsafe.Pointer(cImage))
	}
	if std != "" {
		cStd = C.CString(std)
		defer C.free(unsafe.Pointer(cStd))
	}

	var handlePtr unsafe.Pointer
	if cb != nil {
		handle := cgo.NewHandle(cb)
		handlePtr = unsafe.Pointer(handle)
	}

	cGadget := C.create_button_bridge(
		cLabel,
		cImage,
		C.int(x),
		C.int(y),
		C.int(w),
		C.int(h),
		cStd,
		C.Window(parent),
		getGadgetHandle(popParent),
		handlePtr,
	)
	if cGadget == nil {
		return nil
	}
	return &Button{baseGadget{handle: cGadget}}
}

func CreateTextButton(label string, x, y, w, h int, cb GadgetCallback) *Button {
	var cLabel *C.char
	if label != "" {
		cLabel = C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
	}

	var handlePtr unsafe.Pointer
	if cb != nil {
		handle := cgo.NewHandle(cb)
		handlePtr = unsafe.Pointer(handle)
	}

	cGadget := C.create_text_button_bridge(
		cLabel,
		C.int(x),
		C.int(y),
		C.int(w),
		C.int(h),
		handlePtr,
	)
	if cGadget == nil {
		return nil
	}
	return &Button{baseGadget{handle: cGadget}}
}

func CreateStdButton(std string, x, y int, cb GadgetCallback) *Button {
	var cStd *C.char
	if std != "" {
		cStd = C.CString(std)
		defer C.free(unsafe.Pointer(cStd))
	}

	var handlePtr unsafe.Pointer
	if cb != nil {
		handle := cgo.NewHandle(cb)
		handlePtr = unsafe.Pointer(handle)
	}

	cGadget := C.create_std_button_bridge(
		cStd,
		C.int(x),
		C.int(y),
		handlePtr,
	)
	if cGadget == nil {
		return nil
	}
	return &Button{baseGadget{handle: cGadget}}
}

func CreateImageButton(image string, x, y, w, h int, cb GadgetCallback) *Button {
	var cImage *C.char
	if image != "" {
		cImage = C.CString(image)
		defer C.free(unsafe.Pointer(cImage))
	}

	var handlePtr unsafe.Pointer
	if cb != nil {
		handle := cgo.NewHandle(cb)
		handlePtr = unsafe.Pointer(handle)
	}

	cGadget := C.create_image_button_bridge(
		cImage,
		C.int(x),
		C.int(y),
		C.int(w),
		C.int(h),
		handlePtr,
	)
	if cGadget == nil {
		return nil
	}
	return &Button{baseGadget{handle: cGadget}}
}

func (b *Button) ChangeLabel(label string) {
	if b == nil || b.handle == nil {
		return
	}
	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))
	C.Epplet_change_button_label(b.handle, cLabel)
}

func (b *Button) ChangeImage(image string) {
	if b == nil || b.handle == nil {
		return
	}
	cImage := C.CString(image)
	defer C.free(unsafe.Pointer(cImage))
	C.Epplet_change_button_image(b.handle, cImage)
}

// -----------------------------------------------------------------------------
// Textbox Constructor & Type-Safe Methods
// -----------------------------------------------------------------------------

func CreateTextbox(image, contents string, x, y, w, h int, size int, cb GadgetCallback) *Textbox {
	var cImage, cContents *C.char
	if image != "" {
		cImage = C.CString(image)
		defer C.free(unsafe.Pointer(cImage))
	}
	if contents != "" {
		cContents = C.CString(contents)
		defer C.free(unsafe.Pointer(cContents))
	}

	var handlePtr unsafe.Pointer
	if cb != nil {
		handle := cgo.NewHandle(cb)
		handlePtr = unsafe.Pointer(handle)
	}

	cGadget := C.create_textbox_bridge(
		cImage,
		cContents,
		C.int(x),
		C.int(y),
		C.int(w),
		C.int(h),
		C.char(size),
		handlePtr,
	)
	if cGadget == nil {
		return nil
	}
	return &Textbox{baseGadget{handle: cGadget}}
}

func (tb *Textbox) Contents() string {
	if tb == nil || tb.handle == nil {
		return ""
	}
	cStr := C.Epplet_textbox_contents(tb.handle)
	if cStr == nil {
		return ""
	}
	return C.GoString(cStr)
}

func (tb *Textbox) Reset() {
	if tb == nil || tb.handle == nil {
		return
	}
	C.Epplet_reset_textbox(tb.handle)
}

func (tb *Textbox) ChangeContents(newContents string) {
	if tb == nil || tb.handle == nil {
		return
	}
	cStr := C.CString(newContents)
	defer C.free(unsafe.Pointer(cStr))
	C.Epplet_change_textbox(tb.handle, cStr)
}

func (tb *Textbox) InsertContents(newContents string) {
	if tb == nil || tb.handle == nil {
		return
	}
	cStr := C.CString(newContents)
	defer C.free(unsafe.Pointer(cStr))
	C.Epplet_textbox_insert(tb.handle, cStr)
}

// -----------------------------------------------------------------------------
// DrawingArea Constructor & Type-Safe Methods
// -----------------------------------------------------------------------------

func CreateDrawingArea(x, y, w, h int) *DrawingArea {
	cGadget := C.Epplet_create_drawingarea(C.int(x), C.int(y), C.int(w), C.int(h))
	if cGadget == nil {
		return nil
	}
	return &DrawingArea{baseGadget{handle: cGadget}}
}

func (da *DrawingArea) DrawingAreaWindow() Window {
	if da == nil || da.handle == nil {
		return 0
	}
	return Window(C.Epplet_get_drawingarea_window(da.handle))
}

// -----------------------------------------------------------------------------
// Slider Constructors & Type-Safe Methods
// -----------------------------------------------------------------------------

func CreateHSlider(x, y, length, min, max, step, jump int, val *int, cb GadgetCallback) *Slider {
	var handlePtr unsafe.Pointer
	if cb != nil {
		handle := cgo.NewHandle(cb)
		handlePtr = unsafe.Pointer(handle)
	}

	cGadget := C.create_hslider_bridge(
		C.int(x), C.int(y), C.int(length),
		C.int(min), C.int(max), C.int(step), C.int(jump),
		(*C.int)(unsafe.Pointer(val)),
		handlePtr,
	)
	if cGadget == nil {
		return nil
	}
	return &Slider{baseGadget{handle: cGadget}}
}

func CreateVSlider(x, y, length, min, max, step, jump int, val *int, cb GadgetCallback) *Slider {
	var handlePtr unsafe.Pointer
	if cb != nil {
		handle := cgo.NewHandle(cb)
		handlePtr = unsafe.Pointer(handle)
	}

	cGadget := C.create_vslider_bridge(
		C.int(x), C.int(y), C.int(length),
		C.int(min), C.int(max), C.int(step), C.int(jump),
		(*C.int)(unsafe.Pointer(val)),
		handlePtr,
	)
	if cGadget == nil {
		return nil
	}
	return &Slider{baseGadget{handle: cGadget}}
}

func (s *Slider) HSliderClicked() bool {
	if s == nil || s.handle == nil {
		return false
	}
	return C.Epplet_get_hslider_clicked(s.handle) != 0
}

func (s *Slider) VSliderClicked() bool {
	if s == nil || s.handle == nil {
		return false
	}
	return C.Epplet_get_vslider_clicked(s.handle) != 0
}

// -----------------------------------------------------------------------------
// ToggleButton Constructor
// -----------------------------------------------------------------------------

func CreateToggleButton(label, pixmap string, x, y, w, h int, val *int, cb GadgetCallback) *ToggleButton {
	var cLabel, cPixmap *C.char
	if label != "" {
		cLabel = C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
	}
	if pixmap != "" {
		cPixmap = C.CString(pixmap)
		defer C.free(unsafe.Pointer(cPixmap))
	}

	var handlePtr unsafe.Pointer
	if cb != nil {
		handle := cgo.NewHandle(cb)
		handlePtr = unsafe.Pointer(handle)
	}

	cGadget := C.create_togglebutton_bridge(
		cLabel, cPixmap,
		C.int(x), C.int(y), C.int(w), C.int(h),
		(*C.int)(unsafe.Pointer(val)),
		handlePtr,
	)
	if cGadget == nil {
		return nil
	}
	return &ToggleButton{baseGadget{handle: cGadget}}
}

// -----------------------------------------------------------------------------
// Popup & PopupButton Constructors & Type-Safe Methods
// -----------------------------------------------------------------------------

func CreatePopup() *Popup {
	cGadget := C.Epplet_create_popup()
	if cGadget == nil {
		return nil
	}
	return &Popup{baseGadget{handle: cGadget}}
}

func (p *Popup) AddEntry(label, pixmap string, cb GadgetCallback) {
	if p == nil || p.handle == nil {
		return
	}
	var cLabel, cPixmap *C.char
	if label != "" {
		cLabel = C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
	}
	if pixmap != "" {
		cPixmap = C.CString(pixmap)
		defer C.free(unsafe.Pointer(cPixmap))
	}

	var handlePtr unsafe.Pointer
	if cb != nil {
		handle := cgo.NewHandle(cb)
		handlePtr = unsafe.Pointer(handle)
	}

	C.add_popup_entry_bridge(p.handle, cLabel, cPixmap, handlePtr)
}

func (p *Popup) AddSizedEntry(label, pixmap string, w, h int, cb GadgetCallback) {
	if p == nil || p.handle == nil {
		return
	}
	var cLabel, cPixmap *C.char
	if label != "" {
		cLabel = C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
	}
	if pixmap != "" {
		cPixmap = C.CString(pixmap)
		defer C.free(unsafe.Pointer(cPixmap))
	}

	var handlePtr unsafe.Pointer
	if cb != nil {
		handle := cgo.NewHandle(cb)
		handlePtr = unsafe.Pointer(handle)
	}

	C.add_sized_popup_entry_bridge(p.handle, cLabel, cPixmap, C.int(w), C.int(h), handlePtr)
}

func (p *Popup) RemoveEntry(entryNum int) {
	if p == nil || p.handle == nil {
		return
	}
	C.Epplet_remove_popup_entry(p.handle, C.int(entryNum))
}

func (p *Popup) EntryNum() int {
	if p == nil || p.handle == nil {
		return 0
	}
	return int(C.Epplet_popup_entry_num(p.handle))
}

func (p *Popup) Pop(win Window) {
	if p == nil || p.handle == nil {
		return
	}
	C.Epplet_pop_popup(p.handle, C.Window(win))
}

func CreatePopupButton(label, image string, x, y, w, h int, std string, popup *Popup) *PopupButton {
	var cLabel, cImage, cStd *C.char
	if label != "" {
		cLabel = C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
	}
	if image != "" {
		cImage = C.CString(image)
		defer C.free(unsafe.Pointer(cImage))
	}
	if std != "" {
		cStd = C.CString(std)
		defer C.free(unsafe.Pointer(cStd))
	}

	var popupHandle C.Epplet_gadget
	if popup != nil {
		popupHandle = popup.handle
	}

	cGadget := C.Epplet_create_popupbutton(
		cLabel, cImage,
		C.int(x), C.int(y), C.int(w), C.int(h),
		cStd, popupHandle,
	)
	if cGadget == nil {
		return nil
	}
	return &PopupButton{baseGadget{handle: cGadget}}
}

func (pb *PopupButton) ChangePopup(popup *Popup) {
	if pb == nil || pb.handle == nil {
		return
	}
	var popupHandle C.Epplet_gadget
	if popup != nil {
		popupHandle = popup.handle
	}
	C.Epplet_change_popbutton_popup(pb.handle, popupHandle)
}

func (pb *PopupButton) ChangeLabel(label string) {
	if pb == nil || pb.handle == nil {
		return
	}
	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))
	C.Epplet_change_popbutton_label(pb.handle, cLabel)
}

// -----------------------------------------------------------------------------
// Image Constructor & Type-Safe Methods
// -----------------------------------------------------------------------------

func CreateImage(x, y, w, h int, image string) *ImageGadget {
	var cImage *C.char
	if image != "" {
		cImage = C.CString(image)
		defer C.free(unsafe.Pointer(cImage))
	}

	cGadget := C.Epplet_create_image(C.int(x), C.int(y), C.int(w), C.int(h), cImage)
	if cGadget == nil {
		return nil
	}
	return &ImageGadget{baseGadget{handle: cGadget}}
}

func (img *ImageGadget) Change(w, h int, image string) {
	if img == nil || img.handle == nil {
		return
	}
	var cImage *C.char
	if image != "" {
		cImage = C.CString(image)
		defer C.free(unsafe.Pointer(cImage))
	}
	C.Epplet_change_image(img.handle, C.int(w), C.int(h), cImage)
}

func (img *ImageGadget) MoveChange(x, y, w, h int, image string) {
	if img == nil || img.handle == nil {
		return
	}
	var cImage *C.char
	if image != "" {
		cImage = C.CString(image)
		defer C.free(unsafe.Pointer(cImage))
	}
	C.Epplet_move_change_image(img.handle, C.int(x), C.int(y), C.int(w), C.int(h), cImage)
}

// -----------------------------------------------------------------------------
// Label Constructor & Type-Safe Methods
// -----------------------------------------------------------------------------

func CreateLabel(x, y int, label string, size int) *Label {
	var cLabel *C.char
	if label != "" {
		cLabel = C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
	}

	cGadget := C.Epplet_create_label(C.int(x), C.int(y), cLabel, C.char(size))
	if cGadget == nil {
		return nil
	}
	return &Label{baseGadget{handle: cGadget}}
}

func (lbl *Label) Change(label string) {
	if lbl == nil || lbl.handle == nil {
		return
	}
	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))
	C.Epplet_change_label(lbl.handle, cLabel)
}

func (lbl *Label) MoveChange(x, y int, label string) {
	if lbl == nil || lbl.handle == nil {
		return
	}
	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))
	C.Epplet_move_change_label(lbl.handle, C.int(x), C.int(y), cLabel)
}

// -----------------------------------------------------------------------------
// Progress Bar (HBar / VBar) Constructors
// -----------------------------------------------------------------------------

func CreateHBar(x, y, w, h, dir int, val *int) *ProgressBar {
	cGadget := C.Epplet_create_hbar(
		C.int(x), C.int(y), C.int(w), C.int(h),
		C.char(dir), (*C.int)(unsafe.Pointer(val)),
	)
	if cGadget == nil {
		return nil
	}
	return &ProgressBar{baseGadget{handle: cGadget}}
}

func CreateVBar(x, y, w, h, dir int, val *int) *ProgressBar {
	cGadget := C.Epplet_create_vbar(
		C.int(x), C.int(y), C.int(w), C.int(h),
		C.char(dir), (*C.int)(unsafe.Pointer(val)),
	)
	if cGadget == nil {
		return nil
	}
	return &ProgressBar{baseGadget{handle: cGadget}}
}
