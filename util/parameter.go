package util

import "encoding/xml"

type Base struct {
	FromUserName CDATAText
	ToUserName   CDATAText
	CreateTime   int64
	MsgType      CDATAText
}

type CDATAText struct {
	Text string `xml:",innerxml"`
}

type TextMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	Content CDATAText
}
