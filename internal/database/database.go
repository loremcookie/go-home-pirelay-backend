/*
This is a slightly changed version of the database.go file of th gowebapp boilerplate github.com/josephspurrier/gowebapp
*/

package database

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"time"
)

var db *bolt.DB //Store BoltDB object

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

	//Open connection
	db, err = bolt.Open(database.Path, 0600, &bolt.Options{Timeout: database.Timeout})
	if err != nil {
		return err
	}
	return nil
}

//Update stores structure in key.
//Takes the bucket name a key and the struct to store.
func Update(bucketName string, key string, dataStruct interface{}) error {
	err := db.Update(func(tx *bolt.Tx) error {
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

//View gets structure from key.
//Takes bucket name, key and dataStruct to store data in.
func View(bucketName string, key string, dataStruct interface{}) error {
	err := db.View(func(tx *bolt.Tx) error {
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

//Delete deletes key and key data.
//Takes bucket name and key to delete the given key.
func Delete(bucketName string, key string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		return b.Delete([]byte(key))
	})
	return err
}

//GetAll returns a slice of maps of all keys and values of a bucket
func GetAll(bucket string) ([]map[string]interface{}, error) {
	//Make map slice to append key value pairs to
	var keyValList []map[string]interface{}

	err := db.View(func(tx *bolt.Tx) error {
		//Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucket))

		//Range over all keys in the USER bucket
		return b.ForEach(func(k, v []byte) error {
			keyValList = append(keyValList, map[string]interface{}{
				"key": string(k),
				"val": string(v),
			})

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	//Return slice of key value pairs of the bucket
	return keyValList, nil
}

//GetAllBuckets returns all top level buckets in a string slice
func GetAllBuckets() ([]string, error) {
	//Make slice to store buckets in
	var buckets []string

	err := db.View(func(tx *bolt.Tx) error {
		//Loop through top level bucket
		return tx.ForEach(func(bucket []byte, _ *bolt.Bucket) error {
			//Append buckets to bucket list
			buckets = append(buckets, string(bucket))
			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return buckets, nil
}

//Reset resets the database to be empty.
//WARNING: This function deletes all data stored in the database
func Reset() error {
	var err error

	//Call Get All Bucket function to get all buckets in the top level bucket
	buckets, err := GetAllBuckets()
	if err != nil {
		return err
	}

	err = db.View(func(tx *bolt.Tx) error {
		//Loop through list of buckets to delete them
		for _, bucket := range buckets {
			//Delete bucket
			e := tx.DeleteBucket([]byte(bucket))
			if err != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

//Close releases all database resources. All transactions must be closed before closing the database.
func Close() error {
	var err error

	//Close database connection
	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}
