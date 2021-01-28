package main

//客户端

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接失败，错误：", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdout)
	for {
		input, _ := inputReader.ReadString('\n') //读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "q" {
			return //如果输入q就退出
		}
		_, err = conn.Write([]byte(inputInfo)) //发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("接受失败，错误：", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
