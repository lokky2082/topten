package main

import (
		"fmt"
		"github.com/lokky2082/topten/top"
)

func main() {
	str:= "cat and dog one dog two cats and one man"
  fmt.Println(top.Top(str))
}