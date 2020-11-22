package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
)

func main() {
	pngFile, err := os.Open("./images/asset.png")
	if err != nil {
		panic(err)
	}
	defer pngFile.Close()

	img, _, err := image.Decode(pngFile)
	if err != nil {
		panic(err)
	}

	size := img.Bounds().Size()

	goFile, err := os.Create("./internal/imgdata/asset.go")
	if err != nil {
		panic(err)
	}
	defer goFile.Close()

	goFile.WriteString(fmt.Sprintf(`package imgdata

import "../img"

var asset1 = img.Asset{
	Width:  %d,
	Height: %d,
	Data: []byte("`, size.X, size.Y))

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			if (x > 0 || y > 0) && (y*size.X+x)%14 == 0 {
				goFile.WriteString(`" +
		"`)
			}
			c565 := RGBATo565(img.At(x, y))
			//goFile.WriteString(fmt.Sprintf("%d,%d,", byte(c565), byte(c565 >> 8)))
			goFile.WriteString(fmt.Sprintf("\\%03o\\%03o", byte(c565), byte(c565>>8)))
		}
	}

	goFile.WriteString(`"),
}
`)
}

func RGBATo565(c color.Color) uint16 {
	r, g, b, _ := c.RGBA()
	return uint16((r & 0xF800) +
		((g & 0xFC00) >> 5) +
		((b & 0xF800) >> 11))
}
