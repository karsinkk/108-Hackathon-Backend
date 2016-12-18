package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "adminuser8"
	lol, _ := strconv.ParseInt(strings.Split(str, "adminuser")[1], 10, 64)
	fmt.Println(lol)

}
