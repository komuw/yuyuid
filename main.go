package main

import (
	"crypto/rand"
	"fmt"
)

const (
	RESERVED_NCS       byte = 0x80 //Reserved for NCS compatibility
	RFC_4122           byte = 0x40 //Specified in RFC 4122
	RESERVED_MICROSOFT byte = 0x20 //Reserved for Microsoft compatibility
	RESERVED_FUTURE    byte = 0x00 // Reserved for future definition.
)

type UUID [16]byte

func (u UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

func (u *UUID) setVariant(variant byte) {
	switch variant {
	case RESERVED_NCS:
		u[8] &= 0x7F
	case RFC_4122:
		u[8] &= 0x3F
		u[8] |= 0x80
	case RESERVED_MICROSOFT:
		u[8] &= 0x1F
		u[8] |= 0xC0
	case RESERVED_FUTURE:
		u[8] &= 0x1F
		u[8] |= 0xE0
	}
}

func (u *UUID) setVersion(version byte) {
	u[6] = (u[6] & 0x0F) | (version << 4)
}

func Uuid4() UUID {

	var uuid UUID
	var version byte = 4

	//Read reads up to len(uuid) bytes into uuid.
	//It returns the number of bytes read (0 <= n <= len(p)) and any error encountered
	_, err := rand.Read(uuid[:])
	if err != nil {
		panic(err)
	}

	uuid.setVariant(RFC_4122)
	uuid.setVersion(version)
	return uuid
}

/*func Uuid5() {
	//to be continued
}*/

func main() {
	u := Uuid4()
	fmt.Println(u)
}
