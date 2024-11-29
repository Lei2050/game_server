//AUTO GENERATED. DO NOT EDIT !!!
package out

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

//type Cmd = int

const (
	CMD_CS_LOGIN               Cmd = 101001
	CMD_CS_CREATE              Cmd = 101002
	CMD_CS_GM                  Cmd = 101004
	CMD_CS_READY               Cmd = 101005
	CMD_CS_HERO_DRAW           Cmd = 102001
	CMD_CS_HERO_LEVEL_UP       Cmd = 102002
	CMD_CS_HERO_STAR_UP        Cmd = 102003
	CMD_CS_HERO_LIST           Cmd = 102004
	CMD_CS_BAG_LIST            Cmd = 103001
	CMD_CS_EXPAND_BAG          Cmd = 103002
	CMD_CS_USE_SINGLEPICKBOX   Cmd = 103003
	CMD_CS_SALE_ITEM           Cmd = 103004
	CMD_CS_DRAW                Cmd = 104001
	CMD_CS_DRAW_REWARD         Cmd = 104002
	CMD_CS_RANK_LIST           Cmd = 105001
	CMD_CS_RANK_LIKE           Cmd = 105002
	CMD_CS_MAIL_LIST           Cmd = 106001
	CMD_CS_READ_MAIL           Cmd = 106002
	CMD_CS_RECEIVE_MAIL        Cmd = 106003
	CMD_CS_DELETE_MAIL         Cmd = 106004
	CMD_CS_GET_SHOP            Cmd = 107001
	CMD_CS_SHOP_BUY            Cmd = 107002
	CMD_CS_SHOP_REFRESH        Cmd = 107003
	CMD_CS_TOWER_FIGHT         Cmd = 108001
	CMD_CS_TOWER_LIKE          Cmd = 108002
	CMD_CS_GUILD_CREATE        Cmd = 109001
	CMD_CS_GUILD_DISMISS       Cmd = 109002
	CMD_CS_GUILD_APPLY_JOIN    Cmd = 109003
	CMD_CS_GUILD_APPLY_CHECK   Cmd = 109004
	CMD_CS_GUILD_LEAVE         Cmd = 109005
	CMD_CS_GUILD_KICK          Cmd = 109006
	CMD_CS_GUILD_MESSAGE       Cmd = 109007
	CMD_CS_GET_GUILDLIST       Cmd = 109008
	CMD_CS_GET_FRIENDLIST      Cmd = 110001
	CMD_CS_FRIEND_RECOMMEND    Cmd = 110002
	CMD_CS_FRIEND_SEARCH       Cmd = 110003
	CMD_CS_FRIEND_APPLY        Cmd = 110004
	CMD_CS_FRIEND_APPLY_HANDLE Cmd = 110005
	CMD_CS_ADD_BLACKLIST       Cmd = 110006
	CMD_CS_BATTLE_HEART_BEAT   Cmd = 150000
	CMD_CS_ENTER_STAGE         Cmd = 150001
	CMD_CS_BATTLE_LOGIN        Cmd = 150002
	CMD_CS_BATTLE_READY        Cmd = 150003
	CMD_CS_BATTLE_GM           Cmd = 150004
	CMD_CS_MOVE                Cmd = 150005
	CMD_CS_QUICK_TAKE_ROOM     Cmd = 150009
	CMD_CS_COLLECTION_RESOURCE Cmd = 150010
	CMD_CS_CAST_SKILL          Cmd = 150016
	CMD_CS_SYNC_DAMAGE         Cmd = 150017
	CMD_CS_RELOAD_GUN          Cmd = 150018
	CMD_CS_SHIFT_GUN           Cmd = 150020
	CMD_CS_STOP_MOVE           Cmd = 150021
	CMD_CS_STAGE_RESTART       Cmd = 150022
	CMD_CS_PICK_ABILITY        Cmd = 150035
	CMD_CS_TASK_LIST           Cmd = 150036
	CMD_CS_SHOT_MOVE           Cmd = 150038
	CMD_CS_ENTER_SCENE         Cmd = 150040
	CMD_CS_SKILL_CAST_ON_POS   Cmd = 150041
)

type CSMessageFactory func() *CSMessage

