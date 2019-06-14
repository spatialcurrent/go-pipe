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

// SliceIterator iterates over an array or slice of values.
type SliceIterator struct {
	values reflect.Value
	cursor int
}

// NewSliceIterator returns a new SliceIterator.
func NewSliceIterator(values interface{}) (*SliceIterator, error) {
	v := reflect.ValueOf(values)
	if k := v.Type().Kind(); k != reflect.Array && k != reflect.Slice {
		return nil, &ErrInvalidKind{Value: v.Type(), Expected: []reflect.Kind{reflect.Array, reflect.Slice}}
	}
	return &SliceIterator{values: v, cursor: 0}, nil
}

func (si *SliceIterator) Next() (interface{}, error) {
	if si.cursor >= si.values.Len() {
		return nil, io.EOF
	}
	obj := si.values.Index(si.cursor).Interface()
	si.cursor += 1
	return obj, nil
}
