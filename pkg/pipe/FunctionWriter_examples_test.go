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

// This examples hows how to use a function writer to call a callback with each object.
func ExampleFunctionWriter() {

	w := NewFunctionWriter(func(object interface{}) error {
		fmt.Println(object) // print to stdout
		return nil
	})

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

	// Output: a
	//b
	//c
}
