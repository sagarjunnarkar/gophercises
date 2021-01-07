package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/learning_go/gophercises/cyoa"
)

func main() {
	file := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Println(*file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", story)
}
