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

func produceHits(iter, routines int) <-chan int {
	hitsChan := make(chan int)
	routinesArray := SpreadEvenly(iter, routines)

	for i := 0; i < routines; i++ {
		iterationsForRoutine := routinesArray[i]
		go func() {
			hitsChan <- McCount(iterationsForRoutine)
		}()
	}

	return hitsChan
}

func EstimatePiConcurrently(iter, routines int) float64 {
	hitsChan := produceHits(iter, routines)

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
	concurrency := 4
	iterations := 1 << 25

	sequential := func() interface{} {
		return EstimatePiSequentially(iterations)
	}

	parallel := func() interface{} {
		return EstimatePiConcurrently(iterations, concurrency)
	}

	seqDuration, seqResult := TimeCall(sequential)
	parDuration, parResult := TimeCall(parallel)

	fmt.Println(fmt.Sprintf("sequential result: %f executed in %v", seqResult, seqDuration))
	fmt.Println(fmt.Sprintf("parallel result: %f executed in %v", parResult, parDuration))

}
