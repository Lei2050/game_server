package msg

import (
	"github.com/Lei2050/game_server/protocol/in"
	"github.com/Lei2050/game_server/protocol/out"

	"google.golang.org/protobuf/proto"
)

type Msg struct {
	H MsgHead
	D interface{} //可能proto.Message，也可能是[]byte
}

func (m *Msg) Msg() proto.Message {
	pm, _ := m.D.(proto.Message)
	return pm
}

func (m *Msg) Data() []byte {
	d, _ := m.D.([]byte)
	return d
}

func (m *Msg) OutMessage() *out.CSMessage {
	pm, _ := m.D.(*out.CSMessage)
	return pm
}

func (m *Msg) InMessage() *in.InMessage {
	pm, _ := m.D.(*in.InMessage)
	return pm
}

func (m *Msg) PbMessage() proto.Message {
	pm, _ := m.D.(proto.Message)
	return pm
}

type OutMsg struct {
	H OutMsgHead
	D interface{} //可能proto.Message，也可能是[]byte
}

func (o *OutMsg) Msg() proto.Message {
	pm, _ := o.D.(proto.Message)
	return pm
}

func (o *OutMsg) Data() []byte {
	d, _ := o.D.([]byte)
	return d
}
