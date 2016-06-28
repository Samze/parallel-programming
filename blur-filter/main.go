package main

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func loadPng(fileName string) (image.Image, error) {
	pngFile, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	img, err := png.Decode(pngFile)

	if err != nil {
		return nil, err
	}

	return img, nil
}

func main() {
	arg := os.Args[1]

	img, err := loadPng(arg)

	bounds := img.Bounds()
	imgRGBA := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(imgRGBA, imgRGBA.Bounds(), img, bounds.Min, draw.Src)

	blur := SimpleBlur{1}

	newImage := blur.Blur(*imgRGBA)

	check(err)

	ShowImage(&newImage)
}
