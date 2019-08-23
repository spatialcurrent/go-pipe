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
	"strings"
)

// This examples hows how to use a builder.
func ExampleBuilder() {

	input := []interface{}{"a", "b", "c", 1, 2, 3, false, true}

	it, err := NewSliceIterator(input)
	if err != nil {
		panic(err)
	}

	w := NewSliceWriterWithValues([]string{})

	b := NewBuilder().
		Input(it).
		Filter(func(inputObject interface{}) (bool, error) {
			// filter to only include strings
			_, ok := inputObject.(string)
			return ok, nil
		}).
		Output(w)

	err = b.Run()
	if err != nil {
		panic(err)
	}
	// the slice writer preserves the type of the initial values.
	fmt.Println(strings.Join(w.Values().([]string), "\n"))
	// Output: a
	//b
	//c
}

// This examples hows how to use a builder.
func ExampleBuilder_sliceToSet() {

	// the initial values with duplicates
	input := []interface{}{"a", "b", "c", "a", "b"}

	it, err := NewSliceIterator(input)
	if err != nil {
		panic(err)
	}

	w := NewSetWriter()

	err = NewBuilder().Input(it).Output(w).Run()
	if err != nil {
		panic(err)
	}

	values := w.SliceInterface() // get values written as slice of type []interface{}

	// Sort the returned values
	sort.Slice(values, func(i, j int) bool {
		return fmt.Sprint(values[i]) < fmt.Sprint(values[j])
	})

	for _, value := range values {
		fmt.Println(value)
	}
	// Output: a
	//b
	//c
}
