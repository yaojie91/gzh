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

	"github.com/gomodule/redigo/redis"
)

func init() {
	redisConf := config.Conf.Redis
	AppID := config.Conf.WxKey.AppID
	AppSecret := config.Conf.WxKey.AppSecret
	freshTokenTicker := time.NewTicker(7000 * time.Second)

	redisClient := &redis.Pool{
		MaxIdle:     redisConf.MaxIdle,
		MaxActive:   redisConf.MaxActive,
		IdleTimeout: redisConf.IdleTimeout,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", redisConf.Host)
			if err != nil {
				return nil, err
			}
			return con, err
		},
	}

	go func() {
		for range freshTokenTicker.C {
			token, err := requestToken(AppID, AppSecret)
			if err != nil {
				panic(err)
			}
			rc := redisClient.Get()
			rc.Do("Set", "access-token", token)
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
	fmt.Println("-----------", jMap)
	return jMap["access_token"].(string), nil
}
