package main

import (
	"sync"

	bus "github.com/Lei2050/lei-bus/v2"
	netapi "github.com/Lei2050/lei-net/api"
	log "github.com/Lei2050/lei-utils/log"

	"github.com/Lei2050/game_server/aconst"
	"github.com/Lei2050/game_server/msg"
)

type ClientMgr struct {
	TcpServer netapi.TcpServerer

	sync.RWMutex
	//<uid, lei.TcpConnectioner>
	//Clients map[uint64]lei.TcpConnectioner
	//<uid, *Client>
	Clients map[uint64]*Client
	//<conn_id, *Client>
	UnauthClients map[int]*Client
	//<conn_id, uid>
	ConnIdUid map[int]uint64
}

func NewClientMgr(ser netapi.TcpServerer) *ClientMgr {
	return &ClientMgr{
		TcpServer:     ser,
		Clients:       make(map[uint64]*Client),
		ConnIdUid:     make(map[int]uint64),
		UnauthClients: make(map[int]*Client),
	}
}

func (cm *ClientMgr) GetUidByConnId(id int) uint64 {
	cm.RLock()
	defer cm.RUnlock()
	return cm.ConnIdUid[id]
}

func (cm *ClientMgr) GetUidByConn(conn netapi.TcpConnectioner) uint64 {
	cm.RLock()
	defer cm.RUnlock()
	return cm.GetUidByConnId(conn.Id())
}

func (cm *ClientMgr) AddUnauthClient(conn netapi.TcpConnectioner) *Client {
	client := &Client{State: CLIENT_STATE_DEFAULT, Conn: conn}
	cm.Lock()
	cm.UnauthClients[conn.Id()] = client
	cm.Unlock()
	return client
}

func (cm *ClientMgr) GetUnauthClient(id int) *Client {
	cm.RLock()
	defer cm.RUnlock()
	return cm.UnauthClients[id]
}

func (cm *ClientMgr) GetClientByConnId(id int) *Client {
	cm.RLock()
	defer cm.RUnlock()
	uid := cm.ConnIdUid[id]
	if uid <= 0 {
		return nil
	}
	return cm.Clients[uid]
}

func (cm *ClientMgr) ClientLogin(uid uint64, id int) *Client {
	log.Debug("ClientLogin uid:%d, connId:%d", uid, id)

	cm.Lock()
	defer cm.Unlock()
	client, e := cm.UnauthClients[id]
	if !e {
		return nil
	}
	delete(cm.UnauthClients, id)

	client.State = CLIENT_STATE_LOGINED
	cm.Clients[uid] = client
	cm.ConnIdUid[id] = uid
	return client
}

func (cm *ClientMgr) SendMsgToClient(uid uint64, msg interface{}) {
	cm.RLock()
	defer cm.RUnlock()
	if client, exit := cm.Clients[uid]; exit && client != nil {
		client.Conn.Write(msg)
	}
}

func (cm *ClientMgr) SendMsgToCliConn(id int, msg interface{}) {
	cm.RLock()
	defer cm.RUnlock()
	if uid, e := cm.ConnIdUid[id]; e {
		if client := cm.Clients[uid]; client != nil {
			client.Conn.Write(msg)
		}
	} else {
		if client := cm.UnauthClients[id]; client != nil {
			client.Conn.Write(msg)
		}
	}
}

func (cm *ClientMgr) Kickoff(id int) {
	client, exist := cm.UnauthClients[id]
	if exist {
		client.Conn.Close()
		delete(cm.UnauthClients, id)
	}

	uid, exist := cm.ConnIdUid[id]
	if exist {
		if client, exist := cm.Clients[uid]; exist {
			client.Conn.Close()
			delete(cm.Clients, uid)
		}
		delete(cm.ConnIdUid, id)
	}
}

func (cm *ClientMgr) onConnectionClose(id int) {
	cm.RLock()
	uid, exist := cm.ConnIdUid[id]
	cm.RUnlock()

	if exist { //在线状态
		cm.RLock()
		client := cm.Clients[uid]
		cm.RUnlock()
		if client != nil {
			client.NotifyGameClose()
		}

		cm.Lock()
		delete(cm.ConnIdUid, id)
		delete(cm.Clients, uid)
		cm.Unlock()
		return
	}

	cm.Lock()
	client := cm.UnauthClients[id]
	delete(cm.UnauthClients, id)
	cm.Unlock()
	client.NotifyGameClose()
}

func SendToGame(h msg.MsgHead, oneofmsg interface{}) {
	//如果配置多个serverId，说明是多game的场景的。只要发到对应的Game进程就行，Game进程会将消息转发到对应的game上。
	//所以这里在配置中随便取一个ServerId就行。
	serverId := uint32(G_Config.ServerId[0])
	h.Src = bus.GetBusId(1, aconst.AppTypeGateway, serverId, 1).Uint32()
	h.Dst = bus.GetBusId(1, aconst.AppTypeGame, serverId, 1).Uint32()
	G_Bus.SendMsgToBusId(
		bus.BusId(h.Dst),
		&msg.Msg{H: h, D: oneofmsg},
	)
}
