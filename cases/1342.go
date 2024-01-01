package cases

func ReduceToZeroStep(v int) (step int) {
	for {
		if v == 0 {
			break
		} else if v&1 == 1 {
			step += 1
			v -= 1
		} else {
			step += v&1 + 1
			v = v >> 1
		}
	}

	return step
}
