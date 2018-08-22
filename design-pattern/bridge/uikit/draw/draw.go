package draw

type (
	Point struct {
		X float64
		Y float64
	}

	Size struct {
		Width  float64
		Height float64
	}

	Rect struct {
		Location Point
		Size
	}
)

type Drawer interface {
	DrawEllipseInRect(Rect) error
}
