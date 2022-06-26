package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Create("./hello.txt")
	if err != nil {
		fmt.Printf("failed to create file \n: %v", err)
		return
	}

	fmt.Println("Done!")
}
