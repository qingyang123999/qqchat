package common

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestParams struct {
	Method      string                 `json:"method"`
	Path        string                 `json:"path"`
	QueryParams map[string][]string    `json:"query_params"`
	FormParams  map[string][]string    `json:"form_params"`
	JSONParams  map[string]interface{} `json:"json_params"`
	Headers     map[string][]string    `json:"headers"`
	RawBody     string                 `json:"raw_body"`
}

// 获取gin框架c *gin.Context 中所有的请求参数和内容。不管是什么方式请求的
func GetAllRequestParams(c *gin.Context) *RequestParams {
	// 复制请求体以便重复读取
	bodyBytes, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	params := &RequestParams{
		Method:      c.Request.Method,
		Path:        c.Request.URL.Path,
		QueryParams: c.Request.URL.Query(),
		FormParams:  make(map[string][]string),
		JSONParams:  make(map[string]interface{}),
		Headers:     make(map[string][]string),
		RawBody:     string(bodyBytes),
	}

	// 处理表单参数
	if c.Request.Method == http.MethodPost ||
		c.Request.Method == http.MethodPut ||
		c.Request.Method == http.MethodPatch {
		if strings.Contains(c.ContentType(), "form-urlencoded") {
			if err := c.Request.ParseForm(); err == nil {
				for k, v := range c.Request.PostForm {
					params.FormParams[k] = v
				}
			}
		} else if strings.Contains(c.ContentType(), "multipart/form-data") {
			if err := c.Request.ParseMultipartForm(32 << 20); err == nil {
				for k, v := range c.Request.MultipartForm.Value {
					params.FormParams[k] = v
				}
			}
		}
	}

	// 处理JSON参数
	if strings.Contains(c.ContentType(), "application/json") && len(bodyBytes) > 0 {
		_ = json.Unmarshal(bodyBytes, &params.JSONParams)
	}

	// 处理请求头
	for k, v := range c.Request.Header {
		params.Headers[k] = v
	}

	return params
}
