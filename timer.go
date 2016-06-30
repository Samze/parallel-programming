package parallelgo

import "time"

type fn func() interface{}

func TimeCall(exec fn) (time.Duration, interface{}) {
	before := time.Now()
	result := exec()
	return time.Since(before), result
}
