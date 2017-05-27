package yzgo

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
	"net/url"
)

const (
	// OpenAPI url
	OpenAPI = "https://open.youzan.com/api"
	//TimeFormatter Youzan Timestamp format
	TimeFormatter = "2006-01-02 15:04:05"
)

// Util
func getMd5String(src string) string {
	h := md5.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}

// YZClient standard for Youzan Open API Client
type YZClient struct {
	AppID     string
	AppSecret string
	IsOAuth bool
	AccessToken string
}

func (c *YZClient) getSign(secret string, params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	plainText := secret
	for _, key := range keys {
		plainText += key + params[key]
	}
	plainText += secret
	return getMd5String(plainText)
}

func (c *YZClient) getSignedParams(params map[string]string) map[string]string {
	timestamp := time.Now().Format(TimeFormatter)
	format := "json"
	appID := c.AppID
	appSecret := c.AppSecret
	v := "1.0"
	signMethod := "md5"

	paramsMap := map[string]string{
		"timestamp":   timestamp,
		"format":      format,
		"app_id":      appID,
		"v":           v,
		"sign_method": signMethod,
	}

	for key, value := range params {
		paramsMap[key] = value
	}
	sign := c.getSign(appSecret, paramsMap)
	paramsMap["sign"] = sign

	return paramsMap
}

func (c *YZClient) sendRequest(rawURL string, method string, params map[string]string, files map[string]string) (*http.Response, error) {
	httpClient := &http.Client{}
	// TODO support files

	var req *http.Request
	var err error
	if "GET" == strings.ToUpper(method) {
		httpURL := rawURL
		if len(params) > 0 {
			httpURL += "?"
			for key, value := range params {
				httpURL += url.QueryEscape(key)
				httpURL += "="
				httpURL += url.QueryEscape(value)
				httpURL += "&"
			}
			httpURL = strings.TrimRight(httpURL, "&")
		}
		// println(httpURL)
		req, err = http.NewRequest("GET", httpURL, nil)

	} else if "POST" == strings.ToUpper(method) {
		jsonData, _ := json.Marshal(params)
		dataReader := bytes.NewReader(jsonData)
		req, err = http.NewRequest("POST", rawURL, dataReader)
		req.Header.Set("Content-Type", "application/json")
		println(jsonData)
	} else {
		panic(errors.New("not support method"))
	}

	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "X-YZ-Client 1.0.0 - Golang")
	return httpClient.Do(req)
}

// Invoke Youcan API
func (c *YZClient) Invoke(apiName string, version string, method string, params map[string]string, files map[string]string) string {
	httpURL := OpenAPI
	apiParts := strings.Split(apiName, ".")
	service := strings.Join(apiParts[0:len(apiParts)-1], ".")
	action := apiParts[len(apiParts)-1]

	var paramsMap map[string]string
	if c.IsOAuth {
		httpURL += "/oauthentry"
		paramsMap = params
		paramsMap["access_token"] = c.AccessToken
	} else {
		httpURL += "/entry"
		paramsMap = c.getSignedParams(params)
	}

	httpURL += "/" + service + "/" + version + "/" + action

	resp, err := c.sendRequest(httpURL, method, paramsMap, files)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		println("Invoke failed: " + resp.Status)
		return ""
	}
	return string(content)
}
