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

	"github.com/stretchr/testify/assert"
)

func TestLimitediterator(t *testing.T) {

	si, err := NewSliceIterator([]interface{}{"a", "b", "c"})
	assert.Nil(t, err)

	li := NewLimitedIterator(si, 2)
	assert.Equal(t, 0, li.Count())
	assert.Equal(t, 2, li.Limit())

	obj, err := li.Next()
	assert.Nil(t, err)
	assert.Equal(t, "a", obj)
	assert.Equal(t, 1, li.Count())

	obj, err = li.Next()
	assert.Nil(t, err)
	assert.Equal(t, "b", obj)
	assert.Equal(t, 2, li.Count())

	// Should return io.EOF to indicate the reader is finished
	obj, err = li.Next()
	assert.Equal(t, io.EOF, err)
	assert.Nil(t, obj)

	// Should still return io.EOF to indicate the reader is finished
	obj, err = li.Next()
	assert.Equal(t, io.EOF, err)
	assert.Nil(t, obj)
}
