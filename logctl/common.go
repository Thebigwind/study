package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func ContainString(arr []string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func ExecShell(s string) (error, string) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("exec_shell:\n", err.Error())
		return err, ""
	}
	//fmt.Printf("%s", out.String())
	return err, strings.Trim(out.String(), "\n")
}
