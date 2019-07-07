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

// This examples hows how to use a channel iterator to read from a channel.
func ExampleChannelIterator() {

	c := make(chan interface{}, 1000)

	c <- "a"
	c <- "b"
	c <- "c"
	close(c)

	it, err := NewChannelIterator(c)
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
	// Output: a
	//b
	//c
}
