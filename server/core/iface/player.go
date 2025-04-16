package iface

import "github.com/aceld/zinx/ziface"

// 玩家
type Player struct {
	ID   string
	Conn ziface.IConnection
}
