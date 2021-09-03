package best_time_to_buy_and_sell_stock

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  []int
	expect int
}{
	{
		[]int{7, 1, 5, 3, 6, 4},
		5,
	},
	{
		[]int{7, 6, 4, 3, 1},
		0,
	},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			assert.Equal(t, tt.expect, maxProfit(tt.input))
		})
	}
}
