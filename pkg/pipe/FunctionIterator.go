// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

// FunctionIterator provides a simple wrapper over a callback function to allow it to iterate.
type FunctionIterator struct {
	callback func() (interface{}, error)
}

func NewFunctionIterator(callback func() (interface{}, error)) *FunctionIterator {
	return &FunctionIterator{callback: callback}
}

func (fi *FunctionIterator) Next() (interface{}, error) {
	return fi.callback()
}
