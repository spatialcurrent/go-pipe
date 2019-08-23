// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"io"
)

//
// LimitedIterator returns an iterator that reads up to a given number of objects from the underlying reader.
// Once the maximum number of objects has been read or the underlying iterator has returned io.EOF, Next will return io.EOF.
// If the underlying iterator returns a non-nil, non-EOF error, Next will return that error.
// This approach is similiar to the io.LimitedReader.
//
//	- https://godoc.org/io#LimitedReader
type LimitedIterator struct {
	iterator Iterator
	limit    int
	count    int
}

// NewLimitedIterator returns a new LimitIterator that serialy iterators through the given iterators.
func NewLimitedIterator(iterator Iterator, limit int) *LimitedIterator {
	return &LimitedIterator{
		iterator: iterator,
		limit:    limit,
		count:    0,
	}
}

// Count returns the current count of objects returned.
func (li *LimitedIterator) Count() int {
	return li.count
}

// Limit returns the set limit.
func (li *LimitedIterator) Limit() int {
	return li.limit
}

func (li *LimitedIterator) Next() (interface{}, error) {
	if li.iterator == nil {
		return nil, io.EOF
	}
	if li.count == li.limit {
		return nil, io.EOF
	}
	obj, err := li.iterator.Next()
	li.count++
	return obj, err
}
