package utils

import (
	"crypto/sha1"
)

// CheckPassword 验证提供的密码是否合法
func CheckPassword(pwdPost, pwdDB, salt string) bool {
	return pwdPost == pwdDB
	// sccrte := PwdSha1(fmt.Sprintf("%s-%s", pwdPost, salt))
	// if sccrte == pwdDB {
	// 	return true
	// }
	// return false
}

// GenerateToken GenerateToken
func GenerateToken(s string) string {
	s1 := sha1.New()
	s1.Write([]byte(s))
	return string(s1.Sum(nil))
}

// PwdSha1 PwdSha1
func PwdSha1(s string) string {
	s1 := sha1.New()
	s1.Write([]byte(s))
	return string(s1.Sum(nil))
}
