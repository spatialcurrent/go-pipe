// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestBuilderInputLimit(t *testing.T) {

	input := []interface{}{"a", "b", 1, 2, 3, true, false}
	expected := input[:2]

	it, err := NewSliceIterator(input)
	assert.Nil(t, err)
	assert.NotNil(t, it)

	w := NewSliceWriter()
	assert.NotNil(t, w)

	b := NewBuilder().Input(it).Output(w).InputLimit(2)
	assert.NotNil(t, b)

	err = b.Run()
	assert.Nil(t, err)
	output := w.Values()
	assert.Equal(t, expected, output)

}

func TestBuilderOutputLimit(t *testing.T) {

	input := []interface{}{"a", "b", 1, 2, 3, true, false}
	expected := input[:4]

	it, err := NewSliceIterator(input)
	assert.Nil(t, err)
	assert.NotNil(t, it)

	w := NewSliceWriter()
	assert.NotNil(t, w)

	b := NewBuilder().Input(it).Output(w).OutputLimit(4)
	assert.NotNil(t, b)

	err = b.Run()
	assert.Nil(t, err)
	output := w.Values()
	assert.Equal(t, expected, output)

}

func TestBuilderTransform(t *testing.T) {

	input := []interface{}{"a", "b", 1, 2, 3, true, false}
	expected := []interface{}{"v: a", "v: b", "v: 1", "v: 2", "v: 3", "v: true", "v: false"}

	it, err := NewSliceIterator(input)
	assert.Nil(t, err)
	assert.NotNil(t, it)

	w := NewSliceWriter()
	assert.NotNil(t, w)

	b := NewBuilder().
		Input(it).
		Transform(func(inputObject interface{}) (interface{}, error) {
			return fmt.Sprintf("v: %v", inputObject), nil
		}).
		Output(w)
	assert.NotNil(t, b)

	err = b.Run()
	assert.Nil(t, err)
	output := w.Values()
	assert.Equal(t, expected, output)

}

func TestBuilderTransformError(t *testing.T) {

	input := []interface{}{"a", "b", 1, 2, 3, true, false}
	expectedError := "error transforming input: found value 1 of type int, expected type string"
	expectedValues := []interface{}{"v: a", "v: b"}

	it, err := NewSliceIterator(input)
	assert.Nil(t, err)
	assert.NotNil(t, it)

	w := NewSliceWriter()
	assert.NotNil(t, w)

	b := NewBuilder().
		Input(it).
		Transform(func(inputObject interface{}) (interface{}, error) {
			if _, ok := inputObject.(string); ok {
				return fmt.Sprintf("v: %v", inputObject), nil
			}
			return nil, fmt.Errorf("found value %v of type %T, expected type string", inputObject, inputObject)
		}).
		Error(func(err error) error {
			return errors.Wrap(err, "error transforming input")
		}).
		Output(w)
	assert.NotNil(t, b)

	err = b.Run()
	if assert.NotNil(t, err) {
		assert.Equal(t, expectedError, err.Error())
	}
	output := w.Values()
	assert.Equal(t, expectedValues, output)

}

func TestSliceToSet(t *testing.T) {

	input := []interface{}{"a", "b", "c", "a"}
	expectedValues := map[interface{}]struct{}{
		"a": struct{}{},
		"b": struct{}{},
		"c": struct{}{},
	}

	it, err := NewSliceIterator(input)
	assert.Nil(t, err)
	assert.NotNil(t, it)

	w := NewSetWriter()
	assert.NotNil(t, w)

	b := NewBuilder().Input(it).Output(w)
	assert.NotNil(t, b)

	err = b.Run()
	assert.Nil(t, err)
	output := w.Values()
	assert.Equal(t, expectedValues, output)

}
