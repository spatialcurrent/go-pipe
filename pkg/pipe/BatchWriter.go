// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

// BatchWriter contains the WriteObjects and Flush functions for writing a batch of objects.
type BatchWriter interface {
	WriteObjects(objects interface{}) error
	Flush() error
}
