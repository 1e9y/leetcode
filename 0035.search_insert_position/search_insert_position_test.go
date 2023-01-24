package search_insert_position

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nums   []int
	target int
	expect int
}{
	{
		nums:   []int{1, 3, 5, 6},
		target: 5,
		expect: 2,
	},
	{
		nums:   []int{1, 3, 5, 6},
		target: 2,
		expect: 1,
	},
	{
		nums:   []int{1, 3, 5, 6},
		target: 7,
		expect: 4,
	},
}

func TestSearchInsert(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tt.expect, searchInsert(tt.nums, tt.target))
		})
	}
}
