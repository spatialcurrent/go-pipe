// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

// Writer contains the WriteObject and Flush functions for writing objects.
type Writer interface {
	WriteObject(object interface{}) error
	Flush() error
}
