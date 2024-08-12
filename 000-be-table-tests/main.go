package main

import (
	"fmt"
	"learn-go-v-three/000-be-table-tests/tokenizer"
)

func main() {
	ts := tokenizer.Tokenize("Let's go to the beach")
	fmt.Println(ts)
}

/*
to run

go test ./...

or

go test -v ./...

*/
