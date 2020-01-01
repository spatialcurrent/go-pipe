// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

// TransactionWriter opens up a transaction to the underlying writer
// and closes the underlying writer after the objects are written.
type TransactionWriter struct {
	open  func() (Writer, error)
	mutex *sync.RWMutex
}

// NewBufferWriter returns a new BufferWriter with the given capacity.
func NewTransactionWriter(open func() (Writer, error)) (*TransactionWriter, error) {
	if open == nil {
		return nil, errors.New("cannot create TransactionWriter: open is nil")
	}
	tw := &TransactionWriter{
		open:  open,
		mutex: &sync.RWMutex{},
	}
	return tw, nil
}

func (tw *TransactionWriter) WriteObject(object interface{}) error {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	w, err := tw.open()
	if err != nil {
		return fmt.Errorf("error opening writer: %w", err)
	}

	err = w.WriteObject(object)
	if err != nil {
		return fmt.Errorf("error writing object: %w", err)
	}

	err = w.Flush()
	if err != nil {
		return fmt.Errorf("error flushing writer: %w", err)
	}

	if closer, ok := w.(interface{ Close() error }); ok {
		err = closer.Close()
		if err != nil {
			return fmt.Errorf("error closing writer: %w", err)
		}
	}

	return nil
}

func (tw *TransactionWriter) WriteObjects(objects interface{}) error {

	values := reflect.ValueOf(objects)
	if !values.IsValid() {
		return fmt.Errorf("objects %#v is not valid", objects)
	}
	if values.Kind() != reflect.Array && values.Kind() != reflect.Slice {
		return fmt.Errorf("objects is type %T, expecting kind array or slice", objects)
	}
	if values.IsNil() {
		return fmt.Errorf("objects %#v is nil", objects)
	}

	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	w, err := tw.open()
	if err != nil {
		return fmt.Errorf("error opening writer: %w", err)
	}

	if bw, ok := w.(BatchWriter); ok {
		err := bw.WriteObjects(objects)
		if err != nil {
			return fmt.Errorf("error writing objects: %w", err)
		}
	} else {
		for i := 0; i < values.Len(); i++ {
			err := w.WriteObject(values.Index(i).Interface())
			if err != nil {
				return fmt.Errorf("error writing object %d: %w", i, err)
			}
		}
	}

	err = w.Flush()
	if err != nil {
		return fmt.Errorf("error flushing writer: %w", err)
	}

	if closer, ok := w.(interface{ Close() error }); ok {
		err := closer.Close()
		if err != nil {
			return fmt.Errorf("error closing writer: %w", err)
		}
	}

	return nil
}

func (tw *TransactionWriter) Flush() error {
	return nil
}

func (tw *TransactionWriter) Close() error {
	return nil
}
