package quarrystruct

import (
	"fmt"
)

const (

	// Version - of the  file
	Version = 44
)

// MyValue - of the  file
var (
	MyValue int = 0
)

func init() {

	fmt.Println("init 1 struct module quarrystruct")
}

func init() {

	MyValue = 52
	fmt.Println("init 1 struct module quarrystruct")
}
