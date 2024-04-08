# F-Error

With this package it is possible to store more information about the errors that occur in the code, such as the error message, the layer where the error occured and a message for the user.

## Usage

To import the package use `go get github.com/fabiokaelin/ferror`.

```go
package main

import (
	"errors"
	"fmt"

	"github.com/fabiokaelin/ferror"
)

func main() {
	_, err1 := retrunNormalError()

	ferr1 := ferrors.FromError(err1)
	if ferr1 != nil {
		fmt.Println(ferr1)
	}

	err2 := ferrors.New("this is a ferror")
	err2.SetLayer("db")
	err2.SetUserMsg("Unable to persist the data")
	err2.SetKind("db connection")
	err2.SetInternal("db connection closed unexpectedly")
	fmt.Println(err2)
}

func retrunNormalError() (string, error) {
	return "", errors.New("this is a normal error")
}
```

```log
{Error: 'this is a normal error', Layer: 'db', UserMsg: '', Kind: '', Internal: ''}
{Error: 'this is a ferror', Layer: 'db', UserMsg: 'Unable to persist the data', Kind: 'db connection', Internal: 'db connection closed unexpectedly'}
```
