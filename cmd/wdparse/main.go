package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"../../internal/dm"
	"github.com/bcicen/jstream"
)

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

func kv(j []byte) {
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
