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

	"github.com/pkg/errors"
)

// SetWriter contains the WriteObject and Flush functions for writing objects as keys to a set.
type SetWriter struct {
	values reflect.Value
	mutex  *sync.RWMutex
}

func NewSetWriter() *SetWriter {
	return &SetWriter{
		values: reflect.ValueOf(map[interface{}]struct{}{}),
		mutex:  &sync.RWMutex{},
	}
}

func NewSetWriterWithValues(initialValues interface{}) *SetWriter {
	return &SetWriter{
		values: reflect.ValueOf(initialValues),
		mutex:  &sync.RWMutex{},
	}
}

func (sw *SetWriter) WriteObject(object interface{}) error {
	sw.mutex.Lock()
	if object == nil {
		sw.values.SetMapIndex(reflect.Zero(sw.values.Type().Key()), reflect.Zero(sw.values.Type().Elem()))
	} else {
		sw.values.SetMapIndex(reflect.ValueOf(object), reflect.Zero(sw.values.Type().Elem()))
	}
	sw.mutex.Unlock()
	return nil
}

func (sw *SetWriter) WriteObjects(objects interface{}) error {
	v := reflect.ValueOf(objects)
	if !v.IsValid() {
		return errors.Errorf("objects %#v is not valid", objects)
	}
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return errors.Errorf("objects is type %T, expecting kind array or slice", objects)
	}
	if v.IsNil() {
		return errors.Errorf("objects %#v is nil", objects)
	}
	sw.mutex.Lock()
	for i := 0; i < v.Len(); i++ {
		sw.values.SetMapIndex(v.Index(i), reflect.Zero(sw.values.Type().Elem()))
	}
	sw.mutex.Unlock()
	return nil
}

// Flush has no effect for SetWriter as all objects are immediately written to the underlying set.
func (sw *SetWriter) Flush() error {
	return nil
}

// Reset creates a new underlying set from the type of the original set.
func (sw *SetWriter) Reset() {
	sw.values = reflect.MakeSlice(reflect.TypeOf(sw.values), 0, 0)
}

func (sw *SetWriter) SliceInterface() []interface{} {
	keys := sw.values.MapKeys()
	values := make([]interface{}, 0, len(keys))
	for _, key := range keys {
		values = append(values, key.Interface())
	}
	return values
}

func (sw *SetWriter) SliceType() interface{} {
	keys := sw.values.MapKeys()
	values := reflect.MakeSlice(reflect.SliceOf(sw.values.Type().Key()), 0, len(keys))
	for _, key := range keys {
		values = reflect.Append(values, key)
	}
	return values.Interface()
}

func (sw *SetWriter) Values() interface{} {
	return sw.values.Interface()
}

func (sw *SetWriter) Iterator() *SetIterator {
	return &SetIterator{
		it:   sw.values.MapRange(),
		done: false,
	}
}
