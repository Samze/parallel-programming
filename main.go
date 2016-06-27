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

func EstimatePiConcurrently(iter, routines int) float64 {
	hitsChan := make(chan int)

	iterPerRoutine := iter / routines

	for i := 0; i < routines; i++ {
		go func() {
			hitsChan <- McCount(iterPerRoutine)
		}()
	}

	hits := 0
	for i := 0; i < routines; i++ {
		hits = hits + <-hitsChan
	}

	return 4.0 * float64(hits) / float64(iter)
}

func EstimatePiSequentially(iter int) float64 {
	return 4.0 * float64(McCount(iter)) / float64(iter)
}

func main() {
	piSeq := EstimatePiSequentially(1 << 25)
	piPar := EstimatePiConcurrently(1<<25, 4)
	fmt.Println(fmt.Sprintf("seq %f", piSeq))
	fmt.Println(fmt.Sprintf("par %f", piPar))
}