package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math"
	"mouse-test/object"
)

func circleIntersectionWithCursorPosition(circle *object.Circle, cursorPosition *object.Point) *object.Point {
	dx := cursorPosition.X - circle.Center.X
	dy := cursorPosition.Y - circle.Center.Y
	A := dx*dx + dy*dy
	C := -float64(circle.Radius)*float64(circle.Radius)
	det := - 4 * A * C

	t := (math.Sqrt(det)) / (2 * A)
	intersection1 :=
		&object.Point{X: circle.Center.X + t*dx, Y: circle.Center.Y + t*dy}

	return intersection1
}

func update(screen *ebiten.Image) error {
	x, y := ebiten.CursorPosition()
	squarePosition := circleIntersectionWithCursorPosition(&object.Circle{Center: object.Point{
		X: float64(100),
		Y: float64(100),
	}, Radius: int64(50)}, &object.Point{X: float64(x), Y: float64(y)})
	objects := []object.Object{&object.Square{BaseObject: object.BaseObject{Background: color.White}, TopLeft: *squarePosition, Width: 1, Height: 1}}
	for _, o := range objects {
		if _, ok := o.(object.Object); ok {
			o.Draw(screen)
		}
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "hello world"); err != nil {
		panic(err)
	}
}
