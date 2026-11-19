package logger

import (
	"errors"
	"fmt"
	"time"
)

var errNotFound = errors.New("Входящая строка в logger.mod пустая")

func Info(text string) (string, error) {
	if text != "" {
		output_text := fmt.Sprintf("%s | [INFO] - %s", time.Now().Format("02.01.2006 15:04:05"), text)
		return output_text, nil
	} else {
		return "", errNotFound
	}
}

func Warn(text string) (string, error) {
	if text != "" {
		output_text := fmt.Sprintf("%s | [INFO] - %s", time.Now().Format("02.01.2006 15:04:05"), text)
		return output_text, nil
	} else {
		return "", errNotFound
	}
}

func Error(text string) (string, error) {
	if text != "" {
		output_text := fmt.Sprintf("%s | [INFO] - %s", time.Now().Format("02.01.2006 15:04:05"), text)
		return output_text, nil
	} else {
		return "", errNotFound
	}
}
