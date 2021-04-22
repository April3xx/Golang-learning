// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	test()
	password := []byte("Kjkszpjkjkszpjkj") //16 char long
	plaintext := []byte("exampleplaintext")

	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(password)
	if err != nil {
		fmt.Println("error on top")
		panic(err)

	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext)) //--> already in byte slice array
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println("error at the buttom")
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	fmt.Printf("%x\n", ciphertext)

	mode = cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	fmt.Printf("%s\n", ciphertext)
}
