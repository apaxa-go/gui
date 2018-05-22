package cocoa

import "testing"

func TestFont_Size(t *testing.T) {
	const name = "San Francisco"
	const size = 13
	f, ok := NewFont(name, size)
	defer f.Release()
	if !ok {
		t.Error("unable to create font")
	}
	if s := f.Size(); s != size {
		t.Errorf("expected size %v, but got %v", size, s)
	}
}
