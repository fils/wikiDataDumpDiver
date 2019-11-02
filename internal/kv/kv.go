package kv

import (
	"encoding/json"
	"fmt"

	"../../internal/dm"

	"github.com/boltdb/bolt"
)

// Load places the wikidata elements in a kv store..  not sure why yet  ;)
//  ref:  https://github.com/fils/samplesEarth/blob/master/cmd/webui/main.go
func Load(j []byte) {
	n := dm.WikiData{}

	// need to convert this map value interface back to JSON  :(
	json.Unmarshal(j, &n)
	fmt.Println("place holder for a KV writer?")
	fmt.Println(n.Aliases)
	fmt.Println(n.Type)
	fmt.Println(n.Descriptions)
	fmt.Println(n.Labels)

	DBCon.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("wd"))
		err := b.Put([]byte(n.ID), j)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		return err
	})
}
