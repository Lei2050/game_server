package msg

const (
	MSG_HEAD_LEN    = 28
	UINT32_BYTE_LEN = 4
)

//内部协议的包头
type MsgHead struct {
	//Len uint32
	Cmd uint32
	Uid uint64
	Pa1 uint32 //自定义参数，需要的时候自己填
	Src uint32 //消息来自busid
	Dst uint32 //消息发至busid
	Hop uint32 //该消息当前经过了几跳，转发一次加一跳
}

//big endian，大端解码编码，高位数据位于数组的低索引位置。
func DecodeUint32(data []byte) uint32 {
	return (uint32(data[0]) << 24) | (uint32(data[1]) << 16) | (uint32(data[2]) << 8) | (uint32(data[3]))
}

func EncodeUint32(n uint32, data []byte) {
	data[0] = byte((n >> 24) & 0xFF)
	data[1] = byte((n >> 16) & 0xFF)
	data[2] = byte((n >> 8) & 0xFF)
	data[3] = byte(n & 0xFF)
}

func DecodeUint64(data []byte) uint64 {
	return (uint64(data[0]) << 56) | (uint64(data[1]) << 48) | (uint64(data[2]) << 40) | (uint64(data[3]) << 32) | (uint64(data[4]) << 24) | (uint64(data[5]) << 16) | (uint64(data[6]) << 8) | uint64(data[7])
}

func EncodeUint64(n uint64, data []byte) {
	data[0] = byte((n >> 56) & 0xFF)
	data[1] = byte((n >> 48) & 0xFF)
	data[2] = byte((n >> 40) & 0xFF)
	data[3] = byte((n >> 32) & 0xFF)
	data[4] = byte((n >> 24) & 0xFF)
	data[5] = byte((n >> 16) & 0xFF)
	data[6] = byte((n >> 8) & 0xFF)
	data[7] = byte(n & 0xFF)
}

func EncodeMsgHead(buf []byte, ph *MsgHead) bool {
	if len(buf) < MSG_HEAD_LEN {
		return false
	}
	//EncodeUint32(ph.Len, buf[0:4])
	//EncodeUint32(ph.Cmd, buf[4:8])
	//EncodeUint64(ph.Uid, buf[8:16])
	//EncodeUint32(ph.Sid, buf[16:])
	EncodeUint32(ph.Cmd, buf[0:4])
	EncodeUint64(ph.Uid, buf[4:12])
	EncodeUint32(ph.Pa1, buf[12:16])
	EncodeUint32(ph.Src, buf[16:20])
	EncodeUint32(ph.Dst, buf[20:24])
	EncodeUint32(ph.Hop, buf[24:28])
	return true
}

func DecodeMsgHead(buf []byte) (*MsgHead, bool) {
	if len(buf) < MSG_HEAD_LEN {
		return nil, false
	}
	ph := &MsgHead{
		//Len: DecodeUint32(buf[:4]),
		//Cmd: DecodeUint32(buf[4:8]),
		//Uid: DecodeUint64(buf[8:16]),
		//Sid: DecodeUint32(buf[16:]),
		Cmd: DecodeUint32(buf[0:4]),
		Uid: DecodeUint64(buf[4:12]),
		Pa1: DecodeUint32(buf[12:16]),
		Src: DecodeUint32(buf[16:20]),
		Dst: DecodeUint32(buf[20:24]),
		Hop: DecodeUint32(buf[24:28]),
	}
	return ph, true
}

const (
	OUT_MSG_HEAD_LEN = 4
)

//外部协议的包头
type OutMsgHead struct {
	//Len uint32
	Cmd uint32
}

func EncodeOutMsgHead(buf []byte, ph *OutMsgHead) bool {
	if len(buf) < OUT_MSG_HEAD_LEN {
		return false
	}
	//EncodeUint32(ph.Len, buf[0:4])
	//EncodeUint32(ph.Cmd, buf[4:8])
	//EncodeUint64(ph.Uid, buf[8:16])
	//EncodeUint32(ph.Sid, buf[16:])
	EncodeUint32(ph.Cmd, buf[0:4])
	return true
}

func DecodeOutMsgHead(buf []byte) (*OutMsgHead, bool) {
	if len(buf) < OUT_MSG_HEAD_LEN {
		return nil, false
	}
	ph := &OutMsgHead{
		//Len: DecodeUint32(buf[:4]),
		//Cmd: DecodeUint32(buf[4:8]),
		//Uid: DecodeUint64(buf[8:16]),
		//Sid: DecodeUint32(buf[16:]),
		Cmd: DecodeUint32(buf[0:4]),
	}
	return ph, true
}
