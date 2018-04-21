package common

import (
	// "sync"
)

var moduleList []*Module = make([]*Module, 10)

type Module struct {
	id uint8
	name string
	cmdMin uint8
	cmdMax uint8
	in chan<- []byte
	out <-chan []byte
}

func GetModule(id uint8) (*Module){
	for _, v := range moduleList {
		if v.id == id {
			return v
		} 
	}

	return nil
}

func PushModule(m *Module) {
	moduleList = append(moduleList, m)
}