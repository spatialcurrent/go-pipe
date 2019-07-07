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
	"sort"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestMapIterator(t *testing.T) {

	input := map[string]struct{}{
		"a": struct{}{},
		"b": struct{}{},
		"c": struct{}{},
	}
	expected := []string{"a", "b", "c"}
	output := []string{}

	it, err := NewMapIterator(input)
	assert.Nil(t, err)

	obj, err := it.Next()
	assert.Nil(t, err)
	output = append(output, fmt.Sprint(obj))

	obj, err = it.Next()
	assert.Nil(t, err)
	output = append(output, fmt.Sprint(obj))

	obj, err = it.Next()
	assert.Nil(t, err)
	output = append(output, fmt.Sprint(obj))

	// Should return io.EOF to indicate the reader is finished
	_, err = it.Next()
	assert.Equal(t, io.EOF, err)

	// Should still return io.EOF to indicate the reader is finished
	_, err = it.Next()
	assert.Equal(t, io.EOF, err)

	sort.Strings(output)
	assert.Equal(t, expected, output)
}
