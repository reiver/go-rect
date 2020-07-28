package rect_test

import (
	"github.com/reiver/go-rect"

	"image"
	"math/rand"
	"time"

	"testing"
)

func TestRGBA_At(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	for testNumber:=0; testNumber<10; testNumber++ {

		var x,y int
		{
			x = randomness.Int()
			if 0 == randomness.Int()%2 {
				x = -x
			}

			y = randomness.Int()
			if 0 == randomness.Int()%2 {
				y = -y
			}
		}

		var width,height int
		{
			width  = 1 + randomness.Intn(1000)

			height = 1 + randomness.Intn(1000)
		}

		var r,b,g,a uint8
		{
			r = uint8(randomness.Intn(256))
			g = uint8(randomness.Intn(256))
			b = uint8(randomness.Intn(256))
			a = uint8(randomness.Intn(256))
		}

		var rectangle rect.RGBA = rect.RGBA{
			X:x,
			Y:y,

			Width:width,
			Height:height,

			R:r,
			G:g,
			B:b,
			A:a,
		}

		var img image.Image = rectangle

		for yy:=y; yy<(y+height); yy++ {
			for xx:=x; xx<(x+width); xx++ {

				eR := uint32(r) * (0xffff/0xff)
				eG := uint32(g) * (0xffff/0xff)
				eB := uint32(b) * (0xffff/0xff)
				eA := uint32(a) * (0xffff/0xff)

				aR, aG, aB, aA := img.At(xx,yy).RGBA()

				if eR != aR || eG != aG || eB != aB || eA != aA {
					t.Errorf("For test #%d, the actual color was not what was expected.", testNumber)
					t.Logf("(xx,yy)=(%d,%d)", x,y)
					t.Logf("rgba(%d,%d,%d,%d)", r,g,b,a)
					t.Logf("EXPECTED (r,g,b,a)=(%d,%d,%d,%d)", eR, eG, eB, eA)
					t.Logf("ACTUAL   (r,g,b,a)=(%d,%d,%d,%d)", aR, aG, aB, aA)
					continue
				}
			}
		}

		for subTestNumber:=0; subTestNumber<30; subTestNumber++ {
			var xx,yy int
			{
				xx = randomness.Int()
				if 0 == randomness.Int()%2 {
					xx = img.Bounds().Max.X + xx
				} else {
					xx = img.Bounds().Min.X - 1 - xx
				}

				yy = randomness.Int()
				if 0 == randomness.Int()%2 {
					yy = img.Bounds().Max.Y + yy
				} else {
					yy = img.Bounds().Min.Y - 1 - yy
				}
			}

			var eR, eG, eB, eA uint32 = 0,0,0,0

			aR, aG, aB, aA := img.At(xx,yy).RGBA()

			if eR != aR || eG != aG || eB != aB || eA != aA {
				t.Errorf("For test #%d, the actual color was not what was expected.", testNumber)
				t.Logf("(xx,yy)=(%d,%d)", xx,yy)
				t.Logf("rgba(%d,%d,%d,%d)", r,g,b,a)
				t.Logf("rectangle (x,y)=(%d,%d) (width,height)=(%d,%d)", x,y, width,height)
				t.Logf("EXPECTED (r,g,b,a)=(%d,%d,%d,%d)", eR, eG, eB, eA)
				t.Logf("ACTUAL   (r,g,b,a)=(%d,%d,%d,%d)", aR, aG, aB, aA)
				continue
			}
		}
	}
}
