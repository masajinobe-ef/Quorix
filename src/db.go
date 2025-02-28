package src

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

func InitDB() *badger.DB {
	db, err := badger.Open(badger.DefaultOptions("./chat-history"))
	if err != nil {
		fmt.Println("Error of Database:", err)
		return nil
	}
	return db
}

func SaveMessage(db *badger.DB, key, value string) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
}

func ReadMessage(db *badger.DB, key string) (string, error) {
	var value []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(nil)
		return err
	})
	return string(value), err
}
