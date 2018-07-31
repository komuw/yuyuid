# yuyuid
[![Build Status](https://travis-ci.com/komuw/yuyuid.svg?branch=master)](https://travis-ci.com/komuw/yuyuid)
[![codecov](https://codecov.io/gh/komuw/yuyuid/branch/master/graph/badge.svg)](https://codecov.io/gh/komuw/yuyuid)
[![GoDoc](https://godoc.org/github.com/komuw/yuyuid?status.svg)](https://godoc.org/github.com/komuw/yuyuid)
[![Go Report Card](https://goreportcard.com/badge/github.com/komuw/yuyuid)](https://goreportcard.com/report/github.com/komuw/yuyuid)     

Go uuid library.        

# Installation
`go get github.com/komuw/yuyuid`

# Usage
```go
package main

import (
	"fmt"

	"github.com/komuw/yuyuid"
)

func main() {
	UUID4 := yuyuid.UUID4()
	fmt.Println("UUID4 is", UUID4)

	UUID5 := yuyuid.UUID5(yuyuid.NamespaceDNS, "SomeName")
	fmt.Println("UUID5", UUID5)
}
```


# FAQ
                
Why?         
* golang doesn't have an inbuilt uuid lib
* pet project to help me learn golang                                

Is it done yet?             
* Yes
* Only uuid versions 3, 4 and 5 are implemented. I don't think I'll ever implement uuid version 1
