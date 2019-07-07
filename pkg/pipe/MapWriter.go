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

// MapWriter contains the WriteObject and Flush functions for writing objects as keys to a map
type MapWriter struct {
	values reflect.Value
	mutex  *sync.RWMutex
}

func NewMapWriter() *MapWriter {
	return &MapWriter{
		values: reflect.ValueOf(map[interface{}]struct{}{}),
		mutex:  &sync.RWMutex{},
	}
}

func NewMapWriterWithValues(initialValues interface{}) *MapWriter {
	return &MapWriter{
		values: reflect.ValueOf(initialValues),
		mutex:  &sync.RWMutex{},
	}
}

func (mw *MapWriter) WriteObject(object interface{}) error {
	mw.mutex.Lock()
	if object == nil {
		mw.values.SetMapIndex(reflect.Zero(mw.values.Type().Key()), reflect.Zero(mw.values.Type().Elem()))
	} else {
		mw.values.SetMapIndex(reflect.ValueOf(object), reflect.Zero(mw.values.Type().Elem()))
	}
	mw.mutex.Unlock()
	return nil
}

func (mw *MapWriter) Flush() error {
	return nil
}

// Resets the writer and clears all existing values.
func (mw *MapWriter) Reset() {
	mw.values = reflect.MakeSlice(reflect.TypeOf(mw.values), 0, 0)
}

func (mw *MapWriter) SliceInterface() []interface{} {
	keys := mw.values.MapKeys()
	values := make([]interface{}, 0, len(keys))
	for _, key := range keys {
		values = append(values, key.Interface())
	}
	return values
}

func (mw *MapWriter) SliceType() interface{} {
	keys := mw.values.MapKeys()
	values := reflect.MakeSlice(reflect.SliceOf(mw.values.Type().Key()), 0, len(keys))
	for _, key := range keys {
		values = reflect.Append(values, key)
	}
	return values.Interface()
}

func (mw *MapWriter) Values() interface{} {
	return mw.values.Interface()
}

func (mw *MapWriter) Iterator() *MapIterator {
	return &MapIterator{
		it:   mw.values.MapRange(),
		done: false,
	}
}
