[![CircleCI](https://circleci.com/gh/spatialcurrent/go-pipe/tree/main.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-pipe/tree/main)
[![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-pipe?style=flat-square)](https://goreportcard.com/report/github.com/spatialcurrent/go-pipe)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/spatialcurrent/go-pipe)](https://pkg.go.dev/github.com/spatialcurrent/go-pipe)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-pipe/blob/master/LICENSE)

# go-pipe

# Description

**go-pipe** is a simple library for piping data from iterators to writers.  Pipe uses the following interfaces, which allows reading and writing of objects to any number of inputs and outputs.

**Iterator**

The Next method will return io.EOF once it has read all the objects from the input.

```go
type Iterator interface {
	Next() (interface{}, error)
}
```

**Writer**

Writer is a writer that accepts input as individual objects.

```go
type Writer interface {
	WriteObject(object interface{}) error
	Flush() error
}
```

**BatchWriter**

BatchWriter is a writer that accepts input in batches as an array or slice of objects.

```go
type Writer interface {
	WriteObjects(objects interface{}) error
	Flush() error
}
```

**go-pipe** includes concrete structs for writing to channels, slice, functions, and maps.  It is also used in [railgun](https://github.com/spatialcurrent/railgun) project along with [go-simple-serializer](https://github.com/spatialcurrent/go-simple-serializer) to process objects from files.

# Usage

**Go**

Install the package with:

```shell
go get -d github.com/spatialcurrent/go-pipe/...
```

You can import **go-pipe** as a library with:

```go
import (
  "github.com/spatialcurrent/go-pipe/pkg/pipe"
)
```

See [pipe](https://pkg.go.dev/github.com/spatialcurrent/go-pipe/pkg/pipe) in the docs for information on how to use Go API.

# Examples

See the many [examples](https://pkg.go.dev/github.com/spatialcurrent/go-pipe/pkg/pipe/#pkg-examples) in the docs or the tests.

# Testing

Run test using `make test` or (`bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-pipe/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
