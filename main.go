package main

import (
	"https/route"

	"github.com/fvbock/endless"
)

func main() {
	router := route.InitRouter()
	//router.GET("/hello", func(c *gin.Context) {
	//	c.JSON(200, "hello world")
	//})
	//router.GET("/", func(c *gin.Context) {
	//	signature := c.Query("signature")
	//	timestamp := c.Query("timestamp")
	//	nonce := c.Query("nonce")
	//	echostr := c.Query("echostr")
	//
	//	sig := util.SignatureGen(timestamp, nonce)
	//	if sig == signature {
	//
	//	}
	//})

	endless.ListenAndServe(":8080", router)
}
