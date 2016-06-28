package main

import (
	"image"
	"image/color"
)

type BlurFilter interface {
	Blur(image.Image) image.Image
}

type SimpleBlur struct {
	radius int
}

func (s *SimpleBlur) Blur(img image.RGBA) image.RGBA {
	rect := img.Bounds()

	newImg := image.NewRGBA(image.Rect(0, 0, rect.Dx(), rect.Dy()))

	for x := rect.Min.X; x <= rect.Max.X; x++ {
		for y := rect.Min.Y; y <= rect.Max.Y; y++ {
			s.blurAt(&img, newImg, x, y)
		}
	}
	return *newImg
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

func (s *SimpleBlur) blurAt(img *image.RGBA, newImage *image.RGBA, x, y int) {
	minX := clamp(x-s.radius, 0, img.Rect.Max.X)
	maxX := clamp(x+s.radius, 0, img.Rect.Max.X)
	minY := clamp(y-s.radius, 0, img.Rect.Max.Y)
	maxY := clamp(y+s.radius, 0, img.Rect.Max.Y)

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

	(*newImage).SetRGBA(x, y, color.RGBA{uint8(newR), uint8(newG), uint8(newB), uint8(newA)})
}
