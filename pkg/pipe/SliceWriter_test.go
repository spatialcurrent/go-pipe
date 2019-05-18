// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestSliceWriter(t *testing.T) {

	w := NewSliceWriter()

	err := w.WriteObject("a")
	assert.Nil(t, err)

	err = w.WriteObject("b")
	assert.Nil(t, err)

	err = w.WriteObject("3")
	assert.Nil(t, err)

	err = w.WriteObject(1)
	assert.Nil(t, err)

	err = w.WriteObject(2)
	assert.Nil(t, err)

	err = w.WriteObject(3)
	assert.Nil(t, err)

	values := w.Values()
	assert.Equal(t, []interface{}{"a", "b", "3", 1, 2, 3}, values)
}
