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
	}
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
	}
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
	}
}

// Transform sets the transform for the pipeline.
func (b *Builder) Error(e func(err error) error) *Builder {
	return &Builder{
		input:       b.input,
		output:      b.output,
		transform:   b.transform,
		error:       e,
		filter:      b.filter,
		inputLimit:  b.inputLimit,
		outputLimit: b.outputLimit,
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
	}
}

// Run runs the pipeline.
func (b *Builder) Run() error {
	if b.input != nil && b.output != nil {
		return IteratorToWriter(b.input, b.output, b.transform, b.error, b.filter, b.inputLimit, b.outputLimit)
	}
	return nil
}
