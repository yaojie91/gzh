package controller

import (
	. "https/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckSig(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	sig := SignatureGen(Token, timestamp, nonce)
	if sig == signature {
		c.JSON(http.StatusOK, ResultSuccess(echostr))
		return
	}
	c.JSON(http.StatusOK, ResultError(ERROR))
}