var g_CSMessageFactories = map[Cmd]CSMessageFactory{
	CMD_CS_LOGIN: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsLogin{CsLogin: &CS_Login{}}}
	},
	CMD_CS_CREATE: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsCreate{CsCreate: &CS_Create{}}}
	},
	CMD_CS_GM: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsGm{CsGm: &CS_Gm{}}}
	},
	CMD_CS_READY: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsReady{CsReady: &CS_Ready{}}}
	},
	CMD_CS_HERO_DRAW: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsHeroDraw{CsHeroDraw: &CS_HeroDraw{}}}
	},
	CMD_CS_HERO_LEVEL_UP: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsHeroLevelUp{CsHeroLevelUp: &CS_HeroLevelUp{}}}
	},
	CMD_CS_HERO_STAR_UP: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsHeroStarUp{CsHeroStarUp: &CS_HeroStarUp{}}}
	},
	CMD_CS_HERO_LIST: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsHeroList{CsHeroList: &CS_HeroList{}}}
	},
	CMD_CS_BAG_LIST: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsBagList{CsBagList: &CS_BagList{}}}
	},
	CMD_CS_EXPAND_BAG: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsExpandBag{CsExpandBag: &CS_ExpandBag{}}}
	},
	CMD_CS_USE_SINGLEPICKBOX: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsUseSinglepickbox{CsUseSinglepickbox: &CS_UseSinglePickBox{}}}
	},
	CMD_CS_SALE_ITEM: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsSaleItem{CsSaleItem: &CS_SaleItem{}}}
	},
	CMD_CS_DRAW: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsDraw{CsDraw: &CS_Draw{}}}
	},
	CMD_CS_DRAW_REWARD: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsDrawReward{CsDrawReward: &CS_DrawReward{}}}
	},
	CMD_CS_RANK_LIST: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsRankList{CsRankList: &CS_GetRankList{}}}
	},
	CMD_CS_RANK_LIKE: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsRankLike{CsRankLike: &CS_RankLike{}}}
	},
	CMD_CS_MAIL_LIST: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsMailList{CsMailList: &CS_GetMailList{}}}
	},
	CMD_CS_READ_MAIL: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsReadMail{CsReadMail: &CS_ReadMail{}}}
	},
	CMD_CS_RECEIVE_MAIL: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsReceiveMail{CsReceiveMail: &CS_ReceiveMail{}}}
	},
	CMD_CS_DELETE_MAIL: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsDeleteMail{CsDeleteMail: &CS_DeleteMail{}}}
	},
	CMD_CS_GET_SHOP: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsGetShop{CsGetShop: &CS_GetShop{}}}
	},
	CMD_CS_SHOP_BUY: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsShopBuy{CsShopBuy: &CS_ShopBuy{}}}
	},
	CMD_CS_SHOP_REFRESH: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsShopRefresh{CsShopRefresh: &CS_RefreshShop{}}}
	},
	CMD_CS_TOWER_FIGHT: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsTowerFight{CsTowerFight: &CS_TowerFight{}}}
	},
	CMD_CS_TOWER_LIKE: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsTowerLike{CsTowerLike: &CS_TowerLike{}}}
	},
	CMD_CS_GUILD_CREATE: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsGuildCreate{CsGuildCreate: &CS_GuildCreate{}}}
	},
	CMD_CS_GUILD_DISMISS: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsGuildDismiss{CsGuildDismiss: &CS_GuildDismiss{}}}
	},
	CMD_CS_GUILD_APPLY_JOIN: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsGuildApplyJoin{CsGuildApplyJoin: &CS_GuildApplyJoin{}}}
	},
	CMD_CS_GUILD_APPLY_CHECK: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsGuildApplyCheck{CsGuildApplyCheck: &CS_GuildApplyCheck{}}}
	},
	CMD_CS_GUILD_LEAVE: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsGuildLeave{CsGuildLeave: &CS_GuildLeave{}}}
	},
	CMD_CS_GUILD_KICK: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsGuildKick{CsGuildKick: &CS_GuildKick{}}}
	},
	CMD_CS_GUILD_MESSAGE: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsGuildMessage{CsGuildMessage: &CS_GuildMessage{}}}
	},
	CMD_CS_GET_GUILDLIST: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsGetGuildlist{CsGetGuildlist: &CS_GetGuildList{}}}
	},
	CMD_CS_GET_FRIENDLIST: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsGetFriendlist{CsGetFriendlist: &CS_GetFriendList{}}}
	},
	CMD_CS_FRIEND_RECOMMEND: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsFriendRecommend{CsFriendRecommend: &CS_FriendRecommend{}}}
	},
	CMD_CS_FRIEND_SEARCH: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsFriendSearch{CsFriendSearch: &CS_FriendSearch{}}}
	},
	CMD_CS_FRIEND_APPLY: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsFriendApply{CsFriendApply: &CS_FriendApply{}}}
	},
	CMD_CS_FRIEND_APPLY_HANDLE: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsFriendApplyHandle{CsFriendApplyHandle: &CS_FriendApplyHandle{}}}
	},
	CMD_CS_ADD_BLACKLIST: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsAddBlacklist{CsAddBlacklist: &CS_AddBalckList{}}}
	},
	CMD_CS_BATTLE_HEART_BEAT: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsBattleHeartBeat{CsBattleHeartBeat: &CS_BattleHeartBeat{}}}
	},
	CMD_CS_ENTER_STAGE: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsEnterStage{CsEnterStage: &CS_EnterStage{}}}
	},
	CMD_CS_BATTLE_LOGIN: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsBattleLogin{CsBattleLogin: &CS_BattleLogin{}}}
	},
	CMD_CS_BATTLE_READY: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsBattleReady{CsBattleReady: &CS_BattleReady{}}}
	},
	CMD_CS_BATTLE_GM: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsBattleGm{CsBattleGm: &CS_BattleGm{}}}
	},
	CMD_CS_MOVE: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsMove{CsMove: &CS_Move{}}}
	},
	CMD_CS_QUICK_TAKE_ROOM: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsQuickTakeRoom{CsQuickTakeRoom: &CS_QuickTakeRoom{}}}
	},
	CMD_CS_COLLECTION_RESOURCE: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsCollectionResource{CsCollectionResource: &CS_CollectionResource{}}}
	},
	CMD_CS_CAST_SKILL: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsCastSkill{CsCastSkill: &CS_CastSkill{}}}
	},
	CMD_CS_SYNC_DAMAGE: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsSyncDamage{CsSyncDamage: &CS_SyncDamage{}}}
	},
	CMD_CS_RELOAD_GUN: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsReloadGun{CsReloadGun: &CS_ReloadGun{}}}
	},
	CMD_CS_SHIFT_GUN: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsShiftGun{CsShiftGun: &CS_ShiftGun{}}}
	},
	CMD_CS_STOP_MOVE: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsStopMove{CsStopMove: &CS_StopMove{}}}
	},
	CMD_CS_STAGE_RESTART: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsStageRestart{CsStageRestart: &CS_StageRestart{}}}
	},
	CMD_CS_PICK_ABILITY: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsPickAbility{CsPickAbility: &CS_PickAbility{}}}
	},
	CMD_CS_TASK_LIST: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsTaskList{CsTaskList: &CS_TaskList{}}}
	},
	CMD_CS_SHOT_MOVE: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsShotMove{CsShotMove: &CS_ShotMove{}}}
	},
	CMD_CS_ENTER_SCENE: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsEnterScene{CsEnterScene: &CS_EnterScene{}}}
	},
	CMD_CS_SKILL_CAST_ON_POS: func() *CSMessage {
		return &CSMessage{OneOf: &CSMessage_CsSkillCastOnPos{CsSkillCastOnPos: &CS_SkillCastOnPos{}}}
	},
}

