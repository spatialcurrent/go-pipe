// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package pipe includes interfaces and concerete classes for piping objects from inputs to outputs.
// See the examples below and tests for usage.
package pipe

import (
	"fmt"
	"os"
)

var (
	FilterNotNil = func(object interface{}) (bool, error) {
		return object != nil, nil
	}
	FilterString = func(object interface{}) (bool, error) {
		_, ok := object.(string)
		return ok, nil
	}
	WriterStdout = NewFunctionWriter(func(object interface{}) error {
		_, err := fmt.Println(object)
		return err
	})
	WriterStderr = NewFunctionWriter(func(object interface{}) error {
		_, err := fmt.Fprintln(os.Stderr, object)
		return err
	})
)
