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

func TestReadAll(t *testing.T) {
	in := []interface{}{"a", "b", "3", 1, 2, 3}
	it, err := NewSliceIterator(in)
	require.NoError(t, err)
	out, err := ReadAll(it)
	require.NoError(t, err)
	assert.Equal(t, in, out)
}