func GetCSMessage(cmd Cmd) *CSMessage {
	if f := g_CSMessageFactories[cmd]; f != nil {
		return f()
	}
	return nil
}

type unmarshalfunc func([]byte) (proto.Message, error)

var g_Unmarshals = map[Cmd]unmarshalfunc{
	CMD_CS_LOGIN: func(data []byte) (proto.Message, error) {
		m := &CS_Login{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_CREATE: func(data []byte) (proto.Message, error) {
		m := &CS_Create{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_GM: func(data []byte) (proto.Message, error) {
		m := &CS_Gm{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_READY: func(data []byte) (proto.Message, error) {
		m := &CS_Ready{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_HERO_DRAW: func(data []byte) (proto.Message, error) {
		m := &CS_HeroDraw{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_HERO_LEVEL_UP: func(data []byte) (proto.Message, error) {
		m := &CS_HeroLevelUp{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_HERO_STAR_UP: func(data []byte) (proto.Message, error) {
		m := &CS_HeroStarUp{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_HERO_LIST: func(data []byte) (proto.Message, error) {
		m := &CS_HeroList{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_BAG_LIST: func(data []byte) (proto.Message, error) {
		m := &CS_BagList{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_EXPAND_BAG: func(data []byte) (proto.Message, error) {
		m := &CS_ExpandBag{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_USE_SINGLEPICKBOX: func(data []byte) (proto.Message, error) {
		m := &CS_UseSinglePickBox{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_SALE_ITEM: func(data []byte) (proto.Message, error) {
		m := &CS_SaleItem{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_DRAW: func(data []byte) (proto.Message, error) {
		m := &CS_Draw{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_DRAW_REWARD: func(data []byte) (proto.Message, error) {
		m := &CS_DrawReward{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_RANK_LIST: func(data []byte) (proto.Message, error) {
		m := &CS_GetRankList{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_RANK_LIKE: func(data []byte) (proto.Message, error) {
		m := &CS_RankLike{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_MAIL_LIST: func(data []byte) (proto.Message, error) {
		m := &CS_GetMailList{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_READ_MAIL: func(data []byte) (proto.Message, error) {
		m := &CS_ReadMail{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_RECEIVE_MAIL: func(data []byte) (proto.Message, error) {
		m := &CS_ReceiveMail{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_DELETE_MAIL: func(data []byte) (proto.Message, error) {
		m := &CS_DeleteMail{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_GET_SHOP: func(data []byte) (proto.Message, error) {
		m := &CS_GetShop{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_SHOP_BUY: func(data []byte) (proto.Message, error) {
		m := &CS_ShopBuy{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_SHOP_REFRESH: func(data []byte) (proto.Message, error) {
		m := &CS_RefreshShop{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_TOWER_FIGHT: func(data []byte) (proto.Message, error) {
		m := &CS_TowerFight{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_TOWER_LIKE: func(data []byte) (proto.Message, error) {
		m := &CS_TowerLike{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_GUILD_CREATE: func(data []byte) (proto.Message, error) {
		m := &CS_GuildCreate{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_GUILD_DISMISS: func(data []byte) (proto.Message, error) {
		m := &CS_GuildDismiss{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_GUILD_APPLY_JOIN: func(data []byte) (proto.Message, error) {
		m := &CS_GuildApplyJoin{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_GUILD_APPLY_CHECK: func(data []byte) (proto.Message, error) {
		m := &CS_GuildApplyCheck{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_GUILD_LEAVE: func(data []byte) (proto.Message, error) {
		m := &CS_GuildLeave{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_GUILD_KICK: func(data []byte) (proto.Message, error) {
		m := &CS_GuildKick{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_GUILD_MESSAGE: func(data []byte) (proto.Message, error) {
		m := &CS_GuildMessage{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_GET_GUILDLIST: func(data []byte) (proto.Message, error) {
		m := &CS_GetGuildList{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_GET_FRIENDLIST: func(data []byte) (proto.Message, error) {
		m := &CS_GetFriendList{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_FRIEND_RECOMMEND: func(data []byte) (proto.Message, error) {
		m := &CS_FriendRecommend{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_FRIEND_SEARCH: func(data []byte) (proto.Message, error) {
		m := &CS_FriendSearch{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_FRIEND_APPLY: func(data []byte) (proto.Message, error) {
		m := &CS_FriendApply{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_FRIEND_APPLY_HANDLE: func(data []byte) (proto.Message, error) {
		m := &CS_FriendApplyHandle{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_ADD_BLACKLIST: func(data []byte) (proto.Message, error) {
		m := &CS_AddBalckList{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_BATTLE_HEART_BEAT: func(data []byte) (proto.Message, error) {
		m := &CS_BattleHeartBeat{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_ENTER_STAGE: func(data []byte) (proto.Message, error) {
		m := &CS_EnterStage{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_BATTLE_LOGIN: func(data []byte) (proto.Message, error) {
		m := &CS_BattleLogin{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_BATTLE_READY: func(data []byte) (proto.Message, error) {
		m := &CS_BattleReady{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_BATTLE_GM: func(data []byte) (proto.Message, error) {
		m := &CS_BattleGm{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_MOVE: func(data []byte) (proto.Message, error) {
		m := &CS_Move{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_QUICK_TAKE_ROOM: func(data []byte) (proto.Message, error) {
		m := &CS_QuickTakeRoom{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_COLLECTION_RESOURCE: func(data []byte) (proto.Message, error) {
		m := &CS_CollectionResource{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_CAST_SKILL: func(data []byte) (proto.Message, error) {
		m := &CS_CastSkill{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_SYNC_DAMAGE: func(data []byte) (proto.Message, error) {
		m := &CS_SyncDamage{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_RELOAD_GUN: func(data []byte) (proto.Message, error) {
		m := &CS_ReloadGun{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_SHIFT_GUN: func(data []byte) (proto.Message, error) {
		m := &CS_ShiftGun{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_STOP_MOVE: func(data []byte) (proto.Message, error) {
		m := &CS_StopMove{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_STAGE_RESTART: func(data []byte) (proto.Message, error) {
		m := &CS_StageRestart{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_PICK_ABILITY: func(data []byte) (proto.Message, error) {
		m := &CS_PickAbility{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_TASK_LIST: func(data []byte) (proto.Message, error) {
		m := &CS_TaskList{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_SHOT_MOVE: func(data []byte) (proto.Message, error) {
		m := &CS_ShotMove{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_ENTER_SCENE: func(data []byte) (proto.Message, error) {
		m := &CS_EnterScene{}
		if err := proto.Unmarshal(data, m); err != nil {
			return nil, err
		}
		return m, nil
	},
	CMD_CS_SKILL_CAST_ON_POS: func(data []byte) (proto.Message, error) {
		m := &CS_SkillCastOnPos{}
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
