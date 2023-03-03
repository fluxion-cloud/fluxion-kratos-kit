package logging

import "github.com/go-kratos/kratos/v2/log"

type Adapter log.Logger
type Logger log.Helper

func NewLogger(adapter Adapter) *Logger {
	return (*Logger)(log.NewHelper(adapter))
}
