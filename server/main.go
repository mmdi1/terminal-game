package main

import (
	"fmt"
	"terminal/server/route"

	"github.com/aceld/zinx/znet"
)

func main() {
	// 创建Zinx服务器
	server := znet.NewServer()
	// 添加自定义路由
	route.LocalRouter(server)
	// 启动服务器
	fmt.Println("server start...")
	server.Serve()
}
