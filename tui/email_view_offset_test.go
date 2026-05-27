package tui

import (
	"testing"
)

func TestEmailViewOffsets(t *testing.T) {
	ev := &EmailView{}
	ev.SetOffsets(10, 20)

	if ev.rowOffset != 10 {
		t.Errorf("expected rowOffset 10, got %d", ev.rowOffset)
	}
	if ev.columnOffset != 20 {
		t.Errorf("expected columnOffset 20, got %d", ev.columnOffset)
	}
}
