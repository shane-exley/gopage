package gopage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPaginator(t *testing.T) {
	source := []int{1, 2, 3, 4}
	p, err := NewPaginator(source)
	assert.Nil(t, err)
	assert.Equal(t, p.GetPageSize(), DefaultPageSize)

	sl, errsl := p.Page(1)
	assert.Nil(t, errsl)
	assert.NotEqual(t, len(sl.([]int)), 0)
	assert.Equal(t, len(sl.([]int)), 4)
	for i, s := range sl.([]int) {
		assert.Equal(t, s, source[i])
	}

	errps := p.SetPageSize(2)
	assert.Nil(t, errps)
	assert.Equal(t, p.GetPageSize(), 2)
	sl, errsl = p.Page(1)
	assert.Nil(t, errsl)
	assert.NotEqual(t, len(sl.([]int)), 0)
	assert.Equal(t, len(sl.([]int)), 2)
	for i, s := range sl.([]int) {
		assert.Equal(t, s, source[i])
	}

	errps = p.SetPageSize(-10)
	assert.NotNil(t, errps)
	assert.Equal(t, errps, ErrInvalidPageSize)
	ps := p.GetPageSize()
	assert.Equal(t, ps, 2)

	sl, errsl = p.Page(3)
	assert.Equal(t, errsl, ErrOverflow)
	assert.Nil(t, sl)

	d, errd := NewPaginator(1234)
	assert.NotNil(t, errd)
	assert.Equal(t, errd, ErrNotSlice)
	assert.Nil(t, d)

	e, erre := p.Fetch(0, 1)
	assert.Nil(t, erre)
	a, ok := e.([]int)
	assert.True(t, ok)
	assert.NotNil(t, a)
	assert.Equal(t, len(a), 1)
}

func TestPageCountDefault(t *testing.T) {
	var tests = []struct {
		source    []int
		pageCount int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 1},
		{[]int{1, 2, 3}, 1},
		{[]int{1, 2, 3, 4}, 1},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 2, 3, 4, 5, 6}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}, 3},
	}

	for _, test := range tests {
		p, _ := NewPaginator(test.source)
		assert.Equal(t, test.pageCount, p.GetPageCount())
	}
}

func TestPageCount(t *testing.T) {
	source := []int{1, 2, 3, 4, 5, 6}
	p, _ := NewPaginator(source)

	var tests = []struct {
		pageSize  int
		pageCount int
	}{
		{0, 0},
		{1, 6},
		{2, 3},
		{3, 2},
		{4, 2},
		{5, 2},
		{6, 1},
		{7, 1},
		{8, 1},
		{9, 1},
	}

	for _, test := range tests {
		p.SetPageSize(test.pageSize)
		assert.Equal(t, test.pageCount, p.GetPageCount())
	}
}
