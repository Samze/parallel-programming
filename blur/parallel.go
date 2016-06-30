package main

import (
	. "github.com/samze/parallelgo"
	"image"
)

type HorizontalParallelBlur struct {
	Blur
	routines int
}

func (s *HorizontalParallelBlur) BlurImage(img *image.RGBA) *image.RGBA {
	rect := img.Bounds()
	newImg := image.NewRGBA(image.Rect(0, 0, rect.Dx(), rect.Dy()))

	done := s.doParalleBlur(img, newImg)

	for i := 0; i < s.routines; i++ {
		<-done
	}

	return newImg
}

func (s *HorizontalParallelBlur) doParalleBlur(img, newImg *image.RGBA) <-chan bool {
	rect := img.Bounds()

	rows := rect.Max.Y - rect.Min.Y
	rowsPerRoutineList := SpreadEvenly(rows, s.routines)

	done := make(chan bool)

	startPos := 0
	for _, rowsPerRoutine := range rowsPerRoutineList {
		endPos := startPos + rowsPerRoutine
		go s.doChunk(startPos, endPos, img, newImg, done)
		startPos = endPos
	}
	return done
}

func (s *HorizontalParallelBlur) doChunk(start, end int, img, newImg *image.RGBA, done chan bool) {
	for y := start; y < end; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			color := calcNewRGBA(img, x, y, s.radius)
			newImg.SetRGBA(x, y, color)
		}
	}
	done <- true
}
