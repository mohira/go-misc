package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	const jsonStrBob = `{"name": "Bob", "age": 20}`
	const jsonStrTom = `{"name": "Tom", "age": 35}`

	var bob User

	decoder1 := json.NewDecoder(strings.NewReader(jsonStrBob))
	if err := decoder1.Decode(&bob); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", bob) // {Name:Bob Age:20}
	fmt.Printf("%#v\n", bob) // main.User{Name:"Bob", Age:20}

	var tom User
	decoder2 := json.NewDecoder(strings.NewReader(jsonStrTom))
	if err := decoder2.Decode(&tom); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", tom) // {Name:Tom Age:35}
	fmt.Printf("%#v\n", tom) // main.User{Name:"Tom", Age:35}
}
