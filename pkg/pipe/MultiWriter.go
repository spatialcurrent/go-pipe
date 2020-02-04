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

	"github.com/pkg/errors"
)

// MultiWriter creates a writer that duplicates its writes to all the provided writers.
type MultiWriter struct {
	writers []Writer
}

func NewMultiWriter(writers ...Writer) *MultiWriter {
	return &MultiWriter{
		writers: writers,
	}
}

func (mw *MultiWriter) WriteObject(object interface{}) error {
	for i, w := range mw.writers {
		if err := w.WriteObject(object); err != nil {
			return errors.Wrapf(err, "error writing object to writer %d", i)
		}
	}
	return nil
}

func (mw *MultiWriter) WriteObjects(objects interface{}) error {
	for i, w := range mw.writers {
		if bw, ok := w.(BatchWriter); ok {
			if err := bw.WriteObjects(objects); err != nil {
				return errors.Wrapf(err, "error writing objects to writer %d", i)
			}
			continue
		}
		if slc, ok := objects.([]interface{}); ok {
			for _, object := range slc {
				if err := w.WriteObject(object); err != nil {
					return errors.Wrapf(err, "error writing object to writer %d", i)
				}
			}
			continue
		}
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
		for i := 0; i < values.Len(); i++ {
			err := w.WriteObject(values.Index(i).Interface())
			if err != nil {
				return fmt.Errorf("error writing object %d: %w", i, err)
			}
		}
		return nil
	}
	return nil
}

func (mw *MultiWriter) Flush() error {
	for i, w := range mw.writers {
		if err := w.Flush(); err != nil {
			return errors.Wrapf(err, "error flushing writer %d", i)
		}
	}
	return nil
}
