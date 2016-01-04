## Golang uuid library.

# Installation
`go get github.com/komuW/yuyuid`

# Usage
```
import (
	"github.com/komuW/yuyuid"
	"fmt"
)

func main() {
	uuid4 := yuyuid.Uuid4()
	fmt.Println("uuid4 is", uuid4)

	uuid5 := yuyuid.Uuid5(yuyuid.NAMESPACE_DNS, "SomeName")
	fmt.Println("uuid5", uuid5)
}
```


# FAQ
                
Why?         
* golang doesn't have an inbuilt uuid lib
* pet project to help me learn golang                                

Is it done yet?             
* No
* Only uuid versions 3, 4 and 5 are implemented. I don't think I'll ever implement uuid version 1
