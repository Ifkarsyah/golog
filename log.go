package golog

import (
	"fmt"
	c "github.com/Ifkarsyah/golog/color"
	"log"
	"runtime"
)

func Success(msg string) {
	log.Println(commonLog("SUCCESS", c.Green, msg))
}

func Info(msg string) {
	log.Println(commonLog("INFO", c.Blue, msg))
}

func Warn(msg string) {
	log.Println(commonLog("INFO", c.Yellow, msg))
}

func Debug(err error) {
	log.Println(commonError("DEBUG", c.Red, err))
}

func Fatal(err error) {
	log.Fatal(commonError("FATAL", c.Purple, err))
}

func commonError(tag string, color string, err error) string {
	function, file, line, _ := runtime.Caller(2)
	msg := fmt.Sprintf("[%s][%s:%s:%d]: '%s'", tag, baseFile(file), runtime.FuncForPC(function).Name(), line, err)
	colorizedMsg := c.Ize(color, msg)
	return colorizedMsg
}

func commonLog(tag string, color string, msg string) string {
	function, file, line, _ := runtime.Caller(2)
	msg = fmt.Sprintf("[%s][%s:%s:%d]: '%s'", tag, baseFile(file), runtime.FuncForPC(function).Name(), line, msg)
	colorizedMsg := c.Ize(color, msg)
	return colorizedMsg
}
