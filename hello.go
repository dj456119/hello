/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-11-16 11:43:17
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-16 11:53:50
 */
package main

import (
	"bufio"
	"fmt"
	"os"
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
		commitString := fmt.Sprintf("git commit --date=\"%s %d 9:05:20 2021 +0800\" -m \"提交\"", t.Month(), t.Day())
		fmt.Println("Commit string is ", commitString)
	}
}
