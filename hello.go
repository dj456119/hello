/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-11-16 11:43:17
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-16 12:23:20
 */
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	buff, err := os.Open("map.qs")
	if err != nil {
		fmt.Println("Open file error ", err)
	}
	reader := bufio.NewReader(buff)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Read line: ", string(line))
		result := strings.Split(string(line), " ")
		t, err := time.Parse("2006-01-02", result[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		commitString := fmt.Sprintf("git commit --date=\"%s %d 9:05:20 2020 +0800\" -m \"提交\"", t.Month(), t.Day())
		fmt.Println("Commit string is ", commitString)
		Shellout("echo a >> hello.txt")
		Shellout("git add *")
		Shellout(commitString)
		Shellout("git push origin")
	}
}

func Shellout(command string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("err cmd :", command)
		fmt.Println("stdout: ", stdout.String())
		fmt.Println("stderr: ", stderr.String())
		os.Exit(-1)
	}
	fmt.Println("exec cmd done:", command)
	fmt.Println("stdout: ", stdout.String())
	fmt.Println("stderr: ", stderr.String())
}
