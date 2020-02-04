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
	"github.com/stretchr/testify/require"
)

func TestTransactionWriter(t *testing.T) {

	values := make([]interface{}, 0)

	tw, err := NewTransactionWriter(func() (Writer, error) {
		fw := NewFunctionWriter(func(object interface{}) error {
			values = append(values, object)
			return nil
		})
		return fw, nil
	}, nil, false)

	require.NoError(t, err)

	err = tw.WriteObject("a")
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{"a"}, values)
}

func TestTransactionWriterAsync(t *testing.T) {

	values := make([]interface{}, 0)

	tw, err := NewTransactionWriter(func() (Writer, error) {
		fw := NewFunctionWriter(func(object interface{}) error {
			values = append(values, object)
			return nil
		})
		return fw, nil
	}, nil, true)

	require.NoError(t, err)

	err = tw.WriteObject("a")
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{"a"}, values)
}

func TestTransactionWriterCloser(t *testing.T) {

	closed := 0
	values := make([]interface{}, 0)

	tw, err := NewTransactionWriter(
		func() (Writer, error) {
			fw := NewFunctionWriter(func(object interface{}) error {
				values = append(values, object)
				return nil
			})
			return fw, nil
		},
		func(w Writer) error {
			closed += 1
			return nil
		},
		false,
	)

	require.NoError(t, err)

	err = tw.WriteObject("a")
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{"a"}, values)
	assert.Equal(t, 1, closed)
}
