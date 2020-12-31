package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// [info ]颜色为蓝色，[error]颜色为红色
// 使用log.Lshortfile支持显示文件名和行号
var (
	errorlog = log.New(os.Stdout, "\033[31m[error]\003[0m", log.LstdFlags|log.Lshortfile)
	infolog  = log.New(os.Stdout, "\033[34m[info ]\033[0m", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorlog, infolog}
	mu       sync.Mutex
)

// 暴露Error，Errorf，Info，Infof 4个方法
var (
	Error  = errorlog.Println
	Errorf = errorlog.Printf
	Info   = infolog.Print
	Infof  = infolog.Printf
)


// 日志等级
const (
	InfoLevel = iota
	ErrorLevel
	Fatal
)

// 通过控制output来控制日志是否打印
// 如果设置为ErrorLevel，infoLog的输出会被定向到ioutil.Discard，即不打印日志
func SetLevel(level int)  {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level{
		errorlog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		infolog.SetOutput(ioutil.Discard)
	}
}
