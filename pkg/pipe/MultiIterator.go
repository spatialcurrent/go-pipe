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

// Iterator contains the Next function that returns an object until it returns an io.EOF error.
type MultiIterator struct {
	iterators []Iterator
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
				}
			}
		}
		return obj, nil
	}
	return nil, io.EOF
}
