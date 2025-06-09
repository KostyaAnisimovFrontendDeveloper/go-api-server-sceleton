package utils

import (
	"fmt"
	"log"
	"runtime"
	"server-skeleton/api_init"
)

func Dump(expression ...any) {

	_, file, line, ok := runtime.Caller(1)
	for _, expression := range expression {

		fmt.Println(fmt.Sprintf("\n\n==========\n%#v \n------------ \nfile: %s \nline:%d \nisOk:%t\n\n==========\n\n", expression, file, line, ok))
	}
}

func LogFormat(message string, context string) {
	_, file, line, isOk := runtime.Caller(1)

	serviceName := api_init.InitGlobal.Cfg.ServerName
	serviceDomine := api_init.InitGlobal.Cfg.ServerDomain

	ok := "ok"
	if !isOk {
		ok = "error!!!"
	}

	result := fmt.Sprintf(
		"[%s][%s] {%s}-{%d}:%s %s context: %s",
		serviceDomine,
		serviceName,
		file,
		line,
		ok,
		message,
		context,
	)

	log.Println(result)
}
