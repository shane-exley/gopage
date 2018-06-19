package gopage

import (
	"errors"
	"math"
	"reflect"
)

//DefaultPageSize default size of a page applicable if one is not manually set.
//Its value is 10.
const DefaultPageSize = 10

//PageFetcher defines the method for the interface
type PageFetcher interface {
	Page(i int) (interface{}, error)
	Fetch(offset int, limit int) (interface{}, error)
	SetPageSize(s int) error
	GetPageSize() int
	GetPageCount() int
}

//Paginator the primary struct around which all methods are planned
//This struct's instance hold the actual slice data and its meta information. alongwith the pagination options.
type Paginator struct {
	payloadSlice reflect.Value
	payloadLen   int
	pageSize     int
	pageCount    int
}

//ErrNotSlice is returned when the constructor is passed an item which is not a slice.
var ErrNotSlice error = errors.New("Input not a slice")

//ErrinvalidPageSize is returned when the SetPageSize method is passed a non-positive integer
var ErrInvalidPageSize error = errors.New("Page size not a positive integer")

//ErrOverflow is returned when the requested page lies beyond the size of the slice
var ErrOverflow error = errors.New("Page out of bounds of slice")

//NewPaginator This function creates a new instance of the Paginator struct. It sets the default
//page size which can be changed later.
func NewPaginator(payload interface{}) (PageFetcher, error) {
	s := reflect.ValueOf(payload)
	if s.Kind() != reflect.Slice {
		return nil, ErrNotSlice
	}

	pc := math.Ceil(float64(s.Len()) / float64(DefaultPageSize))

	return &Paginator{payloadSlice: s, pageSize: DefaultPageSize, payloadLen: s.Len(), pageCount: int(pc)}, nil
}

//SetPageSize This method sets the page size for the paged retrival of the slice.
func (p *Paginator) SetPageSize(s int) error {
	if s <= 0 {
		return ErrInvalidPageSize
	}
	p.pageSize = s

	pc := math.Ceil(float64(p.payloadLen) / float64(s))
	p.pageCount = int(pc)

	return nil
}

//GetPageSize This method returns the page size of a Paginator
func (p *Paginator) GetPageSize() int {
	return p.pageSize
}

//GetPageCount This method returns the number of pages of a Paginator
func (p *Paginator) GetPageCount() int {
	return p.pageCount
}

//Page This method returns the snapshot of the slice on the ith page.
func (p *Paginator) Page(i int) (interface{}, error) {
	start := (i - 1) * p.pageSize
	end := i * p.pageSize

	if start >= p.payloadLen {
		return nil, ErrOverflow
	} else if end > p.payloadLen {
		return p.payloadSlice.Slice(start, p.payloadLen).Interface(), nil
	}

	return p.payloadSlice.Slice(start, end).Interface(), nil
}

//Fetch It returns the slice for the given offset and limit
func (p *Paginator) Fetch(offset int, limit int) (interface{}, error) {
	start := offset
	end := offset + limit

	if start >= p.payloadLen {
		return nil, ErrOverflow
	} else if end > p.payloadLen {
		return p.payloadSlice.Slice(start, p.payloadLen).Interface(), nil
	}

	return p.payloadSlice.Slice(start, end).Interface(), nil
}
