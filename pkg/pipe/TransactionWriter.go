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
	opener func() (Writer, error)
	closer func(w Writer) error
	mutex  *sync.Mutex
}

// NewTransactionWriter returns a new TransactionWriter with the opener function and optional closer function.
func NewTransactionWriter(opener func() (Writer, error), closer func(w Writer) error) (*TransactionWriter, error) {
	if opener == nil {
		return nil, errors.New("cannot create TransactionWriter: open is nil")
	}
	tw := &TransactionWriter{
		opener: opener,
		closer: closer,
		mutex:  &sync.Mutex{},
	}
	return tw, nil
}

func (tw *TransactionWriter) WriteObject(object interface{}) error {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	w, err := tw.opener()
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

	if tw.closer != nil {
		err = tw.closer(w)
		if err != nil {
			return fmt.Errorf("error closing writer: %w", err)
		}
	} else {
		if closer, ok := w.(interface{ Close() error }); ok {
			err = closer.Close()
			if err != nil {
				return fmt.Errorf("error closing writer: %w", err)
			}
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

	w, err := tw.opener()
	if err != nil {
		return fmt.Errorf("error opening writer: %w", err)
	}

	if bw, ok := w.(BatchWriter); ok {
		err = bw.WriteObjects(objects)
		if err != nil {
			return fmt.Errorf("error writing objects: %w", err)
		}
	} else {
		for i := 0; i < values.Len(); i++ {
			err = w.WriteObject(values.Index(i).Interface())
			if err != nil {
				return fmt.Errorf("error writing object %d: %w", i, err)
			}
		}
	}

	err = w.Flush()
	if err != nil {
		return fmt.Errorf("error flushing writer: %w", err)
	}

	if tw.closer != nil {
		err = tw.closer(w)
		if err != nil {
			return fmt.Errorf("error closing writer: %w", err)
		}
	} else {
		if closer, ok := w.(interface{ Close() error }); ok {
			err = closer.Close()
			if err != nil {
				return fmt.Errorf("error closing writer: %w", err)
			}
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
