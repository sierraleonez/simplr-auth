package utils

import "errors"

func Error(message string) string {
	err := errors.New(message).Error()
	return err
}
