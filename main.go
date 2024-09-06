package main

import (
	"fmt"
	"module1/services"
)
import "github.com/jalaali/go-jalaali"

func main() {

	year, month, day, err := jalaali.ToGregorian(1397, 15, 1)

	if err == nil {
		fmt.Printf("%d-%d-%d\n", year, month, day)
	} else {
		fmt.Printf("%s", err)
	}

	var servcie services.TestService = services.TestService{}
	fmt.Printf("%v", servcie)
}
