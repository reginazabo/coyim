package gdka

import (
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/gotk3/gotk3/gdk"
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/gdki"
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/gliba"
)

type window struct {
	*gliba.Object
	internal *gdk.Window
}

func WrapWindowSimple(v *gdk.Window) *window {
	if v == nil {
		return nil
	}
	return &window{gliba.WrapObjectSimple(v.Object), v}
}

func WrapWindow(v *gdk.Window, e error) (*window, error) {
	return WrapWindowSimple(v), e
}

func UnwrapWindow(v gdki.Window) *gdk.Window {
	if v == nil {
		return nil
	}
	return v.(*window).internal
}

func (v *window) GetDesktop() uint32 {
	return v.internal.GetDesktop()
}

func (v *window) MoveToDesktop(v1 uint32) {
	v.internal.MoveToDesktop(v1)
}
