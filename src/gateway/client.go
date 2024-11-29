package main

import (
	api "github.com/Lei2050/lei-net/api"

	"github.com/Lei2050/game_server/msg"
	"github.com/Lei2050/game_server/protocol/in"
)

const (
	CLIENT_STATE_DEFAULT  = iota //默认
	CLIENT_STATE_LOGINING        //正在登陆
	CLIENT_STATE_LOGINED         //已登陆
	CLIENT_STATE_CREATING        //创角中
	CLIENT_STATE_OFFLINE         //已下线
)

type Client struct {
	State int //const CLIENT_STATE
	Conn  api.TcpConnectioner
}

func (c *Client) NotifyGameClose() {
	if c.State != CLIENT_STATE_DEFAULT {
		//如果不是默认状态，说明Game中已经保存了该链接信息，则要通知清理。
		//否则，还仅仅只是建立了链接，啥事情都没开始干。
		pm := &in.InMessage{
			OneOf: &in.InMessage_GatewayConnClose{
				GatewayConnClose: &in.GatewayConnClose{},
			},
		}
		SendToGame(msg.MsgHead{Cmd: in.CMD_GATEWAY_CONN_CLOSE, Pa1: uint32(c.Conn.Id())}, pm)
	}
}
