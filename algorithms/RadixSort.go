package algorithms

const BASE = 10

func RadixSort(values []int) (ret []int) {
	max := 0
	for _, v := range values {
		if v > max {
			max = v
		}
	}

	exp := 1
	ret = values
	for ; max/exp > 0; exp *= BASE {
		ret = radixPass(ret, exp)
	}
	return
}

func radixPass(values []int, exp int) (ret []int) {
	ret = make([]int, len(values))

	// Count the number of items for each bucket
	buckets := make([]int, BASE)
	for _, v := range values {
		buckets[(v/exp)%BASE]++
	}

	// Modify the bucket count to be cumulative
	for i := 1; i < BASE; i++ {
		buckets[i] += buckets[i-1]
	}

	for i := len(values) - 1; i >= 0; i-- {
		key := (values[i] / exp) % BASE
		ret[buckets[key]-1] = values[i]
		buckets[key]--
	}
	return
}
