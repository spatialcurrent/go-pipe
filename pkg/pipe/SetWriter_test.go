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

func TestSetWriter(t *testing.T) {

	expectedValues := map[interface{}]struct{}{
		"a": struct{}{},
		"b": struct{}{},
		"c": struct{}{},
	}

	w := NewSetWriter()

	err := w.WriteObject("a")
	assert.Nil(t, err)

	err = w.WriteObject("b")
	assert.Nil(t, err)

	err = w.WriteObject("c")
	assert.Nil(t, err)

	err = w.WriteObject("a")
	assert.Nil(t, err)

	values := w.Values() // get values written as slice of type []interface{}

	assert.Equal(t, expectedValues, values)
}
