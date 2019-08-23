// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"fmt"
	"io"
)

// This examples shows how to use a map iterator to iterate through the keys of a map.
func ExampleSetIterator() {

	input := map[string]struct{}{
		"a": struct{}{},
		"b": struct{}{},
		"c": struct{}{},
	}

	it, err := NewSetIterator(input)
	if err != nil {
		panic(err)
	}

	for {
		obj, err := it.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(obj)
	}
}
