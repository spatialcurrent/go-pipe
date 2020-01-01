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

// SyncWriter wraps a mutex around an underlying writer, so writes happen sequentially.
type SyncWriter struct {
	writer Writer
	mutex  *sync.Mutex
}

// NewSyncWriter returns a new SyncWriter.
func NewSyncWriter(writer Writer) *SyncWriter {
	return &SyncWriter{
		writer: writer,
		mutex:  &sync.Mutex{},
	}
}

func (sw *SyncWriter) WriteObject(object interface{}) error {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()
	return sw.writer.WriteObject(object)
}

func (sw *SyncWriter) WriteObjects(objects interface{}) error {
	values := reflect.ValueOf(objects)
	if !values.IsValid() {
		return errors.Errorf("objects %#v is not valid", objects)
	}
	if values.Kind() != reflect.Array && values.Kind() != reflect.Slice {
		return errors.Errorf("objects is type %T, expecting kind array or slice", objects)
	}
	if values.IsNil() {
		return errors.Errorf("objects %#v is nil", objects)
	}
	sw.mutex.Lock()
	defer sw.mutex.Unlock()

	if w, ok := sw.writer.(BatchWriter); ok {
		err := w.WriteObjects(objects)
		if err != nil {
			return errors.Wrapf(err, "error writing objects %#v to underlying writer", objects)
		}
	} else {
		for i := 0; i < values.Len(); i++ {
			err := sw.writer.WriteObject(values.Index(i).Interface())
			if err != nil {
				return errors.Wrapf(err, "error writing object %d of %#v to underlying writer", i, objects)
			}
		}
	}

	return nil
}

func (sw *SyncWriter) Flush() error {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()
	return sw.writer.Flush()
}

func (sw *SyncWriter) Close() error {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()
	if closer, ok := sw.writer.(interface{ Close() error }); ok {
		return closer.Close()
	}
	return nil
}
