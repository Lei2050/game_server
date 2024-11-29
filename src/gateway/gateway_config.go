package main

import (
	bus "github.com/Lei2050/lei-bus/v2"
)

type xmlPprof struct {
	State string `xml:"state,attr"`
	Port  string `xml:"port"`
}

type xmlClientSer struct {
	ListenAddr string
	MaxConn    int
	IdleTime   int //ms
	InBufSize  int
	OutBufSize int
}

type GatewayConfigXml struct {
	Platform uint32 `xml:"platform"`
	//AppType  uint32   `xml:"app_type"`
	ServerId []uint64 `xml:"server_id"`
	ProgInst uint32   `xml:"prog_inst"`

	//Log log.Config `xml:"log"`

	Bus       bus.BusConfig `xml:"bus"`
	ClientSer xmlClientSer  `xml:"client_ser"`

	ClientSerAddr string `xml:"client_ser_addr"`

	Pprof xmlPprof `xml:"pprof"`
	Env   string   `xml:"env"`
}
