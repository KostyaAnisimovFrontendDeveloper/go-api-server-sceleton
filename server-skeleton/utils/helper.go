package utils

import "fmt"

func Dump(expression any) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}
