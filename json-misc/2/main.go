package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// json.RawMessage でパースを遅延させる
func main() {
	const jsonStr1 = `{"type": "foo", "data": {"hoge": 1}}`
	const jsonStr2 = `{"type": "bar", "data": {"fuga": 2}}`

	var v struct {
		Type string          `json:"type"`
		Data json.RawMessage `json:"data"`
	}

	decoder := json.NewDecoder(strings.NewReader(jsonStr1))
	//decoder := json.NewDecoder(strings.NewReader(jsonStr2))

	if err := decoder.Decode(&v); err != nil {
		panic(err)
	}

	// json.RawMessage は []byte型
	fmt.Printf("v.Data=%T\n", v.Data) // json.RawMessage
	fmt.Printf("%+v\n", v)            // {Type:foo Data:[123 34 104 111 103 101 34 58 32 49 125]}

	switch v.Type {
	case "foo":
		var data struct {
			Hoge int `json:"hoge"`
		}

		decoder := json.NewDecoder(bytes.NewReader(v.Data))
		if err := decoder.Decode(&data); err != nil {
			panic(err)
		}

		fmt.Printf("%+v\n", data)
	case "bar":
		var data struct {
			Fuga int `json:"fuga"`
		}

		decoder := json.NewDecoder(bytes.NewReader(v.Data))
		if err := decoder.Decode(&data); err != nil {
			panic(err)
		}

		fmt.Printf("%+v\n", data)
	}
}
