package main

import (
  "time"
  "github.com/boltdb/bolt"
)

func OpenDB() (*bolt.DB, error) {
  db, err := bolt.Open("azurabot.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
  if err != nil {
    return nil, err
  }
  defer db.Close()

  return db, nil
}

// CreateDB create a database file if it if was not exist
func CreateDB() error {
  db,err := OpenDB()
  if err != nil {
    return err
  }

  err = db.Update(func(tx *bolt.Tx) error {
    _, err := tx.CreateBucketIfNotExists([]byte("ChannelDB"))
    if err != nil {
      return err
    }
    return nil
  })
  return err
}

// PutDB ignore o unignore a test channel
func PutDB(channelID, ignored string) error {
  db,err := OpenDB()
  if err != nil {
    return err
  }

  db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("ChannelDB"))
    err := b.Put([]byte(channelID), []byte(ignored))
    return err
  })
  return err
}

// GetDB read if a text channel is ignored
func GetDB(channelID string) string {
  var v []byte

  db,err := OpenDB()
  if err != nil {
    return ""
  }

  db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("ChannelDB"))
    v = b.Get([]byte(channelID))
    return nil
  })
  return string(v)
}
