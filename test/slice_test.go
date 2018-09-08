package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice(t *testing.T) {
	slice := []int{1, 2, 3}

	changeSlice(slice)
	assert.Equal(t, 0, slice[0])
}

func changeSlice(slice []int) {
	slice[0] = 0
}
