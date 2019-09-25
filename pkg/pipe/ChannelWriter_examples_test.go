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

// This examples hows how to use a channel writer to write to a channel.
func ExampleChannelWriter() {

	c := make(chan interface{}, 1000)

	w, err := NewChannelWriter(c)
	if err != nil {
		panic(err)
	}

	err = w.WriteObject("a")
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

	close(c)

	for v := range c {
		fmt.Println(v)
	}
	// Output: a
	//b
	//c
}
