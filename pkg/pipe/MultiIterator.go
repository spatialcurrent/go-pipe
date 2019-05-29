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
// MultiIterator returns an Iterator that's the logical concatenation of the provided input iterators.
// They're read sequentially. Once all iterators have returned EOF, Next will return EOF.
// If any of the iterators return a non-nil, non-EOF error, Next will return that error.
// This approach is similiar to the io.MultiReader.
//
//	- https://godoc.org/io#MultiReader
type MultiIterator struct {
	iterators []Iterator
}

// NewMultiIterator returns a new MultiIterator that serialy iterators through the given iterators.
func NewMultiIterator(iterators ...Iterator) *MultiIterator {
	return &MultiIterator{
		iterators: iterators,
	}
}

func (mi *MultiIterator) Push(it ...Iterator) {
	mi.iterators = append(mi.iterators, it...)
}

func (mi *MultiIterator) Next() (interface{}, error) {
	for len(mi.iterators) > 0 {
		if len(mi.iterators) == 1 {
			if r, ok := mi.iterators[0].(*MultiIterator); ok {
				mi.iterators = r.iterators
				continue
			}
		}
		obj, err := mi.iterators[0].Next()
		if err != nil {
			if err == io.EOF {
				if len(mi.iterators) == 1 {
					mi.iterators = make([]Iterator, 0)
					return obj, io.EOF
				} else {
					mi.iterators = mi.iterators[1:]
					continue
				}
			}
		}
		return obj, nil
	}
	return nil, io.EOF
}
