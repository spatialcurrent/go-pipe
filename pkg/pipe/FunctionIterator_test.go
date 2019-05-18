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

func TestFunctionIterator(t *testing.T) {

	values := []interface{}{"a", "b", 1, 2, 3, true, false}

	cursor := 0
	it := NewFunctionIterator(func() (interface{}, error) {
		if cursor >= len(values) {
			return nil, io.EOF
		}
		obj := values[cursor]
		cursor++
		return obj, nil
	})
	assert.NotNil(t, it)

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
