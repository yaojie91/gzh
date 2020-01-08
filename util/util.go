package util

import (
	"crypto/sha1"
	"fmt"
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
	return fmt.Sprintf("%x", result)
}

func ResultSuccess(data interface{}) Result {
	return Result{
		"errno":  SUCCESS,
		"errmsg": ErrMsg[SUCCESS],
		"data":   data,
	}
}

func ResultError(code int, msg string) Result {
	if msg == "" {
		msg = ErrMsg[code]
	}
	return Result{
		"errno":  code,
		"errmsg": ErrMsg[code],
		"data":   "",
	}
}

func Value2CDATA(v string) CDATAText {
	return CDATAText{"<![CDATA[" + v + "]]>"}
}
