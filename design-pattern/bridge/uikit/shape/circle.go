package shape

import "github.com/mafuyuk/go-training/design-pattern/bridge/uikit/draw"

type Circle struct {
	DrawingContext draw.Drawer
	Center         draw.Point
	Radius         float64
}

func (c *Circle) Draw() error {
	rect := draw.Rect{
		Location: draw.Point{
			X: c.Center.X - c.Radius,
			Y: c.Center.Y - c.Radius,
		},
		Size: draw.Size{
			Width:  2 * c.Radius,
			Height: 2 * c.Radius,
		},
	}

	return c.DrawingContext.DrawEllipseInRect(rect)
}
