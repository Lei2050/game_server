//AUTO GENERATED. DO NOT EDIT !!!
package out

//type Cmd = int

const (
	CMD_SC_LOGIN                       Cmd = 101001
	CMD_SC_CREATE                      Cmd = 101002
	CMD_SC_KICKOFF                     Cmd = 101003
	CMD_SC_GM                          Cmd = 101004
	CMD_SC_READY                       Cmd = 101005
	CMD_SC_HERO_DRAW                   Cmd = 102001
	CMD_SC_HERO_LEVEL_UP               Cmd = 102002
	CMD_SC_HERO_STAR_UP                Cmd = 102003
	CMD_SC_HERO_LIST                   Cmd = 102004
	CMD_SC_BAG_LIST                    Cmd = 103001
	CMD_SC_EXPAND_BAG                  Cmd = 103002
	CMD_SC_USE_SINGLEPICKBOX           Cmd = 103003
	CMD_SC_SALE_ITEM                   Cmd = 1030004
	CMD_SC_RESOURCE_NOTIFY             Cmd = 103005
	CMD_SC_RESOURCE_LIST               Cmd = 103006
	CMD_SC_ERR_RESPONSE                Cmd = 103020
	CMD_SC_STATISTICS_LIST             Cmd = 103030
	CMD_SC_LIMIT_LIST                  Cmd = 103031
	CMD_SC_DRAW_REWARD                 Cmd = 104002
	CMD_SC_RANK_LIST                   Cmd = 105001
	CMD_SC_MAIL_LIST                   Cmd = 106001
	CMD_SC_MAIL_NOTIFY                 Cmd = 106002
	CMD_SC_GET_SHOP                    Cmd = 107001
	CMD_SC_TOWER_INFO                  Cmd = 108001
	CMD_SC_TOWER_FIGHT                 Cmd = 108002
	CMD_SC_GET_GUILDLIST               Cmd = 109008
	CMD_SC_GET_FRIENDLIST              Cmd = 110001
	CMD_SC_FRIEND_RECOMMEND            Cmd = 110002
	CMD_SC_FRIEND_SEARCH               Cmd = 110003
	CMD_SC_BATTLE_HEART_BEAT           Cmd = 150000
	CMD_SC_ENTER_STAGE                 Cmd = 150001
	CMD_SC_BATTLE_LOGIN                Cmd = 150002
	CMD_SC_BATTLE_READY                Cmd = 150003
	CMD_SC_BATTLE_GM                   Cmd = 150004
	CMD_SC_MOVE                        Cmd = 150005
	CMD_SC_ADD_UNIT                    Cmd = 150006
	CMD_SC_REMOVE_UNIT                 Cmd = 150007
	CMD_SC_SYNC_UNIT_INFO              Cmd = 150008
	CMD_SC_QUICK_TAKE_ROOM             Cmd = 150009
	CMD_SC_COLLECTION_RESOURCE         Cmd = 150010
	CMD_SC_STATE_CHANGE                Cmd = 150011
	CMD_SC_AWARDS_GET                  Cmd = 150012
	CMD_SC_MAIN_UNIT                   Cmd = 150013
	CMD_SC_UNIT_ADD_ITEM               Cmd = 150014
	CMD_SC_UNIT_SYNC_ITEM              Cmd = 150015
	CMD_SC_CAST_SKILL                  Cmd = 150016
	CMD_SC_SYNC_DAMAGE                 Cmd = 150017
	CMD_SC_SYNC_GUN_INFO               Cmd = 150019
	CMD_SC_SHIFT_GUN                   Cmd = 150020
	CMD_SC_STAGE_RESTART               Cmd = 150022
	CMD_SC_STAGE_PHASE_TIMING          Cmd = 150023
	CMD_SC_STAGE_MONSTER_COMING_TIMING Cmd = 150024
	CMD_SC_ADD_BUFF                    Cmd = 150025
	CMD_SC_DEL_BUFF                    Cmd = 150026
	CMD_SC_ADD_MAGICFIELD              Cmd = 150027
	CMD_SC_DEL_MAGICFIELD              Cmd = 150028
	CMD_SC_UNIT_INTERRUPT_SKILL        Cmd = 150029
	CMD_SC_TIPS                        Cmd = 150030
	CMD_SC_SYNC_ALL_ITEMS              Cmd = 150031
	CMD_SC_GET_EXP                     Cmd = 150033
	CMD_SC_GET_ABILITY                 Cmd = 150034
	CMD_SC_TASK_LIST                   Cmd = 150036
	CMD_SC_SYNC_TASKS                  Cmd = 150037
	CMD_SC_MONSTER_WAVE_COMING         Cmd = 150039
	CMD_SC_ENTER_SCENE                 Cmd = 150040
	CMD_SC_MONSTER_WAKEUP              Cmd = 150042
	CMD_SC_CAST_BULLET                 Cmd = 150043
)

