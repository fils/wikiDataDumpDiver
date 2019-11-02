package main

import (
	"fmt"

	"gopkg.in/jdkato/prose.v2"
)

//  Would like to see if I could do something like at
// https://medium.com/@errata.ai/prodigy-prose-radically-efficient-machine-teaching-in-go-93389bf2d772

func main() {
	doc, _ := prose.NewDocument("Lebron James plays basketball in Los Angeles.")
	for _, ent := range doc.Entities() {
		fmt.Println(ent.Text, ent.Label)
		// Lebron James PERSON
		// Los Angeles GPE
	}
}
