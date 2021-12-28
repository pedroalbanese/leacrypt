# LEACrypt
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/leacrypt/blob/master/LICENSE.md) 
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/leacrypt/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/leacrypt/releases)
[![GoDoc](https://godoc.org/github.com/pedroalbanese/leacrypt?status.png)](http://godoc.org/github.com/pedroalbanese/leacrypt)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/leacrypt)](https://goreportcard.com/report/github.com/pedroalbanese/leacrypt)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/leacrypt)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/leacrypt)](https://github.com/pedroalbanese/leacrypt/releases)

The Lightweight Encryption Algorithm (also known as LEA) is a 128-bit block cipher developed by South Korea in 2013 to provide confidentiality in high-speed environments such as big data and cloud computing, as well as lightweight environments such as IoT devices and mobile devices.

LEA is one of the cryptographic algorithms approved by the Korean Cryptographic Module Validation Program (KCMVP) and is the national standard of Republic of Korea (KS X 3246). LEA is included in the ISO/IEC 29192-2:2019 standard (Information security - Lightweight cryptography - Part 2: Block ciphers).
### Command-line Lightweight Encryption Algorithm Utility
<pre>Usage of leacrypt:
leacrypt [-d] [-b N] -p "pass" [-i N] [-s "salt"] -f &lt;file.ext&gt;
  -b int
        Key length: 128, 192 or 256. (default 256)
  -d    Decrypt instead Encrypt.
  -f string
        Target file. ('-' for STDIN)
  -i int
        Iterations. (for PBKDF2) (default 1024)
  -k string
        Symmetric key to Encrypt/Decrypt.
  -m    Cipher-based message authentication code.
  -p string
        Password-based key derivation function 2.
  -r    Generate random cryptographic key.
  -s string
        Salt. (for PBKDF2)</pre>

### Example of encryption/decryption:
```sh
./leacrypt -k "" -f plaintext.ext > ciphertext.ext
./leacrypt -d -k $256bitkey -f ciphertext.ext > plaintext.ext
```

## License

This project is licensed under the ISC License.

##### Commercial-Grade Reliability. Copyright (c) 2020-2021 ALBANESE Research Lab.
