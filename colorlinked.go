package rect

import (
	"image"
	"image/color"
)

type ColorLinked struct {
	X int
	Y int

	Width  int
	Height int

	Color color.Color
}

func (receiver ColorLinked) At(x, y int) color.Color {
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

func (receiver ColorLinked) Bounds() image.Rectangle {
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

func (receiver ColorLinked) ColorModel() color.Model {
	return color.NRGBAModel
}

func (receiver ColorLinked) RGBA() (r, g, b, a uint32) {
	return receiver.Color.RGBA()
}
