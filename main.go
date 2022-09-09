package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	in, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := make(map[string]interface{})
	err = json.Unmarshal(in, &input)
	if err != nil {
		log.Fatal(err)
	}
	res := flatten(input)
	fmt.Println(res)
}

func flatten(input map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for k1, v1 := range input {
		switch v1.(type) {
		case map[string]interface{}:
			cur := flatten(v1.(map[string]interface{}))
			for k2, v2 := range cur {
				res[fmt.Sprintf("%s.%s", k1, k2)] = v2
			}
		default:
			res[k1] = v1
		}
	}
	return res
}
