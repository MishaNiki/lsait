package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// DecodeJSONFile ...
func DecodeJSONFile(path string, v interface{}) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data := make([]byte, 2048)

	var lenBuf int
	for {
		len, e := file.Read(data)
		lenBuf += len
		if e == io.EOF {
			break
		}
	}
	err = json.Unmarshal(data[:lenBuf], v)
	if err != nil {
		return err
	}

	return nil
}

// RandString ...
func RandString(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// Ternary ...
func Ternary(statement bool, a, b interface{}) interface{} {
	if statement {
		return a
	}
	return b
}
