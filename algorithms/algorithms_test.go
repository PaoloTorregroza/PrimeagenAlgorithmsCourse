package algorithms

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	assert.True(t, LinearSearch(foo, 69))
	assert.False(t, LinearSearch(foo, 135))
	assert.True(t, LinearSearch(foo, 69420))
	assert.False(t, LinearSearch(foo, 69421))
	assert.True(t, LinearSearch(foo, 1))
	assert.False(t, LinearSearch(foo, 0))
}

func TestBinarySearch(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	assert.True(t, BinarySearch(foo, 69))
	assert.False(t, BinarySearch(foo, 135))
	assert.True(t, BinarySearch(foo, 69420))
	assert.False(t, BinarySearch(foo, 69421))
	assert.True(t, BinarySearch(foo, 1))
	assert.False(t, BinarySearch(foo, 0))
}

func TestTwoCrystalBalls(t *testing.T) {
	idx := rand.Intn(10001)
	data := make([]bool, 10000)

	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	assert.Equal(t, TwoCrystalBalls(data), idx)
	assert.Equal(t, TwoCrystalBalls(make([]bool, 821)), -1)
}
