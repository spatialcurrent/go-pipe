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

// ChannelIterator iterates over a channel of values.
type ChannelIterator struct {
	values reflect.Value
	closed bool
}

// NewChannelIterator returns a new ChannelIterator.
func NewChannelIterator(values interface{}) (*ChannelIterator, error) {
	v := reflect.ValueOf(values)
	if k := v.Type().Kind(); k != reflect.Chan {
		return nil, &ErrInvalidKind{Value: v.Type(), Expected: []reflect.Kind{reflect.Chan}}
	}
	return &ChannelIterator{values: v, closed: false}, nil
}

func (ci *ChannelIterator) Next() (interface{}, error) {
	if ci.closed {
		return nil, io.EOF
	}
	obj, ok := ci.values.Recv()
	if !ok {
		ci.closed = true
		return nil, io.EOF
	}
	return obj.Interface(), nil
}
