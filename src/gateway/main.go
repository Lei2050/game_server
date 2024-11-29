package main

import (
	"flag"
	"fmt"
	"runtime"

	bus "github.com/Lei2050/lei-bus/v2"
	tcp "github.com/Lei2050/lei-net/tcp/v2"
	leiconfig "github.com/Lei2050/lei-utils/config"
	log "github.com/Lei2050/lei-utils/log"

	"github.com/Lei2050/game_server/aconst"
)

var G_Bus *bus.Bus
var G_ClientMgr *ClientMgr

// var G_InMsgHandle = new(InMsgHandle)
var G_Config = new(GatewayConfigXml)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	configFile := flag.String("config", "./config/gateway_config.xml", "")
	flag.Parse()

	if err := leiconfig.LoadXmlConfig(*configFile, G_Config); err != nil {
		panic(fmt.Sprintf("load config %v failed:%v", *configFile, err))
	}

	if len(G_Config.ServerId) <= 0 || G_Config.ServerId[0] <= 0 {
		panic("not config server_ids")
	}

	//log.Init(&G_Config.Log)

	busIds := make([]bus.BusId, len(G_Config.ServerId))
	for i := range busIds {
		busIds[i] = bus.GetBusId(
			1,
			aconst.AppTypeGateway,
			uint32(G_Config.ServerId[i]),
			G_Config.ProgInst,
		)
	}

	G_Bus = bus.NewBusServer(busIds, &G_Config.Bus, &InnerPacketHandle{})

	serConfig := &G_Config.ClientSer
	ser, err := tcp.NewServer(&OutMsgHandle{},
		tcp.Address(serConfig.ListenAddr),
		tcp.MaxConn(serConfig.MaxConn),
		tcp.IdleTime(serConfig.IdleTime),
	)
	if err != nil {
		log.Error("NewServerTcp failed !\n")
		return
	}
	G_ClientMgr = NewClientMgr(ser)

	G_Bus.Start()
	go inSync()

	log.Info("SERVER STARTED !\n")

	G_ClientMgr.TcpServer.Start()
}
