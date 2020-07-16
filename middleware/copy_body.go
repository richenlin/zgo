package middleware

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"zgo/modules/config"
	"zgo/modules/helper"

	"github.com/gin-gonic/gin"
)

// CopyBodyMiddleware 复制 request body 内容
func CopyBodyMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	var maxMemory int64 = 64 << 20 // 64 MB
	if v := config.C.HTTP.MaxContentLength; v > 0 {
		maxMemory = v
	}

	return func(c *gin.Context) {
		// 直接跳过
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		// 跳过multipart/form-data数据
		if method := c.Request.Method; method == http.MethodPost || method == http.MethodPut {
			mediaType, _, _ := mime.ParseMediaType(c.GetHeader("Content-Type"))
			if mediaType == "multipart/form-data" {
				c.Next()
				return
			}
		}

		// 没有body数据
		body, err := c.Request.GetBody()
		if err != nil {
			c.Next()
			return
		}

		var requestBody []byte
		isGzip := false
		safe := &io.LimitedReader{R: body, N: maxMemory}

		if c.GetHeader("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(safe)
			if err == nil {
				isGzip = true
				requestBody, _ = ioutil.ReadAll(reader)
			}
		}

		if !isGzip {
			requestBody, _ = ioutil.ReadAll(safe)
		}

		body.Close()
		bf := bytes.NewBuffer(requestBody)
		body = http.MaxBytesReader(c.Writer, ioutil.NopCloser(bf), maxMemory)
		c.Request.Body = body
		c.Set(helper.ReqBodyKey, requestBody)

		c.Next()
	}
}
