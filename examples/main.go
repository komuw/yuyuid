package main

import (
	"fmt"

	"github.com/komuw/yuyuid"
)

func main() {
	uuid4 := yuyuid.Uuid4()
	fmt.Println("uuid4 is", uuid4)

	uuid5 := yuyuid.Uuid5(yuyuid.NAMESPACE_DNS, "SomeName")
	fmt.Println("uuid5", uuid5)
}
