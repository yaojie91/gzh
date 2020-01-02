package util

const Token = "abcdefgh"

const (
	SUCCESS = 0
	ERROR   = 9000
)

var ErrMsg = map[int]string{
	SUCCESS: "success",
	ERROR:   "system_error",
}

type Result map[string]interface{}
