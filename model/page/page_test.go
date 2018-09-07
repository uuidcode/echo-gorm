package page

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCeil(t *testing.T) {
	fmt.Println(9 / 10)
	fmt.Println(float64(9 / 10))
	fmt.Println(float64(9) / float64(10))
}

func TestNewWithItemCount10(t *testing.T) {
	page := New(0, 9, 10)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1}, page.List)

	page = New(1, 9, 10)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1}, page.List)

	page = New(1, 10, 10)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1}, page.List)

	page = New(1, 11, 10)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1, 2}, page.List)

	page = New(1, 100, 10)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, page.List)

	page = New(1, 101, 10)
	assert.False(t, page.Previous)
	assert.True(t, page.Next)
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, page.List)

	page = New(11, 201, 10)
	assert.True(t, page.Previous)
	assert.True(t, page.Next)
	assert.Equal(t, []int64{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, page.List)

	page = New(12, 201, 10)
	assert.True(t, page.Previous)
	assert.True(t, page.Next)
	assert.Equal(t, []int64{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, page.List)
}

func TestNewWithItemCount20(t *testing.T) {
	page := New(0, 9, 20)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1}, page.List)

	page = New(1, 9, 20)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1}, page.List)

	page = New(1, 10, 20)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1}, page.List)

	page = New(1, 11, 20)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1, 2}, page.List)

	page = New(1, 100, 20)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, page.List)

	page = New(1, 200, 20)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, page.List)

	page = New(1, 201, 20)
	assert.False(t, page.Previous)
	assert.True(t, page.Next)
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, page.List)

	page = New(12, 401, 20)
	assert.False(t, page.Previous)
	assert.True(t, page.Next)
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, page.List)

	page = New(21, 620, 20)
	assert.True(t, page.Previous)
	assert.True(t, page.Next)
	assert.Equal(t, []int64{21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}, page.List)
}

func TestNewWithItemCount5(t *testing.T) {
	page := New(0, 9, 5)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1}, page.List)

	page = New(1, 9, 5)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1}, page.List)

	page = New(1, 10, 5)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1}, page.List)

	page = New(1, 11, 5)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1, 2}, page.List)

	page = New(1, 50, 5)
	assert.False(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{1, 2, 3, 4, 5}, page.List)

	page = New(1, 51, 5)
	assert.False(t, page.Previous)
	assert.True(t, page.Next)
	assert.Equal(t, []int64{1, 2, 3, 4, 5}, page.List)

	page = New(2, 51, 5)
	assert.False(t, page.Previous)
	assert.True(t, page.Next)
	assert.Equal(t, []int64{1, 2, 3, 4, 5}, page.List)

	page = New(6, 100, 5)
	assert.True(t, page.Previous)
	assert.False(t, page.Next)
	assert.Equal(t, []int64{6, 7, 8, 9, 10}, page.List)

	page = New(7, 101, 5)
	assert.True(t, page.Previous)
	assert.True(t, page.Next)
	assert.Equal(t, []int64{6, 7, 8, 9, 10}, page.List)
}
