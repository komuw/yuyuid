package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/rand"
	"fmt"
)

const (
	RESERVED_NCS       byte = 0x80 //Reserved for NCS compatibility
	RFC_4122           byte = 0x40 //Specified in RFC 4122
	RESERVED_MICROSOFT byte = 0x20 //Reserved for Microsoft compatibility
	RESERVED_FUTURE    byte = 0x00 // Reserved for future definition.
)

//The following standard UUIDs are for use with uuid3() or uuid5().
var (
	NAMESPACE_DNS  = UUID{107, 167, 184, 16, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
	NAMESPACE_URL  = UUID{107, 167, 184, 17, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
	NAMESPACE_OID  = UUID{107, 167, 184, 18, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
	NAMESPACE_X500 = UUID{107, 167, 184, 20, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
)

type UUID [16]byte

func (u UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
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

func Uuid3(namespace UUID, name string) UUID {
	var uuid UUID
	var version byte = 3
	hasher := md5.New()
	hasher.Write(namespace[:])
	hasher.Write([]byte(name))
	sum := hasher.Sum(nil)
	copy(uuid[:], sum[:len(uuid)])

	uuid.setVariant(RFC_4122)
	uuid.setVersion(version)
	return uuid
}

func Uuid4() UUID {

	var uuid UUID
	var version byte = 4

	//Read reads up to len(uuid) bytes into uuid.
	//It returns the number of bytes read (0 <= n <= len(uuid)) and any error encountered
	_, err := rand.Read(uuid[:])
	if err != nil {
		panic(err)
	}

	uuid.setVariant(RFC_4122)
	uuid.setVersion(version)
	return uuid
}

func Uuid5(namespace UUID, name string) UUID {
	var uuid UUID
	var version byte = 5
	hasher := sha1.New()
	hasher.Write(namespace[:])
	hasher.Write([]byte(name))
	sum := hasher.Sum(nil)
	copy(uuid[:], sum[:len(uuid)])

	uuid.setVariant(RFC_4122)
	uuid.setVersion(version)
	return uuid
}

func main() {
	u4 := Uuid4()
	fmt.Println("uuid4", u4)

	u3 := Uuid3(NAMESPACE_DNS, "SomeName")
	fmt.Println("uuid3", u3)

	u5 := Uuid5(NAMESPACE_DNS, "SomeName")
	fmt.Println("uuid5", u5)
}
