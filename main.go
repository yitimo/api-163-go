package main

import (
	"fmt"

	"./sayaka"
)

func main() {
	fmt.Printf("Hello~ This is yitimo's 163 api in golang~\n")
	sayaka.Run("127.0.0.1:9999")
}
