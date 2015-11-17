package sqlite3_test

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	_ "github.com/mutecomm/go-sqlcipher"
)

var db *sql.DB

func init() {
	// create DB
	var key = []byte("passphrase")
	tmpdir, err := ioutil.TempDir("", "sqlcipher_test")
	if err != nil {
		panic(err)
	}
	dbname := filepath.Join(tmpdir, "sqlcipher_test")
	db, err = sql.Open("sqlite3", dbname)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(fmt.Sprintf("PRAGMA key = \"x'%s'\";", hex.EncodeToString(key)))
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("PRAGMA cipher_page_size = 4096;")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`
CREATE TABLE KeyValueStore (
  KeyEntry   TEXT NOT NULL UNIQUE,
  ValueEntry TEXT NOT NULL
);`)
	if err != nil {
		panic(err)
	}
	db.Close()
	// open DB for testing
	db, err = sql.Open("sqlite3", dbname)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(fmt.Sprintf("PRAGMA key = \"x'%s'\";", hex.EncodeToString(key)))
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("PRAGMA cipher_page_size = 4096;")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("SELECT count(*) FROM sqlite_master;")
	if err != nil {
		panic(err)
	}
}

var mapping = map[string]string{
	"foo": "one",
	"bar": "two",
	"baz": "three",
}

func TestInsert(t *testing.T) {
	t.Parallel()
	insertValueQuery, err := db.Prepare("INSERT INTO KeyValueStore (KeyEntry, ValueEntry) VALUES (?, ?);")
	if err != nil {
		t.Fatal(err)
	}
	for key, value := range mapping {
		_, err := insertValueQuery.Exec(key, value)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestSelect(t *testing.T) {
	t.Parallel()
	getValueQuery, err := db.Prepare("SELECT ValueEntry FROM KeyValueStore WHERE KeyEntry=?;")
	if err != nil {
		t.Fatal(err)
	}
	for key, value := range mapping {
		var val string
		err := getValueQuery.QueryRow(key).Scan(&val)
		if err != sql.ErrNoRows {
			if err != nil {
				t.Error(err)
			} else if val != value {
				t.Errorf("%s != %s", val, value)
			}
		}
	}
}
