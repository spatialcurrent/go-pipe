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

func TestBuilderInputLimit(t *testing.T) {

	inputValues := []interface{}{"a", "b", 1, 2, 3, true, false}

	it, err := NewSliceIterator(inputValues)
	assert.Nil(t, err)
	assert.NotNil(t, it)

	w := NewSliceWriter()
	assert.NotNil(t, w)

	b := NewBuilder().Input(it).Output(w).InputLimit(2)
	assert.NotNil(t, b)

	err = b.Run()
	assert.Nil(t, err)
	outputValues := w.Values()
	assert.Equal(t, inputValues[:2], outputValues)

}

func TestBuilderOutputLimit(t *testing.T) {

	inputValues := []interface{}{"a", "b", 1, 2, 3, true, false}

	it, err := NewSliceIterator(inputValues)
	assert.Nil(t, err)
	assert.NotNil(t, it)

	w := NewSliceWriter()
	assert.NotNil(t, w)

	b := NewBuilder().Input(it).Output(w).OutputLimit(4)
	assert.NotNil(t, b)

	err = b.Run()
	assert.Nil(t, err)
	outputValues := w.Values()
	assert.Equal(t, inputValues[:4], outputValues)

}
