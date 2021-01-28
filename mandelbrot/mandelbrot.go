// mandelbrot.go
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 4096, 4096
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
	ShowImage(img)

	//fmt.Println("Hello World!")
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func ShowImage(m image.Image) {
	// 需要保存的文件
	imgcounter := 123
	imgfile, _ := os.Create(fmt.Sprintf("%03d.png", imgcounter))
	defer imgfile.Close()

	// 以PNG格式保存文件
	err := png.Encode(imgfile, m)
	if err != nil {
		log.Fatal(err)
	}

}
