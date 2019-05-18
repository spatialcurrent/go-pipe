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
	filter      func(object interface{}) (bool, error)
	outputLimit int
}

// NewBuilder returns a new builder.
func NewBuilder() *Builder {
	return &Builder{
		input:       nil,
		output:      nil,
		transform:   nil,
		filter:      nil,
		outputLimit: -1,
	}
}

// Input sets the input for the pipeline.
func (b *Builder) Input(in Iterator) *Builder {
	return &Builder{
		input:       in,
		output:      b.output,
		transform:   b.transform,
		filter:      b.filter,
		outputLimit: b.outputLimit,
	}
}

// Output sets the output for the pipeline.
func (b *Builder) Output(w Writer) *Builder {
	return &Builder{
		input:       b.input,
		output:      w,
		transform:   b.transform,
		filter:      b.filter,
		outputLimit: b.outputLimit,
	}
}

// Transform sets the transform for the pipeline.
func (b *Builder) Transform(t func(inputObject interface{}) (interface{}, error)) *Builder {
	return &Builder{
		input:       b.input,
		output:      b.output,
		transform:   t,
		filter:      b.filter,
		outputLimit: b.outputLimit,
	}
}

// Filter sets the filter for the pipeline.
func (b *Builder) Filter(f func(object interface{}) (bool, error)) *Builder {
	return &Builder{
		input:       b.input,
		output:      b.output,
		transform:   b.transform,
		filter:      f,
		outputLimit: b.outputLimit,
	}
}

// OutputLimit sets the outputLimit for the pipeline.
func (b *Builder) OutputLimit(outputLimit int) *Builder {
	return &Builder{
		input:       b.input,
		output:      b.output,
		transform:   b.transform,
		filter:      b.filter,
		outputLimit: outputLimit,
	}
}

// Run runs the pipeline.
func (b *Builder) Run() error {
	if b.input != nil && b.output != nil {
		return IteratorToWriter(b.input, b.output, b.transform, b.filter, b.outputLimit)
	}
	return nil
}
