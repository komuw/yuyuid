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
