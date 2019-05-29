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

func TestMultiIterator(t *testing.T) {

	sia, err := NewSliceIterator([]interface{}{"a", "b", "c"})
	assert.Nil(t, err)

	sib, err := NewSliceIterator([]interface{}{1, 2, 3})
	assert.Nil(t, err)

	mi := NewMultiIterator(sia, sib)

	obj, err := mi.Next()
	assert.Nil(t, err)
	assert.Equal(t, "a", obj)

	obj, err = mi.Next()
	assert.Nil(t, err)
	assert.Equal(t, "b", obj)

	obj, err = mi.Next()
	assert.Nil(t, err)
	assert.Equal(t, "c", obj)

	obj, err = mi.Next()
	assert.Nil(t, err)
	assert.Equal(t, 1, obj)

	obj, err = mi.Next()
	assert.Nil(t, err)
	assert.Equal(t, 2, obj)

	obj, err = mi.Next()
	assert.Nil(t, err)
	assert.Equal(t, 3, obj)

	// Should return io.EOF to indicate the reader is finished
	obj, err = mi.Next()
	assert.Equal(t, io.EOF, err)
	assert.Nil(t, obj)

	// Should still return io.EOF to indicate the reader is finished
	obj, err = mi.Next()
	assert.Equal(t, io.EOF, err)
	assert.Nil(t, obj)
}
