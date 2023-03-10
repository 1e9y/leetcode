package coin_change

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	coins  []int
	amount int
	expect int
}{
	{
		[]int{1, 2, 5}, 11, 3,
	},
	{
		[]int{2}, 3, -1,
	},
	{
		[]int{1}, 0, 0,
	},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.coins), func(t *testing.T) {
			assert.Equal(t, tt.expect, coinChange(tt.coins, tt.amount))
		})
	}
}
