package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code <= 299:
		return green
	case code >= 300 && code <= 399:
		return white
	case code >= 400 && code <= 499:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch {
	case method == "GET":
		return blue
	case method == "POST":
		return cyan
	case method == "PUT":
		return yellow
	case method == "DELETE":
		return red
	case method == "PATCH":
		return green
	case method == "HEAD":
		return magenta
	case method == "OPTIONS":
		return white
	default:
		return reset
	}
}

//Logger2 sadsad
func Logger2() gin.HandlerFunc {
	//stdlogger := log.New(os.Stdout, "", 0)
	//errlogger := log.New(os.Stderr, "", 0)

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// save the IP of the requester
		requester := c.Request.Header.Get("X-Real-IP")
		// if the requester-header is empty, check the forwarded-header
		if len(requester) == 0 {
			requester = c.Request.Header.Get("X-Forwarded-For")
		}
		// if the requester is still empty, use the hard-coded address from the socket
		if len(requester) == 0 {
			requester = c.Request.RemoteAddr
		}

		var color string
		code := c.Writer.Status()
		switch {
		case code >= 200 && code <= 299:
			color = green
		case code >= 300 && code <= 399:
			color = white
		case code >= 400 && code <= 499:
			color = yellow
		default:
			color = red
		}
		end := time.Now()
		latency := end.Sub(start)
		glog.Infof("[GIN] %v |%s %3d %s| %12v | %s %4s %s\n%s",
			end.Format("2006/01/02 - 15:04:05"),
			color, code, reset,
			latency,
			requester,
			c.Request.Method, c.Request.URL.Path,
			c.Errors.String(),
		)
	}
}

//Logger sadas
func Logger() gin.HandlerFunc {
	//stdlogger := log.New(os.Stdout, "", 0)
	//errlogger := log.New(os.Stderr, "", 0)

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		//	clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		statusColor := colorForStatus(statusCode)
		methodColor := colorForMethod(method)

		glog.Infof("[GIN] %v |%s %3d %s| %12v | %s |%s  %s %-7s %s\n%s",
			end.Format("2006/01/02 - 15:04:05"),
			statusColor, statusCode, reset,
			latency,
			methodColor, reset, method,
			c.Request.URL.Path,
			c.Errors.String(),
		)
	}
}
