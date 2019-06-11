package widget

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
)

type shadowType int

const (
	shadowAround shadowType = iota
	shadowBottom
	shadowTop
)

func newShadow(typ shadowType) *shadow {
	return &shadow{typ: typ}
}

var _ fyne.Widget = (*shadow)(nil)

type shadow struct {
	baseWidget
	typ shadowType
}

func (s *shadow) CreateRenderer() fyne.WidgetRenderer {
	r := &shadowRenderer{s: s}
	r.createShadows()
	return r
}

func (s *shadow) Hide() {
	s.hide(s)
}

func (s *shadow) MinSize() fyne.Size {
	return s.minSize(s)
}

func (s *shadow) Move(p fyne.Position) {
	s.move(p, s)
}

func (s *shadow) Resize(size fyne.Size) {
	s.resize(size, s)
}

func (s *shadow) Show() {
	s.show(s)
}

type shadowRenderer struct {
	b, bl, br, l, r, t, tl, tr *canvas.Gradient
	minSize                    fyne.Size
	objects                    []fyne.CanvasObject
	s                          *shadow
}

func (r *shadowRenderer) createShadows() {
	switch r.s.typ {
	case shadowBottom:
		r.b = canvas.NewLinearGradient(
			theme.ShadowColor(),
			color.Transparent,
			canvas.GradientDirectionVertical,
		)
		r.objects = []fyne.CanvasObject{r.b}
	case shadowTop:
		r.t = canvas.NewLinearGradient(
			color.Transparent,
			theme.ShadowColor(),
			canvas.GradientDirectionVertical,
		)
		r.objects = []fyne.CanvasObject{r.t}
	case shadowAround:
		r.tl = canvas.NewLinearGradient(
			theme.ShadowColor(),
			color.Transparent,
			canvas.GradientDirectionCircular,
		)
		r.tl.CenterOffset = fyne.NewPos(theme.Padding(), theme.Padding())
		r.t = canvas.NewLinearGradient(
			color.Transparent,
			theme.ShadowColor(),
			canvas.GradientDirectionVertical,
		)
		r.tr = canvas.NewLinearGradient(
			theme.ShadowColor(),
			color.Transparent,
			canvas.GradientDirectionCircular,
		)
		r.tr.CenterOffset = fyne.NewPos(-theme.Padding(), theme.Padding())
		r.r = canvas.NewLinearGradient(
			theme.ShadowColor(),
			color.Transparent,
			canvas.GradientDirectionHorizontal,
		)
		r.br = canvas.NewLinearGradient(
			theme.ShadowColor(),
			color.Transparent,
			canvas.GradientDirectionCircular,
		)
		r.br.CenterOffset = fyne.NewPos(-theme.Padding(), -theme.Padding())
		r.b = canvas.NewLinearGradient(
			theme.ShadowColor(),
			color.Transparent,
			canvas.GradientDirectionVertical,
		)
		r.bl = canvas.NewLinearGradient(
			theme.ShadowColor(),
			color.Transparent,
			canvas.GradientDirectionCircular,
		)
		r.bl.CenterOffset = fyne.NewPos(theme.Padding(), -theme.Padding())
		r.l = canvas.NewLinearGradient(
			color.Transparent,
			theme.ShadowColor(),
			canvas.GradientDirectionHorizontal,
		)
		r.objects = []fyne.CanvasObject{r.tl, r.t, r.tr, r.r, r.br, r.b, r.bl, r.l}
	}
}

func (r *shadowRenderer) ApplyTheme() {
	r.createShadows()
}

func (r *shadowRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *shadowRenderer) Destroy() {
}

func (r *shadowRenderer) Layout(size fyne.Size) {
	if r.tl != nil {
		r.tl.Resize(fyne.NewSize(theme.Padding()*2, theme.Padding()*2))
		r.tl.Move(fyne.NewPos(-theme.Padding()*2, -theme.Padding()*2))
	}
	if r.t != nil {
		r.t.Resize(fyne.NewSize(size.Width, theme.Padding()*2))
		r.t.Move(fyne.NewPos(0, -theme.Padding()*2))
	}
	if r.tr != nil {
		r.tr.Resize(fyne.NewSize(theme.Padding()*2, theme.Padding()*2))
		r.tr.Move(fyne.NewPos(size.Width, -theme.Padding()*2))
	}
	if r.r != nil {
		r.r.Resize(fyne.NewSize(theme.Padding()*2, size.Height))
		r.r.Move(fyne.NewPos(size.Width, 0))
	}
	if r.br != nil {
		r.br.Resize(fyne.NewSize(theme.Padding()*2, theme.Padding()*2))
		r.br.Move(fyne.NewPos(size.Width, size.Height))
	}
	if r.b != nil {
		r.b.Resize(fyne.NewSize(size.Width, theme.Padding()*2))
		r.b.Move(fyne.NewPos(0, size.Height))
	}
	if r.bl != nil {
		r.bl.Resize(fyne.NewSize(theme.Padding()*2, theme.Padding()*2))
		r.bl.Move(fyne.NewPos(-theme.Padding()*2, size.Height))
	}
	if r.l != nil {
		r.l.Resize(fyne.NewSize(theme.Padding()*2, size.Height))
		r.l.Move(fyne.NewPos(-theme.Padding()*2, 0))
	}
}

func (r *shadowRenderer) MinSize() fyne.Size {
	return r.minSize
}

func (r *shadowRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *shadowRenderer) Refresh() {
}
