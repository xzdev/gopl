package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	count := make(map[string]int)
	files := os.Args[1:]
	fmt.Println("Arguments:", files)

	for _, fileName := range files {
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
		}
		for _, line := range strings.Split(string(file), "\n") {
			count[line]++
		}
	}

	fmt.Println("----")
	for name, value := range count {
		if value > 1 {
			fmt.Printf("%d\t%s\n", value, name)
		}
	}
}
