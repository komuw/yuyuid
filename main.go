package main

import (
	"crypto/rand"
	"fmt"
)

type UUID [16]byte

func (u UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

func Uuid4() UUID {

	var uuid UUID

	//Read reads up to len(b) bytes into b.
	//It returns the number of bytes read (0 <= n <= len(p)) and any error encountered
	_, err := rand.Read(uuid[:])
	if err != nil {
		panic(err)
	}

	return uuid
}

/*func Uuid5() {
	//to be continued
}*/

func main() {
	u := Uuid4()
	fmt.Println(u)
}
