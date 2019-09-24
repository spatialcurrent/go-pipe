// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"reflect"
)

// ChannelWriter passes each object to the callback function
type ChannelWriter struct {
	channel reflect.Value
}

func NewChannelWriter(channel interface{}) (*ChannelWriter, error) {
	v := reflect.ValueOf(channel)
	if k := v.Type().Kind(); k != reflect.Chan {
		return nil, &ErrInvalidKind{Value: v.Type(), Expected: []reflect.Kind{reflect.Chan}}
	}
	return &ChannelWriter{channel: v}, nil
}

func (cw *ChannelWriter) WriteObject(object interface{}) error {
	cw.channel.Send(reflect.ValueOf(object))
	return nil
}

func (cw *ChannelWriter) Flush() error {
	return nil
}

func (cw *ChannelWriter) Close() error {
	cw.channel.Close()
	return nil
}
