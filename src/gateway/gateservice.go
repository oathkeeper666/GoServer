package main

import (
	"network"
)

type GateService struct {
	connMgr *network.ConnMgr
}

func NewGateService() (*GateService) {
	return nil
}

func (this *GateService) ConnectBackend() {

}