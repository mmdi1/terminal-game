package iface

import (
	"sync"
	"terminal/shared"
)

// 游戏接口
type Game interface {
	Init(eventBus shared.EventBus, players map[string]*Player) // 初始化游戏
	HandleAction(action shared.GameAction)                     // 处理玩家操作
	GetState() shared.GameState                                // 获取游戏状态
	BroadcastState()                                           // 广播状态
}

// 游戏工厂函数
type GameFactory func() Game

// 游戏工厂注册
var (
	gameFactories = make(map[string]GameFactory)
	factoryMutex  sync.RWMutex
)

// 注册游戏工厂
func RegisterGameFactory(gameType string, factory GameFactory) {
	factoryMutex.Lock()
	defer factoryMutex.Unlock()
	gameFactories[gameType] = factory
}

// 创建游戏实例
func CreateGame(gameType string) Game {
	factoryMutex.RLock()
	defer factoryMutex.RUnlock()
	if factory, exists := gameFactories[gameType]; exists {
		return factory()
	}
	return nil
}
