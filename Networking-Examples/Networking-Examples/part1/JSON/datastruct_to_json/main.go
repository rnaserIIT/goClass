package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//encoding basic data types to JSON strings

	boolJson, _ := json.Marshal(true) // returns as array of bytes, most efficient for sending
	fmt.Println(string(boolJson))
	boolInt, _ := json.Marshal(10)
	fmt.Println(string(boolInt))

	// encoding slice to JSON
	slice := []string{"golang", "python", "c++"}
	sliceJSON, _ := json.Marshal(slice)
	fmt.Println(string(sliceJSON))

	// encoding map to JSON
	map0 := map[string]int{"golang": 100, "python": 50, "c++": 0}
	mapJSON, _ := json.Marshal(map0)
	fmt.Println(string(mapJSON))

	// next we go over marshaling struct types to JSON
	// and how to convert JSON back into the origional types (as in after being returned from a web server)

}
