//AUTO GENERATED. DO NOT EDIT !!!
package in

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

//type Cmd = int

const (
	CMD_KICKOFF                Cmd = 1001
	CMD_GATEWAY_CONN_CLOSE     Cmd = 1002
	CMD_RANK_INSERT            Cmd = 300001
	CMD_RANK_SYNC              Cmd = 300002
	CMD_GM_ADD_RESOURCE        Cmd = 300011
	CMD_GUILD_SYNC             Cmd = 300021
	CMD_GUILD_CREATE           Cmd = 300022
	CMD_GUILD_DISMISS          Cmd = 300023
	CMD_GUILD_APPLY_JOIN       Cmd = 300024
	CMD_GUILD_APPLY_CHECK      Cmd = 300025
	CMD_GUILD_LEAVE            Cmd = 300026
	CMD_GUILD_KICK             Cmd = 300027
	CMD_GUILD_MESSAGE          Cmd = 300028
	CMD_G2R_FRIEND_RECOMMEND   Cmd = 300031
	CMD_R2G_FRIEND_RECOMMEND   Cmd = 300032
	CMD_FRIEND_SEARCH_REQ      Cmd = 300033
	CMD_FRIEND_SEARCH_RES      Cmd = 300034
	CMD_FRIEND_CHECK           Cmd = 300035
	CMD_FRIEND_CHECK_RES       Cmd = 300036
	CMD_ROUTER_NODE_REPORT     Cmd = 300101
	CMD_ROUTER_CHOOSE_REQUEST  Cmd = 300102
	CMD_ROUTER_CHOOSE_RESPONSE Cmd = 300103
	CMD_SB_TAKE_ROOM           Cmd = 380001
	CMD_BS_TAKE_ROOM           Cmd = 380002
)

type InMessageFactory func() *InMessage

var g_InMessageFactories = map[Cmd]InMessageFactory{
	CMD_KICKOFF: func() *InMessage {
		return &InMessage{OneOf: &InMessage_Kickoff{Kickoff: &Kickoff{}}}
	},
	CMD_GATEWAY_CONN_CLOSE: func() *InMessage {
		return &InMessage{OneOf: &InMessage_GatewayConnClose{GatewayConnClose: &GatewayConnClose{}}}
	},
	CMD_RANK_INSERT: func() *InMessage {
		return &InMessage{OneOf: &InMessage_RankInsert{RankInsert: &RankInsert{}}}
	},
	CMD_RANK_SYNC: func() *InMessage {
		return &InMessage{OneOf: &InMessage_RankSync{RankSync: &RankSync{}}}
	},
	CMD_GM_ADD_RESOURCE: func() *InMessage {
		return &InMessage{OneOf: &InMessage_GmAddResource{GmAddResource: &GmAddResource{}}}
	},
	CMD_GUILD_SYNC: func() *InMessage {
		return &InMessage{OneOf: &InMessage_GuildSync{GuildSync: &GuildSync{}}}
	},
	CMD_GUILD_CREATE: func() *InMessage {
		return &InMessage{OneOf: &InMessage_GuildCreate{GuildCreate: &GuildCreate{}}}
	},
	CMD_GUILD_DISMISS: func() *InMessage {
		return &InMessage{OneOf: &InMessage_GuildDismiss{GuildDismiss: &GuildDismiss{}}}
	},
	CMD_GUILD_APPLY_JOIN: func() *InMessage {
		return &InMessage{OneOf: &InMessage_GuildApplyJoin{GuildApplyJoin: &GuildApplyJoin{}}}
	},
	CMD_GUILD_APPLY_CHECK: func() *InMessage {
		return &InMessage{OneOf: &InMessage_GuildApplyCheck{GuildApplyCheck: &GuildApplyCheck{}}}
	},
	CMD_GUILD_LEAVE: func() *InMessage {
		return &InMessage{OneOf: &InMessage_GuildLeave{GuildLeave: &GuildLeave{}}}
	},
	CMD_GUILD_KICK: func() *InMessage {
		return &InMessage{OneOf: &InMessage_GuildKick{GuildKick: &GuildKick{}}}
	},
	CMD_GUILD_MESSAGE: func() *InMessage {
		return &InMessage{OneOf: &InMessage_GuildMessage{GuildMessage: &GuildMessage{}}}
	},
	CMD_G2R_FRIEND_RECOMMEND: func() *InMessage {
		return &InMessage{OneOf: &InMessage_G2RFriendRecommend{G2RFriendRecommend: &G2R_FriendRecommend{}}}
	},
	CMD_R2G_FRIEND_RECOMMEND: func() *InMessage {
		return &InMessage{OneOf: &InMessage_R2GFriendRecommend{R2GFriendRecommend: &R2G_FriendRecommend{}}}
	},
	CMD_FRIEND_SEARCH_REQ: func() *InMessage {
		return &InMessage{OneOf: &InMessage_FriendSearchReq{FriendSearchReq: &FriendSearchReq{}}}
	},
	CMD_FRIEND_SEARCH_RES: func() *InMessage {
		return &InMessage{OneOf: &InMessage_FriendSearchRes{FriendSearchRes: &FriendSearchRes{}}}
	},
	CMD_FRIEND_CHECK: func() *InMessage {
		return &InMessage{OneOf: &InMessage_FriendCheck{FriendCheck: &FriendCheck{}}}
	},
	CMD_FRIEND_CHECK_RES: func() *InMessage {
		return &InMessage{OneOf: &InMessage_FriendCheckRes{FriendCheckRes: &FriendCheckRes{}}}
	},
	CMD_ROUTER_NODE_REPORT: func() *InMessage {
		return &InMessage{OneOf: &InMessage_RouterNodeReport{RouterNodeReport: &RouterNodeReport{}}}
	},
	CMD_ROUTER_CHOOSE_REQUEST: func() *InMessage {
		return &InMessage{OneOf: &InMessage_RouterChooseRequest{RouterChooseRequest: &RouterChooseRequest{}}}
	},
	CMD_ROUTER_CHOOSE_RESPONSE: func() *InMessage {
		return &InMessage{OneOf: &InMessage_RouterChooseResponse{RouterChooseResponse: &RouterChooseResponse{}}}
	},
	CMD_SB_TAKE_ROOM: func() *InMessage {
		return &InMessage{OneOf: &InMessage_SbTakeRoom{SbTakeRoom: &SB_TakeRoom{}}}
	},
	CMD_BS_TAKE_ROOM: func() *InMessage {
		return &InMessage{OneOf: &InMessage_BsTakeRoom{BsTakeRoom: &BS_TakeRoom{}}}
	},
}

