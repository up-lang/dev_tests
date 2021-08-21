package main

import (
	"encoding/json"
	"fmt"
	"github.com/up-lang/parser"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please supply one file as an argument")
		return
	}

	loc := os.Args[1]

	root, err := parser.ParseFromFile(loc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	serialized, err := json.Marshal(root)
	if err != nil {
		fmt.Println("json serialization error:\n" + err.Error())
		return
	}

	fmt.Println(string(serialized))
}
