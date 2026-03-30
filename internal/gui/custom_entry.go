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

type PriceEntry struct {
	widget.BaseWidget

	wholePart *numericalEntry
	fracPart  *numericalEntry
	divider   *canvas.Line
	OnChanged func(whole, frac string)
}

func NewPriceEntry() *PriceEntry {
	p := &PriceEntry{}
	p.ExtendBaseWidget(p)

	p.wholePart = newNumericalEntry(0)
	p.wholePart.SetPlaceHolder("0")

	p.fracPart = newNumericalEntry(2)
	p.fracPart.SetPlaceHolder("00")

	p.wholePart.OnChanged = func(s string) {
		if p.OnChanged != nil {
			p.OnChanged(s, p.fracPart.Text)
		}
	}
	p.fracPart.OnChanged = func(s string) {
		if p.OnChanged != nil {
			p.OnChanged(p.wholePart.Text, s)
		}
	}
	return p
}

func (p *PriceEntry) CreateRenderer() fyne.WidgetRenderer {
	p.ExtendBaseWidget(p)

	divider := canvas.NewLine(color.NRGBA{R: 150, G: 150, B: 150, A: 255})
	divider.StrokeWidth = 1.5
	p.divider = divider

	c := container.NewBorder(
		nil, nil, nil, nil,
		container.New(&priceLayout{}, p.wholePart, divider, p.fracPart),
	)
	return widget.NewSimpleRenderer(c)
}

func (p *PriceEntry) Value() string {
	whole := p.wholePart.Text
	frac := p.fracPart.Text
	if whole == "" {
		whole = "0"
	}
	if frac == "" {
		frac = "00"
	} else if len(frac) == 1 {
		frac = frac + "0"
	}
	return whole + "." + frac
}

func (p *PriceEntry) SetValue(whole, frac string) {
	p.wholePart.SetText(whole)
	if len(frac) > 2 {
		frac = frac[:2]
	}
	p.fracPart.SetText(frac)
}

const (
	priceLayoutDividerGap = float32(8)
	priceLayoutFracWidth  = float32(52)
)

type priceLayout struct{}

func (l *priceLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) < 3 {
		return
	}

	whole := objects[0]
	line := objects[1].(*canvas.Line)
	frac := objects[2]

	fracW := priceLayoutFracWidth
	gap := priceLayoutDividerGap
	wholeW := size.Width - gap - fracW
	if wholeW < 0 {
		wholeW = 0
	}

	// Whole-number field fills the left portion.
	whole.Move(fyne.NewPos(0, 0))
	whole.Resize(fyne.NewSize(wholeW, size.Height))

	// Vertical divider - drawn in the center of the gap.
	divX := wholeW + gap/2
	const vPad = float32(4)
	line.Position1 = fyne.NewPos(divX, vPad)
	line.Position2 = fyne.NewPos(divX, size.Height-vPad)
	line.Refresh()

	// Fractional field occupies the fixed right portion.
	frac.Move(fyne.NewPos(wholeW+gap, 0))
	frac.Resize(fyne.NewSize(fracW, size.Height))
}

func (l *priceLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) < 3 {
		return fyne.NewSize(0, 0)
	}

	wholeMin := objects[0].MinSize()
	fracMin := objects[2].MinSize()

	// Use the fixed frac width if it is wider than the natural minimum.
	fracW := fyne.Max(fracMin.Width, priceLayoutFracWidth)

	return fyne.NewSize(
		wholeMin.Width+priceLayoutDividerGap+fracW,
		fyne.Max(wholeMin.Height, fracMin.Height),
	)
}
