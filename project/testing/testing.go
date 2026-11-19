package testing

import "errors"

var errorLogger = errors.New("Входящая строка в logger.mod пустая")

func CheckLogger(text string, level func(text string) (string, error)) (string, error) {
	output, err := level(text)
	if err != nil {
		return "", errorLogger
	} else {
		return output, nil
	}
}
