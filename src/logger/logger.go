package logger

import (
	"log"
	"os"
	"fmt"
)

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var LevelStr [5]string = [5]string {
	"DEBUG ", "INFO ", "WARNING ", "ERROR ", "FATAL ",
}

var g_logger *Log

func WRITE_DEBUG(format string, v ... interface {}) {
	if g_logger == nil {
		return
	}
	g_logger.setLevel(DEBUG)
	g_logger.writeLog(format, v ...)
}

func WRITE_INFO(format string, v ... interface {}) {
	if g_logger == nil {
		return
	}
	g_logger.setLevel(INFO)
	g_logger.writeLog(format, v ...)
}

func WRITE_WARNING(format string, v ... interface {}) {
	if g_logger == nil {
		return
	}
	g_logger.setLevel(WARNING)
	g_logger.writeLog(format, v ...)
}

func WRITE_ERROR(format string, v ... interface {}) {
	if g_logger == nil {
		return
	}
	g_logger.setLevel(ERROR)
	g_logger.writeLog(format, v ...)
}

func WRITE_FATAL(format string, v ... interface {}) {
	if g_logger == nil {
		return
	}
	g_logger.setLevel(FATAL)
	g_logger.writeLog(format, v ...)
}

type Log struct {
	minLevel uint8
	Level uint8
	logger *log.Logger
	fileName string
	file *os.File
}

func NewLog(name string, level uint8, daemon bool) {
	var file *os.File
	if daemon {
		var err error
		file, err = os.OpenFile(string("../log/" + name + ".log"), os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf("open %s failed.\n", name)
			return
		}
	} else {
		file = os.Stderr
	}
	logger := log.New(file, "DEBUG", log.LstdFlags)
	if logger == nil {
		fmt.Printf("new logger failed.")
		return
	}
	g_logger = &Log {
		minLevel: level,
		Level: DEBUG,
		fileName: name,
		file: file,
		logger: logger,
	}
}

func (this *Log) setLevel(lv uint8) {
	if lv < DEBUG || lv > FATAL {
		return
	}
	this.Level = lv
}

func (this *Log) writeLog(format string, v ... interface {}) {
	if this.Level < this.minLevel {
		return
	}
	this.logger.SetPrefix(LevelStr[this.Level])
	this.logger.Printf(format + "\n", v ...)
}