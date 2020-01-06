package util

import "time"

const Token = "N3U3gOan"

const (
	SUCCESS = 0
	ERROR   = 9000
)

const (
	ReadTimeOut  = 5 * time.Second
	WriteTimeOut = 5 * time.Second
)

var ErrMsg = map[int]string{
	SUCCESS: "success",
	ERROR:   "system_error",
}

type Result map[string]interface{}
