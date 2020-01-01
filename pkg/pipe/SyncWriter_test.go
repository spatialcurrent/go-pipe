// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"sync"
	//"time"
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestSyncWriter(t *testing.T) {

	values := make([]interface{}, 0)

	sw := NewSyncWriter(NewFunctionWriter(func(object interface{}) error {
		values = append(values, object)
		return nil
	}))

	count := 1000

	wg := &sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			err := sw.WriteObject(i)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	assert.Lenf(t, values, count, "invalid length, expecting %d elements", count)
}
