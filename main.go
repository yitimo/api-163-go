package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"./sayaka"
)

func main() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Sayaka can not read the config correctly :(")
		return
	}
	var config map[string]interface{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Sayaka can not read the config correctly :(")
		return
	}
	fmt.Printf("Hello~ This is yitimo's 163 api in golang~\n")
	sayaka.Run(config)
}
