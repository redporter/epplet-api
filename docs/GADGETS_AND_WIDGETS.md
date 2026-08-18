# Gadgets & Widgets Documentation

This document describes the type-safe widget architecture, base `Gadget` interface, and widget catalog in the `epplet-api` Go library.

---

## 1. Type-Safe Gadget Architecture

In the native C Epplet API, all UI widgets are represented as generic untyped `void *` pointers (`Epplet_gadget`).

`epplet-api` wraps widgets into an object-oriented, type-safe architecture:

- **`Gadget` Interface**: Common receiver methods shared across all widgets (`GetX`, `GetY`, `Show`, `Hide`, `Move`, `Destroy`, `DataChanged`, `Draw`).
- **Concrete Gadget Types**: Specialized structs (`*Button`, `*Textbox`, `*Slider`, `*Popup`, `*Label`, `*ProgressBar`, `*DrawingArea`, `*ToggleButton`, `*PopupButton`, `*ImageGadget`) with type-specific receiver methods.

This prevents invalid cross-widget operations (such as calling `Reset()` on a `*Button`) at Go compile-time.

---

## 2. Base `Gadget` Interface

All UI widgets implement the `Gadget` interface:

```go
type Gadget interface {
	GetX() int
	GetY() int
	GetWidth() int
	GetHeight() int
	GetType() GadgetType
	Show()
	Hide()
	Move(x, y int)
	Destroy()
	DataChanged()
	Draw(unOnly, force bool)
	Handle() unsafe.Pointer
}
```

---

## 3. Widget Catalog & Receivers

### A. Buttons (`*Button`)

- **Constructors**:
  - `CreateButton(x, y, w, h int, label, pixmap string, cb GadgetCallback) *Button`
  - `CreateTextButton(label string, x, y, w, h int, cb GadgetCallback) *Button`
  - `CreateStdButton(label, stdType string, x, y int, cb GadgetCallback) *Button`
  - `CreateImageButton(pixmap string, x, y, w, h int, cb GadgetCallback) *Button`
- **Specialized Receiver Methods**:
  - `(b *Button).ChangeLabel(label string)`: Updates button text.
  - `(b *Button).ChangeImage(image string)`: Updates button image.

---

### B. Textbox (`*Textbox`)

- **Constructor**:
  - `CreateTextbox(label, val string, x, y, w, h, maxlen int, cb GadgetCallback) *Textbox`
- **Specialized Receiver Methods**:
  - `(tb *Textbox).Contents() string`: Retrieves input text.
  - `(tb *Textbox).Reset()`: Clears input contents.
  - `(tb *Textbox).ChangeContents(txt string)`: Replaces input contents with `txt`.
  - `(tb *Textbox).InsertContents(txt string)`: Appends `txt` to current contents.

---

### C. Progress Bars (`*ProgressBar`)

- **Constructors**:
  - `CreateHBar(x, y, w, h, dir int, val *int) *ProgressBar`: Horizontal bar.
  - `CreateVBar(x, y, w, h, dir int, val *int) *ProgressBar`: Vertical bar.
- **Data Binding**:
  - `val` points to a host `int` variable (0-100).
  - Call `bar.DataChanged()` after mutating `val` to re-render the progress bar.

---

### D. Sliders (`*Slider`) & Toggle Buttons (`*ToggleButton`)

- **Constructors**:
  - `CreateHSlider(x, y, w, h, step, min, max int, val *int, cb GadgetCallback) *Slider`
  - `CreateVSlider(x, y, w, h, step, min, max int, val *int, cb GadgetCallback) *Slider`
  - `CreateToggleButton(label, image string, x, y, w, h int, val *int, cb GadgetCallback) *ToggleButton`
- **Slider Receivers**:
  - `(s *Slider).HSliderClicked() bool` / `VSliderClicked() bool`

---

### E. Popups (`*Popup`) & PopupButtons (`*PopupButton`)

- **Constructors**:
  - `CreatePopup() *Popup`
  - `CreatePopupButton(label, image string, x, y, w, h int, popup *Popup) *PopupButton`
- **Popup Receivers**:
  - `(p *Popup).AddEntry(label, pixmap string, cb GadgetCallback)`
  - `(p *Popup).AddSizedEntry(label, pixmap string, w, h int, cb GadgetCallback)`
  - `(p *Popup).RemoveEntry(entryNum int)`
  - `(p *Popup).EntryNum() int`
  - `(p *Popup).Pop(win Window)`
- **PopupButton Receivers**:
  - `(pb *PopupButton).ChangePopup(popup *Popup)`
  - `(pb *PopupButton).ChangeLabel(label string)`

---

### F. Image (`*ImageGadget`) & Label (`*Label`)

- **Constructors**:
  - `CreateImage(x, y, w, h int, image string) *ImageGadget`
  - `CreateLabel(x, y int, label string, size int) *Label`
- **Receivers**:
  - `(img *ImageGadget).Change(w, h int, image string)`
  - `(img *ImageGadget).MoveChange(x, y, w, h int, image string)`
  - `(lbl *Label).Change(label string)`
  - `(lbl *Label).MoveChange(x, y int, label string)`

---

### G. DrawingArea (`*DrawingArea`)

- **Constructor**:
  - `CreateDrawingArea(x, y, w, h int) *DrawingArea`
- **Specialized Receiver Methods**:
  - `(da *DrawingArea).DrawingAreaWindow() Window`
  - `(da *DrawingArea).DefaultBindGL() GLXContext`
  - `(da *DrawingArea).MakeCurrent(cx GLXContext)`
  - `(da *DrawingArea).SwapBuffers()`

---

## 4. Code Example: Complete Multi-Widget Dashboard

```go
package main

import (
	"fmt"
	"time"

	"github.com/redporter/epplet-api"
)

var progressVal int = 25

func main() {
	epplet.Init("WidgetDemo", "1.0", "Multi-Widget Dashboard Example", 4, 8, false)

	// 1. Label
	lbl := epplet.CreateLabel(4, 4, "Control Panel", 2)
	lbl.Show()

	// 2. Text Button
	btn := epplet.CreateTextButton("Action", 4, 20, 56, 18, func() {
		fmt.Println("Action button clicked!")
	})
	btn.Show()

	// 3. Textbox
	tb := epplet.CreateTextbox("Input", "Default Text", 4, 42, 56, 18, 0, func() {
		fmt.Println("Submitted text:", tb.Contents())
	})
	tb.Show()

	// 4. Progress Bar & Slider
	bar := epplet.CreateHBar(4, 64, 56, 10, 0, &progressVal)
	bar.Show()

	slider := epplet.CreateHSlider(4, 78, 56, 10, 5, 0, 100, &progressVal, func() {
		bar.DataChanged()
	})
	slider.Show()

	// 5. Popup Menu
	popup := epplet.CreatePopup()
	popup.AddEntry("Option 1", "", func() { fmt.Println("Selected Option 1") })
	popup.AddEntry("Option 2", "", func() { fmt.Println("Selected Option 2") })

	popBtn := epplet.CreatePopupButton("Menu", "", 4, 94, 56, 18, popup)
	popBtn.Show()

	epplet.Show()
	epplet.Loop()
}
```
