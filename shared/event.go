package shared

import "terminal/shared/types"

// GameEvent 表示游戏事件
type GameEvent struct {
	Type     types.EventType
	PlayerID string // 可选，指定玩家
	MsgType  uint32
	Data     []byte
}

// EventBus 定义事件总线接口
type EventBus interface {
	Emit(event GameEvent)
}
