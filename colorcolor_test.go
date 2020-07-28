package rect_test

import (
	"github.com/reiver/go-rect"

	"image/color"

	"testing"
)

func TestColorLinked_colorcolor(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var datum color.Color = rect.ColorLinked{}

	if nil == datum {
		t.Error("This should never happen")
	}
}

func TestRGBA_colorcolor(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var datum color.Color = rect.RGBA{}

	if nil == datum {
		t.Error("This should never happen")
	}
}
