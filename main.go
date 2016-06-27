package main

import (
	"fmt"
	"math/rand"
	"time"
)

func McCount(iter int) int {
	xRan := rand.New(rand.NewSource(time.Now().UnixNano()))
	yRan := rand.New(rand.NewSource(time.Now().UnixNano()))
	hits := 0

	for i := 0; i < iter; i++ {
		x := xRan.Float32()
		y := yRan.Float32()

		if x*x+y*y < 1 {
			hits = hits + 1
		}
	}

	return hits
}

func EstimatePiSequentially(iter int) float64 {
	return 4.0 * float64(McCount(iter)) / float64(iter)
}

func main() {
	pi := EstimatePiSequentially(1 << 20)
	fmt.Printf("%f", pi)
}
