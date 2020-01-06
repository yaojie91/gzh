package controller

import (
	"fmt"
	. "https/util"
	"io/ioutil"
	"net/http"

	"github.com/clbanning/mxj"
	"github.com/gin-gonic/gin"
)

func CheckSig(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	sig := SignatureGen(Token, timestamp, nonce)
	if sig == signature {
		c.Writer.Write([]byte(echostr))
		return
	}
	c.JSON(http.StatusOK, ResultError(ERROR))
}

func ReplyMsg(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("----------", data)
	m, _ := mxj.NewMapXml(data)
	fmt.Println("---------------", m)
}
