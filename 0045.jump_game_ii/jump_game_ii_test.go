package jump_game_ii

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
		[]int{2, 3, 1, 1, 4},
		2,
	},
	{
		[]int{2, 3, 0, 1, 4},
		2,
	},
}

func TestJump(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			assert.Equal(t, tt.expect, jump(tt.input))
		})
	}
}
