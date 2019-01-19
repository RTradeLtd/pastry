package main

import (
	"fmt"
	"log"

	"github.com/uhthomas/pastry"
)

func main() {
	n, err := pastry.NewNode(nil)
	if err != nil {
		log.Fatal(err)
	}
	// if you want to edit messages before they're forwarded
	n.Forward = func(key, b, next []byte) {}
	for {
		m, err := n.Next()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m)
	}
}