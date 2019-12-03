package main

import (
	"fmt"
	"github.com/yossicohn/testgomodule/quarrystruct"
)

func init() {
	fmt.Println("init main module")

	fmt.Println(quarrystruct.Version)
}

func main() {

	fmt.Println("main function")
	fmt.Println(quarrystruct.Version)
}
