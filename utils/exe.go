package utils

import (
	"os/exec"
	"syscall"
)

// ExecCmd 返回结果状态码，终端【标准输出，错误输出】输出，错误信息
func ExecCmd(command string, args ...string) (int, string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	var res int
	if err != nil {
		if ex, ok := err.(*exec.ExitError); ok {
			res = ex.Sys().(syscall.WaitStatus).ExitStatus()
		}
		return res, "", err
	}
	if utf8Output, err := GbkToUtf8(output); err != nil {
		return res, string(output), nil
	} else {
		return res, string(utf8Output), nil
	}
}
