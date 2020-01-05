package main

import (
		"fmt"
		"github.com/lokky2082/topten/top"
)

func main() {
  top10:= "cat and dog one dog two cats and one man"
  fmt.Println(top.Top(top10))
}