# nestederror [![Build Status](https://travis-ci.org/ctit-team/nestederror.svg?branch=master)](https://travis-ci.org/ctit-team/nestederror) [![GoDoc](https://godoc.org/github.com/ctit-team/nestederror?status.svg)](https://godoc.org/github.com/ctit-team/nestederror)

`nestederror` is a Go package for chaining multiple errors together. It's ability is similar to .NET `Exception.InnerException` or Java `Throwable.getCause`.

## Installation

```sh
go get -u github.com/ctit-team/nestederror
```

## Example

```go
package main

import (
    "errors"
    "log"

    "github.com/ctit-team/nestederror"
)

func main() {
    if err := run(); err != nil {
        log.Fatalln(err)
    }
}

func run() error {
    if err := startServer(); err != nil {
        return nestederror.New(err, "failed to start server")
    }

    return nil
}

func startServer() error {
    return errors.New("unable to open port 80")
}
```

Output:

```text
2017/08/08 15:33:57 failed to start server -> unable to open port 80
```
