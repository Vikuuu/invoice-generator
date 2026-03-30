package gui

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type numericalEntry struct {
	widget.Entry
	maxLen int
}

func newNumericalEntry(maxLen int) *numericalEntry {
	entry := &numericalEntry{maxLen: maxLen}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *numericalEntry) TypedRune(r rune) {
	if !((r >= '0' && r <= '9') || r == ',') {
		return
	}
	if e.maxLen > 0 && len([]rune(e.Text)) >= e.maxLen {
		return
	}
	e.Entry.TypedRune(r)
}

func (e *numericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.Atoi(content); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

// Allows the mobile Operating System to open the
// number keyboard when pressed on this entry
func (e *numericalEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}

func (e *numericalEntry) SetNumber(n int64) {
	s := strconv.FormatInt(n, 10)
	e.SetText(s)
}
