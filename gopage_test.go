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
	assert.NotEqual(t, len(sl), 0)
	assert.Equal(t, len(sl), 4)
	for i, s := range sl {
		assert.Equal(t, s.(int), source[i])
	}

	p.SetPageSize(2)
	assert.Equal(t, p.GetPageSize(), 2)
	sl, errsl = p.Page(1)
	assert.Nil(t, errsl)
	assert.NotEqual(t, len(sl), 0)
	assert.Equal(t, len(sl), 2)
	for i, s := range sl {
		assert.Equal(t, s.(int), source[i])
	}

	sl, errsl = p.Page(3)
	assert.Equal(t, errsl, ErrOverflow)
	assert.Equal(t, len(sl), 0)
	assert.Nil(t, sl)
}
