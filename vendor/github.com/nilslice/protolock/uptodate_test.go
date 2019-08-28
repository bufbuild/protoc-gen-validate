package protolock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPermutation(t *testing.T) {
	assert := assert.New(t)
	assert.True(isPermutation([]int{1, 2}, []int{2, 1}, equalPrimitives))
	assert.True(isPermutation([]int{}, []int{}, equalPrimitives))
	assert.False(isPermutation([]int{1, 2}, []int{2, 2}, equalPrimitives))
	assert.False(isPermutation([]int{1, 2}, []int{2, 1, 2}, equalPrimitives))
	assert.False(isPermutation([]int{1, 2}, []int{}, equalPrimitives))
	assert.Panics(func() {
		isPermutation(1, 2, equalPrimitives)
	})
}
