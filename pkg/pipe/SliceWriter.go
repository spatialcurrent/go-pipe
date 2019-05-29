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
	values reflect.Value
	mutex  *sync.RWMutex
}

func NewSliceWriter() *SliceWriter {
	return &SliceWriter{
		values: reflect.ValueOf(make([]interface{}, 0)),
		mutex:  &sync.RWMutex{},
	}
}

func NewSliceWriterWithValues(initialValues interface{}) *SliceWriter {
	return &SliceWriter{
		values: reflect.ValueOf(initialValues),
		mutex:  &sync.RWMutex{},
	}
}

func NewSliceWriterWithCapacity(initialCapacity int) *SliceWriter {
	return &SliceWriter{
		values: reflect.ValueOf(make([]interface{}, 0, initialCapacity)),
		mutex:  &sync.RWMutex{},
	}
}

func (sw *SliceWriter) WriteObject(object interface{}) error {
	sw.mutex.Lock()
	if object == nil {
		sw.values = reflect.Append(sw.values, reflect.Zero(sw.values.Type().Elem()))
	} else {
		sw.values = reflect.Append(sw.values, reflect.ValueOf(object))
	}
	sw.mutex.Unlock()
	return nil
}

func (sw *SliceWriter) Flush() error {
	return nil
}

// Resets the writer and clears all existing values.
func (sw *SliceWriter) Reset() {
	sw.values = reflect.MakeSlice(reflect.TypeOf(sw.values), 0, 0)
}

func (sw *SliceWriter) Values() interface{} {
	return sw.values.Interface()
}

func (sw *SliceWriter) Iterator() *SliceIterator {
	return &SliceIterator{
		values: sw.values,
		cursor: 0,
	}
}
