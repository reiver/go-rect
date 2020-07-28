# go-rect

Package **rect** provides types that represents a rectangle filled in with a single color at a specific location,
that seamlessly works with Go's built-in `"image"`, `"image/color"`, and `"image/draw"` packages.

## Example

Here is an example of usage:
```Go
import "github.com/reiver/go-rect"

// ...

rectangle := rect.RGBA{
	X:x,
	Y:y,

	Width:width,
	Height:height,

	R: r,
	G: g,
	B: b,
	A: a,
}

// ...

draw.Draw(dst, rectangle.Bounds(), rectangle, rectangle.Bounds().Min, draw.Over)
```

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-rect

[![GoDoc](https://godoc.org/github.com/reiver/go-rect?status.svg)](https://godoc.org/github.com/reiver/go-rect)

