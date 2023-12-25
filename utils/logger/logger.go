package logger

import (
	"github.com/yunduan16/micro-service-go-component-log"
	"time"
)

const LOG_MAX_AGE = 120 * time.Minute //日志最大保留时间
func InitLog(FileName string) {
	logObj, setDefaultLoggerErr := log.New(
		log.TimestampFormat("2006-01-02 15:04:05.000000"),
		log.FileName(FileName),
		log.CallerDeep(3),
		log.MaxAge(LOG_MAX_AGE),
		log.RotationTime(LOG_MAX_AGE),
	)
	if setDefaultLoggerErr != nil {
		panic("LoadConfig setDefaultLoggerErr: " + setDefaultLoggerErr.Error())
		return
	}
	log.SetDefaultLogger(logObj)
}
