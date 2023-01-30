package main 

import (
	"fmt"
	"encoding/json"
)

func main() {
	nameAddrMap := make(map[string]string) 
	name := new(string)
	addr := new(string)
	fmt.Println("Please enter your name:")
	fmt.Scan(name)
	fmt.Println("Please enter your address:")
	fmt.Scan(addr)
	nameAddrMap["name"] = *name
	nameAddrMap["address"] = *addr
	jsonMap, _ := json.Marshal(nameAddrMap)
	fmt.Println(string(jsonMap))
}