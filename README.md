# LEACrypt
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/leacrypt/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/leacrypt?status.png)](http://godoc.org/github.com/pedroalbanese/leacrypt)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/leacrypt)](https://goreportcard.com/report/github.com/pedroalbanese/leacrypt)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/leacrypt)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/leacrypt)](https://github.com/pedroalbanese/leacrypt/releases)
### Command-line Lightweight Encryption Algorithm Utility
<pre>Usage of leacrypt:
  -d    Decrypt instead Encrypt.
  -f string
        Target file. ('-' for STDIN)
  -i int
        Iterations. (for PBKDF2) (default 1024)
  -k string
        256-bit key to Encrypt/Decrypt.
  -p string
        Password-based key derivation function2.
  -r    Generate random 256-bit cryptographic key.
  -s string
        Salt. (for PBKDF2)
</pre>

### Example of encryption/decryption:
```sh
./leacrypt -k $256bitkey -f plaintext.ext > ciphertext.ext
./leacrypt -d -k $256bitkey -f ciphertext.ext > plaintext.ext
```

## License

This project is licensed under the ISC License.

##### Commercial-Grade Reliability. Copyright (c) 2020-2021 ALBANESE Research Lab.
