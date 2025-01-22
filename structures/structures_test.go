package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueues(t *testing.T) {
	list := Queue[int]{}

	list.Enqueue(5)
	list.Enqueue(7)
	list.Enqueue(9)

	assert.Equal(t, list.Dequeue(), 5)
	assert.Equal(t, list.length, 2)

	list.Enqueue(11)

	assert.Equal(t, list.Dequeue(), 7)
	assert.Equal(t, list.Dequeue(), 9)
	assert.Equal(t, list.Peek(), 11)
	assert.Equal(t, list.Dequeue(), 11)
	assert.Equal(t, list.Dequeue(), 0)
	assert.Equal(t, list.length, 0)

	list.Enqueue(69)
	assert.Equal(t, list.Peek(), 69)
	assert.Equal(t, list.length, 1)
}
