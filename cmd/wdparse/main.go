package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"../../internal/dm"
	"../../internal/kv"

	"github.com/bcicen/jstream"
	"github.com/boltdb/bolt"
)

func init() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	kv.DBCon = db

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("wd"))
		if err != nil {
			return fmt.Errorf("create bucket issue: %s", err)
		}
		return nil
	})
}

// TODO
// Flesh out the KV store option
// Look at redis search   https://redislabs.com/redis-enterprise/technology/redis-search/

func main() {
	fmt.Println("Wikidata json data dump streaming parser")

	// take arg 1 as the file name to load
	arg := os.Args[1]
	f, err := os.Open(arg)
	if err != nil {
		log.Println(err)
	}

	decoder := jstream.NewDecoder(f, 1) // extract JSON values at a depth level of 1
	for mv := range decoder.Stream() {
		j, err := json.Marshal(mv.Value)
		if err != nil {
			log.Println(err)
		}
		simplePrint(j)
		kv.Load(j)
	}
}

// Just a simple print func for peace of mind
func simplePrint(j []byte) {
	n := dm.WikiData{}

	// need to convert this map value interface back to JSON  :(
	json.Unmarshal(j, &n)
	fmt.Println("--------")
	fmt.Println(n.ID)
	fmt.Println(n.Aliases)
	fmt.Println(n.Type)
	fmt.Println(n.Descriptions)
	fmt.Println(n.Labels)
}
