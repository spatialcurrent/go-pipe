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

func TestFunctionWriter(t *testing.T) {

	values := make([]interface{}, 0)

	w := NewFunctionWriter(func(object interface{}) error {
		values = append(values, object)
		return nil
	})

	err := w.WriteObject("a")
	assert.NoError(t, err)
	assert.Equal(t, values, []interface{}{"a"}, values)

	err = w.WriteObject("b")
	assert.NoError(t, err)
	assert.Equal(t, values, []interface{}{"a", "b"}, values)

	err = w.WriteObject("c")
	assert.NoError(t, err)
	assert.Equal(t, values, []interface{}{"a", "b", "c"}, values)

}
