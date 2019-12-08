package main

import "fmt"

func main() {
	a := map[string]map[string]interface{}{"a": {"b": 1}}
	if a["a"]["b"] != nil {
		fmt.Println(float64(a["a"]["b"]))
	}
}
