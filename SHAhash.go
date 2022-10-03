package main

//Go implements several hash functions in various crypto/* packages.

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
)

func main() {
	s := "My name is Srinivas Nidumolu"
	fmt.Printf("SHA  hash: %x\n", genSHAHash(s))
	s1, s2 := getBase64Str(s)
	fmt.Printf("BASE 64 str and enc string: %v, %v\n", s1, s2)

}

func genSHAHash(str string) string {
	h := sha256.New()

	//Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.
	h.Write([]byte(str))

	/*This gets the finalized hash result as a byte slice.
	The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed. */
	bs := h.Sum(nil)

	fmt.Println(str)
	//fmt.Printf("%x\n", bs)
	return string(bs)
}

func getBase64Str(str string) (string, string) {
	sEnc := b64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println("Base64 encoded :", sEnc)

	//Decoding may return an error, which you can check if you don’t already know the input to be well-formed.

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println("Base64 decode :", string(sDec))
	fmt.Println()

	//This encodes/decodes using a URL-compatible base64 format.

	uEnc := b64.URLEncoding.EncodeToString([]byte(str))
	fmt.Println("Base64 unencode :", uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
	return string(str), sEnc
}
