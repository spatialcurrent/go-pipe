// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"fmt"
	"reflect"
	"sync"
)

// SliceWriter contains the WriteObject and Flush functions for writing objects to an underlying slice.
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

func (sw *SliceWriter) WriteObjects(objects interface{}) error {
	v := reflect.ValueOf(objects)
	if !v.IsValid() {
		return fmt.Errorf("objects %#v is not valid", objects)
	}
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return fmt.Errorf("objects is type %T, expecting kind array or slice", objects)
	}
	if v.IsNil() {
		return fmt.Errorf("objects %#v is nil", objects)
	}
	sw.mutex.Lock()
	for i := 0; i < v.Len(); i++ {
		sw.values = reflect.Append(sw.values, v.Index(i))
	}
	sw.mutex.Unlock()
	return nil
}

func (sw *SliceWriter) Flush() error {
	return nil
}

// Reset creates a new underlying slice from the type of the original slice.
func (sw *SliceWriter) Reset() {
	sw.values = reflect.MakeSlice(reflect.TypeOf(sw.values), 0, 0)
}

func (sw *SliceWriter) Values() interface{} {
	if !sw.values.IsValid() {
		return nil
	}
	return sw.values.Interface()
}

func (sw *SliceWriter) Iterator() *SliceIterator {
	return &SliceIterator{
		values: sw.values,
		cursor: 0,
	}
}
