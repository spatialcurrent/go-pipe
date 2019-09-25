// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package pipe

// Builder helps build a pipeline
type Builder struct {
	input       Iterator
	output      Writer
	transform   func(inputObject interface{}) (interface{}, error)
	error       func(err error) error
	filter      func(object interface{}) (bool, error)
	inputLimit  int
	outputLimit int
	closeOutput bool
}

// NewBuilder returns a new builder.
func NewBuilder() *Builder {
	return &Builder{
		input:       nil,
		output:      nil,
		transform:   nil,
		filter:      nil,
		inputLimit:  -1,
		outputLimit: -1,
		closeOutput: false,
	}
}

// Input sets the input for the pipeline.
func (b *Builder) Input(in Iterator) *Builder {
	return &Builder{
		input:       in,
		output:      b.output,
		transform:   b.transform,
		error:       b.error,
		filter:      b.filter,
		inputLimit:  b.inputLimit,
		outputLimit: b.outputLimit,
		closeOutput: b.closeOutput,
	}
}

// InputF sets the input for the pipeline to a function by wrapping the provided function with pipe.FunctionIterator.
func (b *Builder) InputF(fn func() (interface{}, error)) *Builder {
	return b.Input(NewFunctionIterator(fn))
}

// Output sets the output for the pipeline.
func (b *Builder) Output(w Writer) *Builder {
	return &Builder{
		input:       b.input,
		output:      w,
		transform:   b.transform,
		error:       b.error,
		filter:      b.filter,
		inputLimit:  b.inputLimit,
		outputLimit: b.outputLimit,
		closeOutput: b.closeOutput,
	}
}

// OutputF sets the output for the pipeline to a function by wrapping the provided function with pipe.FunctionWriter.
func (b *Builder) OutputF(fn func(object interface{}) error) *Builder {
	return b.Output(NewFunctionWriter(fn))
}

// Transform sets the transform for the pipeline.
func (b *Builder) Transform(t func(inputObject interface{}) (interface{}, error)) *Builder {
	return &Builder{
		input:       b.input,
		output:      b.output,
		transform:   t,
		error:       b.error,
		filter:      b.filter,
		inputLimit:  b.inputLimit,
		outputLimit: b.outputLimit,
		closeOutput: b.closeOutput,
	}
}

// Error sets the error handler for the pipeline that catches errors from the transform function.
// If the error handler returns nil, then the pipeline continues as normal.
// If the error handler returns the original error (or a new one), then the pipeline bubbles up the error and exits.
func (b *Builder) Error(e func(err error) error) *Builder {
	return &Builder{
		input:       b.input,
		output:      b.output,
		transform:   b.transform,
		error:       e,
		filter:      b.filter,
		inputLimit:  b.inputLimit,
		outputLimit: b.outputLimit,
		closeOutput: b.closeOutput,
	}
}

// Filter sets the filter for the pipeline.
func (b *Builder) Filter(f func(object interface{}) (bool, error)) *Builder {
	return &Builder{
		input:       b.input,
		output:      b.output,
		transform:   b.transform,
		error:       b.error,
		filter:      f,
		inputLimit:  b.inputLimit,
		outputLimit: b.outputLimit,
		closeOutput: b.closeOutput,
	}
}

// InputLimit sets the inputLimit for the pipeline.
func (b *Builder) InputLimit(inputLimit int) *Builder {
	return &Builder{
		input:       b.input,
		output:      b.output,
		transform:   b.transform,
		error:       b.error,
		filter:      b.filter,
		inputLimit:  inputLimit,
		outputLimit: b.outputLimit,
		closeOutput: b.closeOutput,
	}
}

// OutputLimit sets the outputLimit for the pipeline.
func (b *Builder) OutputLimit(outputLimit int) *Builder {
	return &Builder{
		input:       b.input,
		output:      b.output,
		transform:   b.transform,
		error:       b.error,
		filter:      b.filter,
		inputLimit:  b.inputLimit,
		outputLimit: outputLimit,
		closeOutput: b.closeOutput,
	}
}

// CloseOutput sets the closeOutput for the pipeline.
func (b *Builder) CloseOutput(closeOutput bool) *Builder {
	return &Builder{
		input:       b.input,
		output:      b.output,
		transform:   b.transform,
		error:       b.error,
		filter:      b.filter,
		inputLimit:  b.inputLimit,
		outputLimit: b.outputLimit,
		closeOutput: closeOutput,
	}
}

// Run runs the pipeline.
func (b *Builder) Run() error {
	if b.input != nil && b.output != nil {
		return IteratorToWriter(b.input, b.output, b.transform, b.error, b.filter, b.inputLimit, b.outputLimit, b.closeOutput)
	}
	return nil
}
