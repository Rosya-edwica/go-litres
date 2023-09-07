package tools

import "go-litres/telegram"

func CheckErr(err error) {
	if err != nil {
		telegram.ErrorMessageMailing(err.Error())
		panic(err)
	}
}