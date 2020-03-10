package main

import (
	"gzh/route"
	"gzh/util"

	"github.com/fvbock/endless"
)

func main() {
	router := route.InitRouter()
	endless.DefaultReadTimeOut = util.ReadTimeOut
	endless.DefaultWriteTimeOut = util.WriteTimeOut

	endless.ListenAndServe(":8080", router)
}
