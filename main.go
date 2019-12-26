package main

import (
		"fmt"
		"github.com/lokky2082/topten/top"
)

func main() {
	top10:= "cat and dog one dog two cats and one man test1 test2 test3 test4 test5 test6"
  fmt.Println(top.Top(top10))
}