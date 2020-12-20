// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

import (
	"fmt"
	"io"
)

// IteratorToWriter reads objects from the provided iterator and writes them to the provided writer.
// If a transform function is given, then transforms the input objects before writing them.
// If an errorHandler is given, then propogates errors returned by the transformed function through the errorHandler.
// If the errorHandler returns a non-nil error, then processing will halt.
// If a filter is given, the input object, after transformation if applicable, is filtered.
// If the filter returns true and no error, then the object is passed to the writer.
// If the inputLimit >= 0, then reads the given number of objects from the input.
// If the outputLimit >= 0, then writes the given number of objecst to the writer.
func IteratorToWriter(it Iterator, w Writer, transform func(inputObject interface{}) (interface{}, error), errorHandler func(err error) error, filter func(inputObject interface{}) (bool, error), inputLimit int, outputLimit int, closeOutput bool) error {
	if inputLimit == 0 {
		return nil
	}
	if outputLimit == 0 {
		return nil
	}
	inputCount := 0
	outputCount := 0
	for {
		inputCount++
		inputObject, nextError := it.Next()
		if nextError != nil {
			if nextError == io.EOF {
				break
			}
			return fmt.Errorf("error reading next object: %w", nextError)
		}
		if transform != nil {
			outputObject, transformError := transform(inputObject)
			if transformError != nil {
				if errorHandler != nil {
					transformError = errorHandler(transformError)
				}
				if transformError != nil {
					return transformError
				}
			}
			if filter != nil {
				ok, filterError := filter(outputObject)
				if filterError != nil {
					return fmt.Errorf("error grepping object: %w", filterError)
				}
				if !ok {
					continue
				}
			}
			outputCount++
			writeError := w.WriteObject(outputObject)
			if writeError != nil {
				return fmt.Errorf("error writing object to output: %w", writeError)
			}
		} else {
			if filter != nil {
				ok, filterError := filter(inputObject)
				if filterError != nil {
					return fmt.Errorf("error grepping object: %w", filterError)
				}
				if !ok {
					continue
				}
			}
			outputCount++
			writeError := w.WriteObject(inputObject)
			if writeError != nil {
				return fmt.Errorf("error writing object to output: %w", writeError)
			}
		}
		if inputLimit > 0 && inputCount == inputLimit {
			break
		}
		if outputLimit > 0 && outputCount == outputLimit {
			break
		}
	}

	// Flush propogates and calls the underlying writers flush command, if implemented in the concerete struct.
	err := w.Flush()
	if err != nil {
		return fmt.Errorf("error flushing output: %w", err)
	}

	if closeOutput {
		if c, ok := w.(interface{ Close() error }); ok {
			err = c.Close()
			if err != nil {
				return fmt.Errorf("error closing output: %w", err)
			}
		}
	}

	return nil
}
