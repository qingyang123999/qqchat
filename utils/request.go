package utils

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
	"time"
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

// 通过页码也分页大小获取mysql偏移量
func GetPageOffset(page int, pageSize int) int {
	if page < 1 || pageSize < 0 {
		return 0
	} else {
		return (page - 1) * pageSize
	}
}

// 自定义时间类型，用于JSON序列化时格式化时间  2006-01-02 15:04:05这种格式
type CustomTime struct {
	time.Time
}

// 实现 MarshalJSON 接口来自定义时间格式
func (t CustomTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Time.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// 实现 Valuer 接口用于将自定义时间类型存储到数据库
func (t CustomTime) Value() (driver.Value, error) {
	return t.Time, nil
}

// 实现 Scanner 接口用于从数据库读取自定义时间类型
func (t *CustomTime) Scan(value interface{}) error {
	if value == nil {
		t.Time = time.Time{}
		return nil
	}

	if v, ok := value.(time.Time); ok {
		t.Time = v
		return nil
	}

	return fmt.Errorf("cannot scan %T into CustomTime", value)
}

// NewCustomTime 创建一个新的CustomTime实例
// 使用时：
// LoginTime:     models.NewCustomTime(time.Now()),
func NewCustomTime(t time.Time) CustomTime {
	return CustomTime{Time: t}
}

// NewCustomTimeFromNow 创建一个表示当前时间的CustomTime实例
// 使用时，你可以这样写：
// LoginTime:     models.NewCustomTimeFromNow(),
func NewCustomTimeFromNow() CustomTime {
	return CustomTime{Time: time.Now()}
}
