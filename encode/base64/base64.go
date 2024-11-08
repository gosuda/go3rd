package base64

import (
	"encoding/base64"
	"unsafe"
)

func Encode[T ~string | ~[]byte](data T) string {
	switch any((*T)(nil)).(type) {
	case *string:
		v := *(*string)(unsafe.Pointer(&data))
		return base64.StdEncoding.EncodeToString(unsafe.Slice(unsafe.StringData(v), len(v)))
	case *[]byte:
		v := *(*[]byte)(unsafe.Pointer(&data))
		return base64.StdEncoding.EncodeToString(v)
	default:
		panic("unsupported type")
	}
}

func MustDecode[T ~string](base64String T) []byte {
	data, err := Decode(base64String)
	if err != nil {
		panic(err)
	}
	return data
}

// Decode decodes a base64-encoded string.
func Decode[T ~string](base64String T) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(base64String))
}

// IsBase64 validates whether a string is valid base64-encoded.
func IsBase64[T ~string](base64String T) bool {
	_, err := Decode(base64String)
	return err == nil
}
