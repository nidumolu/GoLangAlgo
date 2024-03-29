/* LoadRSAKeys
This programread back the keys and related certificate (via pem) that were writtenusing gob serialization:
refer to program GenPubPrivKeys.go
*/
package main

import (
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func main() {
	var key rsa.PrivateKey
	loadKey("private.key", &key)
	fmt.Printf("Private key primes:\n[0]:%s\n[1]:%s\n", key.Primes[0].String(),
		key.Primes[1].String())
	fmt.Println("Private key exponent:\n", key.D.String())
	var publicKey rsa.PublicKey
	loadKey("public.key", &publicKey)
	fmt.Println("Public key modulus:\n", publicKey.N.String())
	fmt.Println("Public key exponent:\n", publicKey.E)
}
func loadKey(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)
	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()
}
func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error ", err.Error())
	}
}
