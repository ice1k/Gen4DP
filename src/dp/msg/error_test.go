package msg

import "testing"

func TestError(t *testing.T) {
	RaiseFormat("My name is %s.", "Van")
	if GetErrors() != "My name is Van." {
		t.FailNow()
	}
}

func TestWarn(t *testing.T) {
	WarnFormat(
		"Hey buddy, I think you get the wrong door, the leather club's %s.",
		"two blocks down")
	if GetWarnings() !=
		"Hey buddy, I think you get the wrong door, the leather club's two blocks down." {
		t.FailNow()
	}
}
