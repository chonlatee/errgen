package main

import (
	"log"

	"github.com/chonlatee/errgen/errs"
)

func main() {
	err1 := errs.OrderNotFound("order1")
	log.Println(err1)

	err2 := errs.UserNameTooShort("isme", 10)
	log.Println(err2)

	err3 := errs.UserNotFound("isme")
	log.Println(err3)
}
