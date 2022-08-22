package main

import (
	"fmt"
	"mall/api"
)

func main() {
	str := api.List()
	fmt.Println(str)
	//fmt.Println(controller.List())
}
