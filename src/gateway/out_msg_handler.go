package main

import (
	"io"

	api "github.com/Lei2050/lei-net/api"
	tcp "github.com/Lei2050/lei-net/tcp/v2"
	leiio "github.com/Lei2050/lei-utils/io"
	log "github.com/Lei2050/lei-utils/log"

	"github.com/Lei2050/game_server/msg"
	"github.com/Lei2050/game_server/protocol/out"
)

var _ tcp.Protocoler = &OutMsgHandle{}

// 对客户端的通信，没有bus系统，包头信息和内部协议不一样，所以分别处理
type OutMsgHandle struct {
	*msg.OutMsgHandle
}

func (omh *OutMsgHandle) Read(conn api.TcpConnectioner, reader io.Reader) error {
	var uint32Buffer [msg.UINT32_BYTE_LEN]byte
	var err error

	// 先读长度
	err = leiio.ReadFull(reader, uint32Buffer[:])
	if err != nil {
		return err
	}

	payloadSize := packetEndian.Uint32(uint32Buffer[:])
	if payloadSize > MaxPayloadLength {
		return ErrPayloadTooLarge
	}

	packet := NewPacket()
	payload := packet.extendPayload(int(payloadSize))
	err = leiio.ReadFull(reader, payload)
	if err != nil {
		return err
	}

	packet.SetPayloadLen(payloadSize)

	b.packerHandler.Process(conn, packet)

	return nil
}

func (omh *OutMsgHandle) Write(writer io.Writer, raw any) error {}
func (omh *OutMsgHandle) OnConnect(conn api.TcpConnectioner)    {}
func (omh *OutMsgHandle) HeartBeatMsg() any                     {}
func (omh *OutMsgHandle) SetOption(opt tcp.Options)             {}

func (omh *OutMsgHandle) UnpackMsg(conn api.TcpConnectioner, data []byte) (int, error) {
	unpackMsg, pn, err := omh.OutMsgHandle.UnpackMsg(conn, data)
	if err != nil {
		log.Error("UnpackMsg failed:%v", err)
		return pn, err
	}
	omh.Process(conn, unpackMsg.(*msg.OutMsg))
	return pn, nil
}

func (m *OutMsgHandle) Process(conn api.TcpConnectioner, outMsg *msg.OutMsg) {
	log.Debug("id:%d, Process H:%+v, D:%+v", conn.Id(), outMsg.H, outMsg.D)
	switch outMsg.H.Cmd {
	case out.CMD_CS_LOGIN: //请求登录
		m.onCSUserLogin(conn, outMsg)
		return
	case out.CMD_CS_CREATE: //请求创角
		m.onCSUserCreate(conn, outMsg)
		return
	case out.CMD_CS_GM: //客户端输入的gm命令
		m.onCSGm(conn, outMsg)
		return
	}
	//从客户端收到协议，直接转发到game上，这里后续可能需要在包头加入源bus_id等信息。
	//busId := leibus.GetBusId(1, G_Config.AppType, uint32(G_Config.ServerId), G_Config.ProgInst)

	if client := G_ClientMgr.GetClientByConnId(conn.Id()); client == nil || client.State != CLIENT_STATE_LOGINED {
		return
	}

	uid := G_ClientMgr.GetUidByConn(conn)
	SendToGame(msg.MsgHead{
		Cmd: outMsg.H.Cmd,
		Uid: uid,
		Pa1: uint32(conn.Id()),
	}, outMsg.D)
}

func (m *OutMsgHandle) onConnectionClose(id int) {
	G_ClientMgr.onConnectionClose(id)
}

func (m *OutMsgHandle) OnConnect(conn api.TcpConnectioner) {
	conn.RegisterCloseCb(func(conn api.TcpConnectioner) func() {
		if conn == nil {
			return nil
		}
		id := conn.Id()
		return func() {
			m.onConnectionClose(id)
		}
	}(conn))
	G_ClientMgr.AddUnauthClient(conn)
}

func (*OutMsgHandle) OnClose() {}

func (*OutMsgHandle) HeartBeatMsg() interface{} { return nil }
