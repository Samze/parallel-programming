package parallelgo 

func SpreadEvenly(items, routines int) []int {
	itemsPerRoutine := items / routines
	remaining := items % routines

	routineList := make([]int, routines)
	for i := 0; i < routines; i++ {
		if remaining > 0 {
			routineList[i] = itemsPerRoutine + 1
			remaining = remaining - 1
		} else {
			routineList[i] = itemsPerRoutine
		}
	}
	return routineList
}
