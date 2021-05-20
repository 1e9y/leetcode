package maximum_population_year

const minYear = 1950
const maxYear = 2050

func maximumPopulation(logs [][]int) int {
	counts := make([]int, maxYear-minYear+1)

	for i := range logs {
		counts[logs[i][0]-minYear]++
		counts[logs[i][1]-minYear]--
	}

	max := 0
	for i := range counts {
		if i == 0 {
			continue
		}
		counts[i] += counts[i-1]
		if counts[i] > counts[max] {
			max = i
		}
	}

	return minYear + max
}
