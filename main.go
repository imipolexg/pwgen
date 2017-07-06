package main

import (
	"crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
)

var alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUV1234567890"
var punctAlphabet string = "-=\\`!@#$%^&*()_+|~{}[];:'\",./<>?|"

func main() {
	pwlenp := flag.Int("len", 10, "Password length")
	noPunctp := flag.Bool("nopunct", false, "Don't use punctuation")
	showHelp := flag.Bool("help", false, "Show this help")
	flag.Parse()

	if *showHelp {
		flag.Usage()
		return
	}

	if !*noPunctp {
		alphabet += punctAlphabet
	}

	password := pwgen(*pwlenp, alphabet)

	fmt.Println(password)
}

func pwgen(pwlen int, alphabet string) (password string) {
	randData := make([]byte, 2)

	for i := 0; i < pwlen; i++ {
		randInt := -1
		for randInt < 0 || randInt > len(alphabet)-1 {
			_, err := rand.Read(randData)
			if err != nil {
				panic(err)
			}
			randInt = int(binary.BigEndian.Uint16(randData))
		}
		password = password + string(alphabet[randInt])
	}

	return
}
