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

func TestChannelWriter(t *testing.T) {

	c := make(chan interface{}, 1000)

	w := NewChannelWriter(c)

	err := w.WriteObject("a")
	assert.Nil(t, err)

	err = w.WriteObject("b")
	assert.Nil(t, err)

	err = w.WriteObject("c")
	assert.Nil(t, err)

	err = w.WriteObject(1)
	assert.Nil(t, err)

	err = w.WriteObject(2)
	assert.Nil(t, err)

	err = w.WriteObject(3)
	assert.Nil(t, err)

	close(c)

	values := make([]interface{}, 0)
	for v := range c {
		values = append(values, v)
	}
	assert.Equal(t, []interface{}{"a", "b", "c", 1, 2, 3}, values)
}
