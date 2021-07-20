package utils

import "strings"

func ClearCrLR(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "\n", `\n`), "\r", ``)
}
