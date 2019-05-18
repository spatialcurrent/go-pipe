// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"io"
)

import (
	"github.com/pkg/errors"
)

func IteratorToWriter(it Iterator, w Writer, transform func(inputObject interface{}) (interface{}, error), filter func(inputObject interface{}) (bool, error), outputLimit int) error {
	count := 0
	for {
		inputObject, nextError := it.Next()
		if nextError != nil {
			if nextError == io.EOF {
				break
			}
			return errors.Wrap(nextError, "error reading next object")
		}
		if transform != nil {
			outputObject, transformError := transform(inputObject)
			if transformError != nil {
				return errors.Wrap(transformError, "error transforming object")
			}
			if filter != nil {
				ok, filterError := filter(outputObject)
				if filterError != nil {
					return errors.Wrap(filterError, "error grepping object")
				}
				if !ok {
					continue
				}
			}
			count++
			writeError := w.WriteObject(outputObject)
			if writeError != nil {
				return errors.Wrap(writeError, "error writing object to output")
			}
		} else {
			if filter != nil {
				ok, filterError := filter(inputObject)
				if filterError != nil {
					return errors.Wrap(filterError, "error grepping object")
				}
				if !ok {
					continue
				}
			}
			count++
			writeError := w.WriteObject(inputObject)
			if writeError != nil {
				return errors.Wrap(writeError, "error writing object to output")
			}
		}
		if outputLimit >= 0 && count == outputLimit {
			break
		}
	}

	// Flush propogates and calls the underlying writers flush command, if implemented in the concerete struct.
	err := w.Flush()
	if err != nil {
		return errors.Wrap(err, "error flushing output")
	}

	return nil
}
