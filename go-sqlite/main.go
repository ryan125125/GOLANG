package main

import (
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func NewGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return
	}
	if db == nil {
		fmt.Println("create fail")
	} else {
		fmt.Println("create ok")
	}
	db.Close()
	// sql_table := `
	// CREATE TABLE IF NOT EXISTS userinfo(
	//     uid INTEGER PRIMARY KEY AUTOINCREMENT,
	//     username VARCHAR(64) NULL,
	//     departname VARCHAR(64) NULL,
	//     created DATE NULL
	// );
	// `

	// db.Exec(sql_table)
}
