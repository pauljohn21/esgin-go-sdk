package main

import (
	"fmt"

	"esgin-go-sdk/service"
)

func main() {
	service.EsignInitService()

	res := service.DocTemplatesList()
	fmt.Println(res.Data)
}
