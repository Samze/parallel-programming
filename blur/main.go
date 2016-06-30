package main

import (
	"fmt"
	. "github.com/samze/parallelgo"
	"image"
	"image/draw"
	"image/png"
	"os"
	"strconv"
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
	if len(os.Args) != 3 {
		fmt.Println("Requires two arguments. Image path and blur radius")
		os.Exit(1)
	}

	imgPath := os.Args[1]
	blurRadius, err := strconv.Atoi(os.Args[2])

	check(err)

	img, err := loadPng(imgPath)
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
