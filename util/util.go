package util

import (
	"crypto/sha1"
	"sort"
	"strings"
)

func SignatureGen(token string, timestamp string, nonce string) string {
	arr := make([]string, 0)
	arr = append(arr, token, timestamp, nonce)
	sort.Strings(arr)
	tmpArr := strings.Join(arr, "")
	// 处理字符串
	h := sha1.New()
	h.Write([]byte(tmpArr))
	result := h.Sum(nil)
	return string(result)
}

func ResultSuccess(data interface{}) Result {
	return Result{
		"errno":  SUCCESS,
		"errmsg": ErrMsg[SUCCESS],
		"data":   data,
	}
}

func ResultError(code int) Result {
	//if msg == "" {
	//	msg = ErrMsg[code]
	//}
	return Result{
		"errno":  code,
		"errmsg": ErrMsg[code],
		"data":   "",
	}
}
