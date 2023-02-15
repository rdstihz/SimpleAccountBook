package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/rdstihz/SimpleAccountBook/cmd/user/dal"
	user "github.com/rdstihz/SimpleAccountBook/kitex_gen/user/userservice"
	"log"
	"net"
)

func Init() {
	dal.Init()
}

func main() {
	Init()
	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:18881")
	svr := user.NewServer(
		new(UserServiceImpl),
		server.WithServiceAddr(addr),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
