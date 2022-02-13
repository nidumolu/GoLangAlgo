package main

//Go implements several hash functions in various crypto/* packages.

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	s := "Tom and Jerry is a great cartoon !"
	/*The pattern for generating a hash is sha1.New(), sha1.Write(bytes),
	then sha1.Sum([]byte{}). Here we start with a new hash.
	*/

	key := "thisis32bitlongpassphraseimusing"
	// The encryption key, if set, must be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256 modes.
	cryptoText := EncryptAES([]byte(key), s)

	// encrypt value to base64
	//cryptoText := encrypt(key, s)
	fmt.Println("Base64 string encoded with AES:", cryptoText)

	// encrypt base64 crypto to original value
	text := decrypt([]byte(key), cryptoText)

	fmt.Printf("decrypting back from base64 to org text using AES : %s\n", text)

	h := sha1.New()
	hMd5 := md5.New()
	/*Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.*/

	h.Write([]byte(s))
	hMd5.Write([]byte(s))
	/*This gets the finalized hash result as a byte slice.
	The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed.
	*/
	bs := h.Sum(nil)
	//bsMd5 := hMd5.Sum(nil)
	/*SHA1 values are often printed in hex, for example in git commits. Use the %x format verb to convert a hash results to a hex string.
	 */
	fmt.Println(s)
	fmt.Println(" SHA1 hasH :", bs, " \n")
	fmt.Printf("%x\n", bs)
	fmt.Println(" MD5 hasH : ", hMd5, " \n")
	fmt.Printf("%x\n", hMd5.Sum(nil))

}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

func EncryptAES(key []byte, plaintext string) string {
	// create cipher
	c, err := aes.NewCipher(key)
	CheckError(err)

	// allocate space for ciphered data
	out := make([]byte, len(plaintext))

	// encrypt
	c.Encrypt(out, []byte(plaintext))
	// return hex string
	return hex.EncodeToString(out)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
