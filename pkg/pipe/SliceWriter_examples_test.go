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

// This examples shows how to use a slice writer to write values to a slice.
func ExampleSliceWriter() {

	w := NewSliceWriter()

	err := w.WriteObject("a")
	if err != nil {
		panic(err)
	}

	err = w.WriteObject("b")
	if err != nil {
		panic(err)
	}

	err = w.WriteObject("c")
	if err != nil {
		panic(err)
	}

	values := w.Values() // get values written to slice

	fmt.Println(values)
	// Output: [a b c]
}
