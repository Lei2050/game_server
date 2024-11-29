package msg

import (
	pkt "github.com/Lei2050/lei-net/packet/v2"
)

func EncodeMsgHeadToPacket(ph *MsgHead, packet *pkt.Packet) {
	packet.WriteUint32(ph.Cmd)
	packet.WriteUint64(ph.Uid)
	packet.WriteUint32(ph.Pa1)
	packet.WriteUint32(ph.Src)
	packet.WriteUint32(ph.Dst)
	packet.WriteUint32(ph.Hop)
}

func DecodeMsgHeadFromPacket(ph *MsgHead, packet *pkt.Packet) {
	ph.Cmd = packet.ReadUint32()
	ph.Uid = packet.ReadUint64()
	ph.Pa1 = packet.ReadUint32()
	ph.Src = packet.ReadUint32()
	ph.Dst = packet.ReadUint32()
	ph.Hop = packet.ReadUint32()
}
