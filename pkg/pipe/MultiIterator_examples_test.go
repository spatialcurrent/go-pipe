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

// This examples shows how to use a multi iterator to iterate through multiple subiterators.
func ExampleMultiIterator() {

	sia, err := NewSliceIterator([]interface{}{"a", "b", "c"})
	if err != nil {
		panic(err)
	}

	sib, err := NewSliceIterator([]interface{}{1, 2, 3})
	if err != nil {
		panic(err)
	}

	mi := NewMultiIterator(sia, sib)

	obj, err := mi.Next()
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)

	obj, err = mi.Next()
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)

	obj, err = mi.Next()
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)

	obj, err = mi.Next()
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)

	obj, err = mi.Next()
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)

	obj, err = mi.Next()
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)

	// Output: a
	//b
	//c
	//1
	//2
	//3
}
