package msg

import (
	api "github.com/Lei2050/lei-net/api"
	pkt "github.com/Lei2050/lei-net/packet/v2"
	log "github.com/Lei2050/lei-utils/log"
)

type MsgServiceFunc func(api.TcpConnectioner, *Msg)

var _ pkt.PacketHandler = &MsgService{}

type MsgService struct {
	sFuncs map[uint32]MsgServiceFunc
}

func NewMsgService() *MsgService {
	return &MsgService{sFuncs: make(map[uint32]MsgServiceFunc)}
}

func (m *MsgService) RegisterService(cmd uint32, f MsgServiceFunc) {
	m.sFuncs[cmd] = f
}

func (m *MsgService) Has(cmd uint32) bool {
	_, e := m.sFuncs[cmd]
	return e
}

func (m *MsgService) Process(api.TcpConnectioner, *pkt.Packet) {}
func (m *MsgService) ProcessData(rmd *RawMsgData) {
	conn, msg := rmd.Conn, rmd.M
	if f, e := m.sFuncs[msg.H.Cmd]; e && f != nil {
		f(conn, msg)
	}
}

type OutMsgServiceFunc func(api.TcpConnectioner, *OutMsg)

type OutMsgService struct {
	sFuncs map[uint32]OutMsgServiceFunc
}

func NewOutMsgService() *OutMsgService {
	return &OutMsgService{sFuncs: make(map[uint32]OutMsgServiceFunc)}
}

func (m *OutMsgService) RegisterService(cmd uint32, f OutMsgServiceFunc) {
	m.sFuncs[cmd] = f
}

func (m *OutMsgService) Has(cmd uint32) bool {
	_, e := m.sFuncs[cmd]
	return e
}

func (m *OutMsgService) ProcessData(conn api.TcpConnectioner, msg *OutMsg) {
	if f, e := m.sFuncs[msg.H.Cmd]; e && f != nil {
		f(conn, msg)
	} else {
		log.Error("not find handler for cmd:%d", msg.H.Cmd)
	}
}
