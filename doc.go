/*
Package rect provides types that represents a rectangle filled in with a single color at a specific location,
that seamlessly works with Go's built-in "image", "image/color", and "image/draw" packages.

Example

Here is an example of usage:

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
*/
package rect
