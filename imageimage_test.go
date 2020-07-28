package rect_test

import (
	"github.com/reiver/go-rect"

	"image"

	"testing"
)

func TestColorLinked_imageimage(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var datum image.Image = rect.ColorLinked{}

	if nil == datum {
		t.Error("This should never happen")
	}
}

func TestRGBA_imageimage(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var datum image.Image = rect.RGBA{}

	if nil == datum {
		t.Error("This should never happen")
	}
}
