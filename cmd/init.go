package cmd

import (
	"github.com/Bnei-Baruch/go-example/api"
)

func Init() {
	a := api.App{}
	a.InitApp()
}