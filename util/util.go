package util

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"gzh/config"
	"net/http"
	url2 "net/url"
	"sort"
	"strings"
	"time"
)

func init() {
	AppID := config.Conf.WxKey.AppID
	AppSecret := config.Conf.WxKey.AppSecret
	freshTokenTicker := time.NewTicker(7000 * time.Second)
	rc := config.RedisPool.Get()

	token, err := requestToken(AppID, AppSecret)
	if err != nil {
		panic(err)
	}
	_, err = rc.Do("Set", "access_token", token)
	go func() {
		for range freshTokenTicker.C {
			token, err := requestToken(AppID, AppSecret)
			if err != nil {
				panic(err)
			}
			_, err = rc.Do("Set", "access_token", token)
			if err != nil {
				fmt.Println("err: ", err)
			}
			rc.Close()
		}
	}()
}

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

func requestToken(appId, appSecret string) (string, error) {
	u, err := url2.Parse("https://api.weixin.qq.com/cgi-bin/token")
	if err != nil {
		panic(err)
	}

	params := &url2.Values{}
	params.Set("grant_type", "client_credential")
	params.Set("appid", appId)
	params.Set("secret", appSecret)
	u.RawQuery = params.Encode()
	resp, err := http.Get(u.String())
	//关闭资源
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", errors.New("request token err :" + err.Error())
	}
	jMap := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&jMap)
	if err != nil {
		return "", errors.New("request token response json parse err :" + err.Error())
	}
	if jMap["errcode"] == nil || jMap["errcode"] == 0 {
		return jMap["access_token"].(string), nil
	} else {
		//ResultError(jMap["errcode"].(int), jMap["errmsg"].(string))
		return "", errors.New("request token response json parse err :" + jMap["errmsg"].(string))
	}
	//return jMap["access_token"].(string), nil
}
