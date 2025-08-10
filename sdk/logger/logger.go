// Package logger provides a simple and minimalistic logging
package logger

import (
	"log"
)

func Log(msgs ...any) {
	log.Println(msgs...)
}

func Success(msgs ...any) {
	Log(append([]any{" ✔  |"}, msgs...)...)
}

func Error(msgs ...any) {
	Log(append([]any{" 🗙  |"}, msgs...)...)
}

func Errorf(err error) {
	Log(append([]any{" 🗙  |"}, err.Error())...)
}

func Warning(msgs ...any) {
	Log(append([]any{" ⚠   |"}, msgs...)...)
}
