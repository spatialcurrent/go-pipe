// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"io"
)

func ReadAll(it Iterator) ([]interface{}, error) {
	objects := make([]interface{}, 0)
	for {
		object, err := it.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return objects, err
		}
		objects = append(objects, object)
	}
	return objects, nil
}
