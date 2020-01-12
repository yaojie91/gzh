package controller

import (
	"encoding/xml"
	"fmt"
	. "https/util"
	"io/ioutil"
	"time"

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
	c.JSON(SUCCESS, ResultError(ERROR, ""))
}

func HandleTextMsg(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	m, _ := mxj.NewMapXml(data)
	xmlData, ok := m["xml"]
	if !ok {
		c.JSON(SUCCESS, ResultError(InValidMsg, ""))
	}
	msg, ok := xmlData.(map[string]interface{})
	if !ok {
		c.JSON(SUCCESS, ResultError(InValidMsg, ""))
	}
	var text []byte
	text, err := Text(msg)
	if err != nil {
		c.JSON(SUCCESS, ResultError(ERROR, fmt.Sprintf("error: %v", err)))
	}
	c.String(SUCCESS, string(text))
}

func Text(msg map[string]interface{}) ([]byte, error) {
	resp := TextMessage{}
	resp.Content = Value2CDATA(msg["Content"].(string))
	resp.FromUserName = Value2CDATA(msg["ToUserName"].(string))
	resp.ToUserName = Value2CDATA(msg["FromUserName"].(string))
	resp.CreateTime = time.Now().Unix()
	resp.MsgType = Value2CDATA(msg["MsgType"].(string))
	respXml, err := xml.Marshal(resp)
	if err != nil {
		return []byte{}, err
	}
	return respXml, nil
}
