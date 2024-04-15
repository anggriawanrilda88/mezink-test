package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type ErrorMessage struct {
	ErrStatus int
	ErrData   gin.H
}

// parseErrorMessage function splits the input string and populates the ErrorMessage struct
func ParseErrorMessage(err error) ErrorMessage {
	var result ErrorMessage
	result.ErrData = gin.H{}
	parts := strings.Split(err.Error(), "; ")

	for _, part := range parts {
		pair := strings.Split(part, ":")
		if len(pair) != 2 {
			continue
		}
		key := strings.TrimSpace(pair[0])
		value := strings.TrimSpace(pair[1])

		switch key {
		case "errStatus":
			fmt.Sscanf(value, "%d", &result.ErrStatus)
		case "errCode":
			result.ErrData["code"] = value
		case "message":
			result.ErrData["msg"] = value
		}
	}

	return result
}
