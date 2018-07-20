package config

import (
	"encoding/json"
	"os"
	"io/ioutil"
	"fmt"
)

type GatewayConfig struct {
	ServerId string `json: "ServerId"`
	LogPath string `json: "LogPath"`
	ListenAddress string `json: "ListenAddress"`
	Protocol string `json: "Protocol"`
}

type GameConfig struct {
	ServerId string `json: "ServerId"`
	LogPath string `json: "LogPath"`
	ListenAddress string `json: "ListenAddress"`
	Protocol string `json: "Protocol"`
}

var GatewayConf GatewayConfig
var GameConf GameConfig

func LoadGatewayConfig(filePath string) {
	if FromJsonFile(filePath, &GatewayConf) != nil {
		os.Exit(1)
	}
}

func LoadGameConfig(filePath string) {
	if FromJsonFile(filePath, &GameConf) != nil {
		os.Exit(1)
	}
}

func FromJsonStr(jsonStr []byte, v interface{}) error {
	err := json.Unmarshal(jsonStr, v)
	if err != nil {
		fmt.Printf("Unmarshal jsonData failed, err is %v\n", err)
		return err
	}
	return nil
}

func FromJsonFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file %s for read failed, err is %v\n", filePath, err)
		return err
	}
	defer file.Close()
	jsonData, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		fmt.Printf("read file %s for read failed, err is %v\n", filePath, err2)
		return err2
	}

	return FromJsonStr(jsonData, v)
}