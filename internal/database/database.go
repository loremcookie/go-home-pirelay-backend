/*
This is a slightly changed version of the database.go file of th gowebapp boilerplate https://github.com/josephspurrier/gowebapp
*/

package database

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"time"
)

var (
	boltDB       *bolt.DB  //Store boltDB object
	databaseInfo *Database //Store database info
)

//Database type holds information about the boltDB connection like Path to File and Timeout of the boltDB connection.
type Database struct {
	Path    string        //Path to boltDB file
	Timeout time.Duration //Timeout of boltDB connection
}

//Configure must be called before any other function in the database package.
//Makes boltDB connection for all the other function in the database package to use.
//Takes in a Database pointer type for information about boltDB file location and timeout.
//Returns error when unable to connect to given boltDB file.
func Configure(database *Database) error {
	var err error
	//Store database info
	databaseInfo = database
	//Open connection
	boltDB, err = bolt.Open(database.Path, 0600, &bolt.Options{Timeout: database.Timeout})
	if err != nil {
		return err
	}
	return nil
}

// Update stores structure in key.
//Takes the bucket name a key and the struct to store.
func Update(bucketName string, key string, dataStruct interface{}) error {
	err := boltDB.Update(func(tx *bolt.Tx) error {
		// Create the bucket
		bucket, e := tx.CreateBucketIfNotExists([]byte(bucketName))
		if e != nil {
			return e
		}

		// Encode the record
		encodedRecord, e := json.Marshal(dataStruct)
		if e != nil {
			return e
		}

		// Store the record
		if e = bucket.Put([]byte(key), encodedRecord); e != nil {
			return e
		}
		return nil
	})
	return err
}

// View gets structure from key.
//Takes bucket name, key and datastruct to store data in.
func View(bucketName string, key string, dataStruct interface{}) error {
	err := boltDB.View(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		// Retrieve the record
		v := b.Get([]byte(key))
		if len(v) < 1 {
			return bolt.ErrInvalid
		}

		// Decode the record
		e := json.Unmarshal(v, &dataStruct)
		if e != nil {
			return e
		}

		return nil
	})

	return err
}

// Delete deletes key and key data.
//Takes bucket name and key to delete the given key.
func Delete(bucketName string, key string) error {
	err := boltDB.Update(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		return b.Delete([]byte(key))
	})
	return err
}

//Get db struct for custome database operations
func GetDB() *bolt.DB {
	return boltDB
}

//Get info on the current database
func ReadInfo() Database {
	return *databaseInfo
}
