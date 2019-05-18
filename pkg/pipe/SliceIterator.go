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

import (
	"github.com/pkg/errors"
)

// SliceIterator iterates over an array or slice of values.
type SliceIterator struct {
	values reflect.Value
	cursor int
}

func NewSliceIterator(values interface{}) (*SliceIterator, error) {
	v := reflect.ValueOf(values)
	if k := v.Type().Kind(); k != reflect.Array && k != reflect.Slice {
		return nil, errors.New("invalid type of SliceIterator")
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
