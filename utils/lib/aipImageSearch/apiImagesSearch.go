package aipImageSearch

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
)

const (
	accept                    = "*/*"
	expirationPeriodInSeconds = "1800"
	AppID                     = "22780756"
	connection                = "keep-alive"
	bceAuthV1                 = "bce-auth-v1"
	acceptEncoding            = "gzip, deflate"
	host                      = "aip.baidubce.com"
	contentType               = "application/x-www-form-urlencoded"
	Ak                        = "HZi779Cu5zZL2nQ59cxAeLbo"
	Sk                        = "jPqVXs0erjS8OMvd5k0vOl1rzO1vxOvl"
	UrlTest                   = "https://bos.cn-n1.baidubce.com/example/测试?text&text1=测试&text10=test"
)

var defaultHeaderSlice = []string{"content-type", "host", "x-bce-request-id"}

func reqFunc(req *http.Request, timestamp string) {
	req.Header.Set("host", host)
	req.Header.Set("accept-encoding", acceptEncoding)
	req.Header.Set("x-bce-date", timestamp)
	req.Header.Set("connection", connection)
	req.Header.Set("accept", accept)
	req.Header.Set("content-type", contentType)
	req.Header.Set("x-bce-request-id", "xxxx")
	req.Header.Set("authorization", "xxxx")
}

// 前缀字符串
func authStringPrefix(timestamp string) string {
	return bceAuthV1 + "/" + Ak + "/" + timestamp + "/" + expirationPeriodInSeconds
}

// 规范请求
func canonicalRequest(req *http.Request, method, uri string) (canonicalRequest, signedHeaders string, err error) {
	urlStr, err := url.ParseRequestURI(uri)
	if err != nil {
		return canonicalRequest, signedHeaders, err
	}

	canonicalRequest = method + "\n" + urlStr.RawPath + "\n"

	queryValue, err := url.ParseQuery(urlStr.RawQuery)
	if err != nil {
		return canonicalRequest, signedHeaders, err
	}

	canonicalRequest += queryValue.Encode() + "\n"

	var canonicalHeaders string
	for k, v := range defaultHeaderSlice {
		h := req.Header.Get(v)
		if len(defaultHeaderSlice) == k+1 {
			canonicalHeaders = v + ":" + uriEncode(h, true)
			signedHeaders += v
		} else {
			canonicalHeaders = v + ":" + uriEncode(h, true) + "\n"
			signedHeaders += v + ";"
		}

	}
	canonicalRequest += canonicalHeaders
	return canonicalRequest, signedHeaders, err
}

// 派生签名密钥
func signingKey(ak, authStringPrefix string) (sha string) {
	key := []byte(ak)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(authStringPrefix))
	sha = hex.EncodeToString(h.Sum(nil))
	return sha
}

// 签名摘要
func signature(sha, canonicalRequest string) (signature string) {
	key := []byte(sha)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(canonicalRequest))
	signature = hex.EncodeToString(h.Sum(nil))
	return signature
}

func uriEncode(charSequence string, encodeSlash bool) (str string) {
	for _, v := range charSequence {
		if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') || v == '_' || v == '-' || v == '~' || v == '.' {
			str += string(v)
		} else if v == '/' && encodeSlash {
			str += "%2F"
		} else if v == '/' {
			str += string(v)
		} else {
			str += fmt.Sprintf("%x", v)
		}
	}
	return str
}
