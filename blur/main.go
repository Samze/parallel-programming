package main

import (
	"fmt"
	. "github.com/samze/parallelgo"
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

func convertToRGBAImage(img *image.Image) *image.RGBA {
	bounds := (*img).Bounds()
	imgRGBA := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(imgRGBA, imgRGBA.Bounds(), *img, bounds.Min, draw.Src)
	return imgRGBA
}

func main() {
	arg := os.Args[1]

	img, err := loadPng(arg)
	check(err)

	imgRGBA := convertToRGBAImage(&img)

	config := Config{blurRadius}
	seq := func() interface{} {
		blur := &SequentialBlur{config}
		return blur.BlurImage(imgRGBA)
	}

	par := func() interface{} {
		blur := &HorizontalParallelBlur{config, 100}
		return blur.BlurImage(imgRGBA)
	}

	seqDuration, _ := TimeCall(seq)
	parDuration, result := TimeCall(par)

	fmt.Println(fmt.Sprintf("sequential: %v", seqDuration))
	fmt.Println(fmt.Sprintf("parallel: %v", parDuration))

	ShowImage(result.(*image.RGBA))
}
