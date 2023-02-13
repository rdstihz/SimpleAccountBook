package main

import (
	account "github.com/rdstihz/SimpleAccountBook/kitex_gen/account/accountservice"
	"log"
)

func main() {
	svr := account.NewServer(new(AccountImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
