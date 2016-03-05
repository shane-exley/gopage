package gopage

import (
	"sync"
)

type Iterator struct {
	paginator *Paginator
	pageCount int
	mux       sync.Mutex
}

func (p *Paginator) GetIterator() (*Iterator, error) {
	return &Iterator{paginator: p}, nil
}

func (i *Iterator) GetCursorPosition() int {
	return i.pageCount
}

func (i *Iterator) Next() ([]interface{}, error) {
	i.mux.Lock()
	defer i.mux.Unlock()

	i.pageCount++
	return i.paginator.Page(i.pageCount)
}

func (i *Iterator) Prev() ([]interface{}, error) {
	i.mux.Lock()
	defer i.mux.Unlock()

	i.pageCount--
	return i.paginator.Page(i.pageCount)
}
