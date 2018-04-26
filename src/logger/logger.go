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

var logMap map[string]*Log = make(map[string]*Log)

type Log struct {
	Level uint8
	logger *log.Logger
	fileName string
	file *os.File
}

func NewLog(name string, level uint8) (*Log) {
	file, err := os.OpenFile(string("../log/" + name + ".log"), os.O_APPEND | os.O_CREATE, os.ModeAppend)
	if err != nil {
		fmt.Printf("open %s failed.\n", name)
		return nil
	}
	logger := log.New(file, "DEBUG", log.LstdFlags | log.Lshortfile)
	if logger == nil {
		fmt.Printf("new logger failed.")
		return nil
	}
	return &Log {
		Level: level,
		fileName: name,
		file: file,
		logger: logger,
	}
}

func (this *Log) SetLevel(lv uint8) {
	if lv < DEBUG || lv > FATAL {
		return
	}
	this.Level = lv
}

func (this *Log) WriteLog(level uint8, format string, v ... interface {}) {
	if level < DEBUG || level > FATAL {
		return
	}
	this.logger.SetPrefix(LevelStr[level])
	this.logger.Printf(format + "\r\n", v)
}

func GetLog(name string) (*Log) {
	if logMap[name] != nil {
		return logMap[name]
	}
	log := NewLog(name, DEBUG)
	logMap[name] = log
	return log
}

func DeleteLog(name string) {
	log := logMap[name]
	if log != nil {
		log.file.Close()
		delete(logMap, name)
	}
}