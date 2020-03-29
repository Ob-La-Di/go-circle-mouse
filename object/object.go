package object

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type BaseObject struct {
	Background color.Gray16
}

type Object interface {
	Draw(screen *ebiten.Image)
}

type Point struct {
	X, Y float64
}

type Square struct {
	BaseObject
	TopLeft       Point
	Width, Height int64
}


func (s Square) Draw(screen *ebiten.Image) {
	square, _ := ebiten.NewImage(int(s.Width), int(s.Height), ebiten.FilterNearest)
	square.Fill(color.White)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(s.TopLeft.X, s.TopLeft.Y)
	screen.DrawImage(square, opts)
}


type Circle struct {
	BaseObject
	Center Point
	Radius int64
}