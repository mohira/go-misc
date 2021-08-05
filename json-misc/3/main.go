package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type MySlice []string

func (s MySlice) MarshalJSON() ([]byte, error) {
	m := make(map[string]string)

	for i := 0; i < len(s); i += 2 {
		m[s[i]] = s[i+1]
	}

	return json.Marshal(m)
}

func main() {
	s := MySlice{"foo", "bar", "hoge", "fuga"}

	var buf bytes.Buffer

	encoder := json.NewEncoder(&buf)

	if err := encoder.Encode(s); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
