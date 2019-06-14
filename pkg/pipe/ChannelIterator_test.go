// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"io"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestChannelIterator(t *testing.T) {

	c := make(chan interface{}, 10)
	c <- "a"
	c <- "b"
	c <- 1
	c <- 2
	c <- 3
	c <- true
	c <- false
	close(c)

	it, err := NewChannelIterator(c)
	assert.Nil(t, err)

	obj, err := it.Next()
	assert.Nil(t, err)
	assert.Equal(t, "a", obj)

	obj, err = it.Next()
	assert.Nil(t, err)
	assert.Equal(t, "b", obj)

	obj, err = it.Next()
	assert.Nil(t, err)
	assert.Equal(t, 1, obj)

	obj, err = it.Next()
	assert.Nil(t, err)
	assert.Equal(t, 2, obj)

	obj, err = it.Next()
	assert.Nil(t, err)
	assert.Equal(t, 3, obj)

	obj, err = it.Next()
	assert.Nil(t, err)
	assert.Equal(t, true, obj)

	obj, err = it.Next()
	assert.Nil(t, err)
	assert.Equal(t, false, obj)

	// Should return io.EOF to indicate the reader is finished
	obj, err = it.Next()
	assert.Equal(t, io.EOF, err)
	assert.Nil(t, obj)

	// Should still return io.EOF to indicate the reader is finished
	obj, err = it.Next()
	assert.Equal(t, io.EOF, err)
	assert.Nil(t, obj)
}
