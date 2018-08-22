package main

import (
	"github.com/mafuyuk/go-training/design-pattern/bridge/uikit/draw"
	"github.com/mafuyuk/go-training/design-pattern/bridge/uikit/shape"
)

func main() {
	openGL := &draw.OpenGL{}
	direct2D := &draw.Direct2D{}

	circle := &shape.Circle{
		Center: draw.Point{X: 100, Y: 100},
		Radius: 50,
	}

	circle.DrawingContext = openGL
	circle.Draw()

	circle.DrawingContext = direct2D
	circle.Draw()
}
