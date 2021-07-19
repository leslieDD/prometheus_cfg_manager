package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GbkToUtf8 GbkToUtf8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// Utf8ToGbk Utf8ToGbk
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Getmd5(level int, nId int, s ...string) string {
	has := md5.Sum([]byte(fmt.Sprintf("%d_%d_%s", level, nId, strings.Join(s, "_"))))
	return fmt.Sprintf("%x", has)
}
