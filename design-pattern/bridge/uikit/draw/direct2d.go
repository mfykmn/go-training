package draw

import "fmt"

type Direct2D struct{}

func (d *Direct2D) DrawEllipseInRect(r Rect) error {
	fmt.Printf("Direct2D is drawing ellipse in rect %v \n", r)
	return nil
}
