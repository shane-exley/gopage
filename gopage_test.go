package gopage

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
