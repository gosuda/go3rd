package base64

import (
	"encoding/base64"
)

func Encode(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

func MustDecode(base64String string) []byte {
	data, err := Decode(base64String)
	if err != nil {
		panic(err)
	}
	return data
}

func Decode(base64String string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64String)
}

func IsBase64(str string) bool {
	_, err := Decode(str)
	return err == nil
}
