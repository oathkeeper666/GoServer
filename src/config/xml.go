package config

import (
	"encoding/xml"
	"os"
	"io/ioutil"
	"fmt"
)

type Tasks struct {
	T []Task `xml:"task"`
}

type Task struct {
	Id string `xml:"id,attr"`
	Task_type string `xml:"task_type,attr"`
	Condition string `xml:"condition,attr"`
	Time string `xml:"time,attr"`
	Reward string `xml:"reward,attr"`
}

func FromXmlStr(xmlStr []byte) (interface {}) {
	var reselt Tasks
	err := xml.Unmarshal(xmlStr, &reselt)
	if err != nil {
		fmt.Printf("Unmarshal xml string failed, err is %v\n", err)
		return nil
	}
	return &reselt
}

func FromXmlFile(filePath string) (interface {}) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file %s for read failed, err is %v\n", filePath, err)
		return nil
	}
	defer file.Close()
	xmlData, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		fmt.Printf("read file %s for read failed, err is %v\n", filePath, err2)
		return nil
	}

	return FromXmlStr(xmlData)
}