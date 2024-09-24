package panicutils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"net/http"
	"runtime"
	"strings"
)

type HttpException struct {
	StatusCode int                    `json:"status_code,omitempty"`
	ErrorMsg   error                  `json:"error_msg,omitempty"`
	Data       map[string]interface{} `json:"data"`
}

func PanicRecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer handlePanic(c)
		c.Next()
	}
}

func handlePanic(c *gin.Context) {
	StatusCode := http.StatusUnprocessableEntity
	var Data map[string]interface{}

	if err := recover(); err != nil {

		defer printStrace(3) //to get the correct stack depth as runtime stack increase to another func

		var errStr string
		switch v := err.(type) {

		case *HttpException: //Intentionally/Manually Pass Status Code , Error Message , Data
			errStr = v.ErrorMsg.Error()
			StatusCode = v.StatusCode
			Data = v.Data

		case string: //Got unhandled error / not error type
			errStr = v

		case error: //Got Actual Error Data Type
			errStr = v.Error()
			StatusCode = getStatusCode(v) //Actual error checker possible from only direct error . Not from string or v.Error()

		default:
			errStr = fmt.Sprintf("recovered from: %v", v)
		}

		c.JSON(StatusCode, gin.H{
			"status_code": StatusCode,
			"message":     errStr,
			"data":        Data,
		})

	}
}

func getStatusCode(err error) int {
	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, mongo.ErrNoDocuments) {
		return http.StatusNotFound
	}
	return http.StatusUnprocessableEntity
}

func printStrace(skip int) {
	// Capture file, function, and line number
	pc, file, line, ok := runtime.Caller(skip) // This should now capture the correct function
	if ok {
		fn := runtime.FuncForPC(pc)
		funcName := fn.Name()

		// Formatting the file path to get the last portion
		shortFile := file
		if idx := strings.LastIndex(file, "/"); idx != -1 {
			shortFile = file[idx+1:]
		}

		// ANSI color codes
		const (
			ColorRed    = "\033[31m"
			ColorYellow = "\033[33m"
			ColorReset  = "\033[0m" // Reset to default color
		)

		// Print the actual function name and line number
		fmt.Printf(ColorRed+"Panic occurred --> %s "+ColorYellow+"(%s:%d)"+ColorReset+"\n", funcName, shortFile, line)
	}
}
