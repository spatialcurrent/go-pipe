// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"reflect"
	"sync"
)

// Writer contains the WriteObject and Flush functions for writing objects.
type SliceWriter struct {
	values []interface{}
	mutex  *sync.RWMutex
}

func NewSliceWriter() *SliceWriter {
	return &SliceWriter{
		values: make([]interface{}, 0),
		mutex:  &sync.RWMutex{},
	}
}

func NewSliceWriterWithCapacity(capacity int) *SliceWriter {
	return &SliceWriter{
		values: make([]interface{}, 0, capacity),
	}
}

func (sw *SliceWriter) WriteObject(object interface{}) error {
	sw.mutex.Lock()
	sw.values = append(sw.values, object)
	sw.mutex.Unlock()
	return nil
}

func (sw *SliceWriter) Flush() error {
	return nil
}

func (sw *SliceWriter) Values() []interface{} {
	return sw.values
}

func (sw *SliceWriter) Iterator() *SliceIterator {
	return &SliceIterator{
		values: reflect.ValueOf(sw.values),
		cursor: 0,
	}
}
