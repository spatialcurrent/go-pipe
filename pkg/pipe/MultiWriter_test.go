// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiWriter(t *testing.T) {

	a := make([]interface{}, 0)
	b := make([]interface{}, 0)

	mw := NewMultiWriter(
		NewFunctionWriter(func(object interface{}) error {
			a = append(a, object)
			return nil
		}),
		NewFunctionWriter(func(object interface{}) error {
			b = append(b, object)
			return nil
		}),
	)

	expected := make([]interface{}, 0)

	count := 1000

	for i := 0; i < count; i++ {
		expected = append(expected, i)
		//
		err := mw.WriteObject(i)
		if err != nil {
			panic(err)
		}
	}

	assert.Lenf(t, expected, count, "invalid length, expecting %d elements", count)
	assert.Equal(t, expected, a)
	assert.Equal(t, expected, b)
}

func TestMultiWriterInts(t *testing.T) {

	a := make([]int, 0)
	b := make([]int, 0)

	mw := NewMultiWriter(
		NewFunctionWriter(func(object interface{}) error {
			if i, ok := object.(int); ok {
				a = append(a, i)
			}
			return nil
		}),
		NewFunctionWriter(func(object interface{}) error {
			if i, ok := object.(int); ok {
				b = append(b, i)
			}
			return nil
		}),
	)

	expected := make([]int, 0)

	count := 1000

	for i := 0; i < count; i++ {
		expected = append(expected, i)
		//
		err := mw.WriteObject(i)
		if err != nil {
			panic(err)
		}
	}

	assert.Lenf(t, expected, count, "invalid length, expecting %d elements", count)
	assert.Equal(t, expected, a)
	assert.Equal(t, expected, b)
}
