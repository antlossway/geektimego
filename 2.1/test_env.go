package main

import (
	"os"
	"strings"
)

//print out all the environment variables
func main() {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		println(pair[0])
	}
}
