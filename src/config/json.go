package config

import (
	"encoding/json"
	"os"
	"fmt"
)

type ServerConfig struct {
	ServerId string `json: "ServerId"`
	LogPath string `json: "LogPath"`
}

func FromJsonStr(jsonStr string) (interface {}) {
	var conf ServerConfig
	err := json.Unmarshal([]byte(jsonStr), &conf)
	if err != nil {
		fmt.Printf("Unmarshal jsonData failed, err is %v\n", err)
		return nil
	}
	return &conf
}

func FromJsonFile(filePath string) (interface {}) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file %s for read failed, err is %v\n", filePath, err)
		return nil
	}
	jsonData := make([]byte, 1024)
	n, err2 := file.Read(jsonData)
	if err2 != nil {
		fmt.Printf("read file %s for read failed, err is %v\n", filePath, err2)
		return nil
	}

	return FromJsonStr(string(jsonData[:n]))
}