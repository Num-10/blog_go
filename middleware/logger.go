package middleware

import (
	"blog_go/pkg"
	"database/sql/driver"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"time"
	"unicode"
)

var CurrentC *gin.Context

type MysqlLog struct {}

func (log *MysqlLog) Print (values ...interface{})  {
	if CurrentC == nil {
		//脚本SQL不作日志记录
		/*pkg.Logger.WithFields(logrus.Fields{
			"sql": LogFormatter(values...),
		}).Info("cron sql log")*/
	} else {
		sql := CurrentC.GetStringSlice("mysql log collect")

		current_sql := LogFormatter(values...)

		CurrentC.Set("mysql log collect", append(sql, current_sql...))
	}

}

func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(c.Request.Header["Origin"]) > 0 {
			method := c.Request.Method
			// 核心处理方式
			c.Header("Access-Control-Allow-Origin", c.Request.Header["Origin"][0])
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, token")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
			//放行所有OPTIONS方法
			if method == "OPTIONS" {
				c.JSON(http.StatusOK, "Options Request!")
			}
		}

		CurrentC = c
		// 开始时间
		startTime :=  time.Now().UnixNano() / 1e3
		// 处理请求
		c.Next()
		// 结束时间
		endTime :=  time.Now().UnixNano() / 1e3
		// 执行时间
		latencyTime := float64(endTime - startTime) / 1e6
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 请求参数
		reqParams := c.Request.Form
		// 请求参数
		reqHeader := c.Request.Header
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 执行SQL
		sql, _ := c.Get("mysql log collect")

		// 日志格式
		pkg.Logger.Printf("IP: %3s | Method: %3s | Code: %3d | Uri: %3s | 运行时间：%vs | PARAM： %+v | HEADER： %+v  | MYSQL:  %v",
			clientIP,
			reqMethod,
			statusCode,
			reqUri,
			latencyTime,
			reqParams,
			reqHeader,
			sql,
		)
	}
}

var LogFormatter = func(values ...interface{}) (messages []string) {
	if len(values) > 1 {
		var (
			sql             string
			formattedValues []string
			level           = values[0]
			source          = fmt.Sprintf("%v", values[1])
		)

		messages = []string{"{", source, "|"}

		if level == "sql" {
			// duration
			messages = append(messages, fmt.Sprintf("%.2fms", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
			messages = append(messages, "|")
			// sql

			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						if t.IsZero() {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", "0000-00-00 00:00:00"))
						} else {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
						}
					} else if b, ok := value.([]byte); ok {
						if str := string(b); isPrintable(str) {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						switch value.(type) {
						case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
							formattedValues = append(formattedValues, fmt.Sprintf("%v", value))
						default:
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						}
					}
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			}

			// differentiate between $n placeholders or else treat like ?
			if regexp.MustCompile(`\$\d+`).MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range regexp.MustCompile(`\?`).Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength {
						sql += formattedValues[index]
					}
				}
			}

			messages = append(messages, sql)
			messages = append(messages, "|")
			messages = append(messages, fmt.Sprintf(" %v", strconv.FormatInt(values[5].(int64), 10)+" rows affected or returned"))
		}
		messages = append(messages, "} ")
	}

	return
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
