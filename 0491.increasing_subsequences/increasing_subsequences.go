package increasing_subsequences

func findSubsequences(n []int) [][]int {
	output := make([][]int, 0)

	var find func(i int, subpath []int)
	find = func(i int, subpath []int) {
		if len(subpath) > 1 {
			temp := make([]int, len(subpath))
			copy(temp, subpath)
			output = append(output, temp)
		}

		seen := map[int]bool{}
		for j := i; j < len(n); j++ {
			if seen[n[j]] {
				continue
			}
			if len(subpath) == 0 || n[j] >= subpath[len(subpath)-1] {
				seen[n[j]] = true
				find(j+1, append(subpath, n[j]))
			}
		}
	}
	find(0, []int{})

	return output
}
