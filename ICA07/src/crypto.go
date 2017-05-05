package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"golang.org/x/crypto/nacl/secretbox"
	"encoding/hex"
)

const (
	PASSWORD = "6368616e676520746869732070617373776f726420746f206120736563726574"
)

func crypt(message string) []byte {

	secretKeyBytes, err := hex.DecodeString(PASSWORD)
	if err != nil {
		panic(err)
	}

	var secretKey [32]byte
	copy(secretKey[:], secretKeyBytes)

	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}
	// Encrypt the message
	encrypted := secretbox.Seal(nonce[:], []byte(message), &nonce, &secretKey)
	fmt.Print(encrypted)
	return encrypted
}
//
func decrypt(encrypted []byte) string  {

	secretKeyBytes, err := hex.DecodeString(PASSWORD)
	if err != nil {
		panic(err)
	}

	var secretKey [32]byte
	copy(secretKey[:], secretKeyBytes)

	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])
	decrypted, ok := secretbox.Open([]byte{}, encrypted[24:], &decryptNonce, &secretKey)
	if !ok {
		panic("decryption error")
	}
	return string(decrypted)
}
//
//func SaltPassword() [32]byte {
//	SALT_BYTES = 32
//	HASH_BYTES = 64
//	N = 16384 // CPU/MEMORY
//	// r and p must satisfy r * p < 2³⁰
//	R = 8
//	P = 1
//
//	salt := make([]byte, SALT_BYTES)
//	_, err := io.ReadFull(rand.Reader, salt)
//	if err != nil {
//		log.Fatal(err)
//	}
//	hash, err := scrypt.Key([]byte(PASSWORD), salt, N, R, P, HASH_BYTES)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//
//	var secretKey [32]byte
//	copy(secretKey[:], hash)
//	return secretKey
//
//}