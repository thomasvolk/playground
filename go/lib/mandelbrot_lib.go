package mandelbrot

import (
	"image"
	"image/color"
	"math/cmplx"
	"sync"
)

type Mandelbrot struct {
	Xmin       float64
	Ymin       float64
	Step       float64
	Iterations int
	Width      int
	Height     int
}

func (mandelbrot *Mandelbrot) Draw() *image.RGBA {
	var img = image.NewRGBA(image.Rect(0, 0, mandelbrot.Width, mandelbrot.Height))
	var wg sync.WaitGroup
	wg.Add(mandelbrot.Width * mandelbrot.Height)
	for x := 0; x < mandelbrot.Width; x++ {
		for y := 0; y < mandelbrot.Height; y++ {
			go func(image *image.RGBA, px int, py int) {
				defer wg.Done()
				mandelbrot.drawPoint(image, px, py)
			}(img, x, y)
		}
	}
	wg.Wait()
	return img
}

func (mandelbrot *Mandelbrot) drawPoint(img *image.RGBA, x int, y int) {
	color := color.RGBA{0x00, 0x00, 0x00, 0xff}
	point := complex(mandelbrot.Xmin+float64(x)*mandelbrot.Step,
		mandelbrot.Ymin+float64(y)*mandelbrot.Step)
	if cmplx.Abs(point) < 2 {
		nextPoint := 0 + 0i
		for i := 0; i < mandelbrot.Iterations; i++ {
			nextPoint = nextPoint*nextPoint + point
			if cmplx.Abs(nextPoint) < 2 {
				color = mandelbrot.calculateColor(i)
			}
		}
	}
	img.Set(int(x), int(y), color)
}

func (mandelbrot *Mandelbrot) colorStep() int {
	return 255 / mandelbrot.Iterations
}

func (mandelbrot *Mandelbrot) calculateColor(i int) color.RGBA {
	blue := 0
	green := 2 * i * mandelbrot.colorStep()
	red := 0
	if i >= mandelbrot.Iterations/2 {
		blue = (i - mandelbrot.Iterations/2) * mandelbrot.colorStep()
	}
	if i > mandelbrot.Iterations/2 {
		green = 2 * (mandelbrot.Iterations - i) * mandelbrot.colorStep()
	}
	if i <= mandelbrot.Iterations/2 {
		red = 255 - 2*i*mandelbrot.colorStep()
	}
	return color.RGBA{uint8(red), uint8(green), uint8(blue), 0xff}
}