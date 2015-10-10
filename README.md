## go-sqlcipher [![GoDoc](https://godoc.org/github.com/mutecomm/go-sqlcipher?status.png)](http://godoc.org/github.com/mutecomm/go-sqlcipher) [![Build Status](https://travis-ci.org/mutecomm/go-sqlcipher.png)](https://travis-ci.org/mutecomm/go-sqlcipher)

### Note

Currently SQLCipher has a bug which prevents parallel access to an encrypted
database. The automatic tests will fail until the bug is fixed.


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

API documentation can be found here:
http://godoc.org/github.com/mutecomm/go-sqlcipher

Examples can be found under the `./_example` directory


### License

The code of the originating packages is covered by their respective licenses.
See [LICENSE](LICENSE) file for details.
