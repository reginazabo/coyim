package gtk_mock

import "github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/gtki"

type MockGrid struct {
	MockContainer
}

func (*MockGrid) Attach(v1 gtki.Widget, v2, v3, v4, v5 int) {
}
