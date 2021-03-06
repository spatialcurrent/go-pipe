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

// BufferWriter wraps a buffer around an underlying writer.
// Once the buffer reaches capacity, it writes its values to the underlying writer.
// The Flush method will propagate to the underlying writer.
type BufferWriter struct {
	writer   Writer
	values   reflect.Value
	mutex    *sync.RWMutex
	capacity int
}

// NewBufferWriter returns a new BufferWriter with the given capacity.
func NewBufferWriter(writer Writer, capacity int) *BufferWriter {
	return &BufferWriter{
		writer:   writer,
		values:   reflect.ValueOf(make([]interface{}, 0, capacity)),
		mutex:    &sync.RWMutex{},
		capacity: capacity,
	}
}

func (bw *BufferWriter) WriteObject(object interface{}) error {
	bw.mutex.Lock()
	defer bw.mutex.Unlock()
	if object == nil {
		bw.values = reflect.Append(bw.values, reflect.Zero(bw.values.Type().Elem()))
	} else {
		bw.values = reflect.Append(bw.values, reflect.ValueOf(object))
	}
	if bw.values.Len() == bw.capacity {
		if w, ok := bw.writer.(BatchWriter); ok {
			err := w.WriteObjects(bw.values.Interface())
			if err != nil {
				return fmt.Errorf("error writing objects %#v to underlying writer: %w", bw.values.Interface(), err)
			}
		} else {
			for i := 0; i < bw.values.Len(); i++ {
				err := bw.writer.WriteObject(bw.values.Index(i).Interface())
				if err != nil {
					return fmt.Errorf("error writing object %d of %#v to underlying writer: %w", i, bw.values.Interface(), err)
				}
			}
		}
		// reset the buffer
		bw.values = reflect.MakeSlice(bw.values.Type(), 0, bw.capacity)
	}
	return nil
}

func (bw *BufferWriter) WriteObjects(objects interface{}) error {
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
	bw.mutex.Lock()
	defer bw.mutex.Unlock()
	for i := 0; i < v.Len(); i++ {
		bw.values = reflect.Append(bw.values, v.Index(i))
		if bw.values.Len() == bw.capacity {
			if w, ok := bw.writer.(BatchWriter); ok {
				err := w.WriteObjects(bw.values.Interface())
				if err != nil {
					return fmt.Errorf("error writing objects %#v to underlying writer: %w", bw.values.Interface(), err)
				}
			} else {
				for j := 0; j < bw.values.Len(); j++ {
					err := bw.writer.WriteObject(bw.values.Index(j).Interface())
					if err != nil {
						return fmt.Errorf("error writing object %d of %#v to underlying writer: %w", i, bw.values.Interface(), err)
					}
				}
			}
			// reset the buffer
			bw.values = reflect.MakeSlice(bw.values.Type(), 0, bw.capacity)
		}
	}
	return nil
}

func (bw *BufferWriter) Flush() error {
	bw.mutex.Lock()
	defer bw.mutex.Unlock()
	if w, ok := bw.writer.(BatchWriter); ok {
		err := w.WriteObjects(bw.values.Interface())
		if err != nil {
			return fmt.Errorf("error writing objects %#v to underlying writer: %w", bw.values.Interface(), err)
		}
	} else {
		for i := 0; i < bw.values.Len(); i++ {
			err := bw.writer.WriteObject(bw.values.Index(i).Interface())
			if err != nil {
				return fmt.Errorf("error writing object %d of %#v to underlying writer: %w", i, bw.values.Interface(), err)
			}
		}
	}
	// reset the buffer
	bw.values = reflect.MakeSlice(bw.values.Type(), 0, bw.capacity)
	return bw.writer.Flush()
}

func (bw *BufferWriter) Close() error {
	bw.mutex.Lock()
	defer bw.mutex.Unlock()
	if closer, ok := bw.writer.(interface{ Close() error }); ok {
		return closer.Close()
	}
	return nil
}

// Reset creates a new underlying slice from the type of the original slice.
func (bw *BufferWriter) Reset() {
	bw.values = reflect.MakeSlice(bw.values.Type(), 0, bw.capacity)
}
