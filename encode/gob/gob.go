package gob

import (
	"bytes"
	"encoding/gob"
)

func Encode(data interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func MustDecode(encoded []byte, out interface{}) {
	err := Decode(encoded, out)
	if err != nil {
		panic(err)
	}
}

func Decode(encoded []byte, out interface{}) error {
	buffer := bytes.NewBuffer(encoded)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(out)
}
