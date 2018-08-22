package draw

import "fmt"

type OpenGL struct{}

func (o *OpenGL) DrawEllipseInRect(r Rect) error {
	fmt.Printf("OpenGL is drawing ellipse in rect %v \n", r)
	return nil
}
