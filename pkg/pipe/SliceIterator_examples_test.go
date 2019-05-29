// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"fmt"
)

// This examples shows how to use a slice iterator to iterate through a slice of values.
func ExampleSliceIterator() {

	it, err := NewSliceIterator([]interface{}{"a", "b", "c"})
	if err != nil {
		panic(err)
	}

	obj, err := it.Next()
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)

	obj, err = it.Next()
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)

	obj, err = it.Next()
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)

	// Output: a
	//b
	//c
}