func GetInMessage(cmd Cmd) *InMessage {
	if f := g_InMessageFactories[cmd]; f != nil {
		return f()
	}
	return nil
}

type unmarshalfunc func([]byte) (proto.Message, error)

var g_Unmarshals = map[Cmd]unmarshalfunc{
	CMD_KICKOFF: func(data []byte) (proto.Message, error) {
		m := &Kickoff{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_GATEWAY_CONN_CLOSE: func(data []byte) (proto.Message, error) {
		m := &GatewayConnClose{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_RANK_INSERT: func(data []byte) (proto.Message, error) {
		m := &RankInsert{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_RANK_SYNC: func(data []byte) (proto.Message, error) {
		m := &RankSync{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_GM_ADD_RESOURCE: func(data []byte) (proto.Message, error) {
		m := &GmAddResource{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_GUILD_SYNC: func(data []byte) (proto.Message, error) {
		m := &GuildSync{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_GUILD_CREATE: func(data []byte) (proto.Message, error) {
		m := &GuildCreate{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_GUILD_DISMISS: func(data []byte) (proto.Message, error) {
		m := &GuildDismiss{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_GUILD_APPLY_JOIN: func(data []byte) (proto.Message, error) {
		m := &GuildApplyJoin{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_GUILD_APPLY_CHECK: func(data []byte) (proto.Message, error) {
		m := &GuildApplyCheck{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_GUILD_LEAVE: func(data []byte) (proto.Message, error) {
		m := &GuildLeave{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_GUILD_KICK: func(data []byte) (proto.Message, error) {
		m := &GuildKick{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_GUILD_MESSAGE: func(data []byte) (proto.Message, error) {
		m := &GuildMessage{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_G2R_FRIEND_RECOMMEND: func(data []byte) (proto.Message, error) {
		m := &G2R_FriendRecommend{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_R2G_FRIEND_RECOMMEND: func(data []byte) (proto.Message, error) {
		m := &R2G_FriendRecommend{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_FRIEND_SEARCH_REQ: func(data []byte) (proto.Message, error) {
		m := &FriendSearchReq{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_FRIEND_SEARCH_RES: func(data []byte) (proto.Message, error) {
		m := &FriendSearchRes{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_FRIEND_CHECK: func(data []byte) (proto.Message, error) {
		m := &FriendCheck{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_FRIEND_CHECK_RES: func(data []byte) (proto.Message, error) {
		m := &FriendCheckRes{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_ROUTER_NODE_REPORT: func(data []byte) (proto.Message, error) {
		m := &RouterNodeReport{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_ROUTER_CHOOSE_REQUEST: func(data []byte) (proto.Message, error) {
		m := &RouterChooseRequest{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_ROUTER_CHOOSE_RESPONSE: func(data []byte) (proto.Message, error) {
		m := &RouterChooseResponse{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_SB_TAKE_ROOM: func(data []byte) (proto.Message, error) {
		m := &SB_TakeRoom{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_BS_TAKE_ROOM: func(data []byte) (proto.Message, error) {
		m := &BS_TakeRoom{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
}

func Unmarshal(cmd Cmd, data []byte) (proto.Message, error) {
	if f := g_Unmarshals[cmd]; f != nil {
		return f(data)
	}
	return nil, fmt.Errorf("not find unmarshaler of cmd:%d", cmd)
}
