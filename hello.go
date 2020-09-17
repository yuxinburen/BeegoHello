package main

import "github.com/astaxie/beego"

func main() {
	//http web服务默认端口是8080
	//beego.Run()
	//自己指定端口
	//Run函数的参数是一个可变参数
	beego.Run(":8090")
}
