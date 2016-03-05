package gopage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPaginator(t *testing.T) {
	p, err := NewPaginator([]int{1, 2, 3, 4})
	assert.Nil(t, err)
	assert.Equal(t, p.GetPageSize(), DefaultPageSize)
}
