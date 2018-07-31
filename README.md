## Golang uuid library.

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