type SCMessageFactory func() *SCMessage

var g_SCMessageFactories = map[Cmd]SCMessageFactory{
	CMD_SC_LOGIN: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScLogin{ScLogin: &SC_Login{}}}
	},
	CMD_SC_CREATE: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScCreate{ScCreate: &SC_Create{}}}
	},
	CMD_SC_KICKOFF: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScKickoff{ScKickoff: &SC_Kickoff{}}}
	},
	CMD_SC_GM: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScGm{ScGm: &SC_Gm{}}}
	},
	CMD_SC_READY: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScReady{ScReady: &SC_Ready{}}}
	},
	CMD_SC_HERO_DRAW: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScHeroDraw{ScHeroDraw: &SC_HeroDraw{}}}
	},
	CMD_SC_HERO_LEVEL_UP: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScHeroLevelUp{ScHeroLevelUp: &SC_HeroLevelUp{}}}
	},
	CMD_SC_HERO_STAR_UP: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScHeroStarUp{ScHeroStarUp: &SC_HeroStarUp{}}}
	},
	CMD_SC_HERO_LIST: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScHeroList{ScHeroList: &SC_HeroList{}}}
	},
	CMD_SC_BAG_LIST: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScBagList{ScBagList: &SC_BagList{}}}
	},
	CMD_SC_EXPAND_BAG: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScExpandBag{ScExpandBag: &SC_ExpandBag{}}}
	},
	CMD_SC_USE_SINGLEPICKBOX: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScUseSinglepickbox{ScUseSinglepickbox: &SC_UseSinglePickBox{}}}
	},
	CMD_SC_SALE_ITEM: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScSaleItem{ScSaleItem: &SC_SaleItem{}}}
	},
	CMD_SC_RESOURCE_NOTIFY: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScResourceNotify{ScResourceNotify: &SC_ResourceNotify{}}}
	},
	CMD_SC_RESOURCE_LIST: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScResourceList{ScResourceList: &SC_ResourceList{}}}
	},
	CMD_SC_ERR_RESPONSE: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScErrResponse{ScErrResponse: &SC_ErrResponse{}}}
	},
	CMD_SC_STATISTICS_LIST: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScStatisticsList{ScStatisticsList: &SC_StatisticsList{}}}
	},
	CMD_SC_LIMIT_LIST: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScLimitList{ScLimitList: &SC_LimitList{}}}
	},
	CMD_SC_DRAW_REWARD: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScDrawReward{ScDrawReward: &SC_DrawReward{}}}
	},
	CMD_SC_RANK_LIST: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScRankList{ScRankList: &SC_GetRankList{}}}
	},
	CMD_SC_MAIL_LIST: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScMailList{ScMailList: &SC_GetMailList{}}}
	},
	CMD_SC_MAIL_NOTIFY: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScMailNotify{ScMailNotify: &SC_MailNotify{}}}
	},
	CMD_SC_GET_SHOP: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScGetShop{ScGetShop: &SC_GetShop{}}}
	},
	CMD_SC_TOWER_INFO: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScTowerInfo{ScTowerInfo: &SC_GetTowerInfo{}}}
	},
	CMD_SC_TOWER_FIGHT: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScTowerFight{ScTowerFight: &SC_TowerFight{}}}
	},
	CMD_SC_GET_GUILDLIST: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScGetGuildlist{ScGetGuildlist: &SC_GetGuildList{}}}
	},
	CMD_SC_GET_FRIENDLIST: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScGetFriendlist{ScGetFriendlist: &SC_GetFriendList{}}}
	},
	CMD_SC_FRIEND_RECOMMEND: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScFriendRecommend{ScFriendRecommend: &SC_FriendRecommend{}}}
	},
	CMD_SC_FRIEND_SEARCH: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScFriendSearch{ScFriendSearch: &SC_FriendSearch{}}}
	},
	CMD_SC_BATTLE_HEART_BEAT: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScBattleHeartBeat{ScBattleHeartBeat: &SC_BattleHeartBeat{}}}
	},
	CMD_SC_ENTER_STAGE: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScEnterStage{ScEnterStage: &SC_EnterStage{}}}
	},
	CMD_SC_BATTLE_LOGIN: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScBattleLogin{ScBattleLogin: &SC_BattleLogin{}}}
	},
	CMD_SC_BATTLE_READY: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScBattleReady{ScBattleReady: &SC_BattleReady{}}}
	},
	CMD_SC_BATTLE_GM: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScBattleGm{ScBattleGm: &SC_BattleGm{}}}
	},
	CMD_SC_MOVE: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScMove{ScMove: &SC_Move{}}}
	},
	CMD_SC_ADD_UNIT: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScAddUnit{ScAddUnit: &SC_AddUnit{}}}
	},
	CMD_SC_REMOVE_UNIT: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScRemoveUnit{ScRemoveUnit: &SC_RemoveUnit{}}}
	},
	CMD_SC_SYNC_UNIT_INFO: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScSyncUnitInfo{ScSyncUnitInfo: &SC_SyncUnitInfo{}}}
	},
	CMD_SC_QUICK_TAKE_ROOM: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScQuickTakeRoom{ScQuickTakeRoom: &SC_QuickTakeRoom{}}}
	},
	CMD_SC_COLLECTION_RESOURCE: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScCollectionResource{ScCollectionResource: &SC_CollectionResource{}}}
	},
	CMD_SC_STATE_CHANGE: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScStateChange{ScStateChange: &SC_StateChange{}}}
	},
	CMD_SC_AWARDS_GET: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScAwardsGet{ScAwardsGet: &SC_AwardsGet{}}}
	},
	CMD_SC_MAIN_UNIT: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScMainUnit{ScMainUnit: &SC_MainUnit{}}}
	},
	CMD_SC_UNIT_ADD_ITEM: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScUnitAddItem{ScUnitAddItem: &SC_UnitAddItem{}}}
	},
	CMD_SC_UNIT_SYNC_ITEM: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScUnitSyncItem{ScUnitSyncItem: &SC_UnitSyncItem{}}}
	},
	CMD_SC_CAST_SKILL: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScCastSkill{ScCastSkill: &SC_CastSkill{}}}
	},
	CMD_SC_SYNC_DAMAGE: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScSyncDamage{ScSyncDamage: &SC_SyncDamage{}}}
	},
	CMD_SC_SYNC_GUN_INFO: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScSyncGunInfo{ScSyncGunInfo: &SC_SyncGunInfo{}}}
	},
	CMD_SC_SHIFT_GUN: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScShiftGun{ScShiftGun: &SC_ShiftGun{}}}
	},
	CMD_SC_STAGE_RESTART: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScStageRestart{ScStageRestart: &SC_StageRestart{}}}
	},
	CMD_SC_STAGE_PHASE_TIMING: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScStagePhaseTiming{ScStagePhaseTiming: &SC_StagePhaseTiming{}}}
	},
	CMD_SC_STAGE_MONSTER_COMING_TIMING: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScStageMonsterComingTiming{ScStageMonsterComingTiming: &SC_StageMonsterComingTiming{}}}
	},
	CMD_SC_ADD_BUFF: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScAddBuff{ScAddBuff: &SC_AddBuff{}}}
	},
	CMD_SC_DEL_BUFF: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScDelBuff{ScDelBuff: &SC_DelBuff{}}}
	},
	CMD_SC_ADD_MAGICFIELD: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScAddMagicfield{ScAddMagicfield: &SC_AddMagicField{}}}
	},
	CMD_SC_DEL_MAGICFIELD: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScDelMagicfield{ScDelMagicfield: &SC_DelMagicField{}}}
	},
	CMD_SC_UNIT_INTERRUPT_SKILL: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScUnitInterruptSkill{ScUnitInterruptSkill: &SC_UnitInterruptSkill{}}}
	},
	CMD_SC_TIPS: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScTips{ScTips: &SC_Tips{}}}
	},
	CMD_SC_SYNC_ALL_ITEMS: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScSyncAllItems{ScSyncAllItems: &SC_SyncAllItems{}}}
	},
	CMD_SC_GET_EXP: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScGetExp{ScGetExp: &SC_GetExp{}}}
	},
	CMD_SC_GET_ABILITY: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScGetAbility{ScGetAbility: &SC_GetAbility{}}}
	},
	CMD_SC_TASK_LIST: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScTaskList{ScTaskList: &SC_TaskList{}}}
	},
	CMD_SC_SYNC_TASKS: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScSyncTasks{ScSyncTasks: &SC_SyncTask{}}}
	},
	CMD_SC_MONSTER_WAVE_COMING: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScMonsterWaveComing{ScMonsterWaveComing: &SC_MonsterWaveComing{}}}
	},
	CMD_SC_ENTER_SCENE: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScEnterScene{ScEnterScene: &SC_EnterScene{}}}
	},
	CMD_SC_MONSTER_WAKEUP: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScMonsterWakeup{ScMonsterWakeup: &SC_MonsterWakeUp{}}}
	},
	CMD_SC_CAST_BULLET: func() *SCMessage {
		return &SCMessage{OneOf: &SCMessage_ScCastBullet{ScCastBullet: &SC_CastBullet{}}}
	},
}

func GetSCMessage(cmd Cmd) *SCMessage {
	if f := g_SCMessageFactories[cmd]; f != nil {
		return f()
	}
	return nil
}
