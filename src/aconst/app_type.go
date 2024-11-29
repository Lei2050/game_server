package aconst

const (
	AppTypeMin          uint32 = 0 //无效
	AppTypeGateway      uint32 = 1 //网关
	AppTypeGame         uint32 = 2 //游戏业务进程
	AppTypeGameRouter   uint32 = 3 //游戏业务进程
	AppTypeCross        uint32 = 4 //跨服活动服务器
	AppTypeGuild        uint32 = 5 //工会服务器
	AppTypeBattle       uint32 = 6 //战斗服务器
	AppTypeBattleRouter uint32 = 7 //战斗服务器网关
	AppTypeRankServer   uint32 = 8 //排行榜服务器
	AppTypeMax          uint32 = 9 //不能大于这个数
)

var AppName = map[uint32]string{
	AppTypeGateway:      "GATEWAY",
	AppTypeGame:         "GAME",
	AppTypeGameRouter:   "GAME_ROUTER",
	AppTypeCross:        "CROSS",
	AppTypeGuild:        "GUILD",
	AppTypeBattle:       "BATTLE",
	AppTypeBattleRouter: "BATTLE_ROUTER",
}
