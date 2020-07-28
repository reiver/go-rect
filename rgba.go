package rect

import (
	"image"
	"image/color"
)

type RGBA struct {
	X int
	Y int

	Width  int
	Height int

	R uint8
	G uint8
	B uint8
	A uint8
}

func (receiver RGBA) At(x, y int) color.Color {
	if ! (image.Point{X:x,Y:y}).In(receiver.Bounds()) {
		return RGBA{
			X: receiver.X,
			Y: receiver.Y,
			Width:  receiver.Width,
			Height: receiver.Height,
		}
	}

	return receiver
}

func (receiver RGBA) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: receiver.X,
			Y: receiver.Y,
		},
		Max: image.Point{
			X: receiver.X+receiver.Width,
			Y: receiver.Y+receiver.Height,
		},
	}
}

func (receiver RGBA) ColorModel() color.Model {
	return color.NRGBAModel
}

func (receiver RGBA) RGBA() (r, g, b, a uint32) {
	r = uint32(receiver.R) * (0xffff/0xff)
	g = uint32(receiver.G) * (0xffff/0xff)
	b = uint32(receiver.B) * (0xffff/0xff)
	a = uint32(receiver.A) * (0xffff/0xff)

	return r,g,b,a
}
