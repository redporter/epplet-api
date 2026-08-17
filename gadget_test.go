package epplet

import (
	"testing"
)

func TestGadgetInterfaceCompliance(t *testing.T) {
	var _ Gadget = (*Button)(nil)
	var _ Gadget = (*Textbox)(nil)
	var _ Gadget = (*DrawingArea)(nil)
	var _ Gadget = (*Slider)(nil)
	var _ Gadget = (*ToggleButton)(nil)
	var _ Gadget = (*PopupButton)(nil)
	var _ Gadget = (*Popup)(nil)
	var _ Gadget = (*ImageGadget)(nil)
	var _ Gadget = (*Label)(nil)
	var _ Gadget = (*ProgressBar)(nil)
}
