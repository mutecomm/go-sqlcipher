## go-sqlcipher [![GoDoc](https://godoc.org/github.com/mutecomm/go-sqlcipher?status.png)](http://godoc.org/github.com/mutecomm/go-sqlcipher) [![Build Status](https://travis-ci.org/mutecomm/go-sqlcipher.png)](https://travis-ci.org/mutecomm/go-sqlcipher)

### Description

Self-contained Go sqlite3 driver with an AES-256 encrypted sqlite3 database
conforming to the built-in database/sql interface. It is based on:

- stripped-down Go sqlite3 driver: https://github.com/mattn/go-sqlite3
- SQLite extension with AES-256 codec: https://github.com/sqlcipher/sqlcipher
- AES-256 implementation from: https://github.com/libtom/libtomcrypt

SQLite itself is part of SQLCipher.


### Installation

This package can be installed with the go get command:

    go get github.com/mutecomm/go-sqlcipher


### Documentation

To create and open encrypted database files use the following DSN parameters:

```go
key := "2DD29CA851E7B56E4697B0E1F08507293D761A05CE4D1B628663F411A8086D99"
dbname := fmt.Sprintf("test?_pragma_key=%s&_pragma_cipher_page_size=4096", key)
db, _ := sql.Open("sqlite3", dbname)
```

`_pragma_key` is the hex encoded 32 byte key (must be 64 characters long).
`_pragma_cipher_page_size` is the page size of the encrypted database (set if
you want a different value than the default 1024 bytes).

API documentation can be found here:
http://godoc.org/github.com/mutecomm/go-sqlcipher

Examples can be found under the `./_example` directory


### License

The code of the originating packages is covered by their respective licenses.
See [LICENSE](LICENSE) file for details.
