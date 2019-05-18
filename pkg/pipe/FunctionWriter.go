// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

// FunctionWriter passes each object to the callback function
type FunctionWriter struct {
	callback func(object interface{}) error
}

func NewFunctionWriter(callback func(object interface{}) error) *FunctionWriter {
	return &FunctionWriter{
		callback: callback,
	}
}

func (sw *FunctionWriter) WriteObject(object interface{}) error {
	return sw.callback(object)
}

func (sw *FunctionWriter) Flush() error {
	return nil
}
