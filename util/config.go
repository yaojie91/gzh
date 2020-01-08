package util

import "time"

const Token = "N3U3gOan"

const (
	SUCCESS    = 200
	ERROR      = 9000
	InValidMsg = 3001
)

const (
	ReadTimeOut  = 5 * time.Second
	WriteTimeOut = 5 * time.Second
)

var ErrMsg = map[int]string{
	SUCCESS:    "success",
	ERROR:      "system_error",
	InValidMsg: "invalid_message",
}

type Result map[string]interface{}
