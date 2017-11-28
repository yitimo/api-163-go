package main

import (
	"fmt"

	"./sayaka"
)

func main() {
	fmt.Printf("Hello~ This is yitimo's 163 api in golang~\n")
	sayaka.Run("192.168.1.107:9999")
}
