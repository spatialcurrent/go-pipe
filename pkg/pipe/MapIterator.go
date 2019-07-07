// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"io"
	"reflect"
)

// MapIterator iterates over the keys in a map.
type MapIterator struct {
	it   *reflect.MapIter
	done bool
}

// NewMapIterator returns a new MapIterator.
func NewMapIterator(values interface{}) (*MapIterator, error) {
	v := reflect.ValueOf(values)
	if k := v.Type().Kind(); k != reflect.Map {
		return nil, &ErrInvalidKind{Value: v.Type(), Expected: []reflect.Kind{reflect.Map}}
	}
	return &MapIterator{it: v.MapRange()}, nil
}

func (mi *MapIterator) Next() (interface{}, error) {
	if mi.done {
		return nil, io.EOF
	}
	mi.done = !mi.it.Next()
	if mi.done {
		return nil, io.EOF
	}
	return mi.it.Key().Interface(), nil
}
