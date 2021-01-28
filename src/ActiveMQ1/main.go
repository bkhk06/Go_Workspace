package main

import (
    "net"
    "fmt"
    "os"
    "github.com/gpmgo/gopm/modules/goconfig"
    "github.com/go-stomp/stomp"
    "strconv"
)
// 读取配置文件
func getConfigFile(filePath string) (configFile *goconfig.ConfigFile){
    configFile, err := goconfig.LoadConfigFile(filePath)
    if err != nil {
        fmt.Println("load config file error: " + err.Error())
        os.Exit(1)
    }
    return configFile
}

// 使用IP和端口连接到ActiveMQ服务器
// 返回ActiveMQ连接对象
func connActiveMq(host, port string) (stompConn *stomp.Conn){// @todo 实现断开重连
    stompConn, err := stomp.Dial("tcp", net.JoinHostPort(host, port))
    if err != nil {
        fmt.Println("connect to active_mq server service, error: " + err.Error())
        os.Exit(1)
    }

    return stompConn
}
// 将消息发送到ActiveMQ中
func activeMqProducer(c chan string, queue string, conn *stomp.Conn){
    for{
        err := conn.Send(queue, "text/plain", []byte(<-c))
        fmt.Println("send active mq..." + queue)
        if err != nil {
            fmt.Println("active mq message send erorr: " + err.Error())
        }
    }
}


func main() {
    configPath := "./config.ini"
    configFile := getConfigFile(configPath)

    host, err := configFile.GetValue("active_mq", "host")
    if err != nil {
        fmt.Println("get active_mq host error: " + err.Error())
        os.Exit(1)
    }
    port, err:= configFile.GetValue("active_mq", "port")
    if err != nil {
        fmt.Println("get active_mq port error: " + err.Error())
        os.Exit(1)
    }

    queue, err := configFile.GetValue("active_mq", "queue")
    if err != nil {
        fmt.Println("get active_mq queue error: " + err.Error())
        os.Exit(1)
    }

    activeMq := connActiveMq(host, port)
    defer activeMq.Disconnect()
    c := make(chan string)　　// 启动Go routine发送消息
    go activeMqProducer(c, queue, activeMq)

    for i := 0; i < 10000; i ++{        // 发送1万个消息
        c <- "hello world" + strconv.Itoa(i)
    }

}