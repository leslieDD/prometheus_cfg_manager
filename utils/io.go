package utils

import (
	"io/ioutil"
	"os"
)

// RIoutil RIoutil
func RIoutil(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
}

// WIoutil WIoutil
func WIoutil(name, content string) error {
	data := []byte(content)
	return ioutil.WriteFile(name, data, 0644)
}

// WIoutilByte WIoutilByte
func WIoutilByte(name string, content []byte) error {
	data := []byte(content)
	return ioutil.WriteFile(name, data, 0644)
}

// PathExists PathExists
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
	}
	return true, nil
}

// FileExists FileExists
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
