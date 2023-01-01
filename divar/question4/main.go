package main

import (
	"bufio"
	"log"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
func main() {
	in, err:=readInputs()
	if err != nil {
		log.Fatalf("invalid inputs: %v", err)
	}
}

func readInputs() (interface{}, error){
	input, _ := reader.ReadString('\n')
}
