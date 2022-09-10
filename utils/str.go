package utils

import (
	"math/rand"
	"strings"
	"time"
)

func ClearCrLR(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "\n", `\n`), "\r", ``)
}

// RandString 生成随机字符串
func RandString(slen int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	randbytes := make([]byte, slen)
	for i := 0; i < slen; i++ {
		b := r.Intn(26) + 65
		randbytes[i] = byte(b)
	}
	return string(randbytes)
}

func RandIntString(slen int) string {
	str := []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	str_len := len(str)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	randbytes := make([]byte, slen)
	for i := 0; i < slen; i++ {
		b := r.Intn(str_len)
		randbytes[i] = str[b]
	}
	return string(randbytes)
}
