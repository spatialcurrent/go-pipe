// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

// Iterator contains the Next function that returns an object until it returns an io.EOF error.
type Iterator interface {
	Next() (interface{}, error)
}
