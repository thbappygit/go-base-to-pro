package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func dd(v interface{}) {
	data, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(data))
	os.Exit(0)
}

func main() {
	user := map[string]interface{}{
		"name": "Habib",
		"age":  25,
		"lang": "Go",
	}

	dd(user)
}
