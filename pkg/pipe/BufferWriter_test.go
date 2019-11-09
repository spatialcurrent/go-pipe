// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBufferWriter(t *testing.T) {

	sw := NewSliceWriter()
	capacity := 3
	bw := NewBufferWriter(sw, capacity)

	err := bw.WriteObject("a")
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{}, sw.Values())

	err = bw.WriteObject("b")
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{}, sw.Values())

	err = bw.WriteObject("3")
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{"a", "b", "3"}, sw.Values())

	err = bw.WriteObjects([]interface{}{1, 2, 3})
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{"a", "b", "3", 1, 2, 3}, sw.Values())
}
