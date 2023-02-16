package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/rdstihz/SimpleAccountBook/cmd/account/dal"
	"github.com/rdstihz/SimpleAccountBook/kitex_gen/account/accountservice"
	"log"
	"net"
)

func Init() {
	dal.Init()
}

func main() {
	Init()
	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:18882")
	svr := accountservice.NewServer(
		new(AccountImpl),
		server.WithServiceAddr(addr),
	)
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
