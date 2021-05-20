package main

import (
	"fmt"
	helper "github.com/aldelo/common"
	"time"
)

func main() {

	fmt.Println(time.Now().UTC())
	fmt.Println(helper.FormatDateTime(time.Now()))
	fmt.Println(helper.FormatDateTime(time.Now().UTC()))
}
