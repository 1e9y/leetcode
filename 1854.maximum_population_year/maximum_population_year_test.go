package maximum_population_year

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	logs   [][]int
	expect int
}{
	{
		[][]int{
			{1993, 1999},
			{2000, 2010},
		},
		1993,
	},
	{
		[][]int{
			{1950, 1961},
			{1960, 1971},
			{1970, 1981},
		},
		1960,
	},
}

func TestMaximumPopulation(t *testing.T) {
	for _, c := range testCases {
		assert.Equal(t, c.expect, maximumPopulation(c.logs))
	}
}
