// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"fmt"
	"sort"
)

// This examples shows how to use a SetWriter to write values to a set as keys to a map.
func ExampleSetWriter() {

	w := NewSetWriter()

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

	err = w.WriteObject("a")
	if err != nil {
		panic(err)
	}

	values := w.SliceInterface() // get values written as slice of type []interface{}

	// Sort the returned values
	sort.Slice(values, func(i, j int) bool {
		return fmt.Sprint(values[i]) < fmt.Sprint(values[j])
	})

	// Print the values
	fmt.Println(values)
	// Output: [a b c]
}
