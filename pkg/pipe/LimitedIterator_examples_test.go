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

func ExampleLimitedIterator() {

	si, err := NewSliceIterator([]interface{}{"a", "b", "c"})
	if err != nil {
		panic(err)
	}

	li := NewLimitedIterator(si, 2)

	obj, err := li.Next()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)

	obj, err = li.Next()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)

	obj, err = li.Next()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)

	// Output: a
	//b
	//EOF
	//<nil>
}
