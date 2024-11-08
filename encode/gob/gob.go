package gob

import (
	"bytes"
	"encoding/gob"
)

func Encode[T any](data T) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func MustDecode[T any](encoded []byte, out *T) {
	err := Decode(encoded, out)
	if err != nil {
		panic(err)
	}
}

func Decode[T any](encoded []byte, out *T) error {
	buffer := bytes.NewBuffer(encoded)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(out)
}
