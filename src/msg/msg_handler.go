package msg

import (
	"errors"
	"fmt"

	api "github.com/Lei2050/lei-net/api"
	pkt "github.com/Lei2050/lei-net/packet/v2"
	tcp "github.com/Lei2050/lei-net/tcp/v2"
	log "github.com/Lei2050/lei-utils/log"

	"google.golang.org/protobuf/proto"
)

type RawMsgData struct {
	Conn api.TcpConnectioner
	M    Msg
}

func (r *RawMsgData) Cmd() uint32 {
	return r.M.H.Cmd
}

func (r *RawMsgData) Uid() uint64 {
	return r.M.H.Uid
}

func (r *RawMsgData) Msg() Msg {
	return r.M
}

var _ pkt.PacketHandler = &InnerPacketHandle{}

// 服务器内部协议
type InnerPacketHandle struct {
	opt tcp.Options
}

func (m *InnerPacketHandle) Process(api.TcpConnectioner, *pkt.Packet) {

}

func UnpackRawMsgData(conn api.TcpConnectioner, packet *pkt.Packet) RawMsgData {
	var data RawMsgData
	DecodeMsgHeadFromPacket(&data.M.H, packet)
	// TODO 这里需要池化
	bytes := packet.ReadVarBytesI()
	dst := make([]byte, len(bytes))
	copy(dst, bytes)
	data.M.D = dst
	data.Conn = conn
	return data
}

func (mh *InnerPacketHandle) SetOption(opt tcp.Options) {
	mh.opt = opt
}

// 服务器内部协议
type OutMsgHandle struct {
	opt tcp.Options
}

func (omh *OutMsgHandle) UnpackMsg(conn api.TcpConnectioner, data []byte) (interface{}, int, error) {
	if len(data) < UINT32_BYTE_LEN { //数据不够等待读取更多数据
		return nil, 0, nil
	}

	need := int(DecodeUint32(data[:UINT32_BYTE_LEN]))

	if need >= omh.opt.ReadMaxSize {
		return nil, 0, fmt.Errorf("reached Maximum number:%d of read bytes. need:%d", omh.opt.ReadMaxSize, need)
	}

	if need > len(data) { //数据不够等待读取更多数据
		return nil, 0, nil
	}

	h, ok := DecodeOutMsgHead(data[UINT32_BYTE_LEN:need])
	if !ok {
		return nil, 0, errors.New("DecodeMsgHead failed")
	}
	d := make([]byte, len(data)-OUT_MSG_HEAD_LEN)
	copy(d, data[UINT32_BYTE_LEN+OUT_MSG_HEAD_LEN:need])
	return &OutMsg{H: h, D: d}, need, nil
}

func (*OutMsgHandle) PackMsg(conn api.TcpConnectioner, msg interface{}) (data []byte, err error) {
	data, err = nil, nil
	switch realMsg := msg.(type) {
	case []byte:
		data = realMsg
		return
	case *OutMsg:
		var wdata []byte
		switch raw := realMsg.D.(type) {
		case []byte:
			wdata = raw
		case proto.Message:
			d, e := proto.Marshal(raw)
			if e != nil {
				return nil, e
			}
			wdata = d
		default:
			return
		}
		length := len(wdata)
		data = make([]byte, UINT32_BYTE_LEN+OUT_MSG_HEAD_LEN+length)
		EncodeUint32(UINT32_BYTE_LEN+uint32(OUT_MSG_HEAD_LEN+length), data)
		EncodeOutMsgHead(data[UINT32_BYTE_LEN:], realMsg.H)
		copy(data[UINT32_BYTE_LEN+OUT_MSG_HEAD_LEN:], wdata)
		return
	}
	log.Error("Marshal failed, unknow type.")
	return nil, errors.New("unknow type")
}

func (mh *OutMsgHandle) SetOption(opt tcp.Options) {
	mh.opt = opt
}

func IsOutMsg(cmd uint32) bool {
	return 100000 < cmd && cmd < 300000
}

func IsInMsg(cmd uint32) bool {
	return 300000 < cmd && cmd <= 400000
}

func IsBottomCmd(cmd uint32) bool {
	return 1000 < cmd && cmd < 2000
}
