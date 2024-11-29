package main

import (
	"strings"

	api "github.com/Lei2050/lei-net/api"
	log "github.com/Lei2050/lei-utils/log"

	"github.com/Lei2050/game_server/msg"
	"github.com/Lei2050/game_server/protocol/out"

	"google.golang.org/protobuf/proto"
)

func (m *OutMsgHandle) onCSUserLogin(conn api.TcpConnectioner, outMsg *msg.OutMsg) bool {
	//设置相关状态
	client := G_ClientMgr.GetUnauthClient(conn.Id())
	if client == nil {
		log.Error("not find unauth client:%d", conn.Id())
		return false
	}

	if client.State != CLIENT_STATE_DEFAULT {
		log.Error("onCSUserLogin wrong state:%d", client.State)
		return false
	}

	client.State = CLIENT_STATE_LOGINING

	SendToGame(msg.MsgHead{
		Cmd: outMsg.H.Cmd,
		Pa1: uint32(conn.Id()),
	}, outMsg.D)

	return true
}

func (m *OutMsgHandle) onCSUserCreate(conn api.TcpConnectioner, outMsg *msg.OutMsg) bool {
	//设置相关状态
	client := G_ClientMgr.GetUnauthClient(conn.Id())
	if client == nil {
		log.Error("not find unauth client:%d", conn.Id())
		return false
	}

	if client.State != CLIENT_STATE_CREATING {
		log.Error("onCSUserLogin wrong state:%d", client.State)
		return false
	}

	SendToGame(msg.MsgHead{
		Cmd: outMsg.H.Cmd,
		Pa1: uint32(conn.Id()),
	}, outMsg.D)

	return true
}

func (m *OutMsgHandle) onCSGm(conn api.TcpConnectioner, outMsg *msg.OutMsg) bool {
	//设置相关状态
	client := G_ClientMgr.GetClientByConnId(conn.Id())
	if client == nil {
		log.Error("not find client:%d", conn.Id())
		return false
	}

	if client.State != CLIENT_STATE_LOGINED {
		log.Error("onCSGm wrong state:%d != CLIENT_STATE_LOGINED", client.State)
		return false
	}

	//cs := &out.CSMessage{}
	cs := &out.CS_Gm{}
	if err := proto.Unmarshal(outMsg.Data(), cs); err != nil {
		log.Error("onCSGm unmarshal failed h:%+v, err:%+v", outMsg.H, err)
		return false
	}

	//cmdLine := cs.GetCSGm().GetCmdLine()
	cmdLine := cs.GetCmdLine()

	strs := strings.Fields(cmdLine)
	if len(strs) <= 0 {
		log.Debug("empty cmd_line")
		return false
	}

	if strs[0] == "/gw" {
		//gateway中处理该gm命令，就不需要再转发了。
		//TODO...

		return false
	}

	uid := G_ClientMgr.GetUidByConn(conn)
	SendToGame(msg.MsgHead{
		Cmd: outMsg.H.Cmd,
		Uid: uid,
		Pa1: uint32(conn.Id()),
	}, outMsg.D)

	return true
}
