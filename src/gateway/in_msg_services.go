package main

import (
	"github.com/Lei2050/game_server/msg"
	"github.com/Lei2050/game_server/protocol/in"
	"github.com/Lei2050/game_server/protocol/out"

	api "github.com/Lei2050/lei-net/api"
	pkt "github.com/Lei2050/lei-net/packet/v2"
	tcp "github.com/Lei2050/lei-net/tcp/v2"
	log "github.com/Lei2050/lei-utils/log"

	"google.golang.org/protobuf/proto"
)

func onSCUserLogin(m *msg.Msg) bool {
	//pm := &out.SCMessage{}
	sc := &out.SC_Login{}
	if err := proto.Unmarshal(m.Data(), sc); err != nil {
		log.Error("onSCUserLogin unmarshal failed h:%+v, err:%+v", m.H, err)
		return false
	}

	connId := int(m.H.Pa1) //socket id
	log.Debug("onSCUserLogin:%+v", sc)

	//根据服务器登陆结果，设置玩家连接在网关的相关状态
	client := G_ClientMgr.GetUnauthClient(connId)
	if client == nil {
		log.Error("onSCUserLogin not find unauth client:%d", connId)
		return false
	}

	if client.State != CLIENT_STATE_LOGINING {
		log.Error("onSCUserLogin wrong state:%d", client.State)
		return false
	}

	switch sc.GetResult() {
	case out.LoginResult_LOGIN_SUCCESS:
		client.State = CLIENT_STATE_LOGINED
		G_ClientMgr.ClientLogin(m.H.Uid, connId)
	case out.LoginResult_NO_ROLE:
		client.State = CLIENT_STATE_CREATING
	}

	return true
}

func onSCUserCreate(m *msg.Msg) bool {
	//pm := &out.SCMessage{}
	sc := &out.SC_Create{}
	if err := proto.Unmarshal(m.Data(), sc); err != nil {
		log.Error("onSCUserCreate unmarshal failed h:%+v, err:%+v", m.H, err)
		return false
	}

	connId := int(m.H.Pa1) //socket id

	//根据服务器登陆结果，设置玩家连接在网关的相关状态
	client := G_ClientMgr.GetUnauthClient(connId)
	if client == nil {
		log.Error("onSCUserCreate not find unauth client:%d", connId)
		return false
	}

	if client.State != CLIENT_STATE_CREATING {
		log.Error("onSCUserCreate wrong state:%d", client.State)
		return false
	}

	switch sc.GetRet() {
	case 0:
		client.State = CLIENT_STATE_LOGINING
		//G_ClientMgr.ClientLogin(m.H.Uid, connId)
	}

	return true
}

func onReceiveKickoff(inMsg *msg.Msg) {
	//pm := &in.InMessage{}
	pm := &in.Kickoff{}
	if err := proto.Unmarshal(inMsg.Data(), pm); err != nil {
		log.Error("onReceiveKickoff unmarshal failed h:%+v, err:%+v", inMsg.H, err)
		return
	}
	log.Debug("onReceiveKickoff:%+v", pm)

	reason := pm.GetReason()
	id := int(pm.GetGwId())

	sc := &out.SC_Kickoff{
		Reason: reason.Enum(),
	}
	G_ClientMgr.SendMsgToCliConn(id, &msg.OutMsg{H: msg.OutMsgHead{Cmd: out.CMD_SC_KICKOFF}, D: sc})

	G_ClientMgr.Kickoff(id)
}

func inComingRmd(rmd msg.RawMsgData) {
	inMsg := rmd.Msg()
	log.Debug("gateway Process H:%+v, D:%+v", inMsg.H, inMsg.D)
	switch inMsg.H.Cmd {
	case in.CMD_KICKOFF: //踢下线
		onReceiveKickoff(&inMsg)
		return
	case out.CMD_SC_LOGIN: //登录返回
		if !onSCUserLogin(&inMsg) {
			return
		}
	case out.CMD_SC_CREATE: //创角返回
		if !onSCUserCreate(&inMsg) {
			return
		}
	}
	//从其他地方转发到客户端，转换包头
	//转发到客户端，至少需要知道conn_id或者玩家Uid
	if inMsg.H.Pa1 > 0 {
		G_ClientMgr.SendMsgToCliConn(int(inMsg.H.Pa1), &msg.OutMsg{H: msg.OutMsgHead{Cmd: inMsg.H.Cmd}, D: inMsg.D})
	} else if inMsg.H.Uid > 0 {
		G_ClientMgr.SendMsgToClient(inMsg.H.Uid, &msg.OutMsg{H: msg.OutMsgHead{Cmd: inMsg.H.Cmd}, D: inMsg.D})
	} else {
		log.Error("msg transfer to client, but not assign conn_id or uid.")
	}
}

func inSync() {
	defer func() {
		G_Bus.Exit()

		if err := recover(); err != nil {
			panic("Bug")
		}
	}()

	for raw := range G_Bus.C() {
		raw.Process()
	}
}

var _ pkt.PacketHandler = &InnerPacketHandle{}

// 服务器内部协议
type InnerPacketHandle struct {
	opt tcp.Options
}

// 会在inSync同步执行
func (m *InnerPacketHandle) Process(conn api.TcpConnectioner, packet *pkt.Packet) {
	rawMsgData := msg.UnpackRawMsgData(conn, packet)
	inComingRmd(rawMsgData)
}
