package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"log"
	"os"

	"github.com/RyuaNerin/go-krypto/lea"
	"github.com/RyuaNerin/go-krypto/lsh256"
)

var (
	dec    = flag.Bool("d", false, "Decrypt instead Encrypt.")
	file   = flag.String("f", "", "Target file. ('-' for STDIN)")
	iter   = flag.Int("i", 1024, "Iterations. (for PBKDF2)")
	key    = flag.String("k", "", "256-bit key to Encrypt/Decrypt.")
	pbkdf  = flag.String("p", "", "Password-based key derivation function2.")
	random = flag.Bool("r", false, "Generate random 256-bit cryptographic key.")
	salt   = flag.String("s", "", "Salt. (for PBKDF2)")
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("LEA-GCM Encryption Tool - ALBANESE Lab (c) 2020-2021")
		fmt.Println("Lightweight Encryption Algorithm in Galois/Counter Mode\n")
		fmt.Println("Usage of", os.Args[0]+":")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *random == true {
		var key []byte
		var err error
		key = make([]byte, 32)
		_, err = io.ReadFull(rand.Reader, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hex.EncodeToString(key))
		os.Exit(0)
	}

	var keyHex string
	var prvRaw []byte
	if *pbkdf != "" {
		prvRaw = pbkdf2.Key([]byte(*pbkdf), []byte(*salt), *iter, 32, lsh256.New)
		keyHex = hex.EncodeToString(prvRaw)
	} else {
		keyHex = *key
	}
	var key []byte
	var err error
	if keyHex == "" {
		key = make([]byte, 32)
		_, err = io.ReadFull(rand.Reader, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(os.Stderr, "Key=", hex.EncodeToString(key))
	} else {
		key, err = hex.DecodeString(keyHex)
		if err != nil {
			log.Fatal(err)
		}
		if len(key) != 32 {
			log.Fatal(err)
		}
	}

	buf := bytes.NewBuffer(nil)
	var data io.Reader
	if *file == "-" {
		data = os.Stdin
	} else {
		data, _ = os.Open(*file)
	}
	io.Copy(buf, data)
	msg := buf.Bytes()

	c, _ := lea.NewCipher(key)
	aead, _ := cipher.NewGCM(c)

	if *dec == false {
		nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(msg)+aead.Overhead())

		out := aead.Seal(nonce, nonce, msg, nil)
		fmt.Printf("%s", out)

		os.Exit(0)
	}

	if *dec == true {
		nonce, msg := msg[:aead.NonceSize()], msg[aead.NonceSize():]

		out, err := aead.Open(nil, nonce, msg, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", out)

		os.Exit(0)
	}
}
