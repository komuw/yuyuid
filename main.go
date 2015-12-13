package main

import (
	"crypto/rand"
	"fmt"
)

func Uuid4() string {
	//Read reads up to len(b) bytes into b.
	//It returns the number of bytes read (0 <= n <= len(p)) and any error encountered
	b := make([]byte, 16)

	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	hex_random := make([]string, 0)

	for _, value := range b {
		// convert to hexadecimal representation
		hex_rep := fmt.Sprintf("%x", value)
		hex_random = append(hex_random, hex_rep)
	}

	//propely represent any hex values that are 'single digits'
	for key, _ := range hex_random {
		if len(hex_random[key]) == 1 {
			hex_random[key] = "0" + hex_random[key]
		}
	}

	first_part := hex_random[0]
	second_part := hex_random[4] + hex_random[5]
	third_part := hex_random[6] + hex_random[7]
	fourth_part := hex_random[8] + hex_random[9]
	fifth_part := hex_random[10]

	for key, _ := range hex_random {
		if key <= 3 && key > 0 {
			first_part = first_part + hex_random[key]
		} else if key > 10 {
			fifth_part = fifth_part + hex_random[key]
		}
	}

	uuid4 := first_part + "-" + second_part + "-" + third_part + "-" + fourth_part + "-" + fifth_part
	return uuid4
}

func main() {
	u := Uuid4()
	fmt.Println(u)
}
