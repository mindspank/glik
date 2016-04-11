package main

import (
	"fmt"
	"github.com/mattbaird/glik"
)

func main() {
	api := glik.DefaultApi()
	out, err := api.About()
	fmt.Printf("out:%v\n", out)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	listOfApps, err := api.List()
	fmt.Printf("list:%v\n", listOfApps)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}

}
