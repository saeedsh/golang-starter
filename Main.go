package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("FOO:", os.Getenv("TEST_ENV"))
}
