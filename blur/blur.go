package main

import (
	"image"
	"image/color"
)

type BlurFilter interface {
	BlurImage(*image.RGBA) *image.RGBA
}

type Config struct {
	radius int
}

type SequentialBlur struct {
	Config
}

func (s *SequentialBlur) BlurImage(img *image.RGBA) *image.RGBA {
	rect := img.Bounds()

	newImg := image.NewRGBA(image.Rect(0, 0, rect.Dx(), rect.Dy()))

	for x := rect.Min.X; x <= rect.Max.X; x++ {
		for y := rect.Min.Y; y <= rect.Max.Y; y++ {
			color := calcNewRGBA(img, x, y, s.radius)
			newImg.SetRGBA(x, y, color)
		}
	}
	return newImg
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	} else {
		return value
	}
}

func calcNewRGBA(img *image.RGBA, x, y, radius int) color.RGBA {
	minX := clamp(x-radius, 0, img.Rect.Max.X)
	maxX := clamp(x+radius, 0, img.Rect.Max.X)
	minY := clamp(y-radius, 0, img.Rect.Max.Y)
	maxY := clamp(y+radius, 0, img.Rect.Max.Y)

	var totalR, totalG, totalB, totalA uint32
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			rgba := (*img).At(i, j)
			r, g, b, a := rgba.RGBA()
			totalR = totalR + uint32(uint8(r))
			totalG = totalG + uint32(uint8(g))
			totalB = totalB + uint32(uint8(b))
			totalA = totalA + uint32(uint8(a))
		}
	}

	surroundingCellCount := uint32((maxX - minX + 1) * (maxY - minY + 1))

	newR := totalR / surroundingCellCount
	newG := totalG / surroundingCellCount
	newB := totalB / surroundingCellCount
	newA := totalA / surroundingCellCount

	return color.RGBA{uint8(newR), uint8(newG), uint8(newB), uint8(newA)}
}
