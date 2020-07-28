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

## See Also

For other useful Go packages that seamlessly works with Go's built-in `"image"`, `"image/color"`, and `"image/draw"` packages, also see:

* https://github.com/reiver/go-c80
* https://github.com/reiver/go-dast8x8
* https://github.com/reiver/go-dynaimg
* https://github.com/reiver/go-font8x8
* https://github.com/reiver/go-frame256x288
* https://github.com/reiver/go-imagerelocate
* https://github.com/reiver/go-imgstr
* https://github.com/reiver/go-palette2048
* https://github.com/reiver/go-pel
* https://github.com/reiver/go-rgba32
* https://github.com/reiver/go-rgba32sprite8x8
* https://github.com/reiver/go-sprite8x8
* https://github.com/reiver/go-sprite32x32
* https://github.com/reiver/go-spritesheet8x8x256
* https://github.com/reiver/go-spritesheet32x32x256
* https://github.com/reiver/go-text32x36
* https://github.com/reiver/go-tilemap8x8x256x256